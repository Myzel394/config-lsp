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

	analyzeIncludeValues(ctx)

	if len(ctx.diagnostics) == 0 {
		for _, include := range d.Indexes.Includes {
			for _, value := range include.Values {
				for _, path := range value.Paths {
					_, err := parseFile(string(path))

					if err != nil {
						ctx.diagnostics = append(ctx.diagnostics, protocol.Diagnostic{
							Range:   value.LocationRange.ToLSPRange(),
							Message: err.Error(),
						})
					}
				}
			}
		}
	}

	analyzeValuesAreValid(ctx)
	analyzeTokens(ctx)
	analyzeIgnoreUnknownHasNoUnnecessary(ctx)
	analyzeDependents(ctx)
	analyzeBlocks(ctx)
	analyzeMatchBlocks(ctx)
	analyzeHostBlock(ctx)
	analyzeBlocks(ctx)
	analyzeTagOptions(ctx)
	analyzeTagImports(ctx)

	return ctx.diagnostics
}
