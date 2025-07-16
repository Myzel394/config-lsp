package docvalues

import (
	"config-lsp/common"
	"fmt"
	"strconv"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

type FloatNumberValue struct {
	Min *float64
	Max *float64
}

type FloatNumberNotInRangeError struct {
	Min *float64
	Max *float64
}

func (e FloatNumberNotInRangeError) Error() string {
	if e.Min == nil {
		return fmt.Sprintf("This must be at most or less than %f", *e.Max)
	}

	if e.Max == nil {
		return fmt.Sprintf("This must be equal or more than %f", *e.Min)
	}

	return fmt.Sprintf("This must be between %f and %f (inclusive)", *e.Min, *e.Max)
}

func (v FloatNumberValue) GetTypeDescription() []string {
	if v.Min != nil {
		if v.Max != nil {
			return []string{fmt.Sprintf("A number between %f and %f (inclusive)", *v.Min, *v.Max)}
		} else {
			return []string{fmt.Sprintf("A number of at least %f", *v.Min)}
		}
	} else if v.Max != nil {
		return []string{fmt.Sprintf("A number of at most %f", *v.Max)}
	}

	return []string{"A number"}
}

func (v FloatNumberValue) DeprecatedCheckIsValid(value string) []*InvalidValue {
	number, err := strconv.ParseFloat(value, 64)

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
				Err:   FloatNumberNotInRangeError{v.Min, v.Max},
				Start: 0,
				End:   uint32(len(value)),
			},
		}
	}

	return nil
}

func (v FloatNumberValue) FetchCompletions(value string, cursor common.CursorPosition) []protocol.CompletionItem {
	return []protocol.CompletionItem{}
}

func (v FloatNumberValue) DeprecatedFetchHoverInfo(line string, cursor uint32) []string {
	return []string{}
}
