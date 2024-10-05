package docvalues

import (
	protocol "github.com/tliron/glsp/protocol_3_16"
)

type EmptyStringError struct{}

func (e EmptyStringError) Error() string {
	return "This setting may not be empty"
}

type StringTooLongError struct{}

func (e StringTooLongError) Error() string {
	return "This setting is too long"
}

type StringValue struct {
	MaxLength *uint32
}

func (v StringValue) GetTypeDescription() []string {
	return []string{"String"}
}

func (v StringValue) DeprecatedCheckIsValid(value string) []*InvalidValue {
	if value == "" {
		return []*InvalidValue{{
			Err:   EmptyStringError{},
			Start: 0,
			End:   uint32(len(value)),
		},
		}
	}

	if v.MaxLength != nil && uint32(len(value)) > *v.MaxLength {
		return []*InvalidValue{{
			Err:   StringTooLongError{},
			Start: 0,
			End:   uint32(len(value)),
		}}
	}

	return nil
}

func (v StringValue) DeprecatedFetchCompletions(line string, cursor uint32) []protocol.CompletionItem {
	return []protocol.CompletionItem{}
}

func (v StringValue) DeprecatedFetchHoverInfo(line string, cursor uint32) []string {
	return []string{}
}
