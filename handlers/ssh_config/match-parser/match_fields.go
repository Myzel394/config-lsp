package matchparser

import (
	"config-lsp/common"
	"slices"
)

func (m Match) GetEntryAtPosition(position common.Position) *MatchEntry {
	index, found := slices.BinarySearchFunc(
		m.Entries,
		position,
		func(current *MatchEntry, target common.Position) int {
			if current.IsPositionAfterEnd(target) {
				return -1
			}

			if current.IsPositionBeforeStart(target) {
				return 1
			}

			return 0
		},
	)

	if !found {
		return nil
	}

	entry := m.Entries[index]

	return entry
}

func (m Match) FindEntries(name string) []*MatchEntry {
	entries := make([]*MatchEntry, 0, 5)

	for _, entry := range m.Entries {
		if entry.Value.Value == name {
			entries = append(entries, entry)
		}
	}

	return entries
}

func (m Match) GetPreviousEntry(e *MatchEntry) *MatchEntry {
	index := slices.Index(m.Entries, e)

	if index == 0 || index == -1 {
		return nil
	}

	return m.Entries[index-1]
}

func (e MatchEntry) GetValueAtPosition(position common.Position) *MatchValue {
	if e.Values == nil {
		return nil
	}

	index, found := slices.BinarySearchFunc(
		e.Values.Values,
		position,
		func(current *MatchValue, target common.Position) int {
			if current.IsPositionAfterEnd(target) {
				return -1
			}

			if current.IsPositionBeforeStart(target) {
				return 1
			}

			return 0
		},
	)

	if !found {
		return nil
	}

	value := e.Values.Values[index]

	return value
}
