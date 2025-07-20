package lsp

import (
	"config-lsp/handlers/wireguard"
	"config-lsp/handlers/wireguard/handlers"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentCodeAction(context *glsp.Context, params *protocol.CodeActionParams) ([]protocol.CodeAction, error) {
	d := wireguard.DocumentParserMap[params.TextDocument.URI]

	actions := make([]protocol.CodeAction, 0, 2)

	actions = append(actions, handlers.GetKeyGenerationCodeActions(d, params)...)
	actions = append(actions, handlers.GetKeepaliveCodeActions(d, params)...)
	actions = append(actions, handlers.GetAddPeerLikeThisCodeActions(d, params)...)
	actions = append(actions, handlers.GetPropertyKeywordTypoFixes(d, params)...)
	actions = append(actions, handlers.GetGeneratePostDownCodeActions(d, params)...)

	if len(actions) > 0 {
		return actions, nil
	}

	return nil, nil
}
