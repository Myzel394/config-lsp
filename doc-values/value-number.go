package docvalues

import (
	"fmt"
	"strconv"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

type NotANumberError struct{}

func (e NotANumberError) Error() string {
	return "This must be a number"
}

type NumberNotInRangeError struct {
	Min *int
	Max *int
}

func (e NumberNotInRangeError) Error() string {
	if e.Min == nil {
		return fmt.Sprintf("This must be at most or less than %d", *e.Max)
	}

	if e.Max == nil {
		return fmt.Sprintf("This must be equal or more than %d", *e.Min)
	}

	return fmt.Sprintf("This must be between %d and %d (inclusive)", *e.Min, *e.Max)
}

type NumberValue struct {
	Min *int
	Max *int
}

func (v NumberValue) GetTypeDescription() []string {
	if v.Min != nil {
		if v.Max != nil {
			return []string{fmt.Sprintf("A number between %d and %d (inclusive)", *v.Min, *v.Max)}
		} else {
			return []string{fmt.Sprintf("A number of at least %d", *v.Min)}
		}
	} else if v.Max != nil {
		return []string{fmt.Sprintf("A number of at most %d", *v.Max)}
	}

	return []string{"A number"}
}
func (v NumberValue) CheckIsValid(value string) []*InvalidValue {
	number, err := strconv.Atoi(value)

	if err != nil {
		return []*InvalidValue{
			{
				Err:   NotANumberError{},
				Start: 0,
				End:   uint32(len(value)),
			},
		}
	}

	if (v.Min != nil && number < *v.Min) || (v.Max != nil && number > *v.Max) {
		return []*InvalidValue{
			{
				Err:   NumberNotInRangeError{v.Min, v.Max},
				Start: 0,
				End:   uint32(len(value)),
			},
		}
	}

	return nil
}
func (v NumberValue) FetchCompletions(line string, cursor uint32) []protocol.CompletionItem {
	return []protocol.CompletionItem{}
}

func (v NumberValue) FetchHoverInfo(line string, cursor uint32) []string {
	return []string{}
}
