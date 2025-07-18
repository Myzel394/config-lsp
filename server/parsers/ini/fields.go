package ini

import (
	"slices"
)

// Find the section containing the specified line
func (c *Config) FindSectionByLine(line uint32) *Section {
	index, found := slices.BinarySearchFunc(
		c.Sections,
		line,
		func(current *Section, target uint32) int {
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

// Find a property at the specified line
func (c *Config) FindPropertyByLine(line uint32) *Property {
	section := c.FindSectionByLine(line)

	if section == nil {
		return nil
	}

	if property, found := section.Properties.Get(line); found {
		return property.(*Property)
	}

	return nil
}

// Get all properties from `startLine` to `endLine` (inclusive)
func (c Config) GetPropertesInRange(startLine uint32, endLine uint32) []IniPropertyLocationIndex {
	options := make([]IniPropertyLocationIndex, 0, 50)

	for _, section := range c.Sections {
		it := section.Properties.Iterator()
		for it.Next() {
			property := it.Value().(*Property)

			options = append(options, IniPropertyLocationIndex{
				Section:  section,
				property: *property,
			})
		}
	}

	return options
}

// Find the first property with the given name in a section
func (s *Section) FindFirstPropertyByName(name string) (uint32, *Property) {
	it := s.Properties.Iterator()
	for it.Next() {
		line := it.Key().(uint32)
		property := it.Value().(*Property)
		if property.Key.Name == name {
			return line, property
		}
	}

	return 0, nil
}

// Find the last property in a section
func (s *Section) GetLastProperty() *Property {
	if s.Properties.Size() == 0 {
		return nil
	}

	lastLine, _ := s.Properties.Max()
	lastProperty, _ := s.Properties.Get(lastLine)
	return lastProperty.(*Property)
}
