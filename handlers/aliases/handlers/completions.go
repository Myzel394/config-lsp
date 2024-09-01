package handlers

import (
	"config-lsp/handlers/aliases/analyzer"
	"config-lsp/handlers/aliases/ast"
	"config-lsp/handlers/aliases/fetchers"
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
			Label:         alias,
			Kind:          &kind,
			InsertText:    &text,
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
	relativeCursor := cursor - entry.Key.Location.Start.Character

	if value == nil {
		completions = append(completions, getCommandCompletion())
		completions = append(completions, getIncludeCompletion())

		completions = append(completions, getUserCompletions(
			i,
			indexes.NormalizeKey(entry.Key.Value),
			"",
			0,
		)...)

		println("la completions etaient", completions)
		return completions, nil
	}

	switch (*value).(type) {
	case ast.AliasValueUser:
		userValue := (*value).(ast.AliasValueUser)

		return getUserCompletions(
			i,
			indexes.NormalizeKey(entry.Key.Value),
			userValue.Value,
			relativeCursor,
		), nil
	}

	return completions, nil
}

func getCommandCompletion() protocol.CompletionItem {
	kind := protocol.CompletionItemKindKeyword
	textFormat := protocol.InsertTextFormatSnippet
	insertText := "|"

	return protocol.CompletionItem{
		Label:            "|<command>",
		Documentation:    "Pipe the message to command on its standard input. The command is run under the privileges of the daemon's unprivileged account.",
		Kind:             &kind,
		InsertTextFormat: &textFormat,
		InsertText:       &insertText,
	}
}

func getIncludeCompletion() protocol.CompletionItem {
	kind := protocol.CompletionItemKindKeyword
	textFormat := protocol.InsertTextFormatSnippet
	insertText := ":include:"

	return protocol.CompletionItem{
		Label:            ":include:<path>",
		Documentation:    " Include any definitions in file as alias entries. The format of the file is identical to this one.",
		Kind:             &kind,
		InsertTextFormat: &textFormat,
		InsertText:       &insertText,
	}
}

func getUserCompletions(
	i *indexes.AliasesIndexes,
	excludeKey string,
	line string,
	cursor uint32,
) []protocol.CompletionItem {
	users := fetchers.GetAvailableUserValues(i)

	kind := protocol.CompletionItemKindValue

	completions := make([]protocol.CompletionItem, 0)

	for name, user := range users {
		if name == excludeKey {
			continue
		}

		completions = append(completions, protocol.CompletionItem{
			Label:         name,
			Kind:          &kind,
			Documentation: user.Documentation(),
		})
	}

	return completions
}
