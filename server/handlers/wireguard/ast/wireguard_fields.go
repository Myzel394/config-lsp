package ast

import (
	"slices"
)

func (c *WGConfig) FindSectionByLine(line uint32) *WGSection {
	index, found := slices.BinarySearchFunc(
		c.Sections,
		line,
		func(current *WGSection, target uint32) int {
			if target < current.Start.Line {
				return -1
			}

			if target > current.End.Line {
				return 1
			}

			return 0
		},
	)

	if !found {
		return nil
	}

	return c.Sections[index]
}

func (c *WGConfig) FindPropertyByLine(line uint32) *WGProperty {
	section := c.FindSectionByLine(line)

	if section == nil {
		return nil
	}

	return section.Properties[line]
}

func (s *WGSection) FindFirstPropertyByName(name string) *WGProperty {
	for _, property := range s.Properties {
		if property.Key.Name == name {
			return property
		}
	}

	return nil
}
