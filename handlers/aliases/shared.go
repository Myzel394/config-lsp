package aliases

import (
	"config-lsp/handlers/aliases/ast"
	"config-lsp/handlers/aliases/indexes"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

type HostsDocument struct {
	Parser  *ast.AliasesParser
	Indexes *indexes.AliasesIndexes
}

var DocumentParserMap = map[protocol.DocumentUri]*HostsDocument{}
