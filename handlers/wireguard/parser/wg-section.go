package parser

import (
	"fmt"
	"regexp"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

type PropertyNotFoundError struct{}

func (e PropertyNotFoundError) Error() string {
	return "Property not found"
}

type PropertyNotFullyTypedError struct{}

func (e PropertyNotFullyTypedError) Error() string {
	return "Property not fully typed"
}

type WireguardSection struct {
	Name       *string
	StartLine  uint32
	EndLine    uint32
	Properties WireguardProperties
}

func (s WireguardSection) String() string {
	var name string

	if s.Name == nil {
		name = "<nil>"
	} else {
		name = *s.Name
	}

	return fmt.Sprintf("[%s]; %d-%d: %v", name, s.StartLine, s.EndLine, s.Properties)
}

func (s WireguardSection) GetHeaderLineRange() protocol.Range {
	return protocol.Range{
		Start: protocol.Position{
			Line:      s.StartLine,
			Character: 0,
		},
		End: protocol.Position{
			Line:      s.StartLine,
			Character: 99999999,
		},
	}
}

func (s WireguardSection) GetRange() protocol.Range {
	return protocol.Range{
		Start: protocol.Position{
			Line:      s.StartLine,
			Character: 0,
		},
		End: protocol.Position{
			Line:      s.EndLine,
			Character: 99999999,
		},
	}
}

func (s WireguardSection) FetchFirstProperty(name string) (*uint32, *WireguardProperty) {
	for line, property := range s.Properties {
		if property.Key.Name == name {
			return &line, &property
		}
	}

	return nil, nil
}

func (s WireguardSection) ExistsProperty(name string) bool {
	_, property := s.FetchFirstProperty(name)

	return property != nil
}

func (s WireguardSection) GetPropertyByLine(lineNumber uint32) (*WireguardProperty, error) {
	property, found := s.Properties[lineNumber]

	if !found {
		return nil, PropertyNotFoundError{}
	}

	return &property, nil
}

var validHeaderPattern = regexp.MustCompile(`^\s*\[(?P<header>.+?)\]\s*$`)

// Create a new create section
// Return (<name>, <new section>)
func CreateWireguardSection(
	startLine uint32,
	endLine uint32,
	headerLine string,
	props WireguardProperties,
) WireguardSection {
	match := validHeaderPattern.FindStringSubmatch(headerLine)

	var header string

	if match == nil {
		// Still typing it
		header = headerLine[1:]
	} else {
		header = match[1]
	}

	return WireguardSection{
		Name:       &header,
		StartLine:  startLine,
		EndLine:    endLine,
		Properties: props,
	}
}
