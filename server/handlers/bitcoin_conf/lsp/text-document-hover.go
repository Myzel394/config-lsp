package lsp

import (
	"config-lsp/common"
	"config-lsp/handlers/wireguard"
	"config-lsp/handlers/wireguard/handlers"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentHover(
	context *glsp.Context,
	params *protocol.HoverParams,
) (*protocol.Hover, error) {
	d := wireguard.DocumentParserMap[params.TextDocument.URI]
	line := params.Position.Line

	section := d.Config.FindSectionByLine(line)
	property := d.Config.FindPropertyByLine(line)

	index := common.LSPCharacterAsIndexPosition(params.Position.Character)

	if property != nil && section != nil {
		return handlers.GetPropertyHoverInfo(
			d,
			*section,
			*property,
			index,
		)
	}

	if section != nil && section.Start.Line == line {
		return handlers.GetSectionHoverInfo(
			d,
			*section,
		)
	}

	return nil, nil
}
