package lsp

import (
	"config-lsp/common"
	"config-lsp/handlers/aliases"
	"config-lsp/handlers/aliases/ast"
	"config-lsp/handlers/aliases/handlers"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentCompletion(context *glsp.Context, params *protocol.CompletionParams) (any, error) {
	d := aliases.DocumentParserMap[params.TextDocument.URI]
	cursor := common.LSPCharacterAsCursorPosition(params.Position.Character)
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

	if entry.Key == nil || entry.Key.Location.ContainsPosition(cursor) {
		return handlers.GetAliasesCompletions(d.Indexes), nil
	}

	if entry.Separator == nil && entry.Key.Location.IsPositionBeforeEnd(cursor) {
		return nil, nil
	}

	return handlers.GetCompletionsForEntry(
		cursor,
		entry,
		d.Indexes,
	)
}
