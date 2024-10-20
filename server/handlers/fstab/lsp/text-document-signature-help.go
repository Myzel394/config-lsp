package lsp

import (
	"config-lsp/common"
	"config-lsp/handlers/fstab/ast"
	"config-lsp/handlers/fstab/handlers"
	fstab "config-lsp/handlers/fstab/shared"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentSignatureHelp(context *glsp.Context, params *protocol.SignatureHelpParams) (*protocol.SignatureHelp, error) {
	document := fstab.DocumentParserMap[params.TextDocument.URI]

	line := uint32(params.Position.Line)
	cursor := common.LSPCharacterAsCursorPosition(params.Position.Character)

	if _, found := document.Config.CommentLines[line]; found {
		// Comment
		return nil, nil
	}

	entry, found := document.Config.Entries.Get(line)

	if !found {
		return handlers.GetEntrySignatureHelp(nil, cursor), nil
	} else {
		return handlers.GetEntrySignatureHelp(entry.(*ast.FstabEntry), cursor), nil
	}
}
