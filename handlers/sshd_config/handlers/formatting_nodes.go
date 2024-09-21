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

func formatSSHDOption(
	option *ast.SSHDOption,
	options protocol.FormattingOptions,
) []protocol.TextEdit {
	template := formatting.FormatTemplate(
		"%s/t%s",
	)

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

	template := formatting.FormatTemplate(
		"%s/t%s",
	)
	edits = append(edits, protocol.TextEdit{
		Range:   matchBlock.ToLSPRange(),
		NewText: template.Format(options, matchBlock.MatchEntry.Key.Key, formatMatchToString(matchBlock.MatchValue, options)),
	})

	it := matchBlock.Options.Iterator()
	for it.Next() {
		option := it.Value().(*ast.SSHDOption)

		edits = append(edits, formatSSHDOption(option, options)...)
	}

	return edits
}

func formatMatchToString(
	match *matchparser.Match,
	options protocol.FormattingOptions,
) string {
	entriesAsStrings := utils.Map(
		match.Entries,
		func(entry *matchparser.MatchEntry) string {
			return formatMatchEntryToString(entry, options)
		},
	)

	return strings.Join(entriesAsStrings, " ")
}

func formatMatchEntryToString(
	entry *matchparser.MatchEntry,
	options protocol.FormattingOptions,
) string {
	return fmt.Sprintf(
		"%s %s",
		string(entry.Criteria.Type),
		formatMatchValuesToString(entry.Values, options),
	)
}

func formatMatchValuesToString(
	values *matchparser.MatchValues,
	options protocol.FormattingOptions,
) string {
	valuesAsStrings := utils.Map(
		values.Values,
		func(value *matchparser.MatchValue) string {
			return value.Value.Raw
		},
	)

	return strings.Join(valuesAsStrings, ",")
}
