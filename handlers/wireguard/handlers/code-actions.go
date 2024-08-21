package handlers

import (
	wgcommands "config-lsp/handlers/wireguard/commands"
	"config-lsp/handlers/wireguard/parser"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

type CodeActionName string

const (
	CodeActionGeneratePrivateKey   CodeActionName = "generatePrivateKey"
	CodeActionGeneratePresharedKey CodeActionName = "generatePresharedKey"
	CodeActionAddKeepalive         CodeActionName = "addKeepalive"
)

type CodeAction interface {
	RunCommand(*parser.WireguardParser) (*protocol.ApplyWorkspaceEditParams, error)
}

type CodeActionArgs interface{}

type CodeActionGeneratePrivateKeyArgs struct {
	URI  protocol.DocumentUri
	Line uint32
}

func CodeActionGeneratePrivateKeyArgsFromArguments(arguments map[string]any) CodeActionGeneratePrivateKeyArgs {
	return CodeActionGeneratePrivateKeyArgs{
		URI:  arguments["URI"].(protocol.DocumentUri),
		Line: uint32(arguments["Line"].(float64)),
	}
}

func (args CodeActionGeneratePrivateKeyArgs) RunCommand(p *parser.WireguardParser) (*protocol.ApplyWorkspaceEditParams, error) {
	privateKey, err := wgcommands.CreateNewPrivateKey()

	if err != nil {
		return &protocol.ApplyWorkspaceEditParams{}, err
	}

	section, property := p.GetPropertyByLine(args.Line)

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
						Range:   property.GetInsertRange(args.Line),
					},
				},
			},
		},
	}, nil
}

type CodeActionGeneratePresharedKeyArgs struct {
	URI  protocol.DocumentUri
	Line uint32
}

func CodeActionGeneratePresharedKeyArgsFromArguments(arguments map[string]any) CodeActionGeneratePresharedKeyArgs {
	return CodeActionGeneratePresharedKeyArgs{
		URI:  arguments["URI"].(protocol.DocumentUri),
		Line: uint32(arguments["Line"].(float64)),
	}
}

func (args CodeActionGeneratePresharedKeyArgs) RunCommand(p *parser.WireguardParser) (*protocol.ApplyWorkspaceEditParams, error) {
	presharedKey, err := wgcommands.CreatePresharedKey()

	if err != nil {
		return &protocol.ApplyWorkspaceEditParams{}, err
	}

	section, property := p.GetPropertyByLine(args.Line)

	if section == nil || property == nil {
		return nil, nil
	}

	label := "Generate Preshared Key"
	return &protocol.ApplyWorkspaceEditParams{
		Label: &label,
		Edit: protocol.WorkspaceEdit{
			Changes: map[protocol.DocumentUri][]protocol.TextEdit{
				args.URI: {
					{
						NewText: " " + presharedKey,
						Range:   property.GetInsertRange(args.Line),
					},
				},
			},
		},
	}, nil
}

type CodeActionAddKeepaliveArgs struct {
	URI          protocol.DocumentUri
	SectionIndex uint32
}

func CodeActionAddKeepaliveArgsFromArguments(arguments map[string]any) CodeActionAddKeepaliveArgs {
	return CodeActionAddKeepaliveArgs{
		URI:          arguments["URI"].(protocol.DocumentUri),
		SectionIndex: uint32(arguments["SectionIndex"].(float64)),
	}
}

func (args CodeActionAddKeepaliveArgs) RunCommand(p *parser.WireguardParser) (*protocol.ApplyWorkspaceEditParams, error) {
	section := p.Sections[args.SectionIndex]

	label := "Add PersistentKeepalive"
	return &protocol.ApplyWorkspaceEditParams{
		Label: &label,
		Edit: protocol.WorkspaceEdit{
			Changes: map[protocol.DocumentUri][]protocol.TextEdit{
				args.URI: {
					{
						NewText: "PersistentKeepalive = 25\n",
						Range: protocol.Range{
							Start: protocol.Position{
								Line:      section.EndLine + 1,
								Character: 0,
							},
							End: protocol.Position{
								Line:      section.EndLine + 1,
								Character: 0,
							},
						},
					},
				},
			},
		},
	}, nil
}
