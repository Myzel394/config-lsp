package handlers

import (
	"config-lsp/common"
	"strings"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
	"golang.org/x/exp/maps"

	_ "github.com/tliron/commonlog/simple"
)

func TextDocumentCompletion(context *glsp.Context, params *protocol.CompletionParams) (interface{}, error) {
	line, err := common.GetLine(params.TextDocument.URI, int(params.Position.Line))

	if err != nil {
		return [...]protocol.CompletionItem{}, err
	}

	rootOption := getCurrentOption(line, int(params.Position.Character))

	if (rootOption == "") {
		return getRootCompletions(), nil
	} else {
		return getOptionCompletions(rootOption), nil
	}

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

func getCurrentOption(line string, position int) string {
	words := strings.Split(line, " ")

	if len(words) == 0 {
		return ""
	}

	if (position <= len(words[0])) {
		return ""
	}

	return words[0]
}

