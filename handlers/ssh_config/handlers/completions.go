package handlers

import (
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

	availableOptions := make(map[string]docvalues.DocumentationValue, 0)

	for key, option := range fields.Options {
		alreadyExists := d.FindOptionByNameAndBlock(key, parentBlock) != nil

		if !alreadyExists || utils.KeyExists(fields.AllowedDuplicateOptions, key) {
			availableOptions[key] = option
		}
	}

	// Remove all fields that are already present and are not allowed to be duplicated
	for _, info := range d.Config.GetAllOptions() {
		if _, found := fields.AllowedDuplicateOptions[info.Option.Key.Key]; found {
			continue
		}

		delete(availableOptions, info.Option.Key.Key)
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
