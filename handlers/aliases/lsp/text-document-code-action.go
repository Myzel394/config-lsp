package lsp

import (
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentCodeAction(context *glsp.Context, params *protocol.CodeActionParams) ([]protocol.CodeAction, error) {
	// document := hosts.DocumentParserMap[params.TextDocument.URI]
	//
	// actions := make([]protocol.CodeAction, 0, 1)

	return nil, nil
}
