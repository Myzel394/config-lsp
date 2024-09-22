package analyzer

import (
	sshconfig "config-lsp/handlers/ssh_config"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func Analyze(
	d *sshconfig.SSHDocument,
) []protocol.Diagnostic {
	return nil
}
