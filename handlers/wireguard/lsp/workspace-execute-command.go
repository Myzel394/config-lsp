package lsp

import (
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

		p := documentParserMap[args.URI]

		return args.RunCommand(p)
	case string(handlers.CodeActionGeneratePresharedKey):
		args := handlers.CodeActionGeneratePresharedKeyArgsFromArguments(params.Arguments[0].(map[string]any))

		parser := documentParserMap[args.URI]

		return args.RunCommand(parser)
	case string(handlers.CodeActionAddKeepalive):
		args := handlers.CodeActionAddKeepaliveArgsFromArguments(params.Arguments[0].(map[string]any))

		p := documentParserMap[args.URI]

		return args.RunCommand(p)
	}

	return nil, nil
}
