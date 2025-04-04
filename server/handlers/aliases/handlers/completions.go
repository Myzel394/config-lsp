package handlers

import (
	"config-lsp/common"
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
	cursor common.CursorPosition,
	entry *ast.AliasEntry,
	i *indexes.AliasesIndexes,
) ([]protocol.CompletionItem, error) {
	completions := make([]protocol.CompletionItem, 0)

	if entry.Key == nil {
		return completions, nil
	}

	value := GetValueAtPosition(cursor, entry)

	excludedUsers := getUsersFromEntry(entry)

	if value == nil {
		completions = append(completions, getCommandCompletion())
		completions = append(completions, getIncludeCompletion())
		completions = append(completions, getErrorCompletion())

		completions = append(completions, getUserCompletions(
			i,
			excludedUsers,
		)...)

		return completions, nil
	}

	switch value.(type) {
	case ast.AliasValueUser:
		return getUserCompletions(
			i,
			excludedUsers,
		), nil
	case ast.AliasValueError:
		errorValue := value.(ast.AliasValueError)

		isAtErrorCode := errorValue.Code == nil &&
			errorValue.Location.IsPositionAfterStart(cursor) &&
			(errorValue.Message == nil ||
				errorValue.Message.Location.IsPositionBeforeEnd(cursor))

		if isAtErrorCode {
			kind := protocol.CompletionItemKindValue

			detail_4 := "4XX (TempFail)"
			insertText_4 := "400"

			detail_5 := "5XX (PermFail)"
			insertText_5 := "500"

			return []protocol.CompletionItem{
				{
					Label:      "4XX",
					InsertText: &insertText_4,
					Kind:       &kind,
					Detail:     &detail_4,
				},
				{
					Label:      "5XX",
					InsertText: &insertText_5,
					Kind:       &kind,
					Detail:     &detail_5,
				},
			}, nil
		}
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

func getErrorCompletion() protocol.CompletionItem {
	kind := protocol.CompletionItemKindKeyword
	textFormat := protocol.InsertTextFormatSnippet
	insertText := "error:"

	return protocol.CompletionItem{
		Label:            "error:<message>",
		Documentation:    "A status code and message to return. The code must be 3 digits, starting 4XX (TempFail) or 5XX (PermFail). The message must be present and can be freely chosen.",
		Kind:             &kind,
		InsertTextFormat: &textFormat,
		InsertText:       &insertText,
	}
}

func getUserCompletions(
	i *indexes.AliasesIndexes,
	excluded map[string]struct{},
) []protocol.CompletionItem {
	users := fetchers.GetAvailableUserValues(i)

	kind := protocol.CompletionItemKindValue

	completions := make([]protocol.CompletionItem, 0)

	for name, user := range users {
		if _, found := excluded[name]; found {
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

func getUsersFromEntry(
	entry *ast.AliasEntry,
) map[string]struct{} {
	users := map[string]struct{}{
		indexes.NormalizeKey(entry.Key.Value): {},
	}

	if entry.Values != nil {
		for _, value := range entry.Values.Values {
			switch (value).(type) {
			case ast.AliasValueUser:
				userValue := value.(ast.AliasValueUser)

				users[indexes.NormalizeKey(userValue.Value)] = struct{}{}
			}
		}
	}

	return users
}
