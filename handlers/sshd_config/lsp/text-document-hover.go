package lsp

import (
	sshdconfig "config-lsp/handlers/sshd_config"
	"config-lsp/handlers/sshd_config/handlers"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentHover(
	context *glsp.Context,
	params *protocol.HoverParams,
) (*protocol.Hover, error) {
	line := params.Position.Line
	cursor := params.Position.Character

	d := sshdconfig.DocumentParserMap[params.TextDocument.URI]

	option, matchBlock := d.Config.FindOption(line)

	if option == nil || option.Key == nil {
		// Empty line
		return nil, nil
	}

	return handlers.GetHoverInfoForOption(
		option,
		matchBlock,
		line,
		cursor,
	)
}
