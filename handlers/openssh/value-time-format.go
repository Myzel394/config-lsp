package openssh

import (
	docvalues "config-lsp/doc-values"
	"config-lsp/utils"
	"fmt"
	"regexp"
	"strconv"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

var timeFormatCompletionsPattern = regexp.MustCompile(`(?i)^(\d+)([smhdw])$`)
var timeFormatCheckPattern = regexp.MustCompile(`(?i)^(\d+)([smhdw]?)$`)
var isJustDigitsPattern = regexp.MustCompile(`^\d+$`)

type TimeFormatValue struct{}

func (v TimeFormatValue) GetTypeDescription() []string {
	return []string{"Time value"}
}

func (v TimeFormatValue) CheckIsValid(value string) error {
	if !timeFormatCheckPattern.MatchString(value) {
		return docvalues.RegexInvalidError{Regex: timeFormatCheckPattern.String()}
	}

	return nil
}

func calculateInSeconds(value int, unit string) int {
	switch unit {
	case "s":
		return value
	case "m":
		return value * 60
	case "h":
		return value * 60 * 60
	case "d":
		return value * 60 * 60 * 24
	case "w":
		return value * 60 * 60 * 24 * 7
	default:
		return 0
	}
}

func (v TimeFormatValue) FetchCompletions(line string, cursor uint32) []protocol.CompletionItem {
	completions := make([]protocol.CompletionItem, 0)

	if line != "" && !timeFormatCompletionsPattern.MatchString(line) {
		completions = append(
			completions,
			utils.Map(
				[]string{"s", "m", "h", "d", "w"},
				func(unit string) protocol.CompletionItem {
					kind := protocol.CompletionItemKindValue

					unitName := map[string]string{
						"s": "seconds",
						"m": "minutes",
						"h": "hours",
						"d": "days",
						"w": "weeks",
					}[unit]

					var detail string
					value, err := strconv.Atoi(line)

					if err == nil {
						if unit == "s" {
							detail = fmt.Sprintf("%d seconds", value)
						} else {
							detail = fmt.Sprintf("%d %s (%d seconds)", value, unitName, calculateInSeconds(value, unit))
						}
					}

					return protocol.CompletionItem{
						Label:  line + unit,
						Kind:   &kind,
						Detail: &detail,
					}
				},
			)...,
		)
	}

	if line == "" || isJustDigitsPattern.MatchString(line) {
		completions = append(
			completions,
			utils.Map(
				[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				func(index int) protocol.CompletionItem {
					kind := protocol.CompletionItemKindValue

					sortText := strconv.Itoa(index)

					return protocol.CompletionItem{
						Label:    line + strconv.Itoa(index),
						Kind:     &kind,
						SortText: &sortText,
					}
				},
			)...,
		)
	}

	return completions
}
