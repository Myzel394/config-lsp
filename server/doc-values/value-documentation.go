package docvalues

import (
	"config-lsp/common"
	"strings"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

type DocumentationValue struct {
	Documentation string
	Value         DeprecatedValue
}

func (v DocumentationValue) GetTypeDescription() []string {
	return v.Value.GetTypeDescription()
}

func (v DocumentationValue) DeprecatedCheckIsValid(value string) []*InvalidValue {
	return v.Value.DeprecatedCheckIsValid(value)
}

func (v DocumentationValue) DeprecatedFetchCompletions(line string, cursor uint32) []protocol.CompletionItem {
	return v.Value.DeprecatedFetchCompletions(line, cursor)
}

func (v DocumentationValue) FetchCompletions(value string, cursor common.CursorPosition) []protocol.CompletionItem {
	return v.DeprecatedFetchCompletions(
		value,
		common.DeprecatedImprovedCursorToIndex(
			cursor,
			value,
			0,
		),
	)
}

func (v DocumentationValue) DeprecatedFetchHoverInfo(line string, cursor uint32) []string {
	return strings.Split(v.Documentation, "\n")
}
