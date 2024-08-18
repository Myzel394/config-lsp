package wireguard

import (
	"strings"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentHover(
	context *glsp.Context,
	params *protocol.HoverParams,
) (*protocol.Hover, error) {
	parser := documentParserMap[params.TextDocument.URI]

	switch parser.getTypeByLine(params.Position.Line) {
	case LineTypeComment:
		return nil, nil
	case LineTypeEmpty:
		return nil, nil
	case LineTypeHeader:
		fallthrough
	case LineTypeProperty:
		documentation := parser.getHeaderInfo(params.Position.Line, params.Position.Character)

		hover := protocol.Hover{
			Contents: protocol.MarkupContent{
				Kind:  protocol.MarkupKindMarkdown,
				Value: strings.Join(documentation, "\n"),
			},
		}
		return &hover, nil
	}

	return nil, nil
}
