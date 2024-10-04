package analyzer

import (
	"config-lsp/common"
	sshconfig "config-lsp/handlers/ssh_config"
	"config-lsp/handlers/ssh_config/indexes"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

type analyzerContext struct {
	document    *sshconfig.SSHDocument
	diagnostics []protocol.Diagnostic
}

func Analyze(
	d *sshconfig.SSHDocument,
) []protocol.Diagnostic {
	ctx := &analyzerContext{
		document:    d,
		diagnostics: make([]protocol.Diagnostic, 0),
	}

	analyzeStructureIsValid(ctx)

	if len(ctx.diagnostics) > 0 {
		return ctx.diagnostics
	}

	i, indexErrors := indexes.CreateIndexes(*d.Config)

	if len(indexErrors) > 0 {
		return common.ErrsToDiagnostics(indexErrors)
	}

	d.Indexes = i

	analyzeValuesAreValid(ctx)
	analyzeIgnoreUnknownHasNoUnnecessary(ctx)
	analyzeDependents(ctx)
	analyzeBlocks(ctx)
	analyzeMatchBlocks(ctx)
	analyzeHostBlock(ctx)
	analyzeBlocks(ctx)

	return ctx.diagnostics
}
