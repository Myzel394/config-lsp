package indexes

import (
	"config-lsp/common"
	"config-lsp/handlers/aliases/ast"
	"config-lsp/handlers/aliases/shared"
	"strings"
)

type AliasesIndexes struct {
	Keys map[string]*ast.AliasKey
}

func NormalizeKey(key string) string {
	return strings.ToLower(key)
}

func CreateIndexes(parser ast.AliasesParser) (AliasesIndexes, []common.LSPError) {
	errors := make([]common.LSPError, 0)
	indexes := &AliasesIndexes{
		Keys: make(map[string]*ast.AliasKey),
	}

	it := parser.Aliases.Iterator()

	for it.Next() {
		entry := it.Value().(*ast.AliasEntry)

		normalizedAlias := NormalizeKey(entry.Key.Value)

		if existingEntry, found := indexes.Keys[normalizedAlias]; found {
			errors = append(errors, common.LSPError{
				Range: entry.Key.Location,
				Err: shared.DuplicateKeyEntry{
					AlreadyFoundAt: existingEntry.Location.Start.Line,
					Key:            entry.Key.Value,
				},
			})

			continue
		}

		indexes.Keys[normalizedAlias] = entry.Key
	}

	return *indexes, errors
}
