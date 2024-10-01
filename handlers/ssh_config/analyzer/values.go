package analyzer

import (
	"config-lsp/common"
	docvalues "config-lsp/doc-values"
	sshconfig "config-lsp/handlers/ssh_config"
	"config-lsp/handlers/ssh_config/fields"
	"config-lsp/utils"
	"errors"
	"fmt"
)

func analyzeValuesAreValid(
	d *sshconfig.SSHDocument,
) []common.LSPError {
	errs := make([]common.LSPError, 0)

	for _, info := range d.Config.GetAllOptions() {
		option := info.Option
		block := info.Block

		docOption, found := fields.Options[option.Key.Key]

		if !found {
			if d.Indexes.CanOptionBeIgnored(option, block) {
				// Skip
				continue
			}

			errs = append(errs, common.LSPError{
				Range: option.Key.LocationRange,
				Err:   errors.New(fmt.Sprintf("Unknown option: %s", option.Key.Value.Value)),
			})

			continue
		}

		errs = append(
			errs,
			utils.Map(
				docOption.DeprecatedCheckIsValid(option.OptionValue.Value.Value),
				func(invalidValue *docvalues.InvalidValue) common.LSPError {
					err := docvalues.LSPErrorFromInvalidValue(option.Start.Line, *invalidValue)
					err.ShiftCharacter(option.OptionValue.Start.Character)

					return err
				},
			)...,
		)
	}

	return errs
}
