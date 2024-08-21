package handlers

import (
	docvalues "config-lsp/doc-values"
	"config-lsp/handlers/wireguard/fields"
	"config-lsp/handlers/wireguard/parser"
	"fmt"
	"strings"
)

func getPropertyInfo(
	p parser.WireguardProperty,
	cursor uint32,
	section parser.WireguardSection,
) []string {
	if cursor <= p.Key.Location.End {
		options, found := fields.OptionsHeaderMap[*section.Name]

		if !found {
			return []string{}
		}

		option, found := options[p.Key.Name]

		if !found {
			return []string{}
		}

		return strings.Split(option.Documentation, "\n")
	}

	options, found := fields.OptionsHeaderMap[*section.Name]

	if !found {
		return []string{}
	}

	if option, found := options[p.Key.Name]; found {
		return option.GetTypeDescription()
	}

	return []string{}
}

func getSectionInfo(s parser.WireguardSection) []string {
	if s.Name == nil {
		return []string{}
	}

	contents := []string{
		fmt.Sprintf("## [%s]", *s.Name),
		"",
	}

	var option *docvalues.EnumString = nil

	switch *s.Name {
	case "Interface":
		option = &fields.HeaderInterfaceEnum
	case "Peer":
		option = &fields.HeaderPeerEnum
	}

	if option == nil {
		return contents
	}

	contents = append(contents, strings.Split(option.Documentation, "\n")...)

	return contents
}

func GetHoverContent(
	p parser.WireguardParser,
	line uint32,
	cursor uint32,
) []string {
	section := p.GetSectionByLine(line)

	if section == nil {
		return []string{}
	}

	sectionInfo := getSectionInfo(*section)

	property, _ := section.GetPropertyByLine(line)

	if property == nil {
		return sectionInfo
	}

	propertyInfo := getPropertyInfo(*property, cursor, *section)

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
