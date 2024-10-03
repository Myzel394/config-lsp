package analyzer

import (
	"config-lsp/common"
	"config-lsp/handlers/sshd_config"
	"config-lsp/handlers/sshd_config/indexes"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

type analyzerContext struct {
	document    *sshdconfig.SSHDDocument
	diagnostics []protocol.Diagnostic
}

func Analyze(
	d *sshdconfig.SSHDDocument,
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

	d.Indexes = i

	if len(indexErrors) > 0 {
		return common.ErrsToDiagnostics(indexErrors)
	}

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

	analyzeMatchBlocks(ctx)

	return ctx.diagnostics
}
