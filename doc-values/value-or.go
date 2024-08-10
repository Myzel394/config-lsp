package docvalues

import (
	"config-lsp/utils"
	"strings"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

type OrValue struct {
	Values []Value
}

func (v OrValue) GetTypeDescription() []string {
	lines := make([]string, 0)

	for _, subValueRaw := range v.Values {
		subValue := subValueRaw.(Value)
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
func (v OrValue) CheckIsValid(value string) []*InvalidValue {
	errors := make([]*InvalidValue, 0)

	for _, subValue := range v.Values {
		valueErrors := subValue.CheckIsValid(value)

		if len(valueErrors) == 0 {
			return nil
		}

		errors = append(errors, valueErrors...)
	}

	return errors
}
func (v OrValue) FetchCompletions(line string, cursor uint32) []protocol.CompletionItem {
	// Check for special cases
	if len(v.Values) == 2 {
		switch v.Values[0].(type) {
		case KeyEnumAssignmentValue:
			if cursor > 0 {
				// KeyEnumAssignment + other values
				// If there is an separator, we only want to show
				// the values of the KeyEnumAssignment
				keyEnumValue := v.Values[0].(KeyEnumAssignmentValue)

				_, found := utils.FindPreviousCharacter(line, keyEnumValue.Separator, int(cursor-1))

				if found {
					return keyEnumValue.FetchCompletions(line, cursor)
				}
			}
		}
	}

	completions := make([]protocol.CompletionItem, 0)

	for _, subValue := range v.Values {
		completions = append(completions, subValue.FetchCompletions(line, cursor)...)
	}

	return completions
}

func (v OrValue) FetchHoverInfo(line string, cursor uint32) []string {
	for _, subValue := range v.Values {
		valueErrors := subValue.CheckIsValid(line)

		if len(valueErrors) == 0 {
			// Found
			return subValue.FetchHoverInfo(line, cursor)
		}
	}

	return []string{}
}
