package lsp

import (
	"config-lsp/handlers/hosts/tree"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

var documentParserMap = map[protocol.DocumentUri]*tree.HostsParser{}
