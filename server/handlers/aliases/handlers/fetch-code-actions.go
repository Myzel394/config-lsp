package handlers

import (
	"config-lsp/handlers/aliases"
	"config-lsp/handlers/aliases/ast"
	"config-lsp/handlers/aliases/commands"
	"fmt"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

type disabledExplanation struct {
	Reason string
}

func FetchCodeActions(
	d *aliases.AliasesDocument,
	params *protocol.CodeActionParams,
) []protocol.CodeAction {
	line := params.Range.Start.Line

	rawEntry, found := d.Parser.Aliases.Get(line)

	if !found {
		return nil
	}

	entry := rawEntry.(*ast.AliasEntry)

	if entry.Key != nil {
		address := fmt.Sprintf("%s@localhost.localdomain", entry.Key.Value)

		commandID := "aliases." + CodeActionSendTestMail
		command := protocol.Command{
			Title:   fmt.Sprintf("Send a test mail to %s", address),
			Command: string(commandID),
			Arguments: []any{
				CodeActionSendTestMailArgs{
					URI:  params.TextDocument.URI,
					User: entry.Key.Value,
				},
			},
		}
		codeAction := &protocol.CodeAction{
			Title:   fmt.Sprintf("Send a test mail to %s", address),
			Command: &command,
		}

		if !commands.CanSendTestMails() {
			codeAction.Disabled.Reason = "postfix is required to send test mails"
		}

		return []protocol.CodeAction{
			*codeAction,
		}
	}

	return nil
}
