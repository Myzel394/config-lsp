package lsp

import (
	"config-lsp/handlers/sshd_config"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentDidClose(context *glsp.Context, params *protocol.DidCloseTextDocumentParams) error {
	delete(sshdconfig.DocumentParserMap, params.TextDocument.URI)

	return nil
}
