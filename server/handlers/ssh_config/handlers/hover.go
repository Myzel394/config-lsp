package handlers

import (
	"config-lsp/common"
	"config-lsp/handlers/ssh_config/ast"
	"config-lsp/handlers/ssh_config/fields"
	"strings"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func GetHoverInfoForOption(
	option *ast.SSHOption,
	line uint32,
	index common.IndexPosition,
) (*protocol.Hover, error) {
	docValue, found := fields.Options[option.Key.Key]

	if !found {
		return nil, nil
	}

	if option.Key.ContainsPosition(index) {
		optionName := fields.FieldsNameFormattedMap[option.Key.Key]
		contents := []string{
			"## " + optionName,
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

	if option.OptionValue != nil && option.OptionValue.ContainsPosition(index) {
		line := option.OptionValue.Value.Raw
		contents := docValue.Value.DeprecatedFetchHoverInfo(
			line,
			uint32(option.OptionValue.Start.GetRelativeIndexPosition(index)),
		)

		return &protocol.Hover{
			Contents: strings.Join(contents, "\n"),
		}, nil
	}

	return nil, nil
}
