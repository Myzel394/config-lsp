package wireguard

import (
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentCompletion(context *glsp.Context, params *protocol.CompletionParams) (any, error) {
	parser := documentParserMap[params.TextDocument.URI]

	lineNumber := params.Position.Line

	section := parser.getBelongingSectionByLine(lineNumber)
	lineType := parser.getTypeByLine(lineNumber)

	switch lineType {
	case LineTypeComment:
		return nil, nil
	case LineTypeHeader:
		return parser.getRootCompletionsForEmptyLine(), nil
	case LineTypeEmpty:
		if section == nil {
			return parser.getRootCompletionsForEmptyLine(), nil
		}

		return section.getCompletionsForEmptyLine()
	case LineTypeProperty:
		if section == nil {
			return nil, nil
		}

		return section.getCompletionsForPropertyLine(lineNumber, params.Position.Character)
	}

	panic("TextDocumentCompletion: unexpected line type")
}
