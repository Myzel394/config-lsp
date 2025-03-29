package handlers

import (
	"config-lsp/common"
	sshconfig "config-lsp/handlers/ssh_config"
	"config-lsp/handlers/ssh_config/diagnostics"
	"config-lsp/handlers/ssh_config/fields"
	"config-lsp/utils"
	"fmt"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func getKeywordTypoFixes(
	d *sshconfig.SSHDocument,
	params *protocol.CodeActionParams,
) []protocol.CodeAction {
	if common.ServerOptions.NoTypoSuggestions {
		return nil
	}

	line := params.Range.Start.Line

	if typoOption, found := d.Indexes.UnknownOptions[line]; found {
		name := typoOption.Option.Key.Value.Value

		opts := utils.KeysOfMap(fields.Options)
		suggestedOptions := common.FindSimilarItems(fields.CreateNormalizedName(name), opts)

		actions := make([]protocol.CodeAction, 0, len(suggestedOptions))

		kind := protocol.CodeActionKindQuickFix
		for index, normalizedOptionName := range suggestedOptions {
			isPreferred := index == 0
			optionName := fields.FieldsNameFormattedMap[normalizedOptionName]

			actions = append(actions, protocol.CodeAction{
				Title:       fmt.Sprintf("Typo Fix: %s", optionName),
				IsPreferred: &isPreferred,
				Kind:        &kind,
				Diagnostics: []protocol.Diagnostic{
					diagnostics.GenerateUnknownOption(
						typoOption.Option.Key.ToLSPRange(),
						typoOption.Option.Key.Value.Value,
					),
				},
				Edit: &protocol.WorkspaceEdit{
					Changes: map[protocol.DocumentUri][]protocol.TextEdit{
						params.TextDocument.URI: {
							{
								Range:   typoOption.Option.Key.ToLSPRange(),
								NewText: optionName,
							},
						},
					},
				},
			})
		}

		return actions
	}

	return nil
}
