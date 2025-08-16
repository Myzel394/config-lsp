package lsp

import (
	"config-lsp/handlers/wireguard"
	"config-lsp/handlers/wireguard/handlers"
	"strings"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func WorkspaceExecuteCommand(context *glsp.Context, params *protocol.ExecuteCommandParams) (*protocol.ApplyWorkspaceEditParams, error) {
	_, command, _ := strings.Cut(params.Command, ".")

	switch command {
	case string(handlers.CodeActionGeneratePrivateKey):
		args := handlers.CodeActionGeneratePrivateKeyArgsFromArguments(params.Arguments[0].(map[string]any))

		d := wireguard.DocumentParserMap[args.URI]

		return args.RunCommand(d)
	case string(handlers.CodeActionGeneratePresharedKey):
		args := handlers.CodeActionGeneratePresharedKeyArgsFromArguments(params.Arguments[0].(map[string]any))

		d := wireguard.DocumentParserMap[args.URI]

		return args.RunCommand(d)
	case string(handlers.CodeActionCreatePeer):
		args := handlers.CodeActionCreatePeerArgsFromArguments(params.Arguments[0].(map[string]any))

		d := wireguard.DocumentParserMap[args.URI]

		return args.RunCommand(d)
	case string(handlers.CodeActionGenerateDownRule):
		args := handlers.CodeActionGenerateDownRuleArgsFromArguments(params.Arguments[0].(map[string]any))

		d := wireguard.DocumentParserMap[args.URI]
		return args.RunCommand(d)
	}

	return nil, nil
}
