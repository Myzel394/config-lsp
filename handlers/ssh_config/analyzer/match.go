package analyzer

import (
	"config-lsp/common"
	"config-lsp/handlers/ssh_config/fields"
	matchparser "config-lsp/handlers/ssh_config/match-parser"
	"config-lsp/utils"
	"fmt"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func analyzeMatchBlocks(
	ctx *analyzerContext,
) {
	for _, matchBlock := range ctx.document.GetAllMatchBlocks() {
		isValid := isMatchStructureValid(ctx, matchBlock.MatchValue)

		if !isValid {
			continue
		}

		checkMatch(ctx, matchBlock.MatchValue)
	}
}

func isMatchStructureValid(
	ctx *analyzerContext,
	m *matchparser.Match,
) bool {
	isValid := true

	for _, entry := range m.Entries {
		if !utils.KeyExists(fields.MatchSingleOptionCriterias, entry.Criteria.Type) && entry.Value.Value == "" {
			ctx.diagnostics = append(ctx.diagnostics, protocol.Diagnostic{
				Range:    entry.LocationRange.ToLSPRange(),
				Message:  fmt.Sprintf("Argument '%s' requires a value", entry.Criteria.Type),
				Severity: &common.SeverityError,
			})

			isValid = false
		}
	}

	return isValid
}

func checkMatch(
	ctx *analyzerContext,
	m *matchparser.Match,
) {
	// Check single options
	allEntries := m.FindEntries("all")
	if len(allEntries) > 1 {
		ctx.diagnostics = append(ctx.diagnostics, protocol.Diagnostic{
			Range:    allEntries[1].LocationRange.ToLSPRange(),
			Message:  "'all' may only be used once",
			Severity: &common.SeverityError,
		})
	}

	canonicalEntries := m.FindEntries("canonical")
	if len(canonicalEntries) > 1 {
		ctx.diagnostics = append(ctx.diagnostics, protocol.Diagnostic{
			Range:    canonicalEntries[1].LocationRange.ToLSPRange(),
			Message:  "'canonical' may only be used once",
			Severity: &common.SeverityError,
		})
	}

	finalEntries := m.FindEntries("final")
	if len(finalEntries) > 1 {
		ctx.diagnostics = append(ctx.diagnostics, protocol.Diagnostic{
			Range:    finalEntries[1].LocationRange.ToLSPRange(),
			Message:  "'final' may only be used once",
			Severity: &common.SeverityError,
		})
	}

	// Check the `all` argument
	if len(allEntries) == 1 {
		allEntry := allEntries[0]
		previousEntry := m.GetPreviousEntry(allEntry)

		if previousEntry != nil && !utils.KeyExists(fields.MatchAllArgumentAllowedPreviousOptions, previousEntry.Criteria.Type) {
			ctx.diagnostics = append(ctx.diagnostics, protocol.Diagnostic{
				Range:    allEntry.LocationRange.ToLSPRange(),
				Message:  "'all' should either be the first entry or immediately follow 'final' or 'canonical'",
				Severity: &common.SeverityError,
			})
		}
	}
}
