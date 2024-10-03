package lsp

import (
	sshconfig "config-lsp/handlers/ssh_config"
	"config-lsp/handlers/ssh_config/fields"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

var tagOption = fields.CreateNormalizedName("Tag")

func TextDocumentDefinition(context *glsp.Context, params *protocol.DefinitionParams) ([]protocol.Location, error) {
	line := params.Position.Line

	d := sshconfig.DocumentParserMap[params.TextDocument.URI]

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
