package matchparser

import (
	"config-lsp/common"
	"slices"
)

func (m Match) GetEntryByCursor(cursor common.CursorPosition) *MatchEntry {
	index, found := slices.BinarySearchFunc(
		m.Entries,
		cursor,
		func(current *MatchEntry, target common.CursorPosition) int {
			if current.Start.IsAfterCursorPosition(target) {
				return 1
			}

			if current.End.IsBeforeCursorPosition(target) {
				return -1
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

func (c MatchCriteria) IsCursorBetween(cursor uint32) bool {
	return cursor >= c.Start.Character && cursor <= c.End.Character
}

func (e MatchEntry) GetValueByCursor(cursor uint32) *MatchValue {
	if e.Values == nil {
		return nil
	}

	index, found := slices.BinarySearchFunc(
		e.Values.Values,
		cursor,
		func(current *MatchValue, target uint32) int {
			if target < current.Start.Character {
				return 1
			}

			if target > current.End.Character {
				return -1
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

func (v MatchValues) IsCursorBetween(cursor uint32) bool {
	return cursor >= v.Start.Character && cursor <= v.End.Character
}
