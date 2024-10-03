package handlers

import (
	"config-lsp/common"
	docvalues "config-lsp/doc-values"
	sshdconfig "config-lsp/handlers/sshd_config"
	"config-lsp/handlers/sshd_config/ast"
	"config-lsp/handlers/sshd_config/fields"
	"config-lsp/utils"
	"fmt"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func GetRootCompletions(
	d *sshdconfig.SSHDDocument,
	parentMatchBlock *ast.SSHDMatchBlock,
	suggestValue bool,
) ([]protocol.CompletionItem, error) {
	println("getting root completions and the parnet Match block eta:")
	println(fmt.Sprintf("%v", parentMatchBlock))
	kind := protocol.CompletionItemKindField

	availableOptions := make(map[string]docvalues.DocumentationValue, 0)

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
	d *sshdconfig.SSHDDocument,
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
		return option.DeprecatedFetchCompletions("", 0), nil
	}

	// Hello wo|rld
	line := entry.OptionValue.Value.Raw
	// NEW: docvalues index
	return option.DeprecatedFetchCompletions(
		line,
		common.DeprecatedImprovedCursorToIndex(
			cursor,
			line,
			entry.OptionValue.Start.Character,
		),
	), nil
}
