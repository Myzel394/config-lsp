package handlers

import (
	"config-lsp/common"
	docvalues "config-lsp/doc-values"
	sshconfig "config-lsp/handlers/ssh_config"
	"config-lsp/handlers/ssh_config/ast"
	"config-lsp/handlers/ssh_config/fields"
	"config-lsp/utils"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func GetRootCompletions(
	d *sshconfig.SSHDocument,
	parentBlock ast.SSHBlock,
	suggestValue bool,
) ([]protocol.CompletionItem, error) {
	kind := protocol.CompletionItemKindField

	availableOptions := make(map[fields.NormalizedOptionName]docvalues.DocumentationValue, 0)

	for key, option := range fields.Options {
		// Check for duplicates
		if d.DoesOptionExist(key, parentBlock) && !utils.KeyExists(fields.AllowedDuplicateOptions, key) {
			continue
		}

		if parentBlock != nil && parentBlock.GetBlockType() == ast.SSHBlockTypeHost && utils.KeyExists(fields.HostDisallowedOptions, key) {
			continue
		}

		if parentBlock == nil && utils.KeyExists(fields.GlobalDisallowedOptions, key) {
			continue
		}

		availableOptions[key] = option
	}

	return utils.MapMapToSlice(
		availableOptions,
		func(normalizedName fields.NormalizedOptionName, doc docvalues.DocumentationValue) protocol.CompletionItem {
			name := fields.FieldsNameFormattedMap[normalizedName]

			completion := &protocol.CompletionItem{
				Label:         string(name),
				Kind:          &kind,
				Documentation: doc.Documentation,
			}

			if suggestValue {
				format := protocol.InsertTextFormatSnippet
				insertText := string(name) + " " + "${1:value}"

				completion.InsertTextFormat = &format
				completion.InsertText = &insertText
			}

			return *completion
		},
	), nil
}

func GetOptionCompletions(
	d *sshconfig.SSHDocument,
	entry *ast.SSHOption,
	block ast.SSHBlock,
	line uint32,
	cursor common.CursorPosition,
) ([]protocol.CompletionItem, error) {
	option, found := fields.Options[entry.Key.Key]

	if !found {
		return nil, nil
	}

	if entry.Key.Key == matchOption {
		matchBlock := block.(*ast.SSHMatchBlock)
		return getMatchCompletions(
			d,
			cursor,
			matchBlock.MatchValue,
		)
	}
	if entry.Key.Key == tagOption {
		return getTagCompletions(
			d,
			line,
			cursor,
			entry,
		)
	}

	if entry.OptionValue == nil {
		return option.DeprecatedFetchCompletions("", 0), nil
	}

	// Hello wo|rld
	lineValue := entry.OptionValue.Value.Raw
	// NEW: docvalues index
	return option.DeprecatedFetchCompletions(
		lineValue,
		common.DeprecatedImprovedCursorToIndex(
			cursor,
			lineValue,
			entry.OptionValue.Start.Character,
		),
	), nil
}