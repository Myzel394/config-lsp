package handlers

import (
	bitcoinconf "config-lsp/handlers/bitcoin_conf"
	"config-lsp/handlers/bitcoin_conf/commands"
	"fmt"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func GetGenerateRPCAuthCodeActions(
	d *bitcoinconf.BTCDocument,
	params *protocol.CodeActionParams,
) []protocol.CodeAction {
	// Check if rpcuser and rpcpassword are set
	line := params.Range.Start.Line

	section := d.Config.FindSectionByLine(line)
	property := d.Config.FindPropertyByLine(line)

	if property != nil && (property.Key.Name == "rpcuser" || property.Key.Name == "rpcpassword") && commands.IsPythonAvailable() {
		_, userProperty := section.FindFirstPropertyByName("rpcuser")
		_, passwordProperty := section.FindFirstPropertyByName("rpcpassword")

		if userProperty != nil && passwordProperty != nil {
			// Generate RPC auth command
			commandID := "bitcoinconf." + CodeActionGenerateRPCAuth
			command := protocol.Command{
				Title:   fmt.Sprintf("Generate RPC Auth for '%s'", userProperty.Value.Value),
				Command: string(commandID),
				Arguments: []any{
					CodeActionGenerateRPCAuthArgs{
						URI:          params.TextDocument.URI,
						UsernameLine: userProperty.Start.Line,
						PasswordLine: passwordProperty.Start.Line,
					},
				},
			}

			return []protocol.CodeAction{
				{
					Title:   fmt.Sprintf("Generate RPC Auth for '%s'", userProperty.Value.Value),
					Command: &command,
				},
			}
		}

	}

	return nil
}
