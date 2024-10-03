package lsp

import (
	sshconfig "config-lsp/handlers/ssh_config"
	"config-lsp/handlers/ssh_config/handlers"
	"strings"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func WorkspaceExecuteCommand(context *glsp.Context, params *protocol.ExecuteCommandParams) (*protocol.ApplyWorkspaceEditParams, error) {
	_, command, _ := strings.Cut(params.Command, ".")

	switch command {
	case string(handlers.CodeActionAddToUnknown):
		args := handlers.CodeActionAddToUnknownArgsFromArguments(params.Arguments[0].(map[string]any))

		d := sshconfig.DocumentParserMap[args.URI]

		return args.RunCommand(d)
	}

	return nil, nil
}
