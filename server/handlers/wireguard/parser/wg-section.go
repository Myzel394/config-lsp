package parser

import (
	"config-lsp/handlers/wireguard/ast"
	"regexp"
)

type PropertyNotFoundError struct{}

func (e PropertyNotFoundError) Error() string {
	return "Property not found"
}

type PropertyNotFullyTypedError struct{}

func (e PropertyNotFullyTypedError) Error() string {
	return "Property not fully typed"
}

func (s ast.WGSection) FetchFirstProperty(name string) (*uint32, *ast.WGProperty) {
	for line, property := range s.Properties {
		if property.Key.Name == name {
			return &line, &property
		}
	}

	return nil, nil
}

func (s ast.WGSection) ExistsProperty(name string) bool {
	_, property := s.FetchFirstProperty(name)

	return property != nil
}

func (s ast.WGSection) GetPropertyByLine(lineNumber uint32) (*ast.WGProperty, error) {
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
) ast.WGSection {
	match := validHeaderPattern.FindStringSubmatch(headerLine)

	var header string

	if match == nil {
		// Still typing it
		header = headerLine[1:]
	} else {
		header = match[1]
	}

	return ast.WGSection{
		Header:     &header,
		StartLine:  startLine,
		EndLine:    endLine,
		Properties: props,
	}
}
