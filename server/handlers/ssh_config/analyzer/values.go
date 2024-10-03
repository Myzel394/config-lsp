package analyzer

import (
	"config-lsp/common"
	"config-lsp/handlers/ssh_config/fields"
	"fmt"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func analyzeValuesAreValid(
	ctx *analyzerContext,
) {
	for _, info := range ctx.document.Config.GetAllOptions() {
		option := info.Option
		block := info.Block

		_, found := fields.Options[option.Key.Key]

		if !found {
			if ctx.document.Indexes.CanOptionBeIgnored(option, block) {
				// Skip
				continue
			}

			ctx.diagnostics = append(ctx.diagnostics,
				protocol.Diagnostic{
					Range:    option.Key.ToLSPRange(),
					Message:  fmt.Sprintf("Unknown option: %s", option.Key.Value.Value),
					Severity: &common.SeverityError,
				},
			)
			ctx.document.Indexes.UnknownOptions[info.Option.Start.Line] = info

			continue
		}
	}
}
