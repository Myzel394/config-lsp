package handlers

import (
	"config-lsp/common"
	docvalues "config-lsp/doc-values"
	sshdconfig "config-lsp/handlers/sshd_config"
	"config-lsp/handlers/sshd_config/ast"
	"config-lsp/handlers/sshd_config/fields"
	"config-lsp/utils"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

var matchOption = fields.CreateNormalizedName("Match")

func GetRootCompletions(
	d *sshdconfig.SSHDDocument,
	parentMatchBlock *ast.SSHDMatchBlock,
	suggestValue bool,
) ([]protocol.CompletionItem, error) {
	if d.Indexes == nil {
		return nil, nil
	}

	kind := protocol.CompletionItemKindField

	availableOptions := make(map[fields.NormalizedOptionName]docvalues.DocumentationValue, 0)

	for key, option := range fields.Options {
		var exists = false

		// Don't allow duplicates
		if optionsMap, found := d.Indexes.AllOptionsPerName[key]; found {
			if _, found := optionsMap[parentMatchBlock]; found {
				exists = true
			}
		}

		if exists && !utils.KeyExists(fields.AllowedDuplicateOptions, key) {
			continue
		}

		if parentMatchBlock != nil && !utils.KeyExists(fields.MatchAllowedOptions, key) {
			continue
		}

		availableOptions[key] = option
	}

	return utils.MapMapToSlice(
		availableOptions,
		func(normalizedName fields.NormalizedOptionName, doc docvalues.DocumentationValue) protocol.CompletionItem {
			name := fields.FieldsNameFormattedMap[normalizedName]
			completion := &protocol.CompletionItem{
				Label:         name,
				Kind:          &kind,
				Documentation: doc.Documentation,
			}

			if suggestValue {
				format := protocol.InsertTextFormatSnippet
				insertText := name + " " + "${1:value}"

				completion.InsertTextFormat = &format
				completion.InsertText = &insertText
			}

			return *completion
		},
	), nil
}

func GetOptionCompletions(
	d *sshdconfig.SSHDDocument,
	entry *ast.SSHDOption,
	matchBlock *ast.SSHDMatchBlock,
	cursor common.CursorPosition,
) []protocol.CompletionItem {
	key := entry.Key.Key
	option, found := fields.Options[key]

	if !found {
		return nil
	}

	if entry.Key.Key == matchOption {
		return getMatchCompletions(
			d,
			cursor,
			matchBlock.MatchValue,
		)
	}

	if entry.OptionValue == nil {
		return option.DeprecatedFetchCompletions("", 0)
	}

	// token completions
	completions := getTokenCompletions(entry, cursor)

	// Hello wo|rld
	line := entry.OptionValue.Value.Raw
	// NEW: docvalues index
	completions = append(completions, option.DeprecatedFetchCompletions(
		line,
		common.DeprecatedImprovedCursorToIndex(
			cursor,
			line,
			entry.OptionValue.Start.Character,
		),
	)...)

	return completions
}

func getTokenCompletions(
	entry *ast.SSHDOption,
	cursor common.CursorPosition,
) []protocol.CompletionItem {
	completions := make([]protocol.CompletionItem, 0)
	index := common.CursorToCharacterIndex(uint32(cursor))

	if entry.Value.Raw[index] == '%' {
		if tokens, found := fields.OptionsTokensMap[entry.Key.Key]; found {
			for _, token := range tokens {
				description := fields.AvailableTokens[token]
				kind := protocol.CompletionItemKindConstant

				completions = append(completions, protocol.CompletionItem{
					Label:         token,
					Kind:          &kind,
					Documentation: description,
				})
			}
		}
	}

	return completions
}
