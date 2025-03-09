package lsp

import (
	"config-lsp/common"
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
	index := common.LSPCharacterAsIndexPosition(params.Position.Character)

	if _, found := document.Parser.CommentLines[line]; found {
		// Comment
		return nil, nil
	}

	rawEntry, found := document.Parser.Aliases.Get(line)

	if !found {
		return nil, nil
	}

	entry := rawEntry.(*ast.AliasEntry)

	if entry.Key != nil && entry.Key.Location.ContainsPosition(index) {
		text := handlers.GetAliasHoverInfo(*document.Indexes, *entry)

		return &protocol.Hover{
			Contents: text,
		}, nil
	}

	if entry.Values != nil && entry.Values.Location.ContainsPosition(index) {
		value := handlers.GetValueAtPosition(index, entry)

		if value == nil {
			return nil, nil
		}

		contents := []string{}
		contents = append(contents, handlers.GetAliasValueTypeInfo(value)...)
		contents = append(contents, "")
		contents = append(contents, "#### Name")
		contents = append(contents, handlers.GetAliasValueHoverInfo(*document.Indexes, value))

		text := strings.Join(contents, "\n")

		return &protocol.Hover{
			Contents: text,
		}, nil
	}

	return nil, nil
}
