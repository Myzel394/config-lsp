package lsp

import (
	sshdconfig "config-lsp/handlers/sshd_config"
	"config-lsp/handlers/sshd_config/handlers"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentCompletion(context *glsp.Context, params *protocol.CompletionParams) (any, error) {
	line := params.Position.Line
	cursor := params.Position.Character
	_ = cursor

	d := sshdconfig.DocumentParserMap[params.TextDocument.URI]

	if _, found := d.Config.CommentLines[line]; found {
		return nil, nil
	}

	entry, matchBlock := d.Config.FindOption(line)

	if entry == nil {
		// Empty line
		return handlers.GetRootCompletions(
			d,
			matchBlock,
		)
	}

	return nil, nil
}
