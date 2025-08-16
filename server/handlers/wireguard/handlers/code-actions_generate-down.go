package handlers

import (
	"config-lsp/handlers/wireguard"
	"config-lsp/utils"
	"errors"
	"fmt"
	"strings"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

type CodeActionGenerateDownRuleArgs struct {
	URI  protocol.DocumentUri
	Line uint32
}

func CodeActionGenerateDownRuleArgsFromArguments(arguments map[string]any) CodeActionGenerateDownRuleArgs {
	return CodeActionGenerateDownRuleArgs{
		URI:  arguments["URI"].(protocol.DocumentUri),
		Line: uint32(arguments["Line"].(float64)),
	}
}

func (args CodeActionGenerateDownRuleArgs) RunCommand(d *wireguard.WGDocument) (*protocol.ApplyWorkspaceEditParams, error) {
	if utils.BlockUntilIndexesNotNil(d.Indexes) == false {
		return nil, errors.New("Indexes are not ready yet")
	}

	section := d.Config.FindSectionByLine(args.Line)
	property := d.Config.FindPropertyByLine(args.Line)

	if section == nil || property == nil {
		return nil, errors.New("No section or property found at the specified line")
	}

	rules := strings.Split(property.Value.Value, ";")
	invertedRules := generateInvertedRules(rules)

	if len(invertedRules) == 0 {
		return nil, errors.New("No valid iptables rules found to invert")
	}

	var newKeyName string

	switch property.Key.Name {
	case "PreUp":
		newKeyName = "PreDown"
	case "PostUp":
		newKeyName = "PostDown"
	default:
		return nil, fmt.Errorf("unsupported key %q at line %d; only PreUp/PostUp are supported", property.Key.Name, args.Line)

	}

	newRulesString := strings.Join(invertedRules, "; ")
	newPropertyString := fmt.Sprintf("\n%s = %s", newKeyName, newRulesString)

	label := fmt.Sprintf("Generate %s with inverted rules", newKeyName)
	return &protocol.ApplyWorkspaceEditParams{
		Label: &label,
		Edit: protocol.WorkspaceEdit{
			Changes: map[protocol.DocumentUri][]protocol.TextEdit{
				args.URI: {
					{
						Range: protocol.Range{
							Start: property.ToLSPRange().End,
							End:   property.ToLSPRange().End,
						},
						NewText: newPropertyString,
					},
				},
			},
		},
	}, nil
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
