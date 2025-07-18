package handlers

import (
	"config-lsp/common"
	bitcoinconf "config-lsp/handlers/bitcoin_conf"
	"config-lsp/handlers/bitcoin_conf/diagnostics"
	"config-lsp/handlers/bitcoin_conf/fields"
	"config-lsp/utils"
	"fmt"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func GetPropertyKeywordTypoFixes(
	d *bitcoinconf.BTCDocument,
	params *protocol.CodeActionParams,
) []protocol.CodeAction {
	if common.ServerOptions.NoTypoSuggestions {
		return nil
	}

	if utils.BlockUntilIndexesNotNil(d) == false {
		return nil
	}

	line := params.Range.Start.Line

	if typoInfo, found := d.Indexes.UnknownProperties[line]; found {
		opts := utils.KeysOfMap(fields.Options)

		suggestedProperties := common.FindSimilarItems(typoInfo.Property.Key.Name, opts)

		actions := make([]protocol.CodeAction, 0, len(suggestedProperties))

		kind := protocol.CodeActionKindQuickFix
		for index, optionName := range suggestedProperties {
			isPreferred := index == 0

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

	return nil
}
