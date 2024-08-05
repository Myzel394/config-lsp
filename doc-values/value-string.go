package docvalues

import (
	protocol "github.com/tliron/glsp/protocol_3_16"
)

type EmptyStringError struct{}

func (e EmptyStringError) Error() string {
	return "This setting may not be empty"
}

type StringValue struct{}

func (v StringValue) GetTypeDescription() []string {
	return []string{"String"}
}

func (v StringValue) CheckIsValid(value string) []*InvalidValue {
	if value == "" {
		return []*InvalidValue{{
			Err:   EmptyStringError{},
			Start: 0,
			End:   uint32(len(value)),
		},
		}
	}

	return nil
}

func (v StringValue) FetchCompletions(line string, cursor uint32) []protocol.CompletionItem {
	return []protocol.CompletionItem{}
}
