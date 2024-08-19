package wireguard

import (
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentCodeAction(context *glsp.Context, params *protocol.CodeActionParams) ([]protocol.CodeAction, error) {
	parser := documentParserMap[params.TextDocument.URI]

	actions := make([]protocol.CodeAction, 0, 2)

	actions = append(actions, getKeyGenerationCodeActions(params, parser)...)
	actions = append(actions, getKeepaliveCodeActions(params, parser)...)

	if len(actions) > 0 {
		return actions, nil
	}

	return nil, nil
}
