package analyzer

import (
	"config-lsp/handlers/fstab/shared"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

type analyzerContext struct {
	document    *shared.FstabDocument
	diagnostics []protocol.Diagnostic
}

func Analyze(
	document *shared.FstabDocument,
) []protocol.Diagnostic {
	ctx := &analyzerContext{
		document: document,
	}

	analyzeFieldAreFilled(ctx)

	if len(ctx.diagnostics) > 0 {
		return ctx.diagnostics
	}

	analyzeValuesAreValid(ctx)

	if len(ctx.diagnostics) > 0 {
		return ctx.diagnostics
	}

	analyzePassFields(ctx)

	return ctx.diagnostics
}
