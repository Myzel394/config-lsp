package lsp

import (
	"config-lsp/handlers/aliases"
	"config-lsp/handlers/aliases/ast"
	"config-lsp/handlers/aliases/handlers"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentDefinition(context *glsp.Context, params *protocol.DefinitionParams) ([]protocol.Location, error) {
	d := aliases.DocumentParserMap[params.TextDocument.URI]
	character := params.Position.Character
	line := params.Position.Line

	rawEntry, found := d.Parser.Aliases.Get(line)

	if !found {
		return nil, nil
	}

	entry := rawEntry.(*ast.AliasEntry)

	if entry.Values != nil && character >= entry.Values.Location.Start.Character && character <= entry.Values.Location.End.Character {
		rawValue := handlers.GetValueAtCursor(character, entry)

		if rawValue == nil {
			return nil, nil
		}

		return handlers.GetDefinitionLocationForValue(
			*d.Indexes,
			*rawValue,
			params,
		), nil
	}

	return nil, nil
}
