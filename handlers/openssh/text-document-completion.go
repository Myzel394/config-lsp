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
	option, line, err := Parser.FindByLineNumber(uint32(params.Position.Line))

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
			Label:            label,
			Documentation:    common.GetDocumentation(&option),
			InsertText:       &insertText,
			InsertTextFormat: &format,
		}
	}

	return completions
}

func getCompletionsFromValue(value common.Value) []protocol.CompletionItem {
	switch value.(type) {
	case common.EnumValue:
		enumValue := value.(common.EnumValue)
		completions := make([]protocol.CompletionItem, len(value.(common.EnumValue).Values))

		for index, value := range enumValue.Values {
			textFormat := protocol.InsertTextFormatPlainText

			completions[index] = protocol.CompletionItem{
				Label:            value,
				InsertTextFormat: &textFormat,
			}
		}

		return completions
	case common.CustomValue:
		customValue := value.(common.CustomValue)
		val := customValue.FetchValue()

		return getCompletionsFromValue(val)
	case common.ArrayValue:
		arrayValue := value.(common.ArrayValue)

		return getCompletionsFromValue(arrayValue.SubValue)
	case common.OrValue:
		orValue := value.(common.OrValue)

		completions := make([]protocol.CompletionItem, 0)

		for _, subValue := range orValue.Values {
			completions = append(completions, getCompletionsFromValue(subValue)...)
		}

		return completions
	case common.PrefixWithMeaningValue:
		prefixWithMeaningValue := value.(common.PrefixWithMeaningValue)

		return getCompletionsFromValue(prefixWithMeaningValue.SubValue)
	}

	return []protocol.CompletionItem{}
}

func getOptionCompletions(optionName string) []protocol.CompletionItem {
	option := Options[optionName]

	completions := getCompletionsFromValue(option.Value)

	return completions
}
