package lsp

import (
	"config-lsp/common"
	sshconfig "config-lsp/handlers/ssh_config"
	"config-lsp/handlers/ssh_config/fields"
	"config-lsp/handlers/ssh_config/handlers"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

var tagOption = fields.CreateNormalizedName("Tag")

func TextDocumentDefinition(context *glsp.Context, params *protocol.DefinitionParams) ([]protocol.Location, error) {
	line := params.Position.Line
	index := common.LSPCharacterAsIndexPosition(params.Position.Character)
	d := sshconfig.DocumentParserMap[params.TextDocument.URI]

	if include, found := d.Indexes.Includes[line]; found {
		return handlers.GetIncludeOptionLocation(
			include,
			index,
		), nil
	}

	option, _ := d.Config.FindOption(line)

	if option != nil && option.Key.Key == tagOption && option.OptionValue != nil {
		if info, found := d.Indexes.Tags[option.OptionValue.Value.Value]; found {
			return []protocol.Location{
				{
					URI:   params.TextDocument.URI,
					Range: info.Block.ToLSPRange(),
				},
			}, nil
		}
	}

	return nil, nil
}
