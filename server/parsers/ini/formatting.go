package ini

import (
	"config-lsp/common/formatting"
	"strings"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

var propertyTemplate = formatting.FormatTemplate(
	"/!'%s/!'",
)
func FormatProperty(
	property *Property,
	options protocol.FormattingOptions,
) ([]protocol.TextEdit, error) {
	edits := make([]protocol.TextEdit, 0)

	if property.Key.Name == "" || property.Value == nil {
		return edits, nil
	}

	keyText := property.Key.Name
	valueText := strings.ReplaceAll(
			strings.ReplaceAll(
			property.Value.Value,
			"\"",
			"\\\"",
		),
		"'",
		"\\'",
	)

	newValueText := propertyTemplate.Format(options, valueText)
	newText := keyText + " = " + newValueText

	edits = append(edits, protocol.TextEdit{
		Range:   property.ToLSPRange(),
		NewText: newText,
	})

	return edits, nil
}
