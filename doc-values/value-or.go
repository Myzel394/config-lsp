package docvalues

import (
	protocol "github.com/tliron/glsp/protocol_3_16"
	"strings"
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
func (v OrValue) CheckIsValid(value string) error {
	var firstError error = nil

	for _, subValue := range v.Values {
		err := subValue.CheckIsValid(value)

		if err == nil {
			return nil
		} else if firstError == nil {
			firstError = err
		}
	}

	return firstError
}
func (v OrValue) FetchCompletions(line string, cursor uint32) []protocol.CompletionItem {
	completions := make([]protocol.CompletionItem, 0)

	for _, subValue := range v.Values {
		completions = append(completions, subValue.FetchCompletions(line, cursor)...)
	}

	return completions
}
