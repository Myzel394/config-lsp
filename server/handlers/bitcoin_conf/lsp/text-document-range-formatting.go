package lsp

import (
	bitcoinconf "config-lsp/handlers/bitcoin_conf"
	"config-lsp/handlers/bitcoin_conf/handlers"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentRangeFormatting(
	context *glsp.Context,
	params *protocol.DocumentRangeFormattingParams,
) ([]protocol.TextEdit, error) {
	d := bitcoinconf.DocumentParserMap[params.TextDocument.URI]

	return handlers.FormatDocument(
		d,
		params.Range,
		params.Options,
	)
}
