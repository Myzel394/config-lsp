package handlers

import (
	sshdconfig "config-lsp/handlers/sshd_config"
	"config-lsp/utils"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func FetchCodeActions(
	d *sshdconfig.SSHDDocument,
	params *protocol.CodeActionParams,
) []protocol.CodeAction {
	if utils.BlockUntilIndexesNotNil(d.Indexes) == false {
		return nil
	}

	actions := getKeywordTypoFixes(d, params)

	return actions
}
