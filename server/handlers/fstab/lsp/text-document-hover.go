package lsp

import (
	"config-lsp/handlers/fstab/ast"
	"config-lsp/handlers/fstab/handlers"
	"config-lsp/handlers/fstab/shared"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentHover(context *glsp.Context, params *protocol.HoverParams) (*protocol.Hover, error) {
	line := params.Position.Line
	cursor := params.Position.Character

	p := shared.DocumentParserMap[params.TextDocument.URI]

	rawEntry, found := p.Entries.Get(params.Position.Line)

	// Empty line
	if !found {
		return nil, nil
	}

	entry := rawEntry.(*ast.FstabEntry)

	return handlers.GetHoverInfo(
		line,
		cursor,
		entry,
	)
}
