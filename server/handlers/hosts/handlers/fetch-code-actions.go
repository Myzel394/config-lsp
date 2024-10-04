package handlers

import (
	"config-lsp/handlers/hosts"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func GetInlineAliasesCodeAction(
	d hosts.HostsDocument,
	params *protocol.CodeActionParams,
) []protocol.CodeAction {
	line := params.Range.Start.Line

	if duplicateInfo, found := d.Indexes.DoubleIPs[line]; found {
		commandID := "hosts." + CodeActionInlineAliases
		command := protocol.Command{
			Title:   "Inline Aliases",
			Command: string(commandID),
			Arguments: []any{
				CodeActionInlineAliasesArgs{
					URI:      params.TextDocument.URI,
					FromLine: line,
					ToLine:   duplicateInfo.AlreadyFoundAt,
				},
			},
		}

		return []protocol.CodeAction{
			{
				Title:   "Inline Aliases",
				Command: &command,
			},
		}
	}

	return []protocol.CodeAction{}
}
