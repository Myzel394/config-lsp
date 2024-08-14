package wireguard

import (
	protocol "github.com/tliron/glsp/protocol_3_16"
)

var documentParserMap = map[protocol.DocumentUri]*wireguardParser{}
