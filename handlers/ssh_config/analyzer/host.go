package analyzer

import (
	"cmp"
	"config-lsp/common"
	"config-lsp/handlers/ssh_config/ast"
	hostparser "config-lsp/handlers/ssh_config/host-parser"
	"fmt"
	"slices"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func analyzeHostBlock(
	ctx *analyzerContext,
) {
	hosts := make(map[string]*hostparser.HostValue, 0)

	blocks := ctx.document.GetAllHostBlocks()
	slices.SortFunc(
		blocks,
		func(a, b *ast.SSHHostBlock) int {
			return cmp.Compare(a.Start.Line, b.Start.Line)
		},
	)

	for _, block := range blocks {
		if block == nil || block.HostValue == nil {
			continue
		}

		for _, host := range block.HostValue.Hosts {
			if _, found := hosts[host.Value.Value]; found {
				ctx.diagnostics = append(ctx.diagnostics, protocol.Diagnostic{
					Range:    host.ToLSPRange(),
					Message:  fmt.Sprintf("Host %s has already been defined on line %d", host.Value.Value, hosts[host.Value.Value].Start.Line+1),
					Severity: &common.SeverityError,
				})
			} else {
				hosts[host.Value.Value] = host
			}
		}
	}
}
