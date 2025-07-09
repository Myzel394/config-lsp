package docvalues

import (
	"config-lsp/common"
	"config-lsp/utils"
	"strings"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

type OrValue struct {
	Values []DeprecatedValue
}

func (v OrValue) GetTypeDescription() []string {
	lines := make([]string, 0)

	for _, subValueRaw := range v.Values {
		subValue := subValueRaw.(DeprecatedValue)
		subLines := subValue.GetTypeDescription()

		for index, line := range subLines {
			if strings.HasPrefix(line, "\t*") {
				subLines[index] = "\t" + line
			} else {
				subLines[index] = "\t* " + line
			}
		}

		lines = append(lines, subLines...)
	}

	return append(
		[]string{"One of:"},
		lines...,
	)
}
func (v OrValue) DeprecatedCheckIsValid(value string) []*InvalidValue {
	errors := make([]*InvalidValue, 0)

	for _, subValue := range v.Values {
		valueErrors := subValue.DeprecatedCheckIsValid(value)

		if len(valueErrors) == 0 {
			return nil
		}

		errors = append(errors, valueErrors...)
	}

	return errors
}

func (v OrValue) FetchCompletions(value string, cursor common.CursorPosition) []protocol.CompletionItem {
	// Check for special cases
	if len(v.Values) == 2 {
		switch v.Values[0].(type) {
		case KeyEnumAssignmentValue:
			if cursor > 0 {
				// KeyEnumAssignment + other values
				// If there is an separator, we only want to show
				// the values of the KeyEnumAssignment
				keyEnumValue := v.Values[0].(KeyEnumAssignmentValue)

				index := common.DeprecatedImprovedCursorToIndex(cursor, value, 0)

				_, found := utils.FindPreviousCharacter(
					value,
					keyEnumValue.Separator,
					int(index),
				)

				if found {
					return keyEnumValue.FetchCompletions(value, cursor)
				}
			}
		}
	}

	completions := make([]protocol.CompletionItem, 0)

	for _, subValue := range v.Values {
		completions = append(completions, subValue.FetchCompletions(value, cursor)...)
	}

	return completions
}

func (v OrValue) DeprecatedFetchHoverInfo(line string, cursor uint32) []string {
	for _, subValue := range v.Values {
		valueErrors := subValue.DeprecatedCheckIsValid(line)

		if len(valueErrors) == 0 {
			// Found
			return subValue.DeprecatedFetchHoverInfo(line, cursor)
		}
	}

	return []string{}
}
