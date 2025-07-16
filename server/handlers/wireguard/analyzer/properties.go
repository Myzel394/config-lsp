package analyzer

import (
	"config-lsp/common"
	docvalues "config-lsp/doc-values"
	"config-lsp/handlers/wireguard/diagnostics"
	"config-lsp/handlers/wireguard/fields"
	"config-lsp/handlers/wireguard/indexes"
	"config-lsp/parsers/ini"
	"config-lsp/utils"
	"fmt"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func analyzeProperties(
	ctx *analyzerContext,
) {
	for _, section := range ctx.document.Config.Sections {
		normalizedHeaderName := fields.CreateNormalizedName(section.Header.Name)

		existingProperties := make(map[fields.NormalizedName]*ini.Property)

		it := section.Properties.Iterator()
		for it.Next() {
			property := it.Value().(*ini.Property)

			if property.Key.Name == "" {
				ctx.diagnostics = append(ctx.diagnostics, protocol.Diagnostic{
					Message:  "This property is missing a name",
					Range:    property.Key.ToLSPRange(),
					Severity: &common.SeverityError,
				})
			}

			normalizedPropertyName := fields.CreateNormalizedName(property.Key.Name)
			availableOptions := fields.OptionsHeaderMap[normalizedHeaderName]

			// Unknown property
			if !utils.KeyExists(availableOptions, normalizedPropertyName) {
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

				// Duplicate check
			} else if existingProperty, found := existingProperties[normalizedPropertyName]; found {
				ctx.diagnostics = append(ctx.diagnostics, protocol.Diagnostic{
					Message:  fmt.Sprintf("Property '%s' has already been defined on line %d", property.Key.Name, existingProperty.Start.Line+1),
					Severity: &common.SeverityError,
					Range:    existingProperty.ToLSPRange(),
				})

				// Value missing
			} else if property.Value == nil || property.Value.Value == "" {
				ctx.diagnostics = append(ctx.diagnostics, protocol.Diagnostic{
					Message:  "This property is missing a value",
					Range:    property.ToLSPRange(),
					Severity: &common.SeverityError,
				})

				// Check if value is valid
			} else if option, found := availableOptions[normalizedPropertyName]; found {
				existingProperties[normalizedPropertyName] = property

				invalidValues := option.DeprecatedCheckIsValid(property.Value.Value)

				for _, invalidValue := range invalidValues {
					err := docvalues.LSPErrorFromInvalidValue(property.Start.Line, *invalidValue).ShiftCharacter(property.Value.Start.Character)

					ctx.diagnostics = append(ctx.diagnostics, protocol.Diagnostic{
						Range:    err.Range.ToLSPRange(),
						Message:  err.Err.Error(),
						Severity: &common.SeverityError,
					})
				}
			}
		}
	}
}
