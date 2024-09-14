package sshdconfig

import (
	"config-lsp/handlers/sshd_config/ast"
	"config-lsp/handlers/sshd_config/indexes"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

type SSHDocument struct {
	Config  *ast.SSHConfig
	Indexes *indexes.SSHIndexes
}

var DocumentParserMap = map[protocol.DocumentUri]*SSHDocument{}
