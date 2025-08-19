package handlers

import (
	"config-lsp/common"
	bitcoinconf "config-lsp/handlers/bitcoin_conf"
	"config-lsp/handlers/bitcoin_conf/fields"
	"config-lsp/parsers/ini"
	"strings"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func GetPropertyHoverInfo(
	d *bitcoinconf.BTCDocument,
	property *ini.Property,
	index common.IndexPosition,
) (*protocol.Hover, error) {
	if property.Key == nil {
		return nil, nil
	}

	option, found := fields.Options[property.Key.Name]

	if !found {
		return nil, nil
	}

	if property.Key.ContainsPosition(index) {
		topic, _ := fields.TopicForOption[property.Key.Name]
		text := "## " + topic + "\n\n" + option.Documentation

		return &protocol.Hover{
			Contents: protocol.MarkupContent{
				Kind:  protocol.MarkupKindMarkdown,
				Value: text,
			},
		}, nil
	}

	if property.Value != nil && property.Value.ContainsPosition(index) {
		return &protocol.Hover{
			Contents: protocol.MarkupContent{
				Kind:  protocol.MarkupKindMarkdown,
				Value: strings.Join(option.GetTypeDescription(), "\n"),
			},
		}, nil
	}

	return nil, nil
}
