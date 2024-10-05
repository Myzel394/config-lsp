package analyzer

import (
	"config-lsp/common"
	"fmt"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func analyzeTagImports(
	ctx *analyzerContext,
) {
	for name, info := range ctx.document.Indexes.Tags {
		if _, found := ctx.document.Indexes.TagImports[name]; !found {
			var diagnosticRange protocol.Range

			if len(info.Block.MatchValue.Entries) == 1 {
				diagnosticRange = info.Block.MatchOption.ToLSPRange()
			} else {
				diagnosticRange = info.EntryValue.ToLSPRange()
			}

			ctx.diagnostics = append(ctx.diagnostics, protocol.Diagnostic{
				Range:    diagnosticRange,
				Message:  fmt.Sprintf("Tag %s is not used", name),
				Severity: &common.SeverityWarning,
				Tags: []protocol.DiagnosticTag{
					protocol.DiagnosticTagUnnecessary,
				},
			})
		}
	}
}
