package wireguard

import (
	docvalues "config-lsp/doc-values"
	"config-lsp/utils"
	"fmt"
	"slices"
	"strings"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func (p wireguardParser) analyze() []protocol.Diagnostic {
	validCheckErrors := p.checkIfValuesAreValid()

	if len(validCheckErrors) > 0 {
		return validCheckErrors
	}

	diagnostics := []protocol.Diagnostic{}
	diagnostics = append(diagnostics, p.checkForDuplicateProperties()...)
	diagnostics = append(diagnostics, p.analyzeDNSContainsFallback()...)

	return diagnostics
}

func (p wireguardParser) analyzeDNSContainsFallback() []protocol.Diagnostic {
	lineNumber, property := p.fetchPropertyByName("DNS")

	if property == nil {
		return []protocol.Diagnostic{}
	}

	dnsAmount := len(strings.Split(property.Value.Value, ","))

	if dnsAmount == 1 {
		severity := protocol.DiagnosticSeverityWarning
		return []protocol.Diagnostic{
			{
				Message:  "There is one DNS server specified. It is recommended to set up fallback DNS servers",
				Severity: &severity,
				Range: protocol.Range{
					Start: protocol.Position{
						Line:      *lineNumber,
						Character: property.Value.Location.Start,
					},
					End: protocol.Position{
						Line:      *lineNumber,
						Character: property.Value.Location.End,
					},
				},
			},
		}
	}

	return []protocol.Diagnostic{}
}

func (p wireguardParser) checkIfValuesAreValid() []protocol.Diagnostic {
	diagnostics := []protocol.Diagnostic{}

	for _, section := range p.Sections {
		sectionDiagnostics := section.analyzeSection()

		if len(sectionDiagnostics) > 0 {
			diagnostics = append(diagnostics, sectionDiagnostics...)
			continue
		}

		for lineNumber, property := range section.Properties {
			diagnostics = append(
				diagnostics,
				property.analyzeProperty(section, lineNumber)...,
			)
		}
	}

	return diagnostics
}

func (p wireguardSection) analyzeSection() []protocol.Diagnostic {
	diagnostics := []protocol.Diagnostic{}

	if p.Name == nil {
		// No section name
		severity := protocol.DiagnosticSeverityError
		diagnostics = append(diagnostics, protocol.Diagnostic{
			Message:  "This section is missing a name",
			Severity: &severity,
			Range:    p.getRange(),
		})
		return diagnostics
	}

	if _, found := optionsHeaderMap[*p.Name]; !found {
		// Unknown section
		severity := protocol.DiagnosticSeverityError
		diagnostics = append(diagnostics, protocol.Diagnostic{
			Message:  fmt.Sprintf("Unknown section '%s'. It must be one of: [Interface], [Peer]", *p.Name),
			Severity: &severity,
			Range: protocol.Range{
				Start: protocol.Position{
					Line:      p.StartLine,
					Character: 0,
				},
				End: protocol.Position{
					Line:      p.StartLine,
					Character: 99999999,
				},
			},
		})

		return diagnostics
	}

	return diagnostics
}

// Check if the property is valid.
// Returns a list of diagnostics.
// `belongingSection` is the section to which the property belongs. This value is
// expected to be non-nil and expected to be a valid Wireguard section.
func (p wireguardProperty) analyzeProperty(
	belongingSection *wireguardSection,
	propertyLine uint32,
) []protocol.Diagnostic {
	sectionOptions := optionsHeaderMap[*belongingSection.Name]
	option, found := sectionOptions[p.Key.Name]

	if !found {
		// Unknown property
		severity := protocol.DiagnosticSeverityError
		return []protocol.Diagnostic{
			{
				Message:  fmt.Sprintf("Unknown property '%s'", p.Key.Name),
				Severity: &severity,
				Range: protocol.Range{
					Start: protocol.Position{
						Line:      propertyLine,
						Character: p.Key.Location.Start,
					},
					End: protocol.Position{
						Line:      propertyLine,
						Character: p.Key.Location.End,
					},
				},
			},
		}
	}

	if p.Value == nil {
		// No value to check
		severity := protocol.DiagnosticSeverityWarning
		return []protocol.Diagnostic{
			{
				Message:  "Property is missing a value",
				Severity: &severity,
				Range: protocol.Range{
					Start: protocol.Position{
						Line:      propertyLine,
						Character: 0,
					},
					End: protocol.Position{
						Line:      propertyLine,
						Character: 99999999,
					},
				},
			},
		}
	}

	errors := option.CheckIsValid(p.Value.Value)

	return utils.Map(errors, func(err *docvalues.InvalidValue) protocol.Diagnostic {
		severity := protocol.DiagnosticSeverityError
		return protocol.Diagnostic{
			Message:  err.GetMessage(),
			Severity: &severity,
			Range: protocol.Range{
				Start: protocol.Position{
					Line:      propertyLine,
					Character: p.Value.Location.Start + err.Start,
				},
				End: protocol.Position{
					Line:      propertyLine,
					Character: p.Value.Location.Start + err.End,
				},
			},
		}
	})
}

func (p wireguardParser) checkForDuplicateProperties() []protocol.Diagnostic {
	diagnostics := []protocol.Diagnostic{}

	for _, section := range p.Sections {
		diagnostics = append(diagnostics, section.analyzeDuplicateProperties()...)
	}

	return diagnostics
}

func (p wireguardSection) analyzeDuplicateProperties() []protocol.Diagnostic {
	diagnostics := []protocol.Diagnostic{}

	existingProperties := make(map[string]uint32)

	lines := utils.KeysOfMap(p.Properties)
	slices.Sort(lines)

	for _, currentLineNumber := range lines {
		property := p.Properties[currentLineNumber]
		var skipCheck = false

		if p.Name != nil {
			switch *p.Name {
			case "Interface":
				if _, found := interfaceAllowedDuplicateFields[property.Key.Name]; found {
					skipCheck = true
				}
			case "Peer":
				if _, found := peerAllowedDuplicateFields[property.Key.Name]; found {
					skipCheck = true
				}
			}
		}

		if skipCheck {
			continue
		}

		if existingLineNumber, found := existingProperties[property.Key.Name]; found {
			severity := protocol.DiagnosticSeverityError
			diagnostics = append(diagnostics, protocol.Diagnostic{
				Message:  fmt.Sprintf("Property '%s' is already defined on line %d", property.Key.Name, existingLineNumber+1),
				Severity: &severity,
				Range: protocol.Range{
					Start: protocol.Position{
						Line:      currentLineNumber,
						Character: 0,
					},
					End: protocol.Position{
						Line:      currentLineNumber,
						Character: 99999,
					},
				},
			})
		} else {
			existingProperties[property.Key.Name] = currentLineNumber
		}
	}

	return diagnostics
}

func (p wireguardSection) analyzeInterfaceSection() []protocol.Diagnostic {
	diagnostics := []protocol.Diagnostic{}

	return diagnostics
}

func (p wireguardParser) analyzeAllowedIPIsInRange() []protocol.Diagnostic {
	diagnostics := []protocol.Diagnostic{}

	return diagnostics
}