package analyzer

import (
	"config-lsp/common"
	"config-lsp/handlers/sshd_config"
	"config-lsp/handlers/sshd_config/indexes"
	"config-lsp/utils"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func Analyze(
	d *sshdconfig.SSHDocument,
) []protocol.Diagnostic {
	errors := analyzeOptionsAreValid(d)

	if len(errors) > 0 {
		return errsToDiagnostics(errors)
	}

	i, indexErrors := indexes.CreateIndexes(*d.Config)

	d.Indexes = i

	errors = append(errors, indexErrors...)

	if len(errors) > 0 {
		return errsToDiagnostics(errors)
	}

	includeErrors := analyzeIncludeValues(d)

	if len(includeErrors) > 0 {
		errors = append(errors, includeErrors...)
	} else {
		for _, include := range d.Indexes.Includes {
			for _, value := range include.Values {
				for _, path := range value.Paths {
					_, err := parseFile(string(path))

					if err != nil {
						errors = append(errors, common.LSPError{
							Range: value.LocationRange,
							Err:   err,
						})
					}
				}
			}
		}
	}

	errors = append(errors, analyzeMatchBlocks(d)...)

	if len(errors) > 0 {
		return errsToDiagnostics(errors)
	}

	return nil
}

func errsToDiagnostics(errs []common.LSPError) []protocol.Diagnostic {
	return utils.Map(
		errs,
		func(err common.LSPError) protocol.Diagnostic {
			return err.ToDiagnostic()
		},
	)
}
