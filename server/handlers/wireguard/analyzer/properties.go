package analyzer

import (
	"config-lsp/common"
	docvalues "config-lsp/doc-values"
	"config-lsp/handlers/wireguard/ast"
	"config-lsp/handlers/wireguard/diagnostics"
	"config-lsp/handlers/wireguard/fields"
	"config-lsp/handlers/wireguard/indexes"
	"fmt"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func analyzeProperties(
	ctx *analyzerContext,
) {
	for _, section := range ctx.document.Config.Sections {
		normalizedHeaderName := fields.CreateNormalizedName(section.Header.Name)

		// Whether to check if the property is allowed in the section
		checkAllowedProperty := true
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
				checkAllowedProperty = false
			}

			if checkAllowedProperty {
				availableOptions := fields.OptionsHeaderMap[normalizedHeaderName]

				// Duplicate check
				if existingProperty, found := existingProperties[normalizedPropertyName]; found {
					ctx.diagnostics = append(ctx.diagnostics, protocol.Diagnostic{
						Message:  fmt.Sprintf("Property '%s' has already been defined on line %d", property.Key.Name, existingProperty.Start.Line+1),
						Severity: &common.SeverityError,
						Range:    existingProperty.ToLSPRange(),
					})
					// Check if value is valid
				} else if option, found := availableOptions[normalizedPropertyName]; found {
					invalidValues := option.DeprecatedCheckIsValid(property.Value.Value)

					for _, invalidValue := range invalidValues {
						err := docvalues.LSPErrorFromInvalidValue(property.Start.Line, *invalidValue).ShiftCharacter(property.Value.Start.Character)

						ctx.diagnostics = append(ctx.diagnostics, protocol.Diagnostic{
							Range:    err.Range.ToLSPRange(),
							Message:  err.Err.Error(),
							Severity: &common.SeverityError,
						})
					}
					// Unknown property
				} else {
					ctx.diagnostics = append(ctx.diagnostics,
						diagnostics.GenerateUnknownOption(
							property.ToLSPRange(),
							property.Key.Name,
						),
					)

					ctx.document.Indexes.UnknownProperties[property.Key.Start.Line] = indexes.WGIndexPropertyInfo{
						Section:  section,
						Property: property,
					}
				}

				existingProperties[normalizedPropertyName] = property
			}
		}
	}
}
