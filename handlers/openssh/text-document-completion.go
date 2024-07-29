package openssh

import (
	"config-lsp/common"
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
			cursor := params.Position.Character - uint32(len(optionName + Parser.Options.Separator))

			return getOptionCompletions(optionName, line.Value, cursor), nil
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
		kind := protocol.CompletionItemKindField


		completions[index] = protocol.CompletionItem{
			Label:            label,
			Documentation:    common.GetDocumentation(&option),
			InsertText:       &insertText,
			InsertTextFormat: &format,
			Kind: &kind,
		}
	}

	return completions
}

func getCompletionsFromValue(requiredValue common.Value, line string, cursor uint32) []protocol.CompletionItem {
	switch requiredValue.(type) {
	case common.EnumValue:
		enumValue := requiredValue.(common.EnumValue)
		completions := make([]protocol.CompletionItem, len(requiredValue.(common.EnumValue).Values))

		for index, value := range enumValue.Values {
			textFormat := protocol.InsertTextFormatPlainText
			kind := protocol.CompletionItemKindEnum

			completions[index] = protocol.CompletionItem{
				Label:            value,
				InsertTextFormat: &textFormat,
				Kind: &kind,
			}
		}

		return completions
	case common.CustomValue:
		customValue := requiredValue.(common.CustomValue)
		val := customValue.FetchValue()

		return getCompletionsFromValue(val, line, cursor)
	case common.ArrayValue:
		arrayValue := requiredValue.(common.ArrayValue)
		relativePosition, found := common.FindPreviousCharacter(line, arrayValue.Separator, int(cursor - 1))

		if found {
			line = line[uint32(relativePosition):]
			cursor -= uint32(relativePosition)
		}

		return getCompletionsFromValue(arrayValue.SubValue, line, cursor)
	case common.OrValue:
		orValue := requiredValue.(common.OrValue)

		completions := make([]protocol.CompletionItem, 0)

		for _, subValue := range orValue.Values {
			completions = append(completions, getCompletionsFromValue(subValue, line, cursor)...)
		}

		return completions
	case common.PrefixWithMeaningValue:
		prefixWithMeaningValue := requiredValue.(common.PrefixWithMeaningValue)

		return getCompletionsFromValue(prefixWithMeaningValue.SubValue, line, cursor)
	case common.KeyValueAssignmentValue:
		keyValueAssignmentValue := requiredValue.(common.KeyValueAssignmentValue)

		println("keyLine", line, "cursor", cursor)
		relativePosition, found := common.FindPreviousCharacter(line, keyValueAssignmentValue.Separator, int(cursor - 1))

		println("relativePosition", relativePosition)

		if found {
			line = line[uint32(relativePosition):]
			cursor -= uint32(relativePosition)

			return getCompletionsFromValue(keyValueAssignmentValue.Value, line, cursor)
		} else {
			println("giving key")
			return getCompletionsFromValue(keyValueAssignmentValue.Key, line, cursor)
		}
	}

	return []protocol.CompletionItem{}
}

func getOptionCompletions(optionName string, line string, cursor uint32) []protocol.CompletionItem {
	option := Options[optionName]

	completions := getCompletionsFromValue(option.Value, line, cursor)

	return completions
}
