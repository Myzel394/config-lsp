package openssh

import (
	"config-lsp/common"
	docvalues "config-lsp/doc-values"
	"errors"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
	"golang.org/x/exp/maps"

	_ "github.com/tliron/commonlog/simple"
)

func TextDocumentCompletion(context *glsp.Context, params *protocol.CompletionParams) (interface{}, error) {
	optionName, line, err := Parser.FindByLineNumber(uint32(params.Position.Line))

	if err == nil {
		if params.Position.Character < uint32(len(optionName)) {
			return getRootCompletions(), nil
		} else {
			cursor := params.Position.Character - uint32(line.GetCharacterPositions(optionName)[1])

			return getOptionCompletions(optionName, line.Value, cursor), nil
		}
	} else if errors.Is(err, docvalues.LineNotFoundError{}) {
		return getRootCompletions(), nil
	}

	return nil, err
}

func getRootCompletions() []protocol.CompletionItem {
	completions := make([]protocol.CompletionItem, len(Options))

	optionsKey := maps.Keys(Options)
	for index := 0; index < len(maps.Keys(Options)); index++ {
		label := optionsKey[index]
		option := Options[label]
		insertText := label + " " + "${1:value}"

		format := protocol.InsertTextFormatSnippet
		kind := protocol.CompletionItemKindField

		completions[index] = protocol.CompletionItem{
			Label:            label,
			Documentation:    common.GetDocumentation(&option),
			InsertText:       &insertText,
			InsertTextFormat: &format,
			Kind:             &kind,
		}
	}

	return completions
}

func getOptionCompletions(optionName string, line string, cursor uint32) []protocol.CompletionItem {
	option := Options[optionName]

	completions := option.Value.FetchCompletions(line, cursor)

	return completions
}
