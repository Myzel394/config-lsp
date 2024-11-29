package lsp

import (
	"config-lsp/common"
	"config-lsp/handlers/aliases"
	"config-lsp/handlers/aliases/ast"
	"config-lsp/handlers/aliases/handlers"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentRename(context *glsp.Context, params *protocol.RenameParams) (*protocol.WorkspaceEdit, error) {
	d := aliases.DocumentParserMap[params.TextDocument.URI]
	index := common.LSPCharacterAsIndexPosition(params.Position.Character)
	line := params.Position.Line

	rawEntry, found := d.Parser.Aliases.Get(line)

	if !found {
		return nil, nil
	}

	entry := rawEntry.(*ast.AliasEntry)

	if entry.Key.Location.ContainsPosition(index) {
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

	if entry.Values != nil && entry.Values.Location.ContainsPosition(index) {
		rawValue := handlers.GetValueAtPosition(index, entry)

		if rawValue == nil {
			return nil, nil
		}

		switch rawValue.(type) {
		case ast.AliasValueUser:
			userValue := rawValue.(ast.AliasValueUser)

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
