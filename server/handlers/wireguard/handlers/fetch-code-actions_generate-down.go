package handlers

import (
	"config-lsp/handlers/wireguard"
	"config-lsp/handlers/wireguard/fields"
	"config-lsp/parsers/ini"
	"config-lsp/utils"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

var preUp = fields.CreateNormalizedName("PreUp")
var postUp = fields.CreateNormalizedName("PostUp")

func GetGenerateDownRuleCodeActions(
	d *wireguard.WGDocument,
	params *protocol.CodeActionParams,
) []protocol.CodeAction {
	line := uint32(params.Range.Start.Line)

	if utils.BlockUntilIndexesNotNil(d.Indexes) == false {
		return nil
	}

	section := d.Config.FindSectionByLine(line)

	if section == nil {
		return nil
	}

	rawProperty, found := section.Properties.Get(line)

	if !found {
		return nil
	}

	property := rawProperty.(*ini.Property)

	if property.Value != nil && property.Key != nil {
		var newProperty string = ""
		propertyName := fields.CreateNormalizedName(property.Key.Name)

		if propertyName == preUp && d.Indexes.AsymmetricRules[section].PreMissing {
			newProperty = "PreDown"
		} else if propertyName == postUp && d.Indexes.AsymmetricRules[section].PostMissing {
			newProperty = "PostDown"
		}

		if newProperty != "" {
			title := "Generate " + newProperty + " with inverted rules"
			commandID := "wireguard." + CodeActionGenerateDownRule

			command := protocol.Command{
				Title:   title,
				Command: string(commandID),
				Arguments: []any{
					CodeActionGenerateDownRuleArgs{
						URI:  params.TextDocument.URI,
						Line: line,
					},
				},
			}

			return []protocol.CodeAction{
				{
					Title:   title,
					Command: &command,
				},
			}
		}
	}

	return nil
}
