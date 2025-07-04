package handlers

import (
	"config-lsp/handlers/wireguard"
	"config-lsp/utils"
	"errors"
	"strings"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

type CodeActionGeneratePostdownKeyArgs struct {
	URI protocol.DocumentUri
}

func CodeActionGeneratePostdownArgsFromArguments(arguments map[string]any) CodeActionGeneratePostdownKeyArgs {
	return CodeActionGeneratePostdownKeyArgs{
		URI: arguments["URI"].(protocol.DocumentUri),
	}
}

func (args CodeActionGeneratePostdownKeyArgs) RunCommand(d *wireguard.WGDocument) (*protocol.ApplyWorkspaceEditParams, error) {
	if utils.BlockUntilIndexesNotNil(d.Indexes) == false {
		return nil, errors.New("Indexes are not ready yet")
	}

	lastUpPropertyKey := utils.FindBiggestKey(d.Indexes.UpProperties)
	lastUpProperty := d.Indexes.UpProperties[lastUpPropertyKey]

	// TODO: Find out if the user specified multiple PreDown or only one (or maybe don't do this at all)

	rules := findAllIPTableRules(d)
	invertedRules := generateInvertedRules(rules)

	if len(invertedRules) == 0 {
		return nil, errors.New("No valid iptables rules found to invert")
	}

	newRulesString := strings.Join(invertedRules, "; ")
	newPropertyString := "\nPostDown = " + newRulesString + "\n"

	label := "Generate PostDown with inverted rules"
	return &protocol.ApplyWorkspaceEditParams{
		Label: &label,
		Edit: protocol.WorkspaceEdit{
			Changes: map[protocol.DocumentUri][]protocol.TextEdit{
				args.URI: {
					{
						Range: protocol.Range{
							Start: lastUpProperty.Property.Value.ToLSPRange().End,
							End:   lastUpProperty.Property.Value.ToLSPRange().End,
						},
						NewText: newPropertyString,
					},
				},
			},
		},
	}, nil
}

func findAllIPTableRules(
	d *wireguard.WGDocument,
) []string {
	upRules := make([]string, 0)

	for _, info := range d.Indexes.UpProperties {
		if info.Property != nil && info.Property.Value != nil {
			rulesRaw := info.Property.Value.Value
			rules := strings.Split(rulesRaw, ";")

			upRules = append(upRules, rules...)
		}
	}

	return upRules
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
