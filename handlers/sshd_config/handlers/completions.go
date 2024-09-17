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

	availableOptions := make(map[string]docvalues.DocumentationValue)

	if parentMatchBlock == nil {
		availableOptions = fields.Options
	} else {
		for option := range fields.MatchAllowedOptions {
			if opt, found := fields.Options[option]; found {
				availableOptions[option] = opt
			}
		}
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
	cursor uint32,
) ([]protocol.CompletionItem, error) {
	option, found := fields.Options[entry.Key.Value]

	if !found {
		return nil, nil
	}

	if entry.Key.Value == "Match" {
		return getMatchCompletions(
			d,
			matchBlock.MatchValue,
			cursor-matchBlock.MatchEntry.Start.Character,
		)
	}

	line := entry.OptionValue.Value

	if entry.OptionValue == nil {
		return option.FetchCompletions("", 0), nil
	}

	return option.FetchCompletions(line, common.CursorToCharacterIndex(cursor)), nil
}
