package analyzer

import (
	"config-lsp/common"
	docvalues "config-lsp/doc-values"
	sshconfig "config-lsp/handlers/ssh_config"
	"config-lsp/handlers/ssh_config/ast"
	"config-lsp/handlers/ssh_config/fields"
	"config-lsp/utils"
	"errors"
	"fmt"
)

func analyzeStructureIsValid(
	d *sshconfig.SSHDocument,
) []common.LSPError {
	errs := make([]common.LSPError, 0)
	it := d.Config.Options.Iterator()

	for it.Next() {
		entry := it.Value().(ast.SSHEntry)

		switch entry.(type) {
		case *ast.SSHOption:
			errs = append(errs, checkOption(d, entry.(*ast.SSHOption), nil)...)
		case *ast.SSHMatchBlock:
			matchBlock := entry.(*ast.SSHMatchBlock)
			errs = append(errs, checkMatchBlock(d, matchBlock)...)
		case *ast.SSHHostBlock:
			hostBlock := entry.(*ast.SSHHostBlock)
			errs = append(errs, checkHostBlock(d, hostBlock)...)
		}

	}

	return errs
}

func checkOption(
	d *sshconfig.SSHDocument,
	option *ast.SSHOption,
	block ast.SSHBlock,
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

	// Check for values that are not allowed in Host blocks
	if block != nil && block.GetBlockType() == ast.SSHBlockTypeHost {
		if utils.KeyExists(fields.HostDisallowedOptions, option.Key.Key) {
			errs = append(errs, common.LSPError{
				Range: option.Key.LocationRange,
				Err:   errors.New(fmt.Sprintf("Option '%s' is not allowed in Host blocks", option.Key.Key)),
			})
		}
	}

	if option.OptionValue == nil || option.OptionValue.Value.Value == "" {
		errs = append(errs, common.LSPError{
			Range: option.Key.LocationRange,
			Err:   errors.New(fmt.Sprintf("Option '%s' requires a value", option.Key.Key)),
		})
	} else {
		errs = append(errs, checkIsUsingDoubleQuotes(option.OptionValue.Value, option.OptionValue.LocationRange)...)
		errs = append(errs, checkQuotesAreClosed(option.OptionValue.Value, option.OptionValue.LocationRange)...)
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
	d *sshconfig.SSHDocument,
	matchBlock *ast.SSHMatchBlock,
) []common.LSPError {
	errs := make([]common.LSPError, 0)

	it := matchBlock.Options.Iterator()
	for it.Next() {
		option := it.Value().(*ast.SSHOption)

		errs = append(errs, checkOption(d, option, matchBlock)...)
	}

	return errs
}

func checkHostBlock(
	d *sshconfig.SSHDocument,
	hostBlock *ast.SSHHostBlock,
) []common.LSPError {
	errs := make([]common.LSPError, 0)

	it := hostBlock.Options.Iterator()
	for it.Next() {
		option := it.Value().(*ast.SSHOption)

		errs = append(errs, checkOption(d, option, hostBlock)...)
	}

	return errs
}
