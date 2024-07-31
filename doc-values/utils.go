package docvalues

import (
	"config-lsp/utils"
	protocol "github.com/tliron/glsp/protocol_3_16"
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
