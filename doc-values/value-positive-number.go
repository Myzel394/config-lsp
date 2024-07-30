package docvalues

import (
	protocol "github.com/tliron/glsp/protocol_3_16"
	"strconv"
)

type NotANumberError struct{}

func (e NotANumberError) Error() string {
	return "This must be number"
}

type NumberIsNotPositiveError struct{}

func (e NumberIsNotPositiveError) Error() string {
	return "This number must be positive"
}

type PositiveNumberValue struct{}

func (v PositiveNumberValue) GetTypeDescription() []string {
	return []string{"A positive number"}
}
func (v PositiveNumberValue) CheckIsValid(value string) error {
	number, err := strconv.Atoi(value)

	if err != nil {
		return NotANumberError{}
	}

	if number < 0 {
		return NumberIsNotPositiveError{}
	}

	return nil
}
func (v PositiveNumberValue) FetchCompletions(line string, cursor uint32) []protocol.CompletionItem {
	return []protocol.CompletionItem{}
}
