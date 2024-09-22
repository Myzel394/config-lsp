package sshdconfig

import (
	"config-lsp/handlers/sshd_config/ast"
	"config-lsp/handlers/sshd_config/indexes"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

type SSHDDocument struct {
	Config  *ast.SSHDConfig
	Indexes *indexes.SSHDIndexes
}

var DocumentParserMap = map[protocol.DocumentUri]*SSHDDocument{}
