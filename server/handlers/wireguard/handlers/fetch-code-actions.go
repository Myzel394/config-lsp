package handlers

/*
import (
	"config-lsp/handlers/wireguard/ast"
	"config-lsp/handlers/wireguard/commands"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func GetKeepaliveCodeActions(
	p *ast.WGConfig,
	params *protocol.CodeActionParams,
) []protocol.CodeAction {
	line := params.Range.Start.Line

	for index, section := range p.Sections {
		if section.StartLine >= line && line <= section.EndLine && section.Header != nil && *section.Header == "Peer" {
			if section.ExistsProperty("Endpoint") && !section.ExistsProperty("PersistentKeepalive") {
				commandID := "wireguard." + CodeActionAddKeepalive
				command := protocol.Command{
					Title:   "Add PersistentKeepalive",
					Command: string(commandID),
					Arguments: []any{
						CodeActionAddKeepaliveArgs{
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

func GetKeyGenerationCodeActions(
	p *ast.WGConfig,
	params *protocol.CodeActionParams,
) []protocol.CodeAction {
	line := params.Range.Start.Line
	section, property := p.GetPropertyByLine(line)

	if section == nil || property == nil || property.Separator == nil {
		return nil
	}

	switch property.Key.Name {
	case "PrivateKey":
		if !wgcommands.AreWireguardToolsAvailable() {
			return nil
		}

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
*/
