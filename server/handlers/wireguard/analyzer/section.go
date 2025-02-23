package analyzer

import (
	"config-lsp/common"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func analyzeInterfaceSection(ctx *analyzerContext) {
	sections := ctx.document.Indexes.SectionsByName["Interface"]
	if len(sections) > 1 {
		ctx.diagnostics = append(ctx.diagnostics, protocol.Diagnostic{
			Message:  "Only one [Interface] section is allowed",
			Severity: &common.SeverityError,
			Range:    sections[1].Header.ToLSPRange(),
		})
	}
}
