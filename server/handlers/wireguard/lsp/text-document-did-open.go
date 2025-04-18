package lsp

import (
	"config-lsp/common"
	"config-lsp/handlers/wireguard"
	"config-lsp/handlers/wireguard/analyzer"
	"config-lsp/handlers/wireguard/ast"
	"config-lsp/handlers/wireguard/indexes"
	"config-lsp/utils"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentDidOpen(
	context *glsp.Context,
	params *protocol.DidOpenTextDocumentParams,
) error {
	common.ClearDiagnostics(context, params.TextDocument.URI)

	document := &wireguard.WGDocument{
		Config:  ast.NewWGConfig(),
		Indexes: &indexes.WGIndexes{},
	}
	wireguard.DocumentParserMap[params.TextDocument.URI] = document

	errors := document.Config.Parse(params.TextDocument.Text)

	diagnostics := utils.Map(
		errors,
		func(err common.LSPError) protocol.Diagnostic {
			return err.ToDiagnostic()
		},
	)
	diagnostics = append(diagnostics, analyzer.Analyze(document)...)

	if len(diagnostics) > 0 {
		common.SendDiagnostics(context, params.TextDocument.URI, diagnostics)
	}

	return nil
}
