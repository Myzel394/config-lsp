package lsp

import (
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentHover(
	context *glsp.Context,
	params *protocol.HoverParams,
) (*protocol.Hover, error) {
	// line := params.Position.Line
	// cursor := params.Position.Character
	//
	// d := sshdconfig.DocumentParserMap[params.TextDocument.URI]
	//
	// entry, matchBlock := d.Config.FindOption(line)

	return nil, nil
}
