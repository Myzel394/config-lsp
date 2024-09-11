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
) ([]protocol.CompletionItem, error) {
	kind := protocol.CompletionItemKindField
	format := protocol.InsertTextFormatSnippet

	return utils.MapMapToSlice(
		fields.Options,
		func(name string, rawValue docvalues.Value) protocol.CompletionItem {
			doc := rawValue.(docvalues.DocumentationValue)

			insertText := name + " " + "${1:value}"
			return protocol.CompletionItem{
				Label:            name,
				Kind:             &kind,
				Documentation:    doc.Documentation,
				InsertText:       &insertText,
				InsertTextFormat: &format,
			}
		},
	), nil
}
