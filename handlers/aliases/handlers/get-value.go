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
		func(entry ast.AliasValueInterface, pos uint32) int {
			value := entry.GetAliasValue()

			if pos > value.Location.End.Character {
				return -1
			}

			if pos < value.Location.Start.Character {
				return 1
			}

			return 0
		},
	)

	if !found {
		return nil
	}

	return &entry.Values.Values[index]
}
