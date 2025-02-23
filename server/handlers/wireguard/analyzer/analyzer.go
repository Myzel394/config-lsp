package analyzer

import (
	"config-lsp/common"
	"config-lsp/handlers/wireguard"
	"config-lsp/handlers/wireguard/indexes"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

type analyzerContext struct {
	document    *wireguard.WGDocument
	diagnostics []protocol.Diagnostic
}

func Analyze(
	d *wireguard.WGDocument,
) []protocol.Diagnostic {
	ctx := &analyzerContext{
		document:    d,
		diagnostics: make([]protocol.Diagnostic, 0),
	}

	analyzeStructureIsValid(ctx)

	if len(ctx.diagnostics) > 0 {
		return ctx.diagnostics
	}

	i, indexErrors := indexes.CreateIndexes(d.Config)

	if len(indexErrors) > 0 {
		return common.ErrsToDiagnostics(indexErrors)
	}

	d.Indexes = i

	analyzeInterfaceSection(ctx)
	analyzeDNSPropertyContainsFallback(ctx)
	analyzeKeepAlivePropertyIsSet(ctx)
	analyzeSymmetricPropertiesSet(ctx)
	analyzeDuplicateAllowedIPs(ctx)

	return ctx.diagnostics
}
