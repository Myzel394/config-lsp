package bitcoinconf

import (
	"config-lsp/handlers/bitcoin_conf/ast"
	"config-lsp/handlers/bitcoin_conf/indexes"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

type BTCDocument struct {
	Config  *ast.BTCConfig
	Indexes *indexes.BTCIndexes
}

var DocumentParserMap = map[protocol.DocumentUri]*BTCDocument{}
