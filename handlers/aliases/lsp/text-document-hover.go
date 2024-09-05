package lsp

import (
	"config-lsp/handlers/aliases"
	"config-lsp/handlers/aliases/ast"
	"config-lsp/handlers/aliases/handlers"
	"strings"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentHover(
	context *glsp.Context,
	params *protocol.HoverParams,
) (*protocol.Hover, error) {
	document := aliases.DocumentParserMap[params.TextDocument.URI]

	line := params.Position.Line
	character := params.Position.Character

	if _, found := document.Parser.CommentLines[line]; found {
		// Comment
		return nil, nil
	}

	rawEntry, found := document.Parser.Aliases.Get(line)

	if !found {
		return nil, nil
	}

	entry := rawEntry.(*ast.AliasEntry)

	if entry.Key != nil && character >= entry.Key.Location.Start.Character && character <= entry.Key.Location.End.Character {
		text := handlers.GetAliasHoverInfo(*document.Indexes, *entry)

		return &protocol.Hover{
			Contents: text,
		}, nil
	}

	if entry.Values != nil && character >= entry.Values.Location.Start.Character && character <= entry.Values.Location.End.Character {
		value := handlers.GetValueAtCursor(character, entry)

		if value == nil {
			return nil, nil
		}

		contents := []string{}
		contents = append(contents, handlers.GetAliasValueTypeInfo(*value)...)
		contents = append(contents, "")
		contents = append(contents, "#### Value")
		contents = append(contents, handlers.GetAliasValueHoverInfo(*document.Indexes, *value))

		text := strings.Join(contents, "\n")

		return &protocol.Hover{
			Contents: text,
		}, nil
	}

	return nil, nil
}
