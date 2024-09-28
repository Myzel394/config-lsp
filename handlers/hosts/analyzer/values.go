package analyzer

import (
	"config-lsp/common"
	docvalues "config-lsp/doc-values"
	"config-lsp/handlers/hosts/ast"
	"config-lsp/handlers/hosts/fields"
	"config-lsp/utils"
	"errors"
)

func analyzeEntriesSetCorrectly(
	parser ast.HostsParser,
) []common.LSPError {
	err := make([]common.LSPError, 0)

	it := parser.Tree.Entries.Iterator()

	for it.Next() {
		lineNumber := it.Key().(uint32)
		entry := it.Value().(*ast.HostsEntry)

		if entry.IPAddress == nil {
			err = append(err, common.LSPError{
				Range: common.CreateFullLineRange(lineNumber),
				Err:   errors.New("IP Address is required"),
			})
			continue
		}

		if entry.Hostname == nil {
			err = append(err, common.LSPError{
				Range: common.CreateFullLineRange(lineNumber),
				Err:   errors.New("Hostname is required"),
			})
			continue
		}
	}

	return err
}

func analyzeEntriesAreValid(
	parser ast.HostsParser,
) []common.LSPError {
	err := make([]common.LSPError, 0)

	it := parser.Tree.Entries.Iterator()

	for it.Next() {
		entry := it.Value().(*ast.HostsEntry)

		err = append(
			err,
			utils.Map(
				fields.IPAddressField.DeprecatedCheckIsValid(entry.IPAddress.Value.String()),
				func(val *docvalues.InvalidValue) common.LSPError {
					return common.LSPError{
						Range: entry.IPAddress.Location,
						Err:   val.Err,
					}
				},
			)...,
		)

		err = append(
			err,
			utils.Map(
				fields.HostnameField.DeprecatedCheckIsValid(entry.Hostname.Value),
				func(val *docvalues.InvalidValue) common.LSPError {
					return common.LSPError{
						Range: entry.Hostname.Location,
						Err:   val.Err,
					}
				},
			)...,
		)

		for _, alias := range entry.Aliases {
			err = append(
				err,
				utils.Map(
					fields.HostnameField.DeprecatedCheckIsValid(alias.Value),
					func(val *docvalues.InvalidValue) common.LSPError {
						return common.LSPError{
							Range: alias.Location,
							Err:   val.Err,
						}
					},
				)...,
			)
		}
	}

	return err
}
