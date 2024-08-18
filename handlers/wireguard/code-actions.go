package wireguard

import (
	protocol "github.com/tliron/glsp/protocol_3_16"
)

type codeActionName string

const (
	codeActionGeneratePrivateKey codeActionName = "generatePrivateKey"
)

type codeActionGeneratePrivateKeyArgs struct {
	URI  protocol.DocumentUri
	Line uint32
}

func codeActionGeneratePrivateKeyArgsFromArguments(arguments map[string]any) codeActionGeneratePrivateKeyArgs {
	return codeActionGeneratePrivateKeyArgs{
		URI:  arguments["URI"].(protocol.DocumentUri),
		Line: uint32(arguments["Line"].(float64)),
	}
}

func (p wireguardProperty) getInsertRange(line uint32) protocol.Range {
	insertPosition := p.Separator.Location.End
	var length uint32 = 0

	if p.Value != nil {
		// Length of the value
		length = p.Value.Location.End - p.Value.Location.Start
	}

	return protocol.Range{
		Start: protocol.Position{
			Line:      line,
			Character: insertPosition,
		},
		End: protocol.Position{
			Line:      line,
			Character: insertPosition + length,
		},
	}
}

func (p *wireguardParser) runGeneratePrivateKey(args codeActionGeneratePrivateKeyArgs) (*protocol.ApplyWorkspaceEditParams, error) {
	privateKey, err := createNewPrivateKey()

	if err != nil {
		return &protocol.ApplyWorkspaceEditParams{}, err
	}

	section, property := p.getPropertyByLine(args.Line)

	if section == nil || property == nil {
		return nil, nil
	}

	label := "Generate Private Key"
	return &protocol.ApplyWorkspaceEditParams{
		Label: &label,
		Edit: protocol.WorkspaceEdit{
			Changes: map[protocol.DocumentUri][]protocol.TextEdit{
				args.URI: {
					{
						NewText: " " + privateKey,
						Range:   property.getInsertRange(args.Line),
					},
				},
			},
		},
	}, nil
}
