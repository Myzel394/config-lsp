package analyzer

import (
	"config-lsp/common"
	"config-lsp/handlers/aliases"
	"config-lsp/handlers/aliases/ast"
	"config-lsp/handlers/aliases/fetchers"
	"config-lsp/handlers/aliases/indexes"
	"config-lsp/utils"
	"fmt"

	ers "errors"
)

func analyzeValuesAreValid(
	d *aliases.AliasesDocument,
) []common.LSPError {
	errors := make([]common.LSPError, 0)

	it := d.Parser.Aliases.Iterator()

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

		for _, value := range entry.Values.Values {
			newErrors := checkValue(d.Indexes, value)
			newErrors = utils.Map(
				newErrors,
				func(e common.LSPError) common.LSPError {
					startPosition := value.GetAliasValue().Location.Start.Character
					return e.ShiftCharacter(-startPosition)
				},
			)

			errors = append(errors, newErrors...)
		}
	}

	return errors
}

func checkValue(
	i *indexes.AliasesIndexes,
	value ast.AliasValueInterface,
) []common.LSPError {
	switch value.(type) {
	case ast.AliasValueUser:
		aliasValue := value.(ast.AliasValueUser)

		users := fetchers.GetAvailableUserValues(i)

		if _, found := users[aliasValue.Value]; !found {
			return []common.LSPError{{
				Range: aliasValue.Location,
				Err:   ers.New(fmt.Sprintf("User '%s' not found", aliasValue.Value)),
			}}
		}
	}
	return nil
}
