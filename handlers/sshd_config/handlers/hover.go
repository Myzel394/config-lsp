package handlers

import (
	"config-lsp/common"
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
	index common.IndexPosition,
) (*protocol.Hover, error) {
	var docValue *docvalues.DocumentationValue

	// Either root level or in the line of a match block
	if matchBlock == nil || matchBlock.Start.Line == line {
		val := fields.Options[option.Key.Key]
		docValue = &val
	} else {
		if _, found := fields.MatchAllowedOptions[option.Key.Key]; found {
			val := fields.Options[option.Key.Key]
			docValue = &val
		}
	}

	if option.Key.ContainsPosition(index) {
		if docValue != nil {
			contents := []string{
				"## " + option.Key.Key,
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

	if option.OptionValue != nil && option.OptionValue.ContainsPosition(index) {
		line := option.OptionValue.Value.Raw
		contents := docValue.DeprecatedFetchHoverInfo(
			line,
			uint32(option.OptionValue.Start.GetRelativeIndexPosition(index)),
		)

		return &protocol.Hover{
			Contents: strings.Join(contents, "\n"),
		}, nil
	}

	return nil, nil
}
