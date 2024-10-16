package lsp

import (
	"config-lsp/common"
	"config-lsp/handlers/sshd_config"
	"config-lsp/handlers/sshd_config/fields"
	"config-lsp/handlers/sshd_config/handlers"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

var matchOption = fields.CreateNormalizedName("Match")

func TextDocumentSignatureHelp(context *glsp.Context, params *protocol.SignatureHelpParams) (*protocol.SignatureHelp, error) {
	document := sshdconfig.DocumentParserMap[params.TextDocument.URI]

	line := uint32(params.Position.Line)
	cursor := common.LSPCharacterAsCursorPosition(params.Position.Character)

	if _, found := document.Config.CommentLines[line]; found {
		// Comment
		return nil, nil
	}

	option, block := document.Config.FindOption(line)

	if option != nil {
		if option.Key != nil {
			switch option.Key.Key {
			case matchOption:
				return handlers.GetMatchSignatureHelp(
					block,
					cursor,
				), nil
			}
		}

		return handlers.GetOptionSignatureHelp(option, cursor), nil
	} else {
		return handlers.GetOptionSignatureHelp(option, cursor), nil
	}
}
