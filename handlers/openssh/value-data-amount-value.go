package openssh

import (
	docvalues "config-lsp/doc-values"
	"fmt"
	"regexp"
	"strconv"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

var dataAmountCheckPattern = regexp.MustCompile(`(?i)^(\d+)([KMG])$`)

type InvalidDataAmountError struct{}

func (e InvalidDataAmountError) Error() string {
	return "Data amount is invalid. It must be in the form of: <number>[K|M|G]"
}

type DataAmountValue struct{}

func (v DataAmountValue) GetTypeDescription() []string {
	return []string{"Data amount"}
}

func (v DataAmountValue) CheckIsValid(value string) error {
	if !dataAmountCheckPattern.MatchString(value) {
		return InvalidDataAmountError{}
	}

	return nil
}

func calculateLineToKilobyte(value string, unit string) string {
	val, err := strconv.Atoi(value)

	if err != nil {
		return "<unknown>"
	}

	switch unit {
	case "K":
		return strconv.Itoa(val)
	case "M":
		return strconv.Itoa(val * 1000)
	case "G":
		return strconv.Itoa(val * 1000 * 1000)
	default:
		return "<unknown>"
	}
}

func (v DataAmountValue) FetchCompletions(line string, cursor uint32) []protocol.CompletionItem {
	completions := make([]protocol.CompletionItem, 0)

	if line != "" && !dataAmountCheckPattern.MatchString(line) {
		kind := protocol.CompletionItemKindValue

		completions = append(
			completions,
			protocol.CompletionItem{
				Label:         line + "K",
				Kind:          &kind,
				Documentation: line + " kilobytes",
			},
			protocol.CompletionItem{
				Label:         line + "M",
				Kind:          &kind,
				Documentation: fmt.Sprintf("%s megabytes (%s kilobytes)", line, calculateLineToKilobyte(line, "M")),
			},
			protocol.CompletionItem{
				Label:         line + "G",
				Kind:          &kind,
				Documentation: fmt.Sprintf("%s gigabytes (%s kilobytes)", line, calculateLineToKilobyte(line, "G")),
			},
		)
	}

	if line == "" || isJustDigitsPattern.MatchString(line) {
		completions = append(
			completions,
			docvalues.GenerateBase10Completions(line)...,
		)
	}

	return completions
}
