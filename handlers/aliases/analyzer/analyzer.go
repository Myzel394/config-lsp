package analyzer

import (
	"config-lsp/handlers/hosts"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func Analyze(
	d *hosts.HostsDocument,
) []protocol.Diagnostic {
	return nil
}
