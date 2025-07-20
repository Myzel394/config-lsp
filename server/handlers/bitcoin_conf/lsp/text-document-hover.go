package lsp

import (
	"config-lsp/common"
	bitcoinconf "config-lsp/handlers/bitcoin_conf"
	"config-lsp/handlers/bitcoin_conf/handlers"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentHover(
	context *glsp.Context,
	params *protocol.HoverParams,
) (*protocol.Hover, error) {
	d := bitcoinconf.DocumentParserMap[params.TextDocument.URI]
	line := params.Position.Line

	property := d.Config.FindPropertyByLine(line)

	index := common.LSPCharacterAsIndexPosition(params.Position.Character)

	if property != nil {
		return handlers.GetPropertyHoverInfo(
			d,
			property,
			index,
		)
	}

	return nil, nil
}
