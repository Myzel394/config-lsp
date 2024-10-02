package handlers

import (
	"config-lsp/common"
	"config-lsp/common/formatting"
	sshconfig "config-lsp/handlers/ssh_config"
	"config-lsp/handlers/ssh_config/ast"
	"config-lsp/utils"
	"fmt"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func getTagCompletions(
	d *sshconfig.SSHDocument,
	cursor common.CursorPosition,
	entry *ast.SSHOption,
) ([]protocol.CompletionItem, error) {
	return utils.MapMapToSlice(
		d.Indexes.Tags,
		func(name string, block *ast.SSHMatchBlock) protocol.CompletionItem {
			kind := protocol.CompletionItemKindModule
			text := renderMatchBlock(block)
			return protocol.CompletionItem{
				Label: name,
				Kind:  &kind,
				Documentation: protocol.MarkupContent{
					Kind:  protocol.MarkupKindMarkdown,
					Value: fmt.Sprintf("```sshconfig\n%s\n```", text),
				},
			}
		},
	), nil
}

func renderMatchBlock(
	block *ast.SSHMatchBlock,
) string {
	text := ""

	text += "Match " + formatMatchToString(block.MatchValue) + "\n"

	it := block.Options.Iterator()
	for it.Next() {
		option := it.Value().(*ast.SSHOption)
		text += formatOptionToString(option, formatting.DefaultFormattingOptions, blockOptionTemplate) + "\n"
	}

	return text
}
