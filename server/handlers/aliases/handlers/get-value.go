package handlers

import (
	"config-lsp/common"
	"config-lsp/handlers/aliases/ast"
	"slices"
)

func GetValueAtPosition(
	position common.Position,
	entry *ast.AliasEntry,
) *ast.AliasValueInterface {
	if entry.Values == nil || len(entry.Values.Values) == 0 {
		return nil
	}

	index, found := slices.BinarySearchFunc(
		entry.Values.Values,
		position,
		func(rawCurrent ast.AliasValueInterface, target common.Position) int {
			current := rawCurrent.GetAliasValue()

			if current.Location.IsPositionAfterEnd(target) {
				return -1
			}

			if current.Location.IsPositionBeforeStart(target) {
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
