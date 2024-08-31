package analyzer

import (
	"config-lsp/common"
	"config-lsp/handlers/aliases"
	"config-lsp/utils"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func Analyze(
	d *aliases.AliasesDocument,
) []protocol.Diagnostic {
	errors := analyzeValuesAreValid(*d.Parser)

	if len(errors) > 0 {
		return utils.Map(
			errors,
			func(err common.LSPError) protocol.Diagnostic {
				return err.ToDiagnostic()
			},
		)
	}

	errors = append(errors, analyzeDoubleKeys(d)...)
	errors = append(errors, analyzeContainsRequiredKeys(*d)...)

	return utils.Map(
		errors,
		func(err common.LSPError) protocol.Diagnostic {
			return err.ToDiagnostic()
		},
	)
}
