package wireguard

import (
	"config-lsp/handlers/wireguard/ast"
	"config-lsp/handlers/wireguard/indexes"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

type WGDocument struct {
	Config  *ast.WGConfig
	Indexes *indexes.WGIndexes
}

var DocumentParserMap = map[protocol.DocumentUri]*WGDocument{}
