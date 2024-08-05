package docvalues

import (
	protocol "github.com/tliron/glsp/protocol_3_16"
)

type CustomValueContext interface {
	GetIsContext() bool
}

type EmptyValueContext struct{}

func (EmptyValueContext) GetIsContext() bool {
	return true
}

var EmptyValueContextInstance = EmptyValueContext{}

type CustomValue struct {
	FetchValue func(context CustomValueContext) Value
}

func (v CustomValue) GetTypeDescription() []string {
	return []string{"Custom"}
}

func (v CustomValue) CheckIsValid(value string) []*InvalidValue {
	return v.FetchValue(EmptyValueContextInstance).CheckIsValid(value)
}

func (v CustomValue) FetchCompletions(line string, cursor uint32) []protocol.CompletionItem {
	return v.FetchValue(EmptyValueContextInstance).FetchCompletions(line, cursor)
}
