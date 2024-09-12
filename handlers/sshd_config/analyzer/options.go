package analyzer

import (
	"config-lsp/common"
	docvalues "config-lsp/doc-values"
	sshdconfig "config-lsp/handlers/sshd_config"
	"config-lsp/handlers/sshd_config/ast"
	"config-lsp/handlers/sshd_config/fields"
	"config-lsp/utils"
	"errors"
	"fmt"
)

func analyzeOptionsAreValid(
	d *sshdconfig.SSHDocument,
) []common.LSPError {
	errs := make([]common.LSPError, 0)
	it := d.Config.Options.Iterator()

	for it.Next() {
		line := it.Key().(uint32)
		entry := it.Value().(ast.SSHEntry)

		option := entry.GetOption()

		if option.Key != nil {
			docOption, found := fields.Options[option.Key.Value]

			if !found {
				errs = append(errs, common.LSPError{
					Range: option.Key.LocationRange,
					Err:   errors.New(fmt.Sprintf("Unknown option: %s", option.Key.Value)),
				})
				continue
			}

			if option.OptionValue == nil {
				continue
			}

			invalidValues := docOption.CheckIsValid(option.OptionValue.Value)

			errs = append(
				errs,
				utils.Map(
					invalidValues,
					func(invalidValue *docvalues.InvalidValue) common.LSPError {
						err := docvalues.LSPErrorFromInvalidValue(line, *invalidValue)
						err.ShiftCharacter(option.OptionValue.Start.Character)

						return err
					},
				)...,
			)

		}
	}

	return errs
}
