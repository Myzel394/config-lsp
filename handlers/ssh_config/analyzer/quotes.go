package analyzer

import (
	"config-lsp/common"
	commonparser "config-lsp/common/parser"
	"config-lsp/utils"
	"strings"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func analyzeQuotesAreValid(
	ctx *analyzerContext,
) {
	for _, info := range ctx.document.Config.GetAllOptions() {
		checkIsUsingDoubleQuotes(ctx, info.Option.Key.Value, info.Option.Key.LocationRange)
		checkIsUsingDoubleQuotes(ctx, info.Option.OptionValue.Value, info.Option.OptionValue.LocationRange)

		checkQuotesAreClosed(ctx, info.Option.Key.Value, info.Option.Key.LocationRange)
		checkQuotesAreClosed(ctx, info.Option.OptionValue.Value, info.Option.OptionValue.LocationRange)
	}
}

func checkIsUsingDoubleQuotes(
	ctx *analyzerContext,
	value commonparser.ParsedString,
	valueRange common.LocationRange,
) {
	quoteRanges := utils.GetQuoteRanges(value.Raw)
	singleQuotePosition := strings.Index(value.Raw, "'")

	// Single quoe
	if singleQuotePosition != -1 && !quoteRanges.IsCharInside(singleQuotePosition) {
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
