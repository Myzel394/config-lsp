package sshconfig

import (
	"config-lsp/handlers/ssh_config/indexes"
	"config-lsp/handlers/ssh_config/ast"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

type SSHDocument struct {
	Config  *ast.SSHConfig
	Indexes *indexes.SSHIndexes
}

var DocumentParserMap = map[protocol.DocumentUri]*SSHDocument{}

