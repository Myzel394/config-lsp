package analyzer

import (
	"config-lsp/common"
	"config-lsp/handlers/ssh_config/fields"
	"fmt"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

var ignoreUnknownOption = fields.CreateNormalizedName("IgnoreUnknown")

func analyzeIgnoreUnknownHasNoUnnecessary(
	ctx *analyzerContext,
) {
	for _, block := range ctx.document.GetAllBlocks() {
		ignoreUnknown, found := ctx.document.Indexes.IgnoredOptions[block]

		if !found {
			// No `IgnoreUnknown` option specified
			continue
		}

		for optionName, ignoreInfo := range ignoreUnknown.IgnoredOptions {
			info := ctx.document.FindOptionByNameAndBlock(optionName, block)

			if info == nil {
				ctx.diagnostics = append(ctx.diagnostics, protocol.Diagnostic{
					Range:   ignoreInfo.ToLSPRange(),
					Message: fmt.Sprintf("Option %s is not present", optionName),
					Tags: []protocol.DiagnosticTag{
						protocol.DiagnosticTagUnnecessary,
					},
					Severity: &common.SeverityHint,
				})
			}
		}
	}
}
