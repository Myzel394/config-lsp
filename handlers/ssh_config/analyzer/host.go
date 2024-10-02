package analyzer

import (
	"config-lsp/common"
	hostparser "config-lsp/handlers/ssh_config/host-parser"
	"fmt"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func analyzeHostBlock(
	ctx *analyzerContext,
) {
	hosts := make(map[string]*hostparser.HostValue, 0)

	for _, block := range ctx.document.GetAllHostBlocks() {
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
