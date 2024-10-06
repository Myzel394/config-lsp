package handlers

import (
	sshconfig "config-lsp/handlers/ssh_config"
	"config-lsp/handlers/ssh_config/diagnostics"
	"fmt"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func FetchCodeActions(
	d *sshconfig.SSHDocument,
	params *protocol.CodeActionParams,
) []protocol.CodeAction {
	line := params.Range.Start.Line

	if unknownOption, found := d.Indexes.UnknownOptions[line]; found {
		var blockLine *uint32

		if unknownOption.Block != nil {
			blockLineValue := uint32(unknownOption.Block.GetLocation().Start.Line)
			blockLine = &blockLineValue
		}

		commandID := "sshconfig." + CodeActionAddToUnknown
		command := protocol.Command{
			Title:   fmt.Sprintf("Add %s to unknown options", unknownOption.Option.Key.Key),
			Command: string(commandID),
			Arguments: []any{
				codeActionAddToUnknownArgs{
					URI:        params.TextDocument.URI,
					OptionLine: unknownOption.Option.Start.Line,
					BlockLine:  blockLine,
				},
			},
		}
		kind := protocol.CodeActionKindQuickFix
		codeAction := &protocol.CodeAction{
			Title:   fmt.Sprintf("Add %s to unknown options", unknownOption.Option.Key.Key),
			Command: &command,
			Kind:    &kind,
			Diagnostics: []protocol.Diagnostic{
				diagnostics.GenerateUnknownOption(
					unknownOption.Option.Key.ToLSPRange(),
					unknownOption.Option.Key.Value.Value,
				),
			},
		}

		return []protocol.CodeAction{
			*codeAction,
		}
	}

	return nil
}
