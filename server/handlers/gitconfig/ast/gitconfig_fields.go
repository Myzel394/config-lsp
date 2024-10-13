package ast

import (
	"fmt"
	"regexp"
	"slices"
	"strings"
)

func (c *GitConfig) Clear() {
	c.Sections = []*GitSection{}
	c.CommentLines = map[uint32]struct{}{}
}

func (c *GitConfig) FindSection(line uint32) *GitSection {
	index, found := slices.BinarySearchFunc(
		c.Sections,
		line,
		func(current *GitSection, target uint32) int {
			if target > current.End.Line {
				return -1
			}

			if target < current.Start.Line {
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

func (c *GitConfig) FindOption(line uint32) (*GitSection, *GitEntry) {
	section := c.FindSection(line)

	if section == nil {
		return nil, nil
	}

	entry := section.FindOption(line)

	return section, entry
}

func (s *GitSection) FindOption(line uint32) *GitEntry {
	index, found := slices.BinarySearchFunc(
		s.Entries,
		line,
		func(current *GitEntry, target uint32) int {
			if target > current.End.Line {
				return -1
			}

			if target < current.Start.Line {
				return 1
			}

			return 0
		},
	)

	if !found {
		return nil
	}

	return s.Entries[index]
}

var nonWhitespacePattern = regexp.MustCompile(`\S+`)
var deprecatedSectionPattern = regexp.MustCompile(`.+?\..+`)

func (t GitSectionTitle) NormalizedTitle() string {
	entries := nonWhitespacePattern.FindAllString(string(t), -1)

	if entries == nil {
		return string(t)
	}

	if len(entries) == 1 {
		title := entries[0]

		dotEntries := strings.Split(title, ".")

		if len(dotEntries) == 2 {
			// Deprecated title format
			return fmt.Sprintf(`%s "%s"`, strings.ToLower(dotEntries[0]), strings.ToLower(dotEntries[1]))
		}

		return strings.ToLower(title)
	}

	if len(entries) == 2 {
		return fmt.Sprintf(`%s "%s"`, strings.ToLower(entries[0]), entries[1])
	}

	return string(t)
}
