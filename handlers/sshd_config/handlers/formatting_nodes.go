package handlers

import (
	"config-lsp/common/formatting"
	"config-lsp/handlers/sshd_config/ast"
	matchparser "config-lsp/handlers/sshd_config/fields/match-parser"
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

func formatSSHDOption(
	option *ast.SSHDOption,
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

func formatSSHDMatchBlock(
	matchBlock *ast.SSHDMatchBlock,
	options protocol.FormattingOptions,
) []protocol.TextEdit {
	edits := make([]protocol.TextEdit, 0)

	edits = append(edits, protocol.TextEdit{
		Range: matchBlock.MatchEntry.ToLSPRange(),
		NewText: matchTemplate.Format(
			options,
			matchBlock.MatchEntry.Key.Key,
			formatMatchToString(matchBlock.MatchValue),
		),
	})

	it := matchBlock.Options.Iterator()
	for it.Next() {
		option := it.Value().(*ast.SSHDOption)

		edits = append(edits, formatSSHDOption(
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
			return formatMatchEntryToString(entry)
		},
	)

	return strings.Join(entriesAsStrings, " ")
}

func formatMatchEntryToString(
	entry *matchparser.MatchEntry,
) string {
	return fmt.Sprintf(
		"%s %s",
		string(entry.Criteria.Type),
		formatMatchValuesToString(entry.Values),
	)
}

func formatMatchValuesToString(
	values *matchparser.MatchValues,
) string {
	valuesAsStrings := utils.Map(
		values.Values,
		func(value *matchparser.MatchValue) string {
			return value.Value.Raw
		},
	)

	return strings.Join(valuesAsStrings, ",")
}
