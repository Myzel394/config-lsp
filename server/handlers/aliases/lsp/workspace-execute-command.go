package lsp

import (
	"config-lsp/handlers/aliases"
	"config-lsp/handlers/aliases/handlers"
	"strings"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func WorkspaceExecuteCommand(context *glsp.Context, params *protocol.ExecuteCommandParams) (*protocol.ApplyWorkspaceEditParams, error) {
	_, command, _ := strings.Cut(params.Command, ".")

	switch command {
	case string(handlers.CodeActionSendTestMail):
		args := handlers.CodeActionSendTestMailArgsFromArguments(params.Arguments[0].(map[string]any))

		d := aliases.DocumentParserMap[args.URI]

		return args.RunCommand(d)
	}

	return nil, nil
}
