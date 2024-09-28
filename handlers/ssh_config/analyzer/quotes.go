package analyzer

import (
	"config-lsp/common"
	commonparser "config-lsp/common/parser"
	sshconfig "config-lsp/handlers/ssh_config"
	"errors"
	"strings"
)

func analyzeQuotesAreValid(
	d *sshconfig.SSHDocument,
) []common.LSPError {
	errs := make([]common.LSPError, 0)

	for _, info := range d.Config.GetAllOptions() {
		errs = append(errs, checkIsUsingDoubleQuotes(info.Option.Key.Value, info.Option.Key.LocationRange)...)
		errs = append(errs, checkIsUsingDoubleQuotes(info.Option.OptionValue.Value, info.Option.OptionValue.LocationRange)...)

		errs = append(errs, checkQuotesAreClosed(info.Option.Key.Value, info.Option.Key.LocationRange)...)
		errs = append(errs, checkQuotesAreClosed(info.Option.OptionValue.Value, info.Option.OptionValue.LocationRange)...)
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
				Err:   errors.New("ssh_config does not support single quotes. Use double quotes (\") instead."),
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
