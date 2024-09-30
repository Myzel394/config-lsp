package handlers

import (
	"config-lsp/common/formatting"
	"config-lsp/handlers/ssh_config/ast"
	hostparser "config-lsp/handlers/ssh_config/host-parser"
	matchparser "config-lsp/handlers/ssh_config/match-parser"
	"config-lsp/utils"
	"fmt"
	"strings"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

var optionTemplate = formatting.FormatTemplate(
	"%s /!'%s/!'",
)
var matchTemplate = formatting.FormatTemplate(
	"%s %s",
)
var matchOptionTemplate = formatting.FormatTemplate(
	"    %s /!'%s/!'",
)

func formatSSHOption(
	option *ast.SSHOption,
	options protocol.FormattingOptions,
	template formatting.FormatTemplate,
) []protocol.TextEdit {
	var key string

	if option.Key != nil {
		key = option.Key.Key
	} else {
		key = ""
	}

	var value string

	if option.OptionValue != nil {
		value = option.OptionValue.Value.Raw
	} else {
		value = ""
	}

	return []protocol.TextEdit{
		{
			Range:   option.ToLSPRange(),
			NewText: template.Format(options, key, value),
		},
	}
}

func formatSSHMatchBlock(
	textRange protocol.Range,
	matchBlock *ast.SSHMatchBlock,
	options protocol.FormattingOptions,
) []protocol.TextEdit {
	edits := make([]protocol.TextEdit, 0)

	edits = append(edits, protocol.TextEdit{
		Range: matchBlock.GetEntryOption().ToLSPRange(),
		NewText: matchTemplate.Format(
			options,
			matchBlock.GetEntryOption().Key.Key,
			formatMatchToString(matchBlock.MatchValue),
		),
	})

	it := matchBlock.GetOptions().Iterator()
	for it.Next() {
		option := it.Value().(*ast.SSHOption)

		if !(option.Start.Line >= textRange.Start.Line && option.End.Line <= textRange.End.Line) {
			continue
		}

		edits = append(edits, formatSSHOption(
			option,
			options,
			matchOptionTemplate,
		)...)
	}

	return edits
}

func formatMatchToString(
	match *matchparser.Match,
) string {
	entriesAsStrings := utils.Map(
		match.Entries,
		func(entry *matchparser.MatchEntry) string {
			return fmt.Sprintf(
				"%s %s",
				string(entry.Criteria.Type),
				strings.Join(
					utils.Map(
						entry.Values.Values,
						func(value *matchparser.MatchValue) string {
							return value.Value.Raw
						},
					),
					",",
				),
			)
		},
	)

	return strings.Join(entriesAsStrings, " ")
}

func formatSSHHostBlock(
	textRange protocol.Range,
	hostBlock *ast.SSHHostBlock,
	options protocol.FormattingOptions,
) []protocol.TextEdit {
	edits := make([]protocol.TextEdit, 0)

	edits = append(edits, protocol.TextEdit{
		Range: hostBlock.GetEntryOption().ToLSPRange(),
		NewText: matchTemplate.Format(
			options,
			hostBlock.GetEntryOption().Key.Key,
			formatHostToString(hostBlock.HostValue),
		),
	})

	it := hostBlock.GetOptions().Iterator()
	for it.Next() {
		option := it.Value().(*ast.SSHOption)

		if !(option.Start.Line >= textRange.Start.Line && option.End.Line <= textRange.End.Line) {
			continue
		}

		edits = append(edits, formatSSHOption(
			option,
			options,
			matchOptionTemplate,
		)...)
	}

	return edits
}

func formatHostToString(
	host *hostparser.Host,
) string {
	return strings.Join(
		utils.Map(
			host.Hosts,
			func(host *hostparser.HostValue) string {
				return host.Value.Raw
			},
		),
		" ",
	)
}
