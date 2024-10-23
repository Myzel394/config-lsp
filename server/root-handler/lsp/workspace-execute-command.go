package lsp

import (
	aliases "config-lsp/handlers/aliases/lsp"
	hosts "config-lsp/handlers/hosts/lsp"
	sshconfig "config-lsp/handlers/ssh_config/lsp"
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
	case "hosts":
		edit, err = hosts.WorkspaceExecuteCommand(context, params)
	case "aliases":
		edit, err = aliases.WorkspaceExecuteCommand(context, params)
	case "sshconfig":
		edit, err = sshconfig.WorkspaceExecuteCommand(context, params)
	}

	if err != nil {
		return nil, err
	}

	// Seems like `context.Call` is blocking, so we move it to a goroutine
	go context.Call(
		protocol.ServerWorkspaceApplyEdit,
		edit,
		nil,
	)

	return nil, nil
}
