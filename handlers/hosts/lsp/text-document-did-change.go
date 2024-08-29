package lsp

import (
	"config-lsp/common"
	"config-lsp/handlers/hosts"
	"config-lsp/handlers/hosts/analyzer"
	"config-lsp/utils"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentDidChange(
	context *glsp.Context,
	params *protocol.DidChangeTextDocumentParams,
) error {
	content := params.ContentChanges[0].(protocol.TextDocumentContentChangeEventWhole).Text
	common.ClearDiagnostics(context, params.TextDocument.URI)

	document := hosts.DocumentParserMap[params.TextDocument.URI]
	document.Parser.Clear()

	diagnostics := make([]protocol.Diagnostic, 0)
	errors := document.Parser.Parse(content)

	if len(errors) > 0 {
		diagnostics = append(diagnostics, utils.Map(
			errors,
			func(err common.LSPError) protocol.Diagnostic {
				return err.ToDiagnostic()
			},
		)...)
	}

	diagnostics = append(diagnostics, analyzer.Analyze(document)...)

	if len(diagnostics) > 0 {
		common.SendDiagnostics(context, params.TextDocument.URI, diagnostics)
	}

	return nil
}
