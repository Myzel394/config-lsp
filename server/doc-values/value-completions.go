package docvalues

import (
	"config-lsp/common"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

// Manually add extra completions
type CompletionsValue struct {
	Completions []protocol.CompletionItem
	SubValue    DeprecatedValue
}

func (v CompletionsValue) GetTypeDescription() []string {
	return v.SubValue.GetTypeDescription()
}

func (v CompletionsValue) DeprecatedCheckIsValid(value string) []*InvalidValue {
	return v.SubValue.DeprecatedCheckIsValid(value)
}

func (v CompletionsValue) FetchCompletions(value string, cursor common.CursorPosition) []protocol.CompletionItem {
	completions := v.SubValue.FetchCompletions(value, cursor)

	completions = append(completions, v.Completions...)

	return completions
}

func (v CompletionsValue) DeprecatedFetchHoverInfo(line string, cursor uint32) []string {
	return v.SubValue.DeprecatedFetchHoverInfo(line, cursor)
}
