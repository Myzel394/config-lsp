package matchparser

import "slices"

func (m Match) GetEntryByCursor(cursor uint32) *MatchEntry {
	index, found := slices.BinarySearchFunc(
		m.Entries,
		cursor,
		func(entry *MatchEntry, cursor uint32) int {
			if cursor < entry.Start.Character {
				return 1
			}

			if cursor > entry.End.Character {
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
		func(value *MatchValue, cursor uint32) int {
			if cursor < value.Start.Character {
				return 1
			}

			if cursor > value.End.Character {
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
