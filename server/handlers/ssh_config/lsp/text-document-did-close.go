package lsp

import (
	"config-lsp/handlers/ssh_config"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentDidClose(context *glsp.Context, params *protocol.DidCloseTextDocumentParams) error {
	delete(sshconfig.DocumentParserMap, params.TextDocument.URI)

	return nil
}
