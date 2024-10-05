package lsp

import (
	shared "config-lsp/handlers/fstab/shared"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentDidClose(context *glsp.Context, params *protocol.DidCloseTextDocumentParams) error {
	delete(shared.DocumentParserMap, params.TextDocument.URI)

	return nil
}
