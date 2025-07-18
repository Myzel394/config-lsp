package handlers

import (
	"config-lsp/common/formatting"
	"config-lsp/parsers/ini"
	"strings"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

var propertyTemplate = formatting.FormatTemplate("%s = /!'%s/!'")

func formatProperty(
	property ini.Property,
	options protocol.FormattingOptions,
) []protocol.TextEdit {
	key := strings.ToLower(property.Key.Name)

	var value string

	if property.Value != nil {
		value = property.Value.Value
	} else {
		value = ""
	}

	return []protocol.TextEdit{
		{
			Range:   property.ToLSPRange(),
			NewText: propertyTemplate.Format(options, key, value),
		},
	}
}
