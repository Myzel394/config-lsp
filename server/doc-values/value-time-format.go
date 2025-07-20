package docvalues

import (
	"config-lsp/common"
	"config-lsp/utils"
	"fmt"
	"regexp"
	"strconv"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

var timeFormatCompletionsPattern = regexp.MustCompile(`(?i)^(\d+)([smhdw])$`)
var timeFormatCheckPattern = regexp.MustCompile(`(?i)^(\d+)([smhdw]?)$`)
var isJustDigitsPattern = regexp.MustCompile(`^\d+$`)

type InvalidTimeFormatError struct{}

func (e InvalidTimeFormatError) Error() string {
	return "Time format is invalid. It must be in the form of: <number>[s|m|h|d|w]"
}

type TimeFormatValue struct{}

func (v TimeFormatValue) GetTypeDescription() []string {
	return []string{"Time value"}
}

func (v TimeFormatValue) DeprecatedCheckIsValid(value string) []*InvalidValue {
	if !timeFormatCheckPattern.MatchString(value) {
		return []*InvalidValue{
			{
				Err:   InvalidTimeFormatError{},
				Start: 0,
				End:   uint32(len(value)),
			},
		}
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

func (v TimeFormatValue) FetchCompletions(value string, cursor common.CursorPosition) []protocol.CompletionItem {
	completions := make([]protocol.CompletionItem, 0)

	if value != "" && !timeFormatCompletionsPattern.MatchString(value) {
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
					numberValue, err := strconv.Atoi(value)

					if err == nil {
						if unit == "s" {
							detail = fmt.Sprintf("%d seconds", numberValue)
						} else {
							detail = fmt.Sprintf("%d %s (%d seconds)", numberValue, unitName, calculateInSeconds(numberValue, unit))
						}
					}

					return protocol.CompletionItem{
						Label:  value + unit,
						Kind:   &kind,
						Detail: &detail,
					}
				},
			)...,
		)
	}

	if value == "" || isJustDigitsPattern.MatchString(value) {
		completions = append(
			completions,
			GenerateBase10Completions(value)...,
		)
	}

	return completions
}

func (v TimeFormatValue) DeprecatedFetchHoverInfo(line string, cursor uint32) []string {
	return []string{}
}
