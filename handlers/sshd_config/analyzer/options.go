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

func analyzeStructureIsValid(
	d *sshdconfig.SSHDocument,
) []common.LSPError {
	errs := make([]common.LSPError, 0)
	it := d.Config.Options.Iterator()

	for it.Next() {
		entry := it.Value().(ast.SSHDEntry)

		switch entry.(type) {
		case *ast.SSHDOption:
			errs = append(errs, checkOption(entry.(*ast.SSHDOption), false)...)
		case *ast.SSHDMatchBlock:
			matchBlock := entry.(*ast.SSHDMatchBlock)
			errs = append(errs, checkMatchBlock(matchBlock)...)
		}

	}

	return errs
}

func checkOption(
	option *ast.SSHDOption,
	isInMatchBlock bool,
) []common.LSPError {
	errs := make([]common.LSPError, 0)

	if option.Key == nil {
		return errs
	}

	errs = append(errs, checkIsUsingDoubleQuotes(option.Key.Value, option.Key.LocationRange)...)
	errs = append(errs, checkQuotesAreClosed(option.Key.Value, option.Key.LocationRange)...)

	docOption, found := fields.Options[option.Key.Key]

	if !found {
		errs = append(errs, common.LSPError{
			Range: option.Key.LocationRange,
			Err:   errors.New(fmt.Sprintf("Unknown option: %s", option.Key.Key)),
		})

		return errs
	}

	if _, found := fields.MatchAllowedOptions[option.Key.Key]; !found && isInMatchBlock {
		errs = append(errs, common.LSPError{
			Range: option.Key.LocationRange,
			Err:   errors.New(fmt.Sprintf("Option '%s' is not allowed inside Match blocks", option.Key.Key)),
		})

		return errs
	}

	if option.OptionValue == nil || option.OptionValue.Value.Value == "" {
		errs = append(errs, common.LSPError{
			Range: option.Key.LocationRange,
			Err:   errors.New(fmt.Sprintf("Option '%s' requires a value", option.Key.Key)),
		})
	} else {
		errs = append(errs, checkIsUsingDoubleQuotes(option.OptionValue.Value, option.OptionValue.LocationRange)...)
		errs = append(errs, checkQuotesAreClosed(option.OptionValue.Value, option.OptionValue.LocationRange)...)

		invalidValues := docOption.CheckIsValid(option.OptionValue.Value.Value)

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

	if option.Separator == nil || option.Separator.Value.Value == "" {
		errs = append(errs, common.LSPError{
			Range: option.Key.LocationRange,
			Err:   errors.New(fmt.Sprintf("There should be a separator between an option and its value")),
		})
	} else {
		errs = append(errs, checkIsUsingDoubleQuotes(option.Separator.Value, option.Separator.LocationRange)...)
		errs = append(errs, checkQuotesAreClosed(option.Separator.Value, option.Separator.LocationRange)...)
	}

	return errs
}

func checkMatchBlock(
	matchBlock *ast.SSHDMatchBlock,
) []common.LSPError {
	errs := make([]common.LSPError, 0)

	matchOption := matchBlock.MatchOption.OptionValue
	if matchOption != nil {
		invalidValues := fields.Options["Match"].CheckIsValid(matchOption.Value.Value)

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
		option := it.Value().(*ast.SSHDOption)

		errs = append(errs, checkOption(option, true)...)
	}

	return errs
}
