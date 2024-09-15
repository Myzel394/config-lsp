package handlers

import (
	sshdconfig "config-lsp/handlers/sshd_config"
	"config-lsp/handlers/sshd_config/fields"
	match_parser "config-lsp/handlers/sshd_config/fields/match-parser"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func getMatchCompletions(
	d *sshdconfig.SSHDocument,
	match *match_parser.Match,
	cursor uint32,
) ([]protocol.CompletionItem, error) {
	if len(match.Entries) == 0 {
		completions := getMatchCriteriaCompletions()
		completions = append(completions, getMatchAllKeywordCompletion())

		return completions, nil
	}

	entry := match.GetEntryByCursor(cursor)

	if entry == nil || entry.Criteria.IsCursorBetween(cursor) {
		return getMatchCriteriaCompletions(), nil
	}

	return getMatchValueCompletions(entry, cursor), nil
}

func getMatchCriteriaCompletions() []protocol.CompletionItem {
	kind := protocol.CompletionItemKindEnum

	return []protocol.CompletionItem{
		{
			Label: string(match_parser.MatchCriteriaTypeUser),
			Kind:  &kind,
		},
		{
			Label: string(match_parser.MatchCriteriaTypeGroup),
			Kind:  &kind,
		},
		{
			Label: string(match_parser.MatchCriteriaTypeHost),
			Kind:  &kind,
		},
		{
			Label: string(match_parser.MatchCriteriaTypeAddress),
			Kind:  &kind,
		},
		{
			Label: string(match_parser.MatchCriteriaTypeLocalAddress),
			Kind:  &kind,
		},
		{
			Label: string(match_parser.MatchCriteriaTypeLocalPort),
			Kind:  &kind,
		},
		{
			Label: string(match_parser.MatchCriteriaTypeRDomain),
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
	entry *match_parser.MatchEntry,
	cursor uint32,
) []protocol.CompletionItem {
	value := entry.GetValueByCursor(entry.End.Character)

	var line string
	var relativeCursor uint32

	if value != nil {
		line = value.Value
		relativeCursor = cursor - value.Start.Character
	} else {
		line = ""
		relativeCursor = 0
	}

	switch entry.Criteria.Type {
	case match_parser.MatchCriteriaTypeUser:
		return fields.MatchUserField.FetchCompletions(line, relativeCursor)
	case match_parser.MatchCriteriaTypeGroup:
		return fields.MatchGroupField.FetchCompletions(line, relativeCursor)
	case match_parser.MatchCriteriaTypeHost:
		return fields.MatchHostField.FetchCompletions(line, relativeCursor)
	case match_parser.MatchCriteriaTypeAddress:
		return fields.MatchAddressField.FetchCompletions(line, relativeCursor)
	case match_parser.MatchCriteriaTypeLocalAddress:
		return fields.MatchLocalAddressField.FetchCompletions(line, relativeCursor)
	case match_parser.MatchCriteriaTypeLocalPort:
		return fields.MatchLocalPortField.FetchCompletions(line, relativeCursor)
	case match_parser.MatchCriteriaTypeRDomain:
		return fields.MatchRDomainField.FetchCompletions(line, relativeCursor)
	}

	return nil
}
