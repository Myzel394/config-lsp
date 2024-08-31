package analyzer

import (
	"config-lsp/common"
	"config-lsp/handlers/aliases/ast"

	ers "errors"
)

func analyzeValuesAreValid(
	p ast.AliasesParser,
) []common.LSPError {
	errors := make([]common.LSPError, 0)

	it := p.Aliases.Iterator()

	for it.Next() {
		entry := it.Value().(*ast.AliasEntry)

		if entry.Key == nil {
			errors = append(errors, common.LSPError{
				Range: entry.Location,
				Err:   ers.New("An alias is required"),
			})

			continue
		}

		if entry.Separator == nil {
			errors = append(errors, common.LSPError{
				Range: entry.Location,
				Err:   ers.New("A ':' is required as a separator"),
			})

			continue
		}

		if entry.Values == nil || len(entry.Values.Values) == 0 {
			errors = append(errors, common.LSPError{
				Range: entry.Location,
				Err:   ers.New("A value is required"),
			})

			continue
		}
	}

	return errors
}
