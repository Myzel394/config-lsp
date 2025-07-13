package handlers

import (
	"config-lsp/handlers/wireguard"
	"config-lsp/handlers/wireguard/fields"
	"config-lsp/parsers/ini"
	"config-lsp/utils"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func GetGeneratePostDownCodeActions(
	d *wireguard.WGDocument,
	params *protocol.CodeActionParams,
) []protocol.CodeAction {
	postPrePropertyNames := map[fields.NormalizedName]struct{}{
		fields.CreateNormalizedName("PreUp"):  {},
		fields.CreateNormalizedName("PostUp"): {},
	}

	line := params.Range.Start.Line

	section := d.Config.FindSectionByLine(line)

	if section == nil {
		return nil
	}

	rawProperty, found := section.Properties.Get(line)

	if !found {
		return nil
	}

	property := rawProperty.(*ini.Property)

	propertyName := fields.CreateNormalizedName(property.Key.Name)
	if (utils.KeyExists(postPrePropertyNames, propertyName)) && (property.Value != nil) {
		// Only propose this action if no PostDown is already present
		_, postDownProperty := section.FindFirstPropertyByName("PostDown")

		if postDownProperty == nil {
			commandID := "wireguard." + CodeActionGeneratePostDown

			command := protocol.Command{
				Title:   "Generate PostDown with inverted rules",
				Command: string(commandID),
				Arguments: []any{
					CodeActionGeneratePostdownKeyArgs{
						URI: params.TextDocument.URI,
					},
				},
			}

			return []protocol.CodeAction{
				{
					Title:   "Generate PostDown with inverted rules",
					Command: &command,
				},
			}
		}
	}

	return nil
}
