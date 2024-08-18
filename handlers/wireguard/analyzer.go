package wireguard

import (
	"config-lsp/utils"
	"fmt"
	"slices"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func (p wireguardParser) analyze() []protocol.Diagnostic {
	diagnostics := []protocol.Diagnostic{}

	diagnostics = append(diagnostics, p.checkForDuplicateProperties()...)

	return diagnostics
}

func (p wireguardParser) checkForDuplicateProperties() []protocol.Diagnostic {
	diagnostics := []protocol.Diagnostic{}

	for _, section := range p.Sections {
		diagnostics = append(diagnostics, section.analyDuplicateProperties()...)
	}

	return diagnostics
}

func (p wireguardSection) analyDuplicateProperties() []protocol.Diagnostic {
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
