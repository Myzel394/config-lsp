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
		entry := it.Value().(ast.SSHEntry)

		switch entry.(type) {
		case *ast.SSHOption:
			errs = append(errs, checkOption(entry.(*ast.SSHOption), false)...)
		case *ast.SSHMatchBlock:
			matchBlock := entry.(*ast.SSHMatchBlock)
			errs = append(errs, checkMatchBlock(matchBlock)...)
		}

	}

	return errs
}

func checkOption(
	option *ast.SSHOption,
	isInMatchBlock bool,
) []common.LSPError {
	errs := make([]common.LSPError, 0)

	if option.Key != nil {
		docOption, found := fields.Options[option.Key.Value]

		if !found {
			errs = append(errs, common.LSPError{
				Range: option.Key.LocationRange,
				Err:   errors.New(fmt.Sprintf("Unknown option: %s", option.Key.Value)),
			})

			return errs
		}

		if _, found := fields.MatchAllowedOptions[option.Key.Value]; !found && isInMatchBlock {
			errs = append(errs, common.LSPError{
				Range: option.Key.LocationRange,
				Err:   errors.New(fmt.Sprintf("Option '%s' is not allowed inside Match blocks", option.Key.Value)),
			})

			return errs
		}

		if option.OptionValue == nil {
			return errs
		}

		invalidValues := docOption.CheckIsValid(option.OptionValue.Value)

		errs = append(
			errs,
			utils.Map(
				invalidValues,
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

func checkMatchBlock(
	matchBlock *ast.SSHMatchBlock,
) []common.LSPError {
	errs := make([]common.LSPError, 0)

	matchOption := matchBlock.MatchEntry.OptionValue
	if matchOption != nil {
		invalidValues := fields.Options["Match"].CheckIsValid(matchOption.Value)

		errs = append(
			errs,
			utils.Map(
				invalidValues,
				func(invalidValue *docvalues.InvalidValue) common.LSPError {
					err := docvalues.LSPErrorFromInvalidValue(matchBlock.Start.Line, *invalidValue)
					err.ShiftCharacter(matchOption.Start.Character)

					return err
				},
			)...,
		)
	}

	it := matchBlock.Options.Iterator()

	for it.Next() {
		option := it.Value().(*ast.SSHOption)

		errs = append(errs, checkOption(option, true)...)
	}

	return errs
}
