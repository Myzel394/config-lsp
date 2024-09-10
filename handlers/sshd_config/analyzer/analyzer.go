package analyzer

import (
	"config-lsp/handlers/sshd_config"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func Analyze(
	d *sshdconfig.SSHDocument,
) []protocol.Diagnostic {
	return nil
}
