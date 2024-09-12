package handlers

import (
	docvalues "config-lsp/doc-values"
	sshdconfig "config-lsp/handlers/sshd_config"
	"config-lsp/handlers/sshd_config/ast"
	"config-lsp/handlers/sshd_config/fields"
	"config-lsp/utils"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func GetRootCompletions(
	d *sshdconfig.SSHDocument,
	parentMatchBlock *ast.SSHMatchBlock,
	suggestValue bool,
) ([]protocol.CompletionItem, error) {
	kind := protocol.CompletionItemKindField

	return utils.MapMapToSlice(
		fields.Options,
		func(name string, rawValue docvalues.Value) protocol.CompletionItem {
			doc := rawValue.(docvalues.DocumentationValue)

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
	entry *ast.SSHOption,
	cursor uint32,
) ([]protocol.CompletionItem, error) {
	option, found := fields.Options[entry.Key.Value]

	if !found {
		return nil, nil
	}

	if entry.OptionValue == nil {
		return option.FetchCompletions("", 0), nil
	}

	relativeCursor := cursor - entry.OptionValue.Start.Character
	line := entry.OptionValue.Value

	return option.FetchCompletions(line, relativeCursor), nil
}
