package lsp

import (
	"config-lsp/handlers/wireguard/handlers"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentCodeAction(context *glsp.Context, params *protocol.CodeActionParams) ([]protocol.CodeAction, error) {
	p := documentParserMap[params.TextDocument.URI]

	actions := make([]protocol.CodeAction, 0, 2)

	actions = append(actions, handlers.GetKeyGenerationCodeActions(p, params)...)
	actions = append(actions, handlers.GetKeepaliveCodeActions(p, params)...)

	if len(actions) > 0 {
		return actions, nil
	}

	return nil, nil
}
