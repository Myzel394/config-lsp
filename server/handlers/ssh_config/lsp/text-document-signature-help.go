package lsp

import (
	"config-lsp/common"
	sshconfig "config-lsp/handlers/ssh_config"
	"config-lsp/handlers/ssh_config/ast"
	"config-lsp/handlers/ssh_config/fields"
	"config-lsp/handlers/ssh_config/handlers"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

var hostOption = fields.CreateNormalizedName("Host")

func TextDocumentSignatureHelp(context *glsp.Context, params *protocol.SignatureHelpParams) (*protocol.SignatureHelp, error) {
	document := sshconfig.DocumentParserMap[params.TextDocument.URI]

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
					block.(*ast.SSHMatchBlock),
					cursor,
				), nil
			case hostOption:
				return handlers.GetHostSignatureHelp(
					block.(*ast.SSHHostBlock),
					cursor,
				), nil
			}
		}

		return handlers.GetOptionSignatureHelp(option, cursor), nil
	} else {
		return handlers.GetOptionSignatureHelp(option, cursor), nil
	}
}
