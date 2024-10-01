package analyzer

import (
	"config-lsp/common"
	docvalues "config-lsp/doc-values"
	"config-lsp/handlers/ssh_config/fields"
	"config-lsp/utils"
	"fmt"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func analyzeValuesAreValid(
	ctx *analyzerContext,
) {
	for _, info := range ctx.document.Config.GetAllOptions() {
		option := info.Option
		block := info.Block

		docOption, found := fields.Options[option.Key.Key]

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

			continue
		}

		errs := docOption.DeprecatedCheckIsValid(option.OptionValue.Value.Value)
		ctx.diagnostics = append(
			ctx.diagnostics,
			utils.Map(
				errs,
				func(err *docvalues.InvalidValue) protocol.Diagnostic {
					return protocol.Diagnostic{
						Range:    option.OptionValue.ToLSPRange(),
						Message:  err.Err.Error(),
						Severity: &common.SeverityError,
					}
				},
			)...,
		)
	}
}
