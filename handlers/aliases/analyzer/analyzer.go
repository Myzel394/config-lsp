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
	// Double keys must be checked first so
	// that the index is populated for the
	// other checks
	errors := analyzeDoubleKeys(d)
	errors = append(errors, analyzeValuesAreValid(d)...)

	if len(errors) > 0 {
		return utils.Map(
			errors,
			func(err common.LSPError) protocol.Diagnostic {
				return err.ToDiagnostic()
			},
		)
	}

	errors = append(errors, analyzeContainsRequiredKeys(*d)...)
	errors = append(errors, analyzeContainsNoDoubleValues(*d.Parser)...)

	return utils.Map(
		errors,
		func(err common.LSPError) protocol.Diagnostic {
			return err.ToDiagnostic()
		},
	)
}
