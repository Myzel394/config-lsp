package lsp

import (
	"config-lsp/handlers/fstab/handlers"
	"config-lsp/handlers/fstab/parser"
	"config-lsp/handlers/fstab/shared"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentHover(context *glsp.Context, params *protocol.HoverParams) (*protocol.Hover, error) {
	cursor := params.Position.Character

	p := shared.DocumentParserMap[params.TextDocument.URI]

	entry, found := p.GetEntry(params.Position.Line)

	// Empty line
	if !found {
		return nil, nil
	}

	// Comment line
	if entry.Type == parser.FstabEntryTypeComment {
		return nil, nil
	}

	return handlers.GetHoverInfo(entry, cursor)
}
