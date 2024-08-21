package docvalues

import (
	"strings"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

type DocumentationValue struct {
	Documentation string
	Value         Value
}

func (v DocumentationValue) GetTypeDescription() []string {
	return v.Value.GetTypeDescription()
}

func (v DocumentationValue) CheckIsValid(value string) []*InvalidValue {
	return v.Value.CheckIsValid(value)
}

func (v DocumentationValue) FetchCompletions(line string, cursor uint32) []protocol.CompletionItem {
	return v.Value.FetchCompletions(line, cursor)
}

func (v DocumentationValue) FetchHoverInfo(line string, cursor uint32) []string {
	return strings.Split(v.Documentation, "\n")
}
