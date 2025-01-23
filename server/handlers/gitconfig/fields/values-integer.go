package fields

import (
	docvalues "config-lsp/doc-values"
	"fmt"
	"regexp"
	"strconv"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

var isJustDigitsPattern = regexp.MustCompile(`^\d+$`)
var dataAmountCheckPattern = regexp.MustCompile(`(?i)^(\d+)([kmg])$`)

type InvalidIntegerError struct{}

func (e InvalidIntegerError) Error() string {
	return "Integer is invalid. It must be in the form of: <number>[k|m|g]"
}

type IntegerValue struct{}

func (v IntegerValue) GetTypeDescription() []string {
	return []string{"Integer"}
}

func (v IntegerValue) DeprecatedCheckIsValid(value string) []*docvalues.InvalidValue {
	if !dataAmountCheckPattern.MatchString(value) {
		return []*docvalues.InvalidValue{
			{
				Err:   InvalidIntegerError{},
				Start: 0,
				End:   uint32(len(value)),
			},
		}
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
		return strconv.Itoa(val * 1024)
	case "m":
		return strconv.Itoa(val * 1024 * 1024)
	case "g":
		return strconv.Itoa(val * 1024 * 1024 * 1024)
	default:
		return "<unknown>"
	}
}

func (v IntegerValue) DeprecatedFetchCompletions(line string, cursor uint32) []protocol.CompletionItem {
	completions := make([]protocol.CompletionItem, 0)

	if line != "" && !dataAmountCheckPattern.MatchString(line) {
		kind := protocol.CompletionItemKindValue

		completions = append(
			completions,
			protocol.CompletionItem{
				Label:         line + "k",
				Kind:          &kind,
				Documentation: fmt.Sprintf("scale %s by x1024", line),
			},
			protocol.CompletionItem{
				Label:         line + "m",
				Kind:          &kind,
				Documentation: fmt.Sprintf("scale %s by x1024x1024", line),
			},
			protocol.CompletionItem{
				Label:         line + "g",
				Kind:          &kind,
				Documentation: fmt.Sprintf("scale %s by x1024x1024x1024", line),
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

func (v IntegerValue) DeprecatedFetchHoverInfo(line string, cursor uint32) []string {
	return []string{}
}
