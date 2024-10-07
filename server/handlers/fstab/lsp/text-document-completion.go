package lsp

import (
	"config-lsp/common"
	fstabdocumentation "config-lsp/handlers/fstab/documentation"
	"config-lsp/handlers/fstab/handlers"
	"config-lsp/handlers/fstab/parser"
	"config-lsp/handlers/fstab/shared"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentCompletion(context *glsp.Context, params *protocol.CompletionParams) (any, error) {
	p := shared.DocumentParserMap[params.TextDocument.URI]

	entry, found := p.GetEntry(params.Position.Line)

	if !found {
		// Empty line, return spec completions
		return fstabdocumentation.SpecField.DeprecatedFetchCompletions(
			"",
			params.Position.Character,
		), nil
	}

	if entry.Type == parser.FstabEntryTypeComment {
		return nil, nil
	}

	cursor := common.CursorToCharacterIndex(params.Position.Character)
	line := entry.Line

	return handlers.GetCompletion(line, cursor)
}
