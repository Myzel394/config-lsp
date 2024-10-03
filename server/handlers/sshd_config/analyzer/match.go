package analyzer

import (
	"config-lsp/common"
	docvalues "config-lsp/doc-values"
	"config-lsp/handlers/sshd_config/fields"
	"config-lsp/handlers/sshd_config/match-parser"
	"config-lsp/utils"
	"fmt"
	"strings"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func analyzeMatchBlocks(
	ctx *analyzerContext,
) {
	for matchBlock, options := range ctx.document.Indexes.AllOptionsPerName["Match"] {
		option := options[0]
		// Check if the match block has filled out all fields
		if matchBlock == nil || matchBlock.MatchValue == nil || len(matchBlock.MatchValue.Entries) == 0 {
			ctx.diagnostics = append(ctx.diagnostics, protocol.Diagnostic{
				Range:    option.ToLSPRange(),
				Message:  "A match expression is required",
				Severity: &common.SeverityError,
			})

			continue
		}

		for _, entry := range matchBlock.MatchValue.Entries {
			if entry.Values == nil {
				ctx.diagnostics = append(ctx.diagnostics, protocol.Diagnostic{
					Range:    entry.ToLSPRange(),
					Message:  fmt.Sprintf("A value for %s is required", entry.Criteria.Type),
					Severity: &common.SeverityError,
				})

				continue
			}

			analyzeMatchValuesContainsPositiveValue(ctx, entry.Values)

			for _, value := range entry.Values.Values {
				analyzeMatchValueNegation(ctx, value)
				analyzeMatchValueIsValid(ctx, value, entry.Criteria.Type)
			}
		}

		// Check if match blocks are not empty
		if matchBlock.Options.Size() == 0 {
			ctx.diagnostics = append(ctx.diagnostics, protocol.Diagnostic{
				Range:    option.ToLSPRange(),
				Message:  "This match block is empty",
				Severity: &common.SeverityInformation,
				Tags: []protocol.DiagnosticTag{
					protocol.DiagnosticTagUnnecessary,
				},
			})
		}
	}
}

func analyzeMatchValueNegation(
	ctx *analyzerContext,
	value *matchparser.MatchValue,
) {
	positionsAsList := utils.AllIndexes(value.Value.Raw, "!")
	positions := utils.SliceToMap(positionsAsList, struct{}{})

	delete(positions, 0)

	for position := range positions {
		ctx.diagnostics = append(ctx.diagnostics, protocol.Diagnostic{
			Range: common.LocationRange{
				Start: common.Location{
					Line:      value.Start.Line,
					Character: uint32(position) + value.Start.Character,
				},
				End: common.Location{
					Line:      value.End.Line,
					Character: uint32(position) + value.End.Character,
				},
			}.ToLSPRange(),
			Message:  "The negation operator (!) may only occur at the beginning of a value",
			Severity: &common.SeverityError,
		})
	}
}

func analyzeMatchValuesContainsPositiveValue(
	ctx *analyzerContext,
	values *matchparser.MatchValues,
) {
	if len(values.Values) == 0 {
		return
	}

	containsPositive := false

	for _, value := range values.Values {
		if !strings.HasPrefix(value.Value.Value, "!") {
			containsPositive = true
			break
		}
	}

	if !containsPositive {
		ctx.diagnostics = append(ctx.diagnostics, protocol.Diagnostic{
			Range:    values.LocationRange.ToLSPRange(),
			Message:  "At least one positive value is required. A negated match will never produce a positive result by itself",
			Severity: &common.SeverityError,
		})
	}
}

func analyzeMatchValueIsValid(
	ctx *analyzerContext,
	value *matchparser.MatchValue,
	criteria matchparser.MatchCriteriaType,
) {
	if value.Value.Raw == "" {
		return
	}

	docOption := fields.MatchValueFieldMap[criteria]
	invalidValues := docOption.DeprecatedCheckIsValid(value.Value.Raw)

	for _, invalidValue := range invalidValues {
		err := docvalues.LSPErrorFromInvalidValue(value.Start.Line, *invalidValue)
		err.ShiftCharacter(value.Start.Character)

		ctx.diagnostics = append(ctx.diagnostics, protocol.Diagnostic{
			Range:    err.Range.ToLSPRange(),
			Message:  err.Err.Error(),
			Severity: &common.SeverityError,
		})
	}
}
