package handlers

import (
	"config-lsp/handlers/aliases/ast"
	"fmt"
	"slices"
)

func getValueAtCursor(
	cursor uint32,
	entry *ast.AliasEntry,
) *ast.AliasValueInterface {
	if entry.Values == nil || len(entry.Values.Values) == 0 {
		return nil
	}

	println(fmt.Sprintf("Values: %v", entry.Values.Values))
	index, found := slices.BinarySearchFunc(
		entry.Values.Values,
		cursor,
		func(entry ast.AliasValueInterface, pos uint32) int {
			println(fmt.Sprintf("Entry: %v", entry))
			value := entry.(ast.AliasValue)

			if value.Location.End.Character > pos {
				return 1
			}

			if value.Location.Start.Character < pos {
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

