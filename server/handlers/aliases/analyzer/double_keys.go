package analyzer

import (
	"config-lsp/common"
	"config-lsp/handlers/aliases"
	"config-lsp/handlers/aliases/indexes"
)

func analyzeDoubleKeys(
	d *aliases.AliasesDocument,
) []common.LSPError {
	indexes, errors := indexes.CreateIndexes(*d.Parser)

	d.Indexes = &indexes

	if len(errors) > 0 {
		return errors
	}

	return nil
}
