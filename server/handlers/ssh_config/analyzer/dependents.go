package analyzer

import (
	"config-lsp/common"
	"config-lsp/handlers/ssh_config/ast"
	"config-lsp/handlers/ssh_config/fields"
	"fmt"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func analyzeDependents(
	ctx *analyzerContext,
) {
	for _, option := range ctx.document.Config.GetAllOptions() {
		checkIsDependent(ctx, option.Option.Key, option.Block)
	}
}

func checkIsDependent(
	ctx *analyzerContext,
	key *ast.SSHKey,
	block ast.SSHBlock,
) {
	dependentOptions, found := fields.DependentFields[key.Key]

	if !found {
		return
	}

	for _, dependentOption := range dependentOptions {
		if opts, found := ctx.document.Indexes.AllOptionsPerName[dependentOption]; found {
			_, existsInBlock := opts[block]
			_, existsInGlobal := opts[nil]

			if existsInBlock || existsInGlobal {
				continue
			}
		}

		ctx.diagnostics = append(ctx.diagnostics, protocol.Diagnostic{
			Range:    key.LocationRange.ToLSPRange(),
			Message:  fmt.Sprintf("Option '%s' requires option '%s' to be present", key.Key, dependentOption),
			Severity: &common.SeverityError,
		})
	}
}
