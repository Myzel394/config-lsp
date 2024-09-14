package lsp

import (
	sshdconfig "config-lsp/handlers/sshd_config"
	"config-lsp/handlers/sshd_config/handlers"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentDefinition(context *glsp.Context, params *protocol.DefinitionParams) ([]protocol.Location, error) {
	d := sshdconfig.DocumentParserMap[params.TextDocument.URI]
	cursor := params.Position.Character
	line := params.Position.Line

	if include, found := d.Indexes.Includes[line]; found {
		relativeCursor := cursor - include.Option.Option.LocationRange.Start.Character

		return handlers.GetIncludeOptionLocation(include, relativeCursor), nil
	}

	return nil, nil
}
