package lsp

import (
	"config-lsp/common"
	"config-lsp/handlers/fstab/ast"
	fstabdocumentation "config-lsp/handlers/fstab/fields"
	"config-lsp/handlers/fstab/handlers"
	"config-lsp/handlers/fstab/shared"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentCompletion(context *glsp.Context, params *protocol.CompletionParams) (any, error) {
	d := shared.DocumentParserMap[params.TextDocument.URI]
	cursor := common.LSPCharacterAsCursorPosition(params.Position.Character)

	rawEntry, found := d.Config.Entries.Get(params.Position.Line)

	if !found {
		// Empty line, return spec completions
		return fstabdocumentation.SpecField.DeprecatedFetchCompletions(
			"",
			params.Position.Character,
		), nil
	}

	entry := rawEntry.(*ast.FstabEntry)

	return handlers.GetCompletion(entry, cursor)
}
