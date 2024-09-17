package handlers

import (
	docvalues "config-lsp/doc-values"
	"config-lsp/handlers/sshd_config/ast"
	"config-lsp/handlers/sshd_config/fields"
	"strings"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func GetHoverInfoForOption(
	option *ast.SSHDOption,
	matchBlock *ast.SSHDMatchBlock,
	line uint32,
	cursor uint32,
) (*protocol.Hover, error) {
	var docValue *docvalues.DocumentationValue

	// Either root level or in the line of a match block
	if matchBlock == nil || matchBlock.Start.Line == line {
		val := fields.Options[option.Key.Value]
		docValue = &val
	} else {
		if _, found := fields.MatchAllowedOptions[option.Key.Value]; found {
			val := fields.Options[option.Key.Value]
			docValue = &val
		}
	}

	if cursor >= option.Key.Start.Character && cursor <= option.Key.End.Character {
		if docValue != nil {
			contents := []string{
				"## " + option.Key.Value,
				"",
			}
			contents = append(contents, docValue.Documentation)
			contents = append(contents, []string{
				"",
				"---",
				"",
			}...)
			contents = append(contents, []string{
				"### Type",
				"",
			}...)
			contents = append(contents, docValue.GetTypeDescription()...)

			return &protocol.Hover{
				Contents: strings.Join(contents, "\n"),
			}, nil
		}
	}

	if option.OptionValue != nil && cursor >= option.OptionValue.Start.Character && cursor <= option.OptionValue.End.Character {
		relativeCursor := cursor - option.OptionValue.Start.Character
		contents := docValue.FetchHoverInfo(option.OptionValue.Value, relativeCursor)

		return &protocol.Hover{
			Contents: strings.Join(contents, "\n"),
		}, nil
	}

	return nil, nil
}
