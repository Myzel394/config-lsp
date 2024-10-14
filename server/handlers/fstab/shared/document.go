package shared

import (
	"config-lsp/handlers/fstab/ast"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

type FstabDocument struct {
	Config *ast.FstabConfig
}

var DocumentParserMap = map[protocol.DocumentUri]*FstabDocument{}
