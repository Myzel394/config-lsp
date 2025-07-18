package lsp

import (
	bitcoinconf "config-lsp/handlers/bitcoin_conf"
	"config-lsp/handlers/bitcoin_conf/handlers"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentCodeAction(context *glsp.Context, params *protocol.CodeActionParams) ([]protocol.CodeAction, error) {
	d := bitcoinconf.DocumentParserMap[params.TextDocument.URI]

	actions := make([]protocol.CodeAction, 0, 2)

	actions = append(actions, handlers.GetGenerateRPCAuthCodeActions(d, params)...)
	actions = append(actions, handlers.GetPropertyKeywordTypoFixes(d, params)...)

	if len(actions) > 0 {
		return actions, nil
	}

	return nil, nil
}
