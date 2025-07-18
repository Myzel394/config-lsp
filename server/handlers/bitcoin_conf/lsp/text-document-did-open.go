package lsp

import (
	"config-lsp/common"
	bitcoinconf "config-lsp/handlers/bitcoin_conf"
	"config-lsp/handlers/bitcoin_conf/analyzer"
	"config-lsp/handlers/bitcoin_conf/ast"
	"config-lsp/handlers/bitcoin_conf/indexes"
	"config-lsp/utils"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentDidOpen(
	context *glsp.Context,
	params *protocol.DidOpenTextDocumentParams,
) error {
	common.ClearDiagnostics(context, params.TextDocument.URI)

	document := &bitcoinconf.BTCDocument{
		Config:  ast.NewBTCConfig(),
		Indexes: &indexes.BTCIndexes{},
	}
	bitcoinconf.DocumentParserMap[params.TextDocument.URI] = document

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
