package handlers

import (
	"config-lsp/common"
	sshdconfig "config-lsp/handlers/sshd_config"
	"config-lsp/handlers/sshd_config/fields"
	"config-lsp/handlers/sshd_config/match-parser"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func getMatchCompletions(
	d *sshdconfig.SSHDDocument,
	cursor common.CursorPosition,
	match *matchparser.Match,
) ([]protocol.CompletionItem, error) {
	if match == nil || len(match.Entries) == 0 {
		completions := getMatchCriteriaCompletions()
		completions = append(completions, getMatchAllKeywordCompletion())

		return completions, nil
	}

	entry := match.GetEntryAtPosition(cursor)

	if entry == nil || entry.Criteria.ContainsPosition(cursor) {
		return getMatchCriteriaCompletions(), nil
	}

	return getMatchValueCompletions(entry, cursor), nil
}

func getMatchCriteriaCompletions() []protocol.CompletionItem {
	kind := protocol.CompletionItemKindEnum

	return []protocol.CompletionItem{
		{
			Label: string(matchparser.MatchCriteriaTypeUser),
			Kind:  &kind,
		},
		{
			Label: string(matchparser.MatchCriteriaTypeGroup),
			Kind:  &kind,
		},
		{
			Label: string(matchparser.MatchCriteriaTypeHost),
			Kind:  &kind,
		},
		{
			Label: string(matchparser.MatchCriteriaTypeAddress),
			Kind:  &kind,
		},
		{
			Label: string(matchparser.MatchCriteriaTypeLocalAddress),
			Kind:  &kind,
		},
		{
			Label: string(matchparser.MatchCriteriaTypeLocalPort),
			Kind:  &kind,
		},
		{
			Label: string(matchparser.MatchCriteriaTypeRDomain),
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
	case matchparser.MatchCriteriaTypeUser:
		return fields.MatchUserField.DeprecatedFetchCompletions(line, relativeCursor)
	case matchparser.MatchCriteriaTypeGroup:
		return fields.MatchGroupField.DeprecatedFetchCompletions(line, relativeCursor)
	case matchparser.MatchCriteriaTypeHost:
		return fields.MatchHostField.DeprecatedFetchCompletions(line, relativeCursor)
	case matchparser.MatchCriteriaTypeAddress:
		return fields.MatchAddressField.DeprecatedFetchCompletions(line, relativeCursor)
	case matchparser.MatchCriteriaTypeLocalAddress:
		return fields.MatchLocalAddressField.DeprecatedFetchCompletions(line, relativeCursor)
	case matchparser.MatchCriteriaTypeLocalPort:
		return fields.MatchLocalPortField.DeprecatedFetchCompletions(line, relativeCursor)
	case matchparser.MatchCriteriaTypeRDomain:
		return fields.MatchRDomainField.DeprecatedFetchCompletions(line, relativeCursor)
	}

	return nil
}
