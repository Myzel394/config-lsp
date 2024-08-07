package docvalues

import (
	"config-lsp/utils"

	protocol "github.com/tliron/glsp/protocol_3_16"
	"golang.org/x/exp/slices"
)

func GenerateBase10Completions(prefix string) []protocol.CompletionItem {
	kind := protocol.CompletionItemKindValue

	return utils.Map(
		[]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"},
		func(index string) protocol.CompletionItem {
			return protocol.CompletionItem{
				Label: prefix + index,
				Kind:  &kind,
			}
		},
	)
}

func MergeKeyEnumAssignmentMaps(maps ...map[EnumString]Value) map[EnumString]Value {
	existingEnums := make(map[string]interface{})
	result := make(map[EnumString]Value)

	slices.Reverse(maps)

	for _, m := range maps {
		for key, value := range m {
			if _, ok := existingEnums[key.InsertText]; ok {
				continue
			}

			existingEnums[key.InsertText] = nil

			result[key] = value
		}
	}

	return result
}
