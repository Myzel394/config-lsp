package lsp

import (
	"config-lsp/common"
	"config-lsp/handlers/hosts"
	"config-lsp/handlers/hosts/analyzer"
	"config-lsp/handlers/hosts/ast"
	"config-lsp/handlers/hosts/indexes"
	"config-lsp/utils"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentDidOpen(
	context *glsp.Context,
	params *protocol.DidOpenTextDocumentParams,
) error {
	common.ClearDiagnostics(context, params.TextDocument.URI)

	parser := ast.NewHostsParser()
	i := indexes.NewHostsIndexes()
	document := hosts.HostsDocument{
		Parser:  &parser,
		Indexes: &i,
	}
	hosts.DocumentParserMap[params.TextDocument.URI] = &document

	errors := parser.Parse(params.TextDocument.Text)

	diagnostics := utils.Map(
		errors,
		func(err common.LSPError) protocol.Diagnostic {
			return err.ToDiagnostic()
		},
	)

	diagnostics = append(
		diagnostics,
		analyzer.Analyze(&document)...,
	)

	if len(diagnostics) > 0 {
		common.SendDiagnostics(context, params.TextDocument.URI, diagnostics)
	}

	return nil
}
