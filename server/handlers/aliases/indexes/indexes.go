package indexes

import (
	"config-lsp/common"
	"config-lsp/handlers/aliases/ast"
	"config-lsp/handlers/aliases/shared"
	"strings"
)

type AliasesIndexes struct {
	Keys            map[string]*ast.AliasEntry
	UserOccurrences map[string][]*ast.AliasValueUser
}

func NormalizeKey(key string) string {
	return strings.ToLower(key)
}

func CreateIndexes(parser ast.AliasesParser) (AliasesIndexes, []common.LSPError) {
	errors := make([]common.LSPError, 0)
	indexes := &AliasesIndexes{
		Keys:            make(map[string]*ast.AliasEntry),
		UserOccurrences: make(map[string][]*ast.AliasValueUser),
	}

	it := parser.Aliases.Iterator()

	for it.Next() {
		entry := it.Value().(*ast.AliasEntry)

		if entry.Values != nil {
			for _, value := range entry.Values.Values {
				switch value.(type) {
				case ast.AliasValueUser:
					userValue := value.(ast.AliasValueUser)

					indexes.UserOccurrences[userValue.Value] = append(
						indexes.UserOccurrences[userValue.Value],
						&userValue,
					)
				}
			}
		}

		if entry.Key == nil || entry.Key.Value == "" {
			continue
		}

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

		indexes.Keys[normalizedAlias] = entry
	}

	return *indexes, errors
}
