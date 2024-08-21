package handlers

import (
	docvalues "config-lsp/doc-values"
	"config-lsp/handlers/wireguard/fields"
	"config-lsp/handlers/wireguard/parser"
	"config-lsp/utils"
	"context"
	"fmt"
	"net/netip"
	"slices"
	"strings"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func Analyze(
	p parser.WireguardParser,
) []protocol.Diagnostic {
	sectionsErrors := analyzeSections(p.Sections)
	sectionsErrors = append(sectionsErrors, analyzeOnlyOneInterfaceSectionSpecified(p)...)

	if len(sectionsErrors) > 0 {
		return sectionsErrors
	}

	validCheckErrors := checkIfValuesAreValid(p.Sections)

	if len(validCheckErrors) > 0 {
		return validCheckErrors
	}

	diagnostics := make([]protocol.Diagnostic, 0)
	diagnostics = append(diagnostics, analyzeParserForDuplicateProperties(p)...)
	diagnostics = append(diagnostics, analyzeDNSContainsFallback(p)...)
	diagnostics = append(diagnostics, analyzeKeepAliveIsSet(p)...)
	diagnostics = append(diagnostics, analyzeSymmetricPropertiesExist(p)...)
	diagnostics = append(diagnostics, analyzeDuplicateAllowedIPs(p)...)

	return diagnostics
}

func analyzeSections(
	sections []*parser.WireguardSection,
) []protocol.Diagnostic {
	var diagnostics []protocol.Diagnostic

	for _, section := range sections {
		sectionDiagnostics := analyzeSection(*section)

		if len(sectionDiagnostics) > 0 {
			diagnostics = append(diagnostics, sectionDiagnostics...)
		}
	}

	if len(diagnostics) > 0 {
		return diagnostics
	}

	return diagnostics
}

func analyzeOnlyOneInterfaceSectionSpecified(
	p parser.WireguardParser,
) []protocol.Diagnostic {
	var diagnostics []protocol.Diagnostic
	alreadyFound := false

	for _, section := range p.GetSectionsByName("Interface") {
		if alreadyFound {
			severity := protocol.DiagnosticSeverityError
			diagnostics = append(diagnostics, protocol.Diagnostic{
				Message:  "Only one [Interface] section is allowed",
				Severity: &severity,
				Range:    section.GetHeaderLineRange(),
			})
		}

		alreadyFound = true
	}

	return diagnostics
}

func analyzeDNSContainsFallback(
	p parser.WireguardParser,
) []protocol.Diagnostic {
	lineNumber, property := p.FindFirstPropertyByName("DNS")

	if property == nil {
		return []protocol.Diagnostic{}
	}

	dnsAmount := len(strings.Split(property.Value.Value, ","))

	if dnsAmount == 1 {
		severity := protocol.DiagnosticSeverityHint

		return []protocol.Diagnostic{
			{
				Message:  "There is only one DNS server specified. It is recommended to set up fallback DNS servers",
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

func analyzeKeepAliveIsSet(
	p parser.WireguardParser,
) []protocol.Diagnostic {
	var diagnostics []protocol.Diagnostic

	for _, section := range p.GetSectionsByName("Peer") {
		// If an endpoint is set, then we should only check for the keepalive property
		if section.ExistsProperty("Endpoint") && !section.ExistsProperty("PersistentKeepalive") {
			severity := protocol.DiagnosticSeverityHint
			diagnostics = append(diagnostics, protocol.Diagnostic{
				Message:  "PersistentKeepalive is not set. It is recommended to set this property, as it helps to maintain the connection when users are behind NAT",
				Severity: &severity,
				Range:    section.GetRange(),
			})
		}
	}

	return diagnostics
}

// Check if the values are valid.
// Assumes that sections have been analyzed already.
func checkIfValuesAreValid(
	sections []*parser.WireguardSection,
) []protocol.Diagnostic {
	var diagnostics []protocol.Diagnostic

	for _, section := range sections {
		for lineNumber, property := range section.Properties {
			diagnostics = append(
				diagnostics,
				analyzeProperty(property, section, lineNumber)...,
			)
		}
	}

	return diagnostics
}

func analyzeSection(
	s parser.WireguardSection,
) []protocol.Diagnostic {
	var diagnostics []protocol.Diagnostic

	if s.Name == nil {
		// No section name
		severity := protocol.DiagnosticSeverityError
		diagnostics = append(diagnostics, protocol.Diagnostic{
			Message:  "This section is missing a name",
			Severity: &severity,
			Range:    s.GetRange(),
		})
		return diagnostics
	}

	if _, found := fields.OptionsHeaderMap[*s.Name]; !found {
		// Unknown section
		severity := protocol.DiagnosticSeverityError
		diagnostics = append(diagnostics, protocol.Diagnostic{
			Message:  fmt.Sprintf("Unknown section '%s'. It must be one of: [Interface], [Peer]", *s.Name),
			Severity: &severity,
			Range:    s.GetHeaderLineRange(),
		})

		return diagnostics
	}

	return diagnostics
}

// Check if the property is valid.
// Returns a list of diagnostics.
// `belongingSection` is the section to which the property belongs. This value is
// expected to be non-nil and expected to be a valid Wireguard section.
func analyzeProperty(
	p parser.WireguardProperty,
	belongingSection *parser.WireguardSection,
	propertyLine uint32,
) []protocol.Diagnostic {
	sectionOptions := fields.OptionsHeaderMap[*belongingSection.Name]
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
				Range:    p.GetLineRange(propertyLine),
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

func analyzeParserForDuplicateProperties(
	p parser.WireguardParser,
) []protocol.Diagnostic {
	diagnostics := make([]protocol.Diagnostic, 0)

	for _, section := range p.Sections {
		diagnostics = append(diagnostics, analyzeDuplicateProperties(*section)...)
	}

	return diagnostics
}

func analyzeDuplicateProperties(
	s parser.WireguardSection,
) []protocol.Diagnostic {
	var diagnostics []protocol.Diagnostic

	existingProperties := make(map[string]uint32)

	lines := utils.KeysOfMap(s.Properties)
	slices.Sort(lines)

	for _, currentLineNumber := range lines {
		property := s.Properties[currentLineNumber]
		var skipCheck = false

		if s.Name != nil {
			switch *s.Name {
			case "Interface":
				if _, found := fields.InterfaceAllowedDuplicateFields[property.Key.Name]; found {
					skipCheck = true
				}
			case "Peer":
				if _, found := fields.PeerAllowedDuplicateFields[property.Key.Name]; found {
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

type propertyWithLine struct {
	Line     uint32
	Property parser.WireguardProperty
	IpPrefix netip.Prefix
}

func mapAllowedIPsToMasks(p parser.WireguardParser) map[uint8][]propertyWithLine {
	ips := make(map[uint8][]propertyWithLine)

	for _, section := range p.GetSectionsByName("Peer") {
		for lineNumber, property := range section.Properties {
			if property.Key.Name == "AllowedIPs" {
				ipAddress, err := netip.ParsePrefix(property.Value.Value)

				if err != nil {
					// This should not happen...
					continue
				}

				hostBits := uint8(ipAddress.Bits())

				if _, found := ips[uint8(hostBits)]; !found {
					ips[hostBits] = make([]propertyWithLine, 0)
				}

				ips[hostBits] = append(ips[hostBits], propertyWithLine{
					Line:     uint32(lineNumber),
					Property: property,
					IpPrefix: ipAddress,
				})
			}
		}
	}

	return ips
}

// Strategy
// Simply compare the host bits of the IP addresses.
// Use a binary tree to store the host bits.
func analyzeDuplicateAllowedIPs(p parser.WireguardParser) []protocol.Diagnostic {
	diagnostics := make([]protocol.Diagnostic, 0)

	maskedIPs := mapAllowedIPsToMasks(p)
	hostBits := utils.KeysOfMap(maskedIPs)
	slices.Sort(hostBits)

	ipHostSet := utils.CreateIPv4HostSet()

	for _, hostBit := range hostBits {
		ips := maskedIPs[hostBit]

		for _, ipInfo := range ips {
			if ctx, _ := ipHostSet.ContainsIP(ipInfo.IpPrefix); ctx != nil {
				severity := protocol.DiagnosticSeverityError
				definedLine := (*ctx).Value("line").(uint32)

				diagnostics = append(diagnostics, protocol.Diagnostic{
					Message:  fmt.Sprintf("This IP range is already covered on line %d", definedLine),
					Severity: &severity,
					Range: protocol.Range{
						Start: protocol.Position{
							Line:      ipInfo.Line,
							Character: ipInfo.Property.Key.Location.Start,
						},
						End: protocol.Position{
							Line:      ipInfo.Line,
							Character: ipInfo.Property.Value.Location.End,
						},
					},
				})
			} else {
				humanLineNumber := ipInfo.Line + 1
				ctx := context.WithValue(context.Background(), "line", humanLineNumber)

				ipHostSet.AddIP(
					ipInfo.IpPrefix,
					ctx,
				)
			}
		}
	}

	return diagnostics
}

func analyzeSymmetricPropertiesExist(
	p parser.WireguardParser,
) []protocol.Diagnostic {
	diagnostics := make([]protocol.Diagnostic, 0, 4)
	severity := protocol.DiagnosticSeverityHint

	for _, section := range p.GetSectionsByName("Interface") {
		preUpLine, preUpProperty := section.FetchFirstProperty("PreUp")
		preDownLine, preDownProperty := section.FetchFirstProperty("PreDown")

		postUpLine, postUpProperty := section.FetchFirstProperty("PostUp")
		postDownLine, postDownProperty := section.FetchFirstProperty("PostDown")

		if preUpProperty != nil && preDownProperty == nil {
			diagnostics = append(diagnostics, protocol.Diagnostic{
				Message:  "PreUp is set, but PreDown is not. It is recommended to set both properties symmetrically",
				Range:    preUpProperty.GetLineRange(*preUpLine),
				Severity: &severity,
			})
		} else if preUpProperty == nil && preDownProperty != nil {
			diagnostics = append(diagnostics, protocol.Diagnostic{
				Message:  "PreDown is set, but PreUp is not. It is recommended to set both properties symmetrically",
				Range:    preDownProperty.GetLineRange(*preDownLine),
				Severity: &severity,
			})
		}

		if postUpProperty != nil && postDownProperty == nil {
			diagnostics = append(diagnostics, protocol.Diagnostic{
				Message:  "PostUp is set, but PostDown is not. It is recommended to set both properties symmetrically",
				Range:    postUpProperty.GetLineRange(*postUpLine),
				Severity: &severity,
			})
		} else if postUpProperty == nil && postDownProperty != nil {
			diagnostics = append(diagnostics, protocol.Diagnostic{
				Message:  "PostDown is set, but PostUp is not. It is recommended to set both properties symmetrically",
				Range:    postDownProperty.GetLineRange(*postDownLine),
				Severity: &severity,
			})
		}
	}

	return diagnostics
}
