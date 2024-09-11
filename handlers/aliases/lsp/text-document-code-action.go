package lsp

import (
	"config-lsp/handlers/aliases"
	"config-lsp/handlers/aliases/handlers"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentCodeAction(context *glsp.Context, params *protocol.CodeActionParams) ([]protocol.CodeAction, error) {
	d := aliases.DocumentParserMap[params.TextDocument.URI]
	actions := handlers.FetchCodeActions(d, params)

	return actions, nil
}
