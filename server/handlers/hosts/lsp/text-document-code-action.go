package lsp

import (
	"config-lsp/handlers/hosts"
	"config-lsp/handlers/hosts/handlers"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentCodeAction(context *glsp.Context, params *protocol.CodeActionParams) ([]protocol.CodeAction, error) {
	document := hosts.DocumentParserMap[params.TextDocument.URI]

	actions := make([]protocol.CodeAction, 0, 1)

	actions = append(actions, handlers.GetInlineAliasesCodeAction(*document, params)...)

	if len(actions) > 0 {
		return actions, nil
	}

	return nil, nil
}
