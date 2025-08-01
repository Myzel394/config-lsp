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
	key := option.Key.Key

	// Either root level or in the line of a match block
	if matchBlock == nil || matchBlock.Start.Line == line {
		val := fields.Options[key]
		docValue = &val
	} else {
		if _, found := fields.MatchAllowedOptions[key]; found {
			val := fields.Options[key]
			docValue = &val
		}
	}

	if option.Key.ContainsPosition(index) {
		if docValue != nil {
			name := fields.FieldsNameFormattedMap[key]
			contents := []string{
				"## " + name,
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
		line := option.OptionValue.Value.Value
		// `DocumentationValue` only shows the documentation for `DeprecatedFetchHoverInfo`,
		// since the cursor here is now at the value, we use the `Value`'s method instead to
		// get the proper info.
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
