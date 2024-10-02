package handlers

import (
	"config-lsp/common/formatting"
	"config-lsp/handlers/ssh_config/ast"
	"config-lsp/handlers/ssh_config/fields"
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
var blockOptionTemplate = formatting.FormatTemplate(
	"    %s /!'%s/!'",
)
var matchTemplate = formatting.FormatTemplate(
	"%s %s",
)

func formatOption(
	option *ast.SSHOption,
	block ast.SSHBlock,
	options protocol.FormattingOptions,
) []protocol.TextEdit {
	if option.Key == nil || option.OptionValue == nil {
		return nil
	}

	edits := make([]protocol.TextEdit, 0)

	if option.Key.Key == hostOption {
		edits = append(edits, formatHostBlock(block.(*ast.SSHHostBlock), options)...)
	} else if option.Key.Key == matchOption {
		edits = append(edits, formatMatchBlock(block.(*ast.SSHMatchBlock), options)...)
	} else {
		var template formatting.FormatTemplate

		if block == nil {
			template = optionTemplate
		} else if block.GetBlockType() == ast.SSHBlockTypeMatch {
			template = blockOptionTemplate
		} else if block.GetBlockType() == ast.SSHBlockTypeHost {
			template = blockOptionTemplate
		}

		edits = append(edits, formatSSHOption(option, options, template)...)
	}

	return edits
}

func formatHostBlock(
	hostBlock *ast.SSHHostBlock,
	options protocol.FormattingOptions,
) []protocol.TextEdit {
	if hostBlock.HostValue == nil || hostBlock.HostValue.Hosts == nil {
		return nil
	}

	edits := make([]protocol.TextEdit, 0)

	key := fields.FieldsNameFormattedMap[hostBlock.GetEntryOption().Key.Key]
	edits = append(edits, protocol.TextEdit{
		Range: hostBlock.GetEntryOption().ToLSPRange(),
		NewText: matchTemplate.Format(
			options,
			key,
			formatHostToString(hostBlock.HostValue),
		),
	})

	return edits
}

func formatMatchBlock(
	matchBlock *ast.SSHMatchBlock,
	options protocol.FormattingOptions,
) []protocol.TextEdit {
	if matchBlock.MatchValue == nil || matchBlock.MatchValue.Entries == nil {
		return nil
	}

	edits := make([]protocol.TextEdit, 0)

	key := fields.FieldsNameFormattedMap[matchBlock.GetEntryOption().Key.Key]
	edits = append(edits, protocol.TextEdit{
		Range: matchBlock.GetEntryOption().ToLSPRange(),
		NewText: matchTemplate.Format(
			options,
			key,
			formatMatchToString(matchBlock.MatchValue),
		),
	})

	return edits
}

func formatSSHOption(
	option *ast.SSHOption,
	options protocol.FormattingOptions,
	template formatting.FormatTemplate,
) []protocol.TextEdit {
	var key string

	if option.Key != nil {
		if optionName, found := fields.FieldsNameFormattedMap[option.Key.Key]; found {
			key = optionName
		} else {
			key = option.Key.Value.Raw
		}
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
