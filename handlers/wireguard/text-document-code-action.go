package wireguard

import (
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentCodeAction(context *glsp.Context, params *protocol.CodeActionParams) ([]protocol.CodeAction, error) {
	line := params.Range.Start.Line
	parser := documentParserMap[params.TextDocument.URI]

	section, property := parser.getPropertyByLine(line)

	if section == nil || property == nil || property.Separator == nil {
		return nil, nil
	}

	switch property.Key.Name {
	case "PrivateKey":
		if !areWireguardToolsAvailable() {
			return nil, nil
		}

		commandID := "wireguard." + codeActionGeneratePrivateKey
		command := protocol.Command{
			Title:   "Generate Private Key",
			Command: string(commandID),
			Arguments: []any{
				codeActionGeneratePrivateKeyArgs{
					URI:  params.TextDocument.URI,
					Line: line,
				},
			},
		}

		return []protocol.CodeAction{
			{
				Title:   "Generate Private Key",
				Command: &command,
			},
		}, nil
	}

	return nil, nil
}
