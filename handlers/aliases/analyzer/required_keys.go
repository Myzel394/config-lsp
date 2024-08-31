package analyzer

import (
	"config-lsp/common"
	"config-lsp/handlers/aliases"
	"config-lsp/handlers/aliases/indexes"
	"fmt"

	ers "errors"
)

var requiredFields = []string{
	indexes.NormalizeKey("mailer-daemon"),
	indexes.NormalizeKey("hostmaster"),
	indexes.NormalizeKey("postmaster"),
}

func analyzeContainsRequiredKeys(
	d aliases.AliasesDocument,
) []common.LSPError {
	errors := make([]common.LSPError, 0)

	for _, requiredField := range requiredFields {
		if _, found := d.Indexes.Keys[requiredField]; !found {
			errors = append(errors, common.LSPError{
				Range: common.GlobalLocationRange,
				Err:   ers.New(fmt.Sprintf("Please add the alias '%s'. It is required by the aliases file.", requiredField)),
			})
		}
	}

	return errors
}
