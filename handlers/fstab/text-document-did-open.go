package fstab

import (
	"config-lsp/common"
	"config-lsp/utils"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentDidOpen(
	context *glsp.Context,
	params *protocol.DidOpenTextDocumentParams,
) error {
	common.ClearDiagnostics(context, params.TextDocument.URI)

	parser := FstabParser{}
	documentParserMap[params.TextDocument.URI] = &parser

	errors := parser.ParseFromContent(params.TextDocument.Text)
	diagnostics := utils.Map(
		errors,
		func(err common.ParseError) protocol.Diagnostic {
			return err.ToDiagnostic()
		},
	)

	common.SendDiagnostics(context, params.TextDocument.URI, diagnostics)

	return nil
}
