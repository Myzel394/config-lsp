package lsp

import (
	"config-lsp/handlers/aliases"
	"config-lsp/handlers/aliases/ast"
	"config-lsp/handlers/aliases/handlers"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentSignatureHelp(context *glsp.Context, params *protocol.SignatureHelpParams) (*protocol.SignatureHelp, error) {
	document := aliases.DocumentParserMap[params.TextDocument.URI]

	line := params.Position.Line
	character := params.Position.Character

	if _, found := document.Parser.CommentLines[line]; found {
		// Comment
		return nil, nil
	}

	rawEntry, found := document.Parser.Aliases.Get(line)

	if !found {
		return handlers.GetRootSignatureHelp(0), nil
	}

	entry := rawEntry.(*ast.AliasEntry)

	if entry.Key != nil && character >= entry.Key.Location.Start.Character && character <= entry.Key.Location.End.Character {
		return handlers.GetRootSignatureHelp(0), nil
	}

	if entry.Values != nil && character >= entry.Values.Location.Start.Character && character <= entry.Values.Location.End.Character {
		value := handlers.GetValueAtCursor(character, entry)

		if value == nil {
			// For some reason, this does not really work,
			// When we return all, and then a user value is entered
			// and the `GetValueSignatureHelp` is called, still the old
			// signatures with all signature are shown
			// return handlers.GetAllValuesSignatureHelp(), nil

			return nil, nil
		}

		return handlers.GetValueSignatureHelp(*value, character), nil
	}

	return nil, nil
}
