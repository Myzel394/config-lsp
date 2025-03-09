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
		normalizedHeaderName := fields.CreateNormalizedName(section.Header.Name)
		// Whether to check if the property is allowed in the section
		checkAllowedProperty := true

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
			// Do not check as the section is unknown
			checkAllowedProperty = false
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
			existingProperties := make(map[fields.NormalizedName]*ast.WGProperty)

			it := section.Properties.Iterator()
			for it.Next() {
				property := it.Value().(*ast.WGProperty)
				normalizedPropertyName := fields.CreateNormalizedName(property.Key.Name)

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
					options := fields.OptionsHeaderMap[normalizedHeaderName]

					if !utils.KeyExists(options, normalizedPropertyName) {
						ctx.diagnostics = append(ctx.diagnostics, protocol.Diagnostic{
							Message:  fmt.Sprintf("Unknown property '%s'", property.Key.Name),
							Range:    property.Key.ToLSPRange(),
							Severity: &common.SeverityError,
						})
					} else if existingProperty, found := existingProperties[normalizedPropertyName]; found {
						ctx.diagnostics = append(ctx.diagnostics, protocol.Diagnostic{
							Message:  fmt.Sprintf("Property '%s' has already been defined on line %d", property.Key.Name, existingProperty.Start.Line+1),
							Severity: &common.SeverityError,
							Range:    existingProperty.ToLSPRange(),
						})
					}

					existingProperties[normalizedPropertyName] = property
				}
			}
		}
	}
}
