package hosts

import (
	"config-lsp/handlers/hosts/ast"
	"config-lsp/handlers/hosts/indexes"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

type HostsDocument struct {
	Parser  *ast.HostsParser
	Indexes *indexes.HostsIndexes
}

var DocumentParserMap = map[protocol.DocumentUri]*HostsDocument{}
