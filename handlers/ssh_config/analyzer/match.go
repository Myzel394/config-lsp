package analyzer

import (
	"config-lsp/common"
	sshconfig "config-lsp/handlers/ssh_config"
	"config-lsp/handlers/ssh_config/fields"
	matchparser "config-lsp/handlers/ssh_config/match-parser"
	"config-lsp/utils"
	"errors"
	"fmt"
)

func analyzeMatchBlocks(
	d *sshconfig.SSHDocument,
) []common.LSPError {
	errs := make([]common.LSPError, 0)

	for _, matchBlock := range d.GetAllMatchBlocks() {
		structureErrs := isMatchStructureValid(matchBlock.MatchValue)
		errs = append(errs, structureErrs...)

		if len(structureErrs) > 0 {
			continue
		}

		errs = append(errs, checkMatch(matchBlock.MatchValue)...)
	}

	return errs
}

func isMatchStructureValid(
	m *matchparser.Match,
) []common.LSPError {
	errs := make([]common.LSPError, 0)

	for _, entry := range m.Entries {
		if !utils.KeyExists(fields.MatchSingleOptionCriterias, entry.Criteria.Type) && entry.Value.Value == "" {
			errs = append(errs, common.LSPError{
				Range: entry.LocationRange,
				Err:   errors.New(fmt.Sprintf("Argument '%s' requires a value", entry.Criteria.Type)),
			})
		}
	}

	return errs
}

func checkMatch(
	m *matchparser.Match,
) []common.LSPError {
	errs := make([]common.LSPError, 0)

	// Check single options
	allEntries := m.FindEntries("all")
	if len(allEntries) > 1 {
		errs = append(errs, common.LSPError{
			Range: allEntries[1].LocationRange,
			Err:   errors.New("'all' may only be used once"),
		})
	}

	canonicalEntries := m.FindEntries("canonical")
	if len(canonicalEntries) > 1 {
		errs = append(errs, common.LSPError{
			Range: canonicalEntries[1].LocationRange,
			Err:   errors.New("'canonical' may only be used once"),
		})
	}

	finalEntries := m.FindEntries("final")
	if len(finalEntries) > 1 {
		errs = append(errs, common.LSPError{
			Range: finalEntries[1].LocationRange,
			Err:   errors.New("'final' may only be used once"),
		})
	}

	// Check the `all` argument
	if len(allEntries) == 1 {
		allEntry := allEntries[0]
		previousEntry := m.GetPreviousEntry(allEntry)

		if previousEntry != nil && !utils.KeyExists(fields.MatchAllArgumentAllowedPreviousOptions, previousEntry.Criteria.Type) {
			errs = append(errs, common.LSPError{
				Range: allEntry.LocationRange,
				Err:   errors.New("'all' should either be the first entry or immediately follow 'final' or 'canonical'"),
			})
		}
	}

	return errs
}
