package shared

import (
	"config-lsp/handlers/fstab/parser"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

var DocumentParserMap = map[protocol.DocumentUri]*parser.FstabParser{}
