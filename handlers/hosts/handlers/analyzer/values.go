package analyzer

import (
	"config-lsp/common"
	docvalues "config-lsp/doc-values"
	"config-lsp/handlers/hosts/fields"
	"config-lsp/utils"
	"errors"
)

func analyzeEntriesSetCorrectly(
	parser HostsParser,
) []common.LSPError {
	err := make([]common.LSPError, 0)

	for lineNumber, entry := range parser.Tree.Entries {
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
	parser HostsParser,
) []common.LSPError {
	err := make([]common.LSPError, 0)

	for _, entry := range parser.Tree.Entries {
		err = append(
			err,
			utils.Map(
				fields.IPAddressField.CheckIsValid(entry.IPAddress.Value.String()),
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
				fields.HostnameField.CheckIsValid(entry.Hostname.Value),
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
					fields.HostnameField.CheckIsValid(alias.Value),
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
