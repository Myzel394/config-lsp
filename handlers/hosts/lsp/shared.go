package lsp

import (
	"config-lsp/handlers/hosts/handlers/analyzer"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

var documentParserMap = map[protocol.DocumentUri]*analyzer.HostsParser{}
