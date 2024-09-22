package analyzer

import (
	"config-lsp/common"
	commonparser "config-lsp/common/parser"
	sshdconfig "config-lsp/handlers/sshd_config"
	"errors"
	"strings"
)

func analyzeQuotesAreValid(
	d *sshdconfig.SSHDDocument,
) []common.LSPError {
	errs := make([]common.LSPError, 0)

	for _, option := range d.Config.GetAllOptions() {
		errs = append(errs, checkIsUsingDoubleQuotes(option.Key.Value, option.Key.LocationRange)...)
		errs = append(errs, checkIsUsingDoubleQuotes(option.OptionValue.Value, option.OptionValue.LocationRange)...)

		errs = append(errs, checkQuotesAreClosed(option.Key.Value, option.Key.LocationRange)...)
		errs = append(errs, checkQuotesAreClosed(option.OptionValue.Value, option.OptionValue.LocationRange)...)
	}

	return errs
}

func checkIsUsingDoubleQuotes(
	value commonparser.ParsedString,
	valueRange common.LocationRange,
) []common.LSPError {
	singleQuotePosition := strings.Index(value.Raw, "'")

	if singleQuotePosition != -1 {
		return []common.LSPError{
			{
				Range: valueRange,
				Err:   errors.New("sshd_config does not support single quotes. Use double quotes (\") instead."),
			},
		}
	}

	return nil
}

func checkQuotesAreClosed(
	value commonparser.ParsedString,
	valueRange common.LocationRange,
) []common.LSPError {
	if strings.Count(value.Raw, "\"")%2 != 0 {
		return []common.LSPError{
			{
				Range: valueRange,
				Err:   errors.New("There are unclosed quotes here. Make sure all quotes are closed."),
			},
		}
	}

	return nil
}
