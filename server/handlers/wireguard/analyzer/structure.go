package analyzer

import (
	"config-lsp/common"
	"config-lsp/handlers/wireguard/ast"
	"config-lsp/handlers/wireguard/fields"
	"config-lsp/utils"
	"fmt"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func analyzeStructureIsValid(ctx *analyzerContext) {
	for _, section := range ctx.document.Config.Sections {
		// Whether to check if the property is allowed in the section
		checkAllowedProperty := true

		if section.Header.Name == "" {
			ctx.diagnostics = append(ctx.diagnostics, protocol.Diagnostic{
				Message:  "This section is missing a name",
				Range:    section.Header.ToLSPRange(),
				Severity: &common.SeverityError,
			})
		} else if !utils.KeyExists(fields.OptionsHeaderMap, section.Header.Name) {
			ctx.diagnostics = append(ctx.diagnostics, protocol.Diagnostic{
				Message:  fmt.Sprintf("Unknown section '%s'. It must be one of: [Interface], [Peer]", section.Header.Name),
				Range:    section.Header.ToLSPRange(),
				Severity: &common.SeverityError,
			})
			// Do not check as the section is unknown
			checkAllowedProperty = false
		}

		if len(section.Properties) == 0 {
			ctx.diagnostics = append(ctx.diagnostics, protocol.Diagnostic{
				Message:  "This section is empty",
				Range:    section.Header.ToLSPRange(),
				Severity: &common.SeverityInformation,
				Tags: []protocol.DiagnosticTag{
					protocol.DiagnosticTagUnnecessary,
				},
			})
		} else {
			existingProperties := make(map[string]*ast.WGProperty)

			for _, property := range section.Properties {
				if property.Key.Name == "" {
					ctx.diagnostics = append(ctx.diagnostics, protocol.Diagnostic{
						Message:  "This property is missing a name",
						Range:    property.Key.ToLSPRange(),
						Severity: &common.SeverityError,
					})
				}

				if property.Value == nil || property.Value.Value == "" {
					ctx.diagnostics = append(ctx.diagnostics, protocol.Diagnostic{
						Message:  "This property is missing a value",
						Range:    property.ToLSPRange(),
						Severity: &common.SeverityError,
					})
				}

				if checkAllowedProperty {
					options := fields.OptionsHeaderMap[section.Header.Name]

					if !utils.KeyExists(options, property.Key.Name) {
						ctx.diagnostics = append(ctx.diagnostics, protocol.Diagnostic{
							Message:  fmt.Sprintf("Unknown property '%s'", property.Key.Name),
							Range:    property.Key.ToLSPRange(),
							Severity: &common.SeverityError,
						})
					} else if existingProperty, found := existingProperties[property.Key.Name]; found {
						ctx.diagnostics = append(ctx.diagnostics, protocol.Diagnostic{
							Message:  fmt.Sprintf("Property '%s' has already been defined on line %d", property.Key.Name, existingProperty.Start.Line+1),
							Severity: &common.SeverityError,
							Range:    existingProperty.ToLSPRange(),
						})
					}
				}
			}
		}
	}
}
