package lsp

import (
	"config-lsp/common"
	sshconfig "config-lsp/handlers/ssh_config"
	"config-lsp/handlers/ssh_config/ast"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentPrepareRename(context *glsp.Context, params *protocol.PrepareRenameParams) (any, error) {
	d := sshconfig.DocumentParserMap[params.TextDocument.URI]
	index := common.LSPCharacterAsIndexPosition(params.Position.Character)
	line := params.Position.Line

	option, block := d.Config.FindOption(line)

	if option == nil || option.Key == nil {
		// Empty line
		return nil, nil
	}

	if option.OptionValue != nil && option.OptionValue.Value.Value != "" && option.OptionValue.ContainsPosition(index) {
		if option.Key.Key == tagOption {
			return option.OptionValue.ToLSPRange(), nil
		}

		if option.Key.Key == matchOption {
			matchBlock := block.(*ast.SSHMatchBlock)
			entry := matchBlock.MatchValue.GetEntryAtPosition(common.LSPCharacterAsIndexPosition(params.Position.Character))

			if entry == nil {
				return nil, nil
			}

			value := entry.GetValueAtPosition(common.LSPCharacterAsIndexPosition(params.Position.Character))

			if value == nil {
				return nil, nil
			}

			return value.ToLSPRange(), nil
		}
	}

	return nil, nil
}
