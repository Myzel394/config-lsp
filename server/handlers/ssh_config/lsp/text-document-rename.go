package lsp

import (
	"config-lsp/common"
	"config-lsp/common/formatting"
	sshconfig "config-lsp/handlers/ssh_config"
	"config-lsp/handlers/ssh_config/ast"
	"config-lsp/handlers/ssh_config/fields"
	"config-lsp/handlers/ssh_config/handlers"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

var nameTemplate = formatting.FormatTemplate("/!'%s/!'")
var matchOption = fields.CreateNormalizedName("Match")

func TextDocumentRename(context *glsp.Context, params *protocol.RenameParams) (*protocol.WorkspaceEdit, error) {
	d := sshconfig.DocumentParserMap[params.TextDocument.URI]
	index := common.LSPCharacterAsIndexPosition(params.Position.Character)
	line := params.Position.Line

	option, block := d.Config.FindOption(line)

	if option != nil && option.OptionValue != nil && option.OptionValue.Value.Value != "" && option.OptionValue.ContainsPosition(index) {
		newName := nameTemplate.Format(formatting.DefaultFormattingOptions, params.NewName)

		if option.Key.Key == tagOption {
			return handlers.RenameTag(
				params,
				d,
				option.OptionValue.Value.Value,
				newName,
			)
		}

		if option.Key.Key == matchOption {
			matchBlock := block.(*ast.SSHMatchBlock)

			entry := matchBlock.MatchValue.GetEntryAtPosition(index)

			if entry != nil {
				value := entry.GetValueAtPosition(index)

				if value != nil {
					return handlers.RenameTag(
						params,
						d,
						value.Value.Value,
						newName,
					)
				}
			}
		}
	}

	return nil, nil
}
