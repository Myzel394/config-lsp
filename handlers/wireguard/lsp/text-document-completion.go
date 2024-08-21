package lsp

import (
	"config-lsp/handlers/wireguard/handlers"
	"config-lsp/handlers/wireguard/parser"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentCompletion(context *glsp.Context, params *protocol.CompletionParams) (any, error) {
	p := documentParserMap[params.TextDocument.URI]

	lineNumber := params.Position.Line

	section := p.GetBelongingSectionByLine(lineNumber)
	lineType := p.GetTypeByLine(lineNumber)

	switch lineType {
	case parser.LineTypeComment:
		return nil, nil
	case parser.LineTypeHeader:
		return handlers.GetRootCompletionsForEmptyLine(*p)
	case parser.LineTypeEmpty:
		if section.Name == nil {
			// Root completions
			return handlers.GetRootCompletionsForEmptyLine(*p)
		}

		return handlers.GetCompletionsForSectionEmptyLine(*section)
	case parser.LineTypeProperty:
		completions, err := handlers.GetCompletionsForSectionPropertyLine(*section, lineNumber, params.Position.Character)

		if completions == nil && err != nil {
			switch err.(type) {
			// Ignore
			case parser.PropertyNotFullyTypedError:
				return handlers.GetCompletionsForSectionEmptyLine(*section)
			default:
				return nil, err
			}
		}

		return completions, nil
	}

	panic("TextDocumentCompletion: unexpected line type")
}
