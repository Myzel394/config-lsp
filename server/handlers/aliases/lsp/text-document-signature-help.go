package lsp

import (
	"config-lsp/common"
	"config-lsp/handlers/aliases"
	"config-lsp/handlers/aliases/ast"
	"config-lsp/handlers/aliases/handlers"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentSignatureHelp(context *glsp.Context, params *protocol.SignatureHelpParams) (*protocol.SignatureHelp, error) {
	document := aliases.DocumentParserMap[params.TextDocument.URI]

	line := params.Position.Line
	cursor := common.LSPCharacterAsCursorPosition(params.Position.Character)

	if _, found := document.Parser.CommentLines[line]; found {
		// Comment
		return nil, nil
	}

	rawEntry, found := document.Parser.Aliases.Get(line)

	if !found {
		return handlers.GetRootSignatureHelp(0), nil
	}

	entry := rawEntry.(*ast.AliasEntry)

	if entry.Key != nil && entry.Key.Location.ContainsPosition(cursor) {
		return handlers.GetRootSignatureHelp(0), nil
	}

	if entry.Values != nil && entry.Values.Location.ContainsPosition(cursor) {
		value := handlers.GetValueAtPosition(cursor, entry)

		if value == nil || value.GetAliasValue().Value == "" {
			// For some reason, this does not really work,
			// When we return all, and then a user value is entered
			// and the `GetValueSignatureHelp` is called, still the old
			// signatures with all signature are shown
			return handlers.GetAllValuesSignatureHelp(), nil
		}

		return handlers.GetValueSignatureHelp(cursor, value), nil
	}

	return nil, nil
}
