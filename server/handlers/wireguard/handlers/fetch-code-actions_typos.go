package handlers

import (
	"config-lsp/common"
	"config-lsp/handlers/wireguard"
	"config-lsp/handlers/wireguard/diagnostics"
	"config-lsp/handlers/wireguard/fields"
	"config-lsp/utils"
	"fmt"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func GetPropertyKeywordTypoFixes(
	d *wireguard.WGDocument,
	params *protocol.CodeActionParams,
) []protocol.CodeAction {
	if common.ServerOptions.NoTypoSuggestions {
		return nil
	}

	if utils.BlockUntilIndexesNotNil(d.Indexes) == false {
		return nil
	}

	line := params.Range.Start.Line

	if typoInfo, found := d.Indexes.UnknownProperties[line]; found {
		if options, found := fields.OptionsHeaderMap[fields.CreateNormalizedName(typoInfo.Section.Header.Name)]; found {
			normalizedPropertyKey := fields.CreateNormalizedName(typoInfo.Property.Key.Name)
			opts := utils.KeysOfMap(options)

			suggestedProperties := common.FindSimilarItems(normalizedPropertyKey, opts)

			actions := make([]protocol.CodeAction, 0, len(suggestedProperties))

			kind := protocol.CodeActionKindQuickFix
			for index, normalizedPropertyName := range suggestedProperties {
				isPreferred := index == 0
				optionName := fields.AllOptionsFormatted[normalizedPropertyName]

				actions = append(actions, protocol.CodeAction{
					Title:       fmt.Sprintf("Typo Fix: %s", optionName),
					IsPreferred: &isPreferred,
					Kind:        &kind,
					Diagnostics: []protocol.Diagnostic{
						diagnostics.GenerateUnknownOption(
							typoInfo.Property.Key.ToLSPRange(),
							typoInfo.Property.Key.Name,
						),
					},
					Edit: &protocol.WorkspaceEdit{
						Changes: map[protocol.DocumentUri][]protocol.TextEdit{
							params.TextDocument.URI: {
								{
									Range:   typoInfo.Property.Key.ToLSPRange(),
									NewText: optionName,
								},
							},
						},
					},
				})
			}

			return actions
		}
	}

	return nil
}
