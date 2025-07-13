package utils

import protocol "github.com/tliron/glsp/protocol_3_16"

// If you have a list of completions but they don't start with the same word
// (word defined as separated by spaces), then the LSP client will not
// show the completions. This function adds the `substr` to the beginning
// of each completion item.
func AddSubstrToCompletionItems(
	completions []protocol.CompletionItem,
	substr string,
) []protocol.CompletionItem {
	return Map(
		completions,
		func(item protocol.CompletionItem) protocol.CompletionItem {
			newItem := item
			newItem.Label = substr + newItem.Label

			if newItem.InsertText != nil {
				*newItem.InsertText = substr + *newItem.InsertText
			}

			return newItem
		},
	)
}
