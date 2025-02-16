package handlers

import (
	"config-lsp/common"
	sshconfig "config-lsp/handlers/ssh_config"
	"config-lsp/handlers/ssh_config/diagnostics"
	"config-lsp/handlers/ssh_config/fields"
	"fmt"

	"github.com/hbollon/go-edlib"
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

		suggestedOptions := findSimilarOptions(name)

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

// Find options that are similar to the given option name.
// This is used to find typos & suggest the correct option name.
// Once an option is found that has a Damerau-Levenshtein distance of 1, it is immediately returned.
// If not, then the next 2 options of similarity 2, or 3 options of similarity 3 are returned.
// If no options with similarity <= 3 are found, then an empty slice is returned.
func findSimilarOptions(
	optionName string,
) []fields.NormalizedOptionName {
	normalizedOptionName := string(fields.CreateNormalizedName(optionName))

	optionsPerSimilarity := map[uint8][]fields.NormalizedOptionName{
		2: make([]fields.NormalizedOptionName, 0, 2),
		3: make([]fields.NormalizedOptionName, 0, 3),
	}

	for name := range fields.Options {
		normalizedName := string(name)
		similarity := edlib.DamerauLevenshteinDistance(normalizedName, normalizedOptionName)

		switch similarity {
		case 1:
			return []fields.NormalizedOptionName{name}
		case 2:
			optionsPerSimilarity[2] = append(optionsPerSimilarity[2], name)

			if len(optionsPerSimilarity[2]) >= 2 {
				return optionsPerSimilarity[2]
			}
		case 3:
			optionsPerSimilarity[3] = append(optionsPerSimilarity[3], name)

			if len(optionsPerSimilarity[3]) >= 3 {
				return optionsPerSimilarity[3]
			}
		}
	}

	return append(optionsPerSimilarity[2], optionsPerSimilarity[3]...)
}
