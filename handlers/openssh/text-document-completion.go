package handlers

import (
	"config-lsp/common"
	"errors"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
	"golang.org/x/exp/maps"

	_ "github.com/tliron/commonlog/simple"
)

func TextDocumentCompletion(context *glsp.Context, params *protocol.CompletionParams) (interface{}, error) {
	option, line, err := Parser.FindByLineNumber(int(params.Position.Line))

	if err == nil {
		if line.IsCursorAtRootOption(int(params.Position.Character)) {
			return getRootCompletions(), nil
		} else {
			return getOptionCompletions(option), nil
		}
	} else if errors.Is(err, common.LineNotFoundError{}) {
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

		completions[index] = protocol.CompletionItem{
			Label: label,
			Documentation: common.GetDocumentation(&option),
			InsertText: &insertText,
			InsertTextFormat: &format,
		}
	}

	return completions
}

func getOptionCompletions(optionName string) []protocol.CompletionItem {
	option := Options[optionName]

	switch option.Value.(type) {
	case common.EnumValue:
		enumOption := option.Value.(common.EnumValue)
		completions := make([]protocol.CompletionItem, len(option.Value.(common.EnumValue).Values))

		for index, value := range enumOption.Values {
			textFormat := protocol.InsertTextFormatPlainText

			completions[index] = protocol.CompletionItem{
				Label: value,
				InsertTextFormat: &textFormat,
			}
		}

		return completions
	}

	return []protocol.CompletionItem{}
}

