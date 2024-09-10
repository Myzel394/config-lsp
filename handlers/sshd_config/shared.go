package sshdconfig

import (
	"config-lsp/handlers/sshd_config/ast"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

type SSHDocument struct {
	Config  *ast.SSHConfig
}

var DocumentParserMap = map[protocol.DocumentUri]*SSHDocument{}

