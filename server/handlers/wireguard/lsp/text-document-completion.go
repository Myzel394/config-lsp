package lsp

import (
	"config-lsp/handlers/wireguard"
	"config-lsp/handlers/wireguard/handlers"
	"regexp"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

var headerPattern = regexp.MustCompile(`^\s*\[(\w+)?]?`)

func TextDocumentCompletion(context *glsp.Context, params *protocol.CompletionParams) (any, error) {
	d := wireguard.DocumentParserMap[params.TextDocument.URI]

	return handlers.SuggestCompletions(d, params)
}
