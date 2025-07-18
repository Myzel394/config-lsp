package lsp

import (
	bitcoinconf "config-lsp/handlers/bitcoin_conf"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentDidClose(context *glsp.Context, params *protocol.DidCloseTextDocumentParams) error {
	delete(bitcoinconf.DocumentParserMap, params.TextDocument.URI)

	return nil
}
