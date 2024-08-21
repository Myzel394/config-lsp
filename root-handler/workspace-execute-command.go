package roothandler

import (
	wireguard "config-lsp/handlers/wireguard/lsp"
	"strings"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func WorkspaceExecuteCommand(context *glsp.Context, params *protocol.ExecuteCommandParams) (any, error) {
	commandSection, _, _ := strings.Cut(params.Command, ".")

	var edit *protocol.ApplyWorkspaceEditParams
	var err error

	switch commandSection {
	case "wireguard":
		edit, err = wireguard.WorkspaceExecuteCommand(context, params)
	}

	if err != nil {
		return nil, err
	}

	context.Notify(
		"workspace/applyEdit",
		edit,
	)

	return nil, nil
}
