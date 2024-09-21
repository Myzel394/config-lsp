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

func GetRootCompletions(
	d *sshdconfig.SSHDocument,
	parentMatchBlock *ast.SSHDMatchBlock,
	suggestValue bool,
) ([]protocol.CompletionItem, error) {
	kind := protocol.CompletionItemKindField

	availableOptions := make(map[string]docvalues.DocumentationValue, 0)

	if parentMatchBlock == nil {
		for key, option := range fields.Options {
			if d.Indexes != nil && utils.KeyExists(d.Indexes.AllOptionsPerName, key) && !utils.KeyExists(fields.AllowedDuplicateOptions, key) {
				continue
			}

			availableOptions[key] = option
		}
	} else {
		for key := range fields.MatchAllowedOptions {
			if option, found := fields.Options[key]; found {
				if d.Indexes != nil && utils.KeyExists(d.Indexes.AllOptionsPerName, key) && !utils.KeyExists(fields.AllowedDuplicateOptions, key) {
					continue
				}

				availableOptions[key] = option
			}
		}
	}

	// Remove all fields that are already present and are not allowed to be duplicated
	for _, option := range d.Config.GetAllOptions() {
		if _, found := fields.AllowedDuplicateOptions[option.Key.Key]; found {
			continue
		}

		delete(availableOptions, option.Key.Key)
	}

	return utils.MapMapToSlice(
		availableOptions,
		func(name string, doc docvalues.DocumentationValue) protocol.CompletionItem {
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
	d *sshdconfig.SSHDocument,
	entry *ast.SSHDOption,
	matchBlock *ast.SSHDMatchBlock,
	cursor common.CursorPosition,
) ([]protocol.CompletionItem, error) {
	option, found := fields.Options[entry.Key.Key]

	if !found {
		return nil, nil
	}

	if entry.Key.Key == "Match" {
		return getMatchCompletions(
			d,
			cursor,
			matchBlock.MatchValue,
		)
	}

	if entry.OptionValue == nil {
		return option.FetchCompletions("", 0), nil
	}

	// Hello wo|rld
	line := entry.OptionValue.Value.Raw
	// NEW: docvalues index
	return option.FetchCompletions(
		line,
		common.DeprecatedImprovedCursorToIndex(
			cursor,
			line,
			entry.OptionValue.Start.Character,
		),
	), nil
}
