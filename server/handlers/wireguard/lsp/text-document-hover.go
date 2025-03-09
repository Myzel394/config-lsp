package lsp

import (
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentHover(
	context *glsp.Context,
	params *protocol.HoverParams,
) (*protocol.Hover, error) {
	/*
		p := documentParserMap[params.TextDocument.URI]

		switch p.GetTypeByLine(params.Position.Line) {
		case parser.LineTypeComment:
			return nil, nil
		case parser.LineTypeEmpty:
			return nil, nil
		case parser.LineTypeHeader:
			fallthrough
		case parser.LineTypeProperty:
			documentation := handlers.GetHoverContent(
				*p,
				params.Position.Line,
				params.Position.Character,
			)

			hover := protocol.Hover{
				Contents: protocol.MarkupContent{
					Kind:  protocol.MarkupKindMarkdown,
					Value: strings.Join(documentation, "\n"),
				},
			}
			return &hover, nil
		}
	*/

	return nil, nil
}
