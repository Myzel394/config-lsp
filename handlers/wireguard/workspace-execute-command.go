package wireguard

import (
	"strings"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func WorkspaceExecuteCommand(context *glsp.Context, params *protocol.ExecuteCommandParams) (*protocol.ApplyWorkspaceEditParams, error) {
	_, command, _ := strings.Cut(params.Command, ".")

	switch command {
	case string(codeActionGeneratePrivateKey):
		args := codeActionGeneratePrivateKeyArgsFromArguments(params.Arguments[0].(map[string]any))

		parser := documentParserMap[args.URI]

		return parser.runGeneratePrivateKey(args)
	}

	return nil, nil
}
