package analyzer

import (
	"config-lsp/common"
	"config-lsp/handlers/hosts"
	"config-lsp/utils"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func Analyze(document *hosts.HostsDocument) []protocol.Diagnostic {
	errors := analyzeEntriesSetCorrectly(*document.Parser)

	if len(errors) > 0 {
		return utils.Map(
			errors,
			func(err common.LSPError) protocol.Diagnostic {
				return err.ToDiagnostic()
			},
		)
	}

	errors = analyzeEntriesAreValid(*document.Parser)

	if len(errors) > 0 {
		return utils.Map(
			errors,
			func(err common.LSPError) protocol.Diagnostic {
				return err.ToDiagnostic()
			},
		)
	}

	errors = append(errors, analyzeDoubleIPs(document)...)
	errors = append(errors, analyzeDoubleHostNames(document)...)

	return utils.Map(
		errors,
		func(err common.LSPError) protocol.Diagnostic {
			return err.ToDiagnostic()
		},
	)
}
