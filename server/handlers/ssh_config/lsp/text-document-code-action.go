package lsp

import (
	sshconfig "config-lsp/handlers/ssh_config"
	"config-lsp/handlers/ssh_config/handlers"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentCodeAction(context *glsp.Context, params *protocol.CodeActionParams) ([]protocol.CodeAction, error) {
	d := sshconfig.DocumentParserMap[params.TextDocument.URI]
	actions := handlers.FetchCodeActions(d, params)

	return actions, nil
}
