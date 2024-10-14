package handlers

import (
	"config-lsp/common"
	sshconfig "config-lsp/handlers/ssh_config"
	"config-lsp/handlers/ssh_config/fields"
	matchparser "config-lsp/handlers/ssh_config/match-parser"
	"config-lsp/utils"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func getMatchCompletions(
	d *sshconfig.SSHDocument,
	cursor common.CursorPosition,
	match *matchparser.Match,
) []protocol.CompletionItem {
	if match == nil || len(match.Entries) == 0 {
		completions := getMatchCriteriaCompletions()
		completions = append(completions, getMatchAllKeywordCompletion())

		return completions
	}

	entry := match.GetEntryAtPosition(cursor)

	if entry == nil || entry.Criteria.ContainsPosition(cursor) {
		completions := getMatchCriteriaCompletions()

		var showAllArgument = true

		previousEntry := match.GetPreviousEntryFromCursor(cursor)

		if previousEntry != nil && !utils.KeyExists(fields.MatchAllArgumentAllowedPreviousOptions, previousEntry.Criteria.Type) {
			showAllArgument = false
		}

		if showAllArgument {
			completions = append(completions, getMatchAllKeywordCompletion())
		}

		return completions
	}

	return getMatchValueCompletions(entry, cursor)
}

func getMatchCriteriaCompletions() []protocol.CompletionItem {
	kind := protocol.CompletionItemKindEnum

	return []protocol.CompletionItem{
		{
			Label: string(matchparser.MatchCriteriaTypeCanonical),
			Kind:  &kind,
		},
		{
			Label: string(matchparser.MatchCriteriaTypeFinal),
			Kind:  &kind,
		},
		{
			Label: string(matchparser.MatchCriteriaTypeExec),
			Kind:  &kind,
		},
		{
			Label: string(matchparser.MatchCriteriaTypeLocalNetwork),
			Kind:  &kind,
		},
		{
			Label: string(matchparser.MatchCriteriaTypeHost),
			Kind:  &kind,
		},
		{
			Label: string(matchparser.MatchCriteriaTypeOriginalHost),
			Kind:  &kind,
		},
		{
			Label: string(matchparser.MatchCriteriaTypeTagged),
			Kind:  &kind,
		},
		{
			Label: string(matchparser.MatchCriteriaTypeUser),
			Kind:  &kind,
		},
		{
			Label: string(matchparser.MatchCriteriaTypeLocalUser),
			Kind:  &kind,
		},
	}
}

func getMatchAllKeywordCompletion() protocol.CompletionItem {
	kind := protocol.CompletionItemKindKeyword

	return protocol.CompletionItem{
		Label: "all",
		Kind:  &kind,
	}
}

func getMatchValueCompletions(
	entry *matchparser.MatchEntry,
	cursor common.CursorPosition,
) []protocol.CompletionItem {
	value := entry.GetValueAtPosition(cursor)

	var line string
	var relativeCursor uint32

	if value != nil {
		line = value.Value.Raw
		relativeCursor = common.DeprecatedImprovedCursorToIndex(
			cursor,
			line,
			value.Start.Character,
		)
	} else {
		line = ""
		relativeCursor = 0
	}

	switch entry.Criteria.Type {
	case matchparser.MatchCriteriaTypeExec:
		return fields.MatchExecField.DeprecatedFetchCompletions(line, relativeCursor)
	case matchparser.MatchCriteriaTypeLocalNetwork:
		return fields.MatchLocalNetworkField.DeprecatedFetchCompletions(line, relativeCursor)
	case matchparser.MatchCriteriaTypeHost:
		return fields.MatchHostField.DeprecatedFetchCompletions(line, relativeCursor)
	case matchparser.MatchCriteriaTypeOriginalHost:
		return fields.MatchOriginalHostField.DeprecatedFetchCompletions(line, relativeCursor)
	case matchparser.MatchCriteriaTypeTagged:
		return fields.MatchTypeTaggedField.DeprecatedFetchCompletions(line, relativeCursor)
	case matchparser.MatchCriteriaTypeUser:
		return fields.MatchUserField.DeprecatedFetchCompletions(line, relativeCursor)
	case matchparser.MatchCriteriaTypeLocalUser:
		return fields.MatchTypeLocalUserField.DeprecatedFetchCompletions(line, relativeCursor)
	}

	return nil
}
