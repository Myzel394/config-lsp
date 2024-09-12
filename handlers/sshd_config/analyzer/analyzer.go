package analyzer

import (
	"config-lsp/common"
	"config-lsp/handlers/sshd_config"
	"config-lsp/utils"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func Analyze(
	d *sshdconfig.SSHDocument,
) []protocol.Diagnostic {
	errors := analyzeOptionsAreValid(d)

	if len(errors) > 0 {
		return utils.Map(
			errors,
			func(err common.LSPError) protocol.Diagnostic {
				return err.ToDiagnostic()
			},
		)
	}

	return nil
}
