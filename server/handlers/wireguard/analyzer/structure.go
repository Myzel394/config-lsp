package analyzer

import (
	"config-lsp/common"
	"config-lsp/handlers/wireguard/fields"
	"config-lsp/parsers/ini"
	"config-lsp/utils"
	"fmt"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func analyzeStructureIsValid(ctx *analyzerContext) {
	for _, section := range ctx.document.Config.Sections {
		normalizedHeaderName := fields.CreateNormalizedName(section.Header.Name)

		if section.Header.Name == "" {
			ctx.diagnostics = append(ctx.diagnostics, protocol.Diagnostic{
				Message:  "This section is missing a name",
				Range:    section.Header.ToLSPRange(),
				Severity: &common.SeverityError,
			})
		} else if !utils.KeyExists(fields.OptionsHeaderMap, normalizedHeaderName) {
			ctx.diagnostics = append(ctx.diagnostics, protocol.Diagnostic{
				Message:  fmt.Sprintf("Unknown section '%s'. It must be one of: [Interface], [Peer]", section.Header.Name),
				Range:    section.Header.ToLSPRange(),
				Severity: &common.SeverityError,
			})
		}

		if section.Properties.Size() == 0 {
			ctx.diagnostics = append(ctx.diagnostics, protocol.Diagnostic{
				Message:  "This section is empty",
				Range:    section.Header.ToLSPRange(),
				Severity: &common.SeverityInformation,
				Tags: []protocol.DiagnosticTag{
					protocol.DiagnosticTagUnnecessary,
				},
			})
		} else {
			it := section.Properties.Iterator()

			for it.Next() {
				property := it.Value().(*ini.Property)

				if property.Key == nil {
					ctx.diagnostics = append(ctx.diagnostics, protocol.Diagnostic{
						Message:  "This property is missing a key",
						Range:    property.ToLSPRange(),
						Severity: &common.SeverityError,
					})
				}
			}
		}
	}
}
