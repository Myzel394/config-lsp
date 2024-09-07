package handlers

import (
	"config-lsp/handlers/aliases/ast"
	"config-lsp/handlers/aliases/indexes"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func RenameAlias(
	i indexes.AliasesIndexes,
	oldEntry *ast.AliasEntry,
	newName string,
) []protocol.TextEdit {
	occurrences := i.UserOccurrences[indexes.NormalizeKey(oldEntry.Key.Value)]
	changes := make([]protocol.TextEdit, 0, len(occurrences))

	// Own rename
	changes = append(changes, protocol.TextEdit{
		Range:   oldEntry.Key.Location.ToLSPRange(),
		NewText: newName,
	})

	// Other AliasValueUser occurrences
	for _, value := range occurrences {
		changes = append(changes, protocol.TextEdit{
			Range:   value.Location.ToLSPRange(),
			NewText: newName,
		})
	}

	return changes
}
