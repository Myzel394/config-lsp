package handlers

import (
	"config-lsp/handlers/aliases/analyzer"
	"config-lsp/handlers/aliases/ast"
	"config-lsp/handlers/aliases/indexes"
	"fmt"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func GetAliasesCompletions(
	i *indexes.AliasesIndexes,
) []protocol.CompletionItem {
	completions := make([]protocol.CompletionItem, 0)
	aliases := analyzer.RequiredAliases

	kind := protocol.CompletionItemKindValue

	for _, alias := range aliases {
		if i != nil {
			if _, found := i.Keys[alias]; found {
				continue
			}
		}

		text := fmt.Sprintf("%s: ", alias)
		completions = append(completions, protocol.CompletionItem{
			Label: alias,
			Kind: &kind,
			InsertText: &text,
			Documentation: "This alias is required by the aliases file",
		})
	}

	return completions
}

func GetCompletionsForEntry(
	cursor uint32,
	entry *ast.AliasEntry,
	i *indexes.AliasesIndexes,
) ([]protocol.CompletionItem, error) {
	completions := make([]protocol.CompletionItem, 0)

	if entry.Key == nil {
		return completions, nil
	}

	value := getValueAtCursor(cursor, entry)

	println(fmt.Sprintf("Value: %v", value))

	return completions, nil
}

