package lsp

import (
	"config-lsp/common"
	"config-lsp/handlers/aliases"
	"config-lsp/handlers/aliases/ast"
	"config-lsp/handlers/aliases/handlers"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentDefinition(context *glsp.Context, params *protocol.DefinitionParams) ([]protocol.Location, error) {
	d := aliases.DocumentParserMap[params.TextDocument.URI]
	index := common.LSPCharacterAsIndexPosition(params.Position.Character)
	line := params.Position.Line

	rawEntry, found := d.Parser.Aliases.Get(line)

	if !found {
		return nil, nil
	}

	entry := rawEntry.(*ast.AliasEntry)

	if entry.Values != nil && entry.Values.Location.ContainsPosition(index) {
		rawValue := handlers.GetValueAtPosition(index, entry)

		if rawValue == nil {
			return nil, nil
		}

		return handlers.GetDefinitionLocationForValue(
			*d.Indexes,
			rawValue,
			params.TextDocument.URI,
		), nil
	}

	return nil, nil
}
