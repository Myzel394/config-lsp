package analyzer

import (
	"config-lsp/common"
	sshconfig "config-lsp/handlers/ssh_config"
	"config-lsp/handlers/ssh_config/indexes"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func Analyze(
	d *sshconfig.SSHDocument,
) []protocol.Diagnostic {
	errors := analyzeStructureIsValid(d)

	if len(errors) > 0 {
		return common.ErrsToDiagnostics(errors)
	}

	i, indexErrors := indexes.CreateIndexes(*d.Config)

	d.Indexes = i

	errors = append(errors, indexErrors...)

	if len(errors) > 0 {
		return common.ErrsToDiagnostics(errors)
	}

	errors = append(errors, analyzeDependents(d)...)

	return common.ErrsToDiagnostics(errors)
}
