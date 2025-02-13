package analyzer

import (
	"config-lsp/handlers/ssh_config/diagnostics"
	"config-lsp/handlers/ssh_config/fields"
)

func analyzeValuesAreValid(
	ctx *analyzerContext,
) {
	// Check if there are unknown options
	for _, info := range ctx.document.Config.GetAllOptions() {
		option := info.Option
		block := info.Block

		_, found := fields.Options[option.Key.Key]

		if !found {
			if ctx.document.Indexes.CanOptionBeIgnored(option, block) {
				// Skip
				continue
			}

			ctx.diagnostics = append(
				ctx.diagnostics,
				diagnostics.GenerateUnknownOption(
					option.Key.ToLSPRange(),
					option.Key.Value.Value,
				),
			)
			ctx.document.Indexes.UnknownOptions[info.Option.Start.Line] = info

			continue
		}
	}
}
