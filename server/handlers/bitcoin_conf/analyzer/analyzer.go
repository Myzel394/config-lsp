package analyzer

import (
	"config-lsp/common"
	bitcoinconf "config-lsp/handlers/bitcoin_conf"
	"config-lsp/handlers/bitcoin_conf/indexes"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

type analyzerContext struct {
	document    *bitcoinconf.BTCDocument
	diagnostics []protocol.Diagnostic
}

func Analyze(
	d *bitcoinconf.BTCDocument,
) []protocol.Diagnostic {
	ctx := &analyzerContext{
		document:    d,
		diagnostics: make([]protocol.Diagnostic, 0),
	}

	analyzeSectionsNamesAreValid(ctx)

	if len(ctx.diagnostics) > 0 {
		return ctx.diagnostics
	}

	i, indexErrors := indexes.CreateIndexes(d.Config)

	if len(indexErrors) > 0 {
		return common.ErrsToDiagnostics(indexErrors)
	}

	d.Indexes = i

	// Analyze the structure of the Bitcoin configuration
	analyzeProperties(ctx)

	if len(ctx.diagnostics) > 0 {
		return ctx.diagnostics
	}

	return ctx.diagnostics
}
