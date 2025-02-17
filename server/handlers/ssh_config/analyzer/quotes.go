package analyzer

import (
	"config-lsp/common"
	commonparser "config-lsp/common/parser"
	"config-lsp/utils"
	"strings"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func checkIsUsingDoubleQuotes(
	ctx *analyzerContext,
	value commonparser.ParsedString,
	valueRange common.LocationRange,
) {
	quoteRanges := utils.GetQuoteRanges(value.Raw)
	invertedRanges := quoteRanges.GetInvertedRanges(len(value.Raw))

	for _, rang := range invertedRanges {
		text := value.Raw[rang[0]:rang[1]]

		if strings.Contains(text, "'") {
			ctx.diagnostics = append(ctx.diagnostics, protocol.Diagnostic{
				Range:    valueRange.ToLSPRange(),
				Message:  "ssh_config does not support single quotes. Use double quotes (\") instead.",
				Severity: &common.SeverityError,
			})
		}
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
