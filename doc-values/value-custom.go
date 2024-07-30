package docvalues

import (
	protocol "github.com/tliron/glsp/protocol_3_16"
)

type CustomValue struct {
	FetchValue func() Value
}

func (v CustomValue) GetTypeDescription() []string {
	return []string{"Custom"}
}

func (v CustomValue) CheckIsValid(value string) error {
	return v.FetchValue().CheckIsValid(value)
}

func (v CustomValue) FetchCompletions(line string, cursor uint32) []protocol.CompletionItem {
	return v.FetchValue().FetchCompletions(line, cursor)
}
