package handlers

import (
	"config-lsp/handlers/wireguard"
	"config-lsp/handlers/wireguard/commands"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func GetKeyGenerationCodeActions(
	d *wireguard.WGDocument,
	params *protocol.CodeActionParams,
) []protocol.CodeAction {
	if !wgcommands.AreWireguardToolsAvailable() {
		return nil
	}

	line := params.Range.Start.Line

	section := d.Config.FindSectionByLine(line)
	property := d.Config.FindPropertyByLine(line)

	if section == nil || property == nil || property.Separator == nil {
		return nil
	}

	switch property.Key.Name {
	case "PrivateKey":
		commandID := "wireguard." + CodeActionGeneratePrivateKey
		command := protocol.Command{
			Title:   "Generate Private Key",
			Command: string(commandID),
			Arguments: []any{
				CodeActionGeneratePrivateKeyArgs{
					URI:  params.TextDocument.URI,
					Line: line,
				},
			},
		}

		return []protocol.CodeAction{
			{
				Title:   "Generate Private Key",
				Command: &command,
			},
		}
	case "PresharedKey":
		if !wgcommands.AreWireguardToolsAvailable() {
			return nil
		}

		commandID := "wireguard." + CodeActionGeneratePresharedKey
		command := protocol.Command{
			Title:   "Generate PresharedKey",
			Command: string(commandID),
			Arguments: []any{
				CodeActionGeneratePresharedKeyArgs{
					URI:  params.TextDocument.URI,
					Line: line,
				},
			},
		}

		return []protocol.CodeAction{
			{
				Title:   "Generate PresharedKey",
				Command: &command,
			},
		}
	}

	return nil
}
