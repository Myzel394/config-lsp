package analyzer

import (
	"config-lsp/common"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func analyzeBlocks(
	ctx *analyzerContext,
) {
	for _, block := range ctx.document.GetAllBlocks() {
		if block.GetOptions().Size() == 0 {
			ctx.diagnostics = append(ctx.diagnostics, protocol.Diagnostic{
				Range:    block.GetEntryOption().LocationRange.ToLSPRange(),
				Message:  "This block is empty",
				Severity: &common.SeverityHint,
				Tags: []protocol.DiagnosticTag{
					protocol.DiagnosticTagUnnecessary,
				},
			})
		}
	}
}
