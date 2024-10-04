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
