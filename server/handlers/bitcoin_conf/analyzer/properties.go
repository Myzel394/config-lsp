package analyzer

import (
	"config-lsp/common"
	docvalues "config-lsp/doc-values"
	"config-lsp/handlers/bitcoin_conf/fields"
	"config-lsp/handlers/bitcoin_conf/indexes"
	"config-lsp/parsers/ini"
	"config-lsp/utils"
	"fmt"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func analyzeProperties(ctx *analyzerContext) {
	for _, section := range ctx.document.Config.Sections {
		existingProperties := make(map[string]*ini.Property)

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

			name := property.Key.Name

			if existingProperty, found := existingProperties[name]; found && !utils.KeyExists(fields.AllowedDuplicateOptions, name) {
				ctx.diagnostics = append(ctx.diagnostics, protocol.Diagnostic{
					Message:  fmt.Sprintf("Property '%s' has already been defined on line %d", property.Key.Name, existingProperty.Start.Line+1),
					Severity: &common.SeverityError,
					Range:    property.ToLSPRange(),
				})

				// Value missing
			} else if property.Value == nil || property.Value.Value == "" {
				ctx.diagnostics = append(ctx.diagnostics, protocol.Diagnostic{
					Message:  "This property is missing a value",
					Range:    property.ToLSPRange(),
					Severity: &common.SeverityError,
				})

				// Check if value is valid
			} else if option, found := fields.Options[name]; found {
				existingProperties[name] = property

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
				ctx.diagnostics = append(ctx.diagnostics, protocol.Diagnostic{
					Message:  fmt.Sprintf("Unknown property '%s'", property.Key.Name),
					Severity: &common.SeverityWarning,
					Range:    property.ToLSPRange(),
				})

				ctx.document.Indexes.UnknownProperties[property.Key.Start.Line] = indexes.BTCIndexPropertyInfo{
					Section:  section,
					Property: property,
				}
			}
		}
	}
}
