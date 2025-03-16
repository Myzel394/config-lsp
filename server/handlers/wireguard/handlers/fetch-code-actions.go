package handlers

import (
	"config-lsp/handlers/wireguard"
	"config-lsp/handlers/wireguard/commands"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func GetKeepaliveCodeActions(
	d *wireguard.WGDocument,
	params *protocol.CodeActionParams,
) []protocol.CodeAction {
	line := params.Range.Start.Line

	for _, section := range d.Indexes.SectionsByName["Peer"] {
		if section.Start.Line >= line && line <= section.End.Line {
			if section.FindPropertyByName("Endpoint") != nil && section.FindFirstPropertyByName("PersistentKeepalive") == nil {
				var insertionLine uint32
				lastProperty := section.GetLastProperty()

				if lastProperty == nil {
					insertionLine = section.End.Line
				} else {
					insertionLine = lastProperty.End.Line + 1
				}

				return []protocol.CodeAction{
					{
						Title: "Add PersistentKeepalive",
						Edit: &protocol.WorkspaceEdit{
							Changes: map[protocol.DocumentUri][]protocol.TextEdit{
								params.TextDocument.URI: {
									{
										Range: protocol.Range{
											Start: protocol.Position{
												Line:      insertionLine,
												Character: 0,
											},
											End: protocol.Position{
												Line:      insertionLine,
												Character: 0,
											},
										},
										NewText: "PersistentKeepalive = 25\n",
									},
								},
							},
						},
					},
				}
			}
		}
	}

	return nil
}

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
