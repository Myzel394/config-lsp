package handlers

import (
	"config-lsp/common"
	docvalues "config-lsp/doc-values"
	"config-lsp/handlers/wireguard"
	"config-lsp/handlers/wireguard/ast"
	"config-lsp/handlers/wireguard/fields"
	"fmt"
	"strings"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func GetPropertyHoverInfo(
	d *wireguard.WGDocument,
	section ast.WGSection,
	property ast.WGProperty,
	index common.IndexPosition,
) (*protocol.Hover, error) {
	availableOptions, found := fields.OptionsHeaderMap[fields.CreateNormalizedName(section.Header.Name)]

	if !found {
		return nil, nil
	}

	option, found := availableOptions[fields.CreateNormalizedName(property.Key.Name)]

	if !found {
		return nil, nil
	}

	if property.Key.ContainsPosition(index) {
		return &protocol.Hover{
			Contents: protocol.MarkupContent{
				Kind:  protocol.MarkupKindMarkdown,
				Value: option.Documentation,
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

func GetSectionHoverInfo(
	d *wireguard.WGDocument,
	section ast.WGSection,
) (*protocol.Hover, error) {
	var docValue *docvalues.EnumString = nil

	switch section.Header.Name {
	case "Interface":
		docValue = &fields.HeaderInterfaceEnum
	case "Peer":
		docValue = &fields.HeaderPeerEnum
	}

	if docValue == nil {
		return nil, nil
	}

	return &protocol.Hover{
		Contents: protocol.MarkupContent{
			Kind: protocol.MarkupKindMarkdown,
			Value: fmt.Sprintf(
				"## [%s]\n\n%s",
				section.Header.Name,
				docValue.Documentation,
			),
		},
	}, nil
}
