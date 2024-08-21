package lsp

import (
	"config-lsp/handlers/wireguard/parser"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

var documentParserMap = map[protocol.DocumentUri]*parser.WireguardParser{}
