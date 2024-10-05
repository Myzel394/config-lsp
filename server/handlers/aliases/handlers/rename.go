package handlers

import (
	"config-lsp/handlers/aliases/indexes"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func RenameAlias(
	i indexes.AliasesIndexes,
	oldName string,
	newName string,
) []protocol.TextEdit {
	normalizedName := indexes.NormalizeKey(oldName)
	definitionEntry := i.Keys[normalizedName]
	occurrences := i.UserOccurrences[normalizedName]

	changes := make([]protocol.TextEdit, 0, len(occurrences))

	if definitionEntry != nil {
		// Own rename
		changes = append(changes, protocol.TextEdit{
			Range:   definitionEntry.Key.Location.ToLSPRange(),
			NewText: newName,
		})
	}

	// Other AliasValueUser occurrences
	for _, value := range occurrences {
		changes = append(changes, protocol.TextEdit{
			Range:   value.Location.ToLSPRange(),
			NewText: newName,
		})
	}

	return changes
}
