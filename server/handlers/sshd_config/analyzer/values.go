package analyzer

import (
	"config-lsp/handlers/sshd_config/diagnostics"
	"config-lsp/handlers/sshd_config/fields"
)

func analyzeValuesAreValid(
	ctx *analyzerContext,
) {
	// Check if there are unknown options
	for _, info := range ctx.document.Config.GetAllOptions() {
		normalizedName := fields.CreateNormalizedName(info.Option.Key.Value.Value)

		var isUnknown bool = true

		// Check if the option is unknown
		if info.MatchBlock == nil {
			// All options are allowed
			if _, found := fields.Options[normalizedName]; found {
				isUnknown = false
			}
		} else {
			// Only `MatchAllowedOptions` are allowed
			if _, found := fields.MatchAllowedOptions[normalizedName]; found {
				isUnknown = false
			}
		}

		if isUnknown {
			ctx.diagnostics = append(ctx.diagnostics, diagnostics.GenerateUnknownOption(
				info.Option.Key.ToLSPRange(),
				info.Option.Key.Value.Value,
			))

			ctx.document.Indexes.UnknownOptions[info.Option.Start.Line] = info
		}
	}
}
