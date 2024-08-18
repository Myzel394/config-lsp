package wireguard

import (
	docvalues "config-lsp/doc-values"
	"fmt"
	"strings"
)

func (p wireguardProperty) getHoverInfo(cursor uint32, section *wireguardSection) []string {
	if cursor <= p.Key.Location.End {
		options, found := optionsHeaderMap[*section.Name]

		if !found {
			return []string{}
		}

		option, found := options[p.Key.Name]

		if !found {
			return []string{}
		}

		return strings.Split(option.Documentation, "\n")
	}

	options, found := optionsHeaderMap[*section.Name]

	if !found {
		return []string{}
	}

	if option, found := options[p.Key.Name]; found {
		return option.GetTypeDescription()
	}

	return []string{}
}

func (p wireguardParser) getHeaderInfo(line uint32, cursor uint32) []string {
	section := p.getSectionByLine(line)

	if section == nil {
		return []string{}
	}

	sectionInfo := section.getHeaderInfo()

	property, _ := section.findProperty(line)

	if property == nil {
		return sectionInfo
	}

	propertyInfo := property.getHoverInfo(cursor, section)

	if len(propertyInfo) == 0 {
		return sectionInfo
	}

	contents := append(sectionInfo, []string{
		"",
		fmt.Sprintf("### %s", property.Key.Name),
	}...)
	contents = append(contents, propertyInfo...)

	return contents
}

func (p wireguardSection) getHeaderInfo() []string {
	if p.Name == nil {
		return []string{}
	}

	contents := []string{
		fmt.Sprintf("## [%s]", *p.Name),
		"",
	}

	var option *docvalues.EnumString = nil

	switch *p.Name {
	case "Interface":
		option = &headerInterfaceEnum
	case "Peer":
		option = &headerPeerEnum
	}

	if option == nil {
		return contents
	}

	contents = append(contents, strings.Split(option.Documentation, "\n")...)

	return contents
}
