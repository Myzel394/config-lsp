package analyzer

import (
	"config-lsp/common"
	"config-lsp/handlers/bitcoin_conf/fields"
	"config-lsp/parsers/ini"
	"config-lsp/utils"
	"fmt"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

// Check that the chain properties are set exclusively
func analyzeChainProperties(
	ctx *analyzerContext,
) {
	var firstSetAt *ini.Property

	for _, section := range ctx.document.Config.Sections {
		it := section.Properties.Iterator()
		for it.Next() {
			property := it.Value().(*ini.Property)

			name := property.Key.Name

			if utils.KeyExists(fields.ChainOptions, name) {
				if firstSetAt == nil {
					firstSetAt = property
				} else {
					ctx.diagnostics = append(ctx.diagnostics, protocol.Diagnostic{
						Message:  fmt.Sprintf("Chain property has already been set at line %d", firstSetAt.Start.Line+1),
						Range:    property.ToLSPRange(),
						Severity: &common.SeverityError,
					})
				}
			}
		}
	}
}
