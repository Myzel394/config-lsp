package handlers

import (
	bitcoinconf "config-lsp/handlers/bitcoin_conf"
	"config-lsp/handlers/bitcoin_conf/commands"
	"fmt"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

type CodeActionGenerateRPCAuthArgs struct {
	URI          protocol.DocumentUri
	UsernameLine uint32
	PasswordLine uint32
}

func CodeActionGenerateRPCAuthArgsFromArguments(arguments map[string]any) CodeActionGenerateRPCAuthArgs {
	return CodeActionGenerateRPCAuthArgs{
		URI:          arguments["URI"].(protocol.DocumentUri),
		UsernameLine: uint32(arguments["UsernameLine"].(float64)),
		PasswordLine: uint32(arguments["PasswordLine"].(float64)),
	}
}

func (args CodeActionGenerateRPCAuthArgs) RunCommand(d *bitcoinconf.BTCDocument) (*protocol.ApplyWorkspaceEditParams, error) {
	if !commands.IsPythonAvailable() {
		return nil, fmt.Errorf("python is not available, cannot generate RPC auth")
	}

	// Check if rpcuser and rpcpassword are set
	userProperty := d.Config.FindPropertyByLine(args.UsernameLine)
	passwordProperty := d.Config.FindPropertyByLine(args.PasswordLine)

	if userProperty == nil || passwordProperty == nil || userProperty.Key == nil || passwordProperty.Key == nil || userProperty.Key.Name != "rpcuser" || passwordProperty.Key.Name != "rpcpassword" {
		return nil, fmt.Errorf("rpcuser or rpcpassword not found at specified lines")
	}

	username := userProperty.Value.Value
	password := passwordProperty.Value.Value

	rpcAuth, err := commands.GenerateRPCAuth(username, password)

	if err != nil {
		return nil, fmt.Errorf("failed to generate RPC auth: %s", err.Error())
	}

	label := fmt.Sprintf("Generate RPC Auth for user '%s'", username)

	newPropertyText := fmt.Sprintf("rpcauth=%s", rpcAuth)

	// Remove the existing rpcuser and rpcpassword properties
	return &protocol.ApplyWorkspaceEditParams{
		Label: &label,
		Edit: protocol.WorkspaceEdit{
			Changes: map[protocol.DocumentUri][]protocol.TextEdit{
				args.URI: {
					{
						// Remove username property
						Range:   userProperty.IncludeNextLine().ToLSPRange(),
						NewText: "",
					},
					{
						// Remove password property
						// Insert rpcauth
						Range:   passwordProperty.ToLSPRange(),
						NewText: newPropertyText,
					},
				},
			},
		},
	}, nil
}
