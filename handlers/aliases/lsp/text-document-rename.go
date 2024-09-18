package lsp

import (
	"config-lsp/handlers/aliases"
	"config-lsp/handlers/aliases/ast"
	"config-lsp/handlers/aliases/handlers"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentRename(context *glsp.Context, params *protocol.RenameParams) (*protocol.WorkspaceEdit, error) {
	d := aliases.DocumentParserMap[params.TextDocument.URI]
	character := params.Position.Character
	line := params.Position.Line

	rawEntry, found := d.Parser.Aliases.Get(line)

	if !found {
		return nil, nil
	}

	entry := rawEntry.(*ast.AliasEntry)

	if character >= entry.Key.Location.Start.Character && character <= entry.Key.Location.End.Character {
		changes := handlers.RenameAlias(
			*d.Indexes,
			entry.Key.Value,
			params.NewName,
		)

		return &protocol.WorkspaceEdit{
			Changes: map[protocol.DocumentUri][]protocol.TextEdit{
				params.TextDocument.URI: changes,
			},
		}, nil
	}

	if entry.Values != nil && character >= entry.Values.Location.Start.Character && character <= entry.Values.Location.End.Character {
		rawValue := handlers.GetValueAtCursor(character, entry)

		if rawValue == nil {
			return nil, nil
		}

		switch (*rawValue).(type) {
		case ast.AliasValueUser:
			userValue := (*rawValue).(ast.AliasValueUser)

			changes := handlers.RenameAlias(
				*d.Indexes,
				userValue.Value,
				params.NewName,
			)

			return &protocol.WorkspaceEdit{
				Changes: map[protocol.DocumentUri][]protocol.TextEdit{
					params.TextDocument.URI: changes,
				},
			}, nil
		}
	}

	return nil, nil
}
