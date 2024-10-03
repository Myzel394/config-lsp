package handlers

import (
	"config-lsp/handlers/aliases"
	"config-lsp/handlers/aliases/commands"
	"fmt"
	"time"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

type CodeActionName string

const (
	CodeActionSendTestMail CodeActionName = "sendTestMail"
)

type CodeAction interface {
	RunCommand(*aliases.AliasesDocument) (*protocol.ApplyWorkspaceEditParams, error)
}

type CodeActionArgs interface{}

type CodeActionSendTestMailArgs struct {
	URI  protocol.DocumentUri
	User string
}

func CodeActionSendTestMailArgsFromArguments(arguments map[string]interface{}) CodeActionSendTestMailArgs {
	return CodeActionSendTestMailArgs{
		URI:  arguments["URI"].(protocol.DocumentUri),
		User: arguments["User"].(string),
	}
}

func (args CodeActionSendTestMailArgs) RunCommand(d *aliases.AliasesDocument) (*protocol.ApplyWorkspaceEditParams, error) {
	content := fmt.Sprintf(
		`Subject: Test mail from %s

This is a test mail from config-lsp.
It is intended for the user %s.
`,
		time.Now().Format(time.RFC1123),
		args.User,
	)

	address := fmt.Sprintf("%s@localhost.localdomain", args.User)
	commands.SendTestMail(address, content)

	return &protocol.ApplyWorkspaceEditParams{}, nil
}
