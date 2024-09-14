package handlers

import (
	docvalues "config-lsp/doc-values"
	"config-lsp/handlers/sshd_config/ast"
	"config-lsp/handlers/sshd_config/fields"
	"strings"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func GetHoverInfoForOption(
	option *ast.SSHOption,
	matchBlock *ast.SSHMatchBlock,
	cursor uint32,
) (*protocol.Hover, error) {
	var docValue *docvalues.DocumentationValue

	if matchBlock == nil {
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
