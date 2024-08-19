package wireguard

import protocol "github.com/tliron/glsp/protocol_3_16"

func getKeepaliveCodeActions(
	params *protocol.CodeActionParams,
	parser *wireguardParser,
) []protocol.CodeAction {
	line := params.Range.Start.Line

	for index, section := range parser.Sections {
		if section.StartLine >= line && line <= section.EndLine && section.Name != nil && *section.Name == "Peer" {
			if section.fetchFirstProperty("Endpoint") != nil && section.fetchFirstProperty("PersistentKeepalive") == nil {
				commandID := "wireguard." + codeActionAddKeepalive
				command := protocol.Command{
					Title:   "Add PersistentKeepalive",
					Command: string(commandID),
					Arguments: []any{
						codeActionAddKeepaliveArgs{
							URI:          params.TextDocument.URI,
							SectionIndex: uint32(index),
						},
					},
				}

				return []protocol.CodeAction{
					{
						Title:   "Add PersistentKeepalive",
						Command: &command,
					},
				}
			}
		}
	}

	return nil
}

func getKeyGenerationCodeActions(
	params *protocol.CodeActionParams,
	parser *wireguardParser,
) []protocol.CodeAction {
	line := params.Range.Start.Line
	section, property := parser.getPropertyByLine(line)

	if section == nil || property == nil || property.Separator == nil {
		return nil
	}

	switch property.Key.Name {
	case "PrivateKey":
		if !areWireguardToolsAvailable() {
			return nil
		}

		commandID := "wireguard." + codeActionGeneratePrivateKey
		command := protocol.Command{
			Title:   "Generate Private Key",
			Command: string(commandID),
			Arguments: []any{
				codeActionGeneratePrivateKeyArgs{
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
		if !areWireguardToolsAvailable() {
			return nil
		}

		commandID := "wireguard." + codeActionGeneratePresharedKey
		command := protocol.Command{
			Title:   "Generate PresharedKey",
			Command: string(commandID),
			Arguments: []any{
				codeActionGeneratePresharedKeyArgs{
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
