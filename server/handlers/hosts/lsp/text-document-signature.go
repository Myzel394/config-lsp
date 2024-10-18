package lsp

import (
	"config-lsp/common"
	"config-lsp/handlers/hosts"
	"config-lsp/handlers/hosts/ast"
	"config-lsp/handlers/hosts/handlers"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentSignatureHelp(context *glsp.Context, params *protocol.SignatureHelpParams) (*protocol.SignatureHelp, error) {
	document := hosts.DocumentParserMap[params.TextDocument.URI]

	line := uint32(params.Position.Line)
	cursor := common.LSPCharacterAsCursorPosition(params.Position.Character)

	if _, found := document.Parser.CommentLines[line]; found {
		// Comment
		return nil, nil
	}

	entry, found := document.Parser.Tree.Entries.Get(line)

	if !found {
		return handlers.GetEntrySignatureHelp(nil, cursor), nil
	} else {
		return handlers.GetEntrySignatureHelp(entry.(*ast.HostsEntry), cursor), nil
	}
}
