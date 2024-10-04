package lsp

import (
	"config-lsp/common"
	"config-lsp/handlers/fstab/parser"
	"config-lsp/handlers/fstab/shared"
	"config-lsp/utils"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentDidOpen(
	context *glsp.Context,
	params *protocol.DidOpenTextDocumentParams,
) error {
	common.ClearDiagnostics(context, params.TextDocument.URI)

	p := parser.FstabParser{}
	p.Clear()
	shared.DocumentParserMap[params.TextDocument.URI] = &p

	content := params.TextDocument.Text

	diagnostics := make([]protocol.Diagnostic, 0)
	errors := p.ParseFromContent(content)

	if len(errors) > 0 {
		diagnostics = append(diagnostics, utils.Map(
			errors,
			func(err common.ParseError) protocol.Diagnostic {
				return err.ToDiagnostic()
			},
		)...)
	} else {
		diagnostics = append(diagnostics, p.AnalyzeValues()...)
	}

	if len(diagnostics) > 0 {
		common.SendDiagnostics(context, params.TextDocument.URI, diagnostics)
	}

	return nil
}
