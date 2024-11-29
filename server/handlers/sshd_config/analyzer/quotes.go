package analyzer

import (
	"config-lsp/common"
	commonparser "config-lsp/common/parser"
	"strings"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func analyzeQuotesAreValid(
	ctx *analyzerContext,
) {
	for _, option := range ctx.document.Config.GetAllOptions() {
		checkIsUsingDoubleQuotes(ctx, option.Key.Value, option.Key.LocationRange)
		checkIsUsingDoubleQuotes(ctx, option.OptionValue.Value, option.OptionValue.LocationRange)

		checkQuotesAreClosed(ctx, option.Key.Value, option.Key.LocationRange)
		checkQuotesAreClosed(ctx, option.OptionValue.Value, option.OptionValue.LocationRange)
	}
}

func checkIsUsingDoubleQuotes(
	ctx *analyzerContext,
	value commonparser.ParsedString,
	valueRange common.LocationRange,
) {
	if strings.HasPrefix(value.Raw, "'") && strings.HasSuffix(value.Raw, "'") {
		ctx.diagnostics = append(ctx.diagnostics, protocol.Diagnostic{
			Range:    valueRange.ToLSPRange(),
			Message:  "ssh_config does not support single quotes. Use double quotes (\") instead.",
			Severity: &common.SeverityError,
		})
	}
}

func checkQuotesAreClosed(
	ctx *analyzerContext,
	value commonparser.ParsedString,
	valueRange common.LocationRange,
) {
	if strings.Count(value.Raw, "\"")%2 != 0 {
		ctx.diagnostics = append(ctx.diagnostics, protocol.Diagnostic{
			Range:    valueRange.ToLSPRange(),
			Message:  "There are unclosed quotes here. Make sure all quotes are closed.",
			Severity: &common.SeverityError,
		})
	}
}
