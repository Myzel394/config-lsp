package handlers

import (
	"config-lsp/handlers/aliases/ast"
	"slices"
)

func GetValueAtCursor(
	cursor uint32,
	entry *ast.AliasEntry,
) *ast.AliasValueInterface {
	if entry.Values == nil || len(entry.Values.Values) == 0 {
		return nil
	}

	index, found := slices.BinarySearchFunc(
		entry.Values.Values,
		cursor,
		func(current ast.AliasValueInterface, target uint32) int {
			value := current.GetAliasValue()

			if target < value.Location.Start.Character {
				return 1
			}

			if target > value.Location.End.Character {
				return -1
			}

			return 0
		},
	)

	if !found {
		return nil
	}

	return &entry.Values.Values[index]
}
