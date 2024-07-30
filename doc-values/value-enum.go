package docvalues

import (
	"fmt"
	"strings"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

type ValueNotInEnumError struct {
	AvailableValues []string
	ProvidedValue   string
}

func (e ValueNotInEnumError) Error() string {
	return fmt.Sprintf("This value is not valid. Select one from: %s", strings.Join(e.AvailableValues, ","))
}

type EnumValue struct {
	Values []string
	// If `true`, the value MUST be one of the values in the Values array
	// Otherwise an error is shown
	// If `false`, the value is just a hint
	EnforceValues bool
}

func (v EnumValue) GetTypeDescription() []string {
	if len(v.Values) == 1 {
		return []string{fmt.Sprintf("'%s'", v.Values[0])}
	}

	lines := make([]string, len(v.Values)+1)
	lines[0] = "Enum of:"

	for index, value := range v.Values {
		lines[index+1] += "\t* " + value
	}

	return lines
}
func (v EnumValue) CheckIsValid(value string) error {
	if !v.EnforceValues {
		return nil
	}

	for _, validValue := range v.Values {
		if validValue == value {
			return nil
		}

	}

	return ValueNotInEnumError{
		ProvidedValue:   value,
		AvailableValues: v.Values,
	}
}
func (v EnumValue) FetchCompletions(line string, cursor uint32) []protocol.CompletionItem {
	completions := make([]protocol.CompletionItem, len(v.Values))

	for index, value := range v.Values {
		textFormat := protocol.InsertTextFormatPlainText
		kind := protocol.CompletionItemKindEnum

		completions[index] = protocol.CompletionItem{
			Label:            value,
			InsertTextFormat: &textFormat,
			Kind:             &kind,
		}
	}

	return completions
}
