package handlers

import (
	sshdconfig "config-lsp/handlers/sshd_config"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func FetchCodeActions(
	d *sshdconfig.SSHDDocument,
	params *protocol.CodeActionParams,
) []protocol.CodeAction {
	if d.Indexes == nil {
		return nil
	}

	actions := getKeywordTypoFixes(d, params)

	return actions
}
