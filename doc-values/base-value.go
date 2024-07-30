package docvalues

import (
	protocol "github.com/tliron/glsp/protocol_3_16"
)

type Value interface {
	GetTypeDescription() []string
	CheckIsValid(value string) error
	FetchCompletions(line string, cursor uint32) []protocol.CompletionItem
}
