package handlers

import (
	sshconfig "config-lsp/handlers/ssh_config"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func FetchCodeActions(
	d *sshconfig.SSHDocument,
	params *protocol.CodeActionParams,
) []protocol.CodeAction {
	if d.Indexes == nil {
		return nil
	}

	actions := getAddToUnknownCodeAction(d, params)
	actions = append(actions, getKeywordTypoFixes(d, params)...)

	return actions
}
