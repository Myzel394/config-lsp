package handlers

import (
	"config-lsp/common"
	"config-lsp/common/formatting"
	sshconfig "config-lsp/handlers/ssh_config"
	"config-lsp/handlers/ssh_config/ast"
	"fmt"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func getTagCompletions(
	d *sshconfig.SSHDocument,
	line uint32,
	cursor common.CursorPosition,
	entry *ast.SSHOption,
) ([]protocol.CompletionItem, error) {
	completions := make([]protocol.CompletionItem, 0)

	for name, info := range d.Indexes.Tags {
		if info.Block.Start.Line < line {
			continue
		}

		kind := protocol.CompletionItemKindModule
		text := renderMatchBlock(info.Block)
		completions = append(completions, protocol.CompletionItem{
			Label: name,
			Kind:  &kind,
			Documentation: protocol.MarkupContent{
				Kind:  protocol.MarkupKindMarkdown,
				Value: fmt.Sprintf("```sshconfig\n%s\n```", text),
			},
		})
	}

	return completions, nil
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
