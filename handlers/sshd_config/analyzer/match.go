package analyzer

import (
	"config-lsp/common"
	sshdconfig "config-lsp/handlers/sshd_config"
	"config-lsp/handlers/sshd_config/match-parser"
	"config-lsp/utils"
	"errors"
	"fmt"
	"strings"
)

func analyzeMatchBlocks(
	d *sshdconfig.SSHDDocument,
) []common.LSPError {
	errs := make([]common.LSPError, 0)

	for matchBlock, options := range d.Indexes.AllOptionsPerName["Match"] {
		option := options[0]
		// Check if the match block has filled out all fields
		if matchBlock == nil || matchBlock.MatchValue == nil || len(matchBlock.MatchValue.Entries) == 0 {
			errs = append(errs, common.LSPError{
				Range: option.LocationRange,
				Err:   errors.New("A match expression is required"),
			})
			continue
		}

		for _, entry := range matchBlock.MatchValue.Entries {
			if entry.Values == nil {
				errs = append(errs, common.LSPError{
					Range: entry.LocationRange,
					Err:   errors.New(fmt.Sprintf("A value for %s is required", entry.Criteria.Type)),
				})
			} else {
				errs = append(errs, analyzeMatchValuesContainsPositiveValue(entry.Values)...)

				for _, value := range entry.Values.Values {
					errs = append(errs, analyzeMatchValueNegation(value)...)
				}
			}
		}

		// Check if match blocks are not empty
		if matchBlock.Options.Size() == 0 {
			errs = append(errs, common.LSPError{
				Range: option.LocationRange,
				Err:   errors.New("This match block is empty"),
			})
		}
	}

	return errs
}

func analyzeMatchValueNegation(
	value *matchparser.matchparser,
) []common.LSPError {
	errs := make([]common.LSPError, 0)

	positionsAsList := utils.AllIndexes(value.Value.Value, "!")
	positions := utils.SliceToMap(positionsAsList, struct{}{})

	delete(positions, 0)

	for position := range positions {
		errs = append(errs, common.LSPError{
			Range: common.LocationRange{
				Start: common.Location{
					Line:      value.Start.Line,
					Character: uint32(position) + value.Start.Character,
				},
				End: common.Location{
					Line:      value.End.Line,
					Character: uint32(position) + value.End.Character,
				},
			},
			Err: errors.New("The negation operator (!) may only occur at the beginning of a value"),
		})
	}

	return errs
}

func analyzeMatchValuesContainsPositiveValue(
	values *matchparser.MatchValues,
) []common.LSPError {
	if len(values.Values) == 0 {
		return nil
	}

	containsPositive := false

	for _, value := range values.Values {
		if !strings.HasPrefix(value.Value.Value, "!") {
			containsPositive = true
			break
		}
	}

	if !containsPositive {
		return []common.LSPError{
			{
				Range: values.LocationRange,
				Err:   errors.New("At least one positive value is required. A negated match will never produce a positive result by itself"),
			},
		}
	}

	return nil
}
