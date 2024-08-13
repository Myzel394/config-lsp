package wireguard

import (
	"fmt"
	"regexp"
)

type wireguardSection struct {
	StartLine uint32
	EndLine   uint32
	// nil = do not belong to a section
	Name       *string
	Properties wireguardProperties
}

func (s wireguardSection) String() string {
	var name string

	if s.Name == nil {
		name = "//<nil>//"
	} else {
		name = *s.Name
	}

	return fmt.Sprintf("[%s]; %d-%d: %v", name, s.StartLine, s.EndLine, s.Properties)
}

var validHeaderPattern = regexp.MustCompile(`^\s*\[(?P<header>.+?)\]\s*$`)

func createWireguardSection(startLine uint32, endLine uint32, headerLine string, props wireguardProperties) wireguardSection {
	match := validHeaderPattern.FindStringSubmatch(headerLine)

	var header string

	if match == nil {
		// Still typing it
		header = headerLine[1:]
	} else {
		header = match[1]
	}

	return wireguardSection{
		StartLine:  startLine,
		EndLine:    endLine,
		Name:       &header,
		Properties: props,
	}
}
