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
				return 1
			}

			if target > current.End.Line {
				return -1
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

	if property, found := section.Properties.Get(line); found {
		return property.(*WGProperty)
	}

	return nil
}

func (s *WGSection) FindFirstPropertyByName(name string) (uint32, *WGProperty) {
	it := s.Properties.Iterator()
	for it.Next() {
		line := it.Key().(uint32)
		property := it.Value().(*WGProperty)
		if property.Key.Name == name {
			return line, property
		}
	}

	return 0, nil
}

func (s *WGSection) GetLastProperty() *WGProperty {
	if s.Properties.Size() == 0 {
		return nil
	}

	lastLine, _ := s.Properties.Max()
	lastProperty, _ := s.Properties.Get(lastLine)
	return lastProperty.(*WGProperty)
}
