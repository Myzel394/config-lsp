package lsp

import (
	"config-lsp/common"
	"config-lsp/handlers/fstab/analyzer"
	"config-lsp/handlers/fstab/ast"
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

	config := ast.NewFstabConfig()
	d := &shared.FstabDocument{
		Config: config,
	}
	shared.DocumentParserMap[params.TextDocument.URI] = d

	content := params.TextDocument.Text

	diagnostics := make([]protocol.Diagnostic, 0)
	errors := d.Config.Parse(content)

	if len(errors) > 0 {
		diagnostics = append(diagnostics, utils.Map(
			errors,
			func(err common.LSPError) protocol.Diagnostic {
				return err.ToDiagnostic()
			},
		)...)
	} else {
		diagnostics = append(diagnostics, analyzer.Analyze(d)...)
	}

	if len(diagnostics) > 0 {
		common.SendDiagnostics(context, params.TextDocument.URI, diagnostics)
	}

	return nil
}
