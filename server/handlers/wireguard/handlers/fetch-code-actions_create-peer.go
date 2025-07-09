package handlers

import (
	"config-lsp/handlers/wireguard"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func GetAddPeerLikeThisCodeActions(
	d *wireguard.WGDocument,
	params *protocol.CodeActionParams,
) []protocol.CodeAction {
	// First, check if is on peer line

	line := params.Range.Start.Line

	section := d.Config.FindSectionByLine(line)

	// Check if section can be copied
	if section == nil || section.Start.Line != line || section.Header.Name != "Peer" {
		return nil
	}

	// Then add option
	commandID := "wireguard." + CodeActionCreatePeer
	command := protocol.Command{
		Title:   "Create new Peer based on this one",
		Command: string(commandID),
		Arguments: []any{
			CodeActionCreatePeerArgs{
				URI:  params.TextDocument.URI,
				Line: line,
			},
		},
	}
	return []protocol.CodeAction{
		{
			Title:   "Create new Peer based on this one",
			Command: &command,
		},
	}
}
