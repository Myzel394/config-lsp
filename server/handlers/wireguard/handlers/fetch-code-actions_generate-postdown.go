package handlers

import (
	"config-lsp/handlers/wireguard"
	"config-lsp/handlers/wireguard/fields"
	"config-lsp/parsers/ini"
	"config-lsp/utils"
	"strings"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func GetGeneratePostDownCodeActions(
	d *wireguard.WGDocument,
	params *protocol.CodeActionParams,
) []protocol.CodeAction {
	postDownPropertyNames := map[fields.NormalizedName]struct{}{
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
	if (utils.KeyExists(postDownPropertyNames, propertyName)) && (property.Value != nil) {
		// Only propose this action if no PostDown is already present
		_, postDownProperty := section.FindFirstPropertyByName("PostDown")

		if postDownProperty != nil {
			return nil
		}

		// TODO: Find out if the user specified multiple PreDown or only one (or maybe don't do this at all)

		rules := property.Value.Value
		invertedRules := generateInvertedRules(strings.Split(rules, ";"))

		if len(invertedRules) == 0 {
			return nil
		}

		newRulesString := strings.Join(invertedRules, "; ")
		newPropertyString := "\nPostDown = " + newRulesString + "\n"

		return []protocol.CodeAction{
			{
				Title: "Generate PostDown with inverted rules",
				Edit: &protocol.WorkspaceEdit{
					Changes: map[protocol.DocumentUri][]protocol.TextEdit{
						params.TextDocument.URI: {
							{
								Range: protocol.Range{
									Start: property.Value.ToLSPRange().End,
									End:   property.Value.ToLSPRange().End,
								},
								NewText: newPropertyString,
							},
						},
					},
				},
			},
		}
	}

	return nil
}

func generateInvertedRules(rules []string) []string {
	var postDownRules []string

	for _, rule := range rules {
		ipTableRule, err := utils.ParseIpTableRule(rule)

		if err != nil {
			return nil
		}

		inverted := ipTableRule.InvertAction()

		if inverted.Action != utils.IpTableActionDelete {
			// We only want to generate actions for common actions
			return nil
		}

		ruleStringRaw := inverted.String()
		ruleString := strings.ReplaceAll(strings.TrimSpace(ruleStringRaw), "  ", " ")
		postDownRules = append(postDownRules, ruleString)
	}

	return postDownRules
}
