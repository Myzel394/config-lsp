package roothandler

import protocol "github.com/tliron/glsp/protocol_3_16"

var openedFiles = make(map[protocol.DocumentUri]struct{})
