package lsp

import (
	bitcoinconf "config-lsp/handlers/bitcoin_conf"
	"config-lsp/handlers/bitcoin_conf/handlers"
	"strings"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func WorkspaceExecuteCommand(context *glsp.Context, params *protocol.ExecuteCommandParams) (*protocol.ApplyWorkspaceEditParams, error) {
	_, command, _ := strings.Cut(params.Command, ".")

	switch command {
	case string(handlers.CodeActionGenerateRPCAuth):
		args := handlers.CodeActionGenerateRPCAuthArgsFromArguments(params.Arguments[0].(map[string]any))

		d := bitcoinconf.DocumentParserMap[args.URI]

		return args.RunCommand(d)
	}

	return nil, nil
}
