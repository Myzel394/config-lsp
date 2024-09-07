package lsp

import (
	"config-lsp/handlers/aliases"
	"config-lsp/handlers/aliases/ast"
	"config-lsp/handlers/aliases/handlers"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentCompletion(context *glsp.Context, params *protocol.CompletionParams) (any, error) {
	d := aliases.DocumentParserMap[params.TextDocument.URI]
	cursor := params.Position.Character
	line := params.Position.Line

	if _, found := d.Parser.CommentLines[line]; found {
		return nil, nil
	}

	rawEntry, found := d.Parser.Aliases.Get(line)

	if !found {
		// For the key there are no completions available
		return handlers.GetAliasesCompletions(d.Indexes), nil
	}

	entry := rawEntry.(*ast.AliasEntry)

	if entry.Key == nil {
		return handlers.GetAliasesCompletions(d.Indexes), nil
	}

	if cursor >= entry.Key.Location.Start.Character && cursor <= entry.Key.Location.End.Character {
		return handlers.GetAliasesCompletions(d.Indexes), nil
	}

	if entry.Separator == nil && cursor > entry.Key.Location.End.Character {
		return nil, nil
	}

	if cursor > entry.Separator.End.Character {
		return handlers.GetCompletionsForEntry(
			cursor,
			entry,
			d.Indexes,
		)
	}

	return nil, nil
}