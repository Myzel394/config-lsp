package handlers

import (
	"config-lsp/handlers/wireguard"
	"config-lsp/handlers/wireguard/ast"
	wgcommands "config-lsp/handlers/wireguard/commands"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

type CodeActionName string

const (
	CodeActionGeneratePrivateKey   CodeActionName = "generatePrivateKey"
	CodeActionGeneratePresharedKey CodeActionName = "generatePresharedKey"
)

type CodeAction interface {
	RunCommand(*ast.WGConfig) (*protocol.ApplyWorkspaceEditParams, error)
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

func (args CodeActionGeneratePrivateKeyArgs) RunCommand(d *wireguard.WGDocument) (*protocol.ApplyWorkspaceEditParams, error) {
	privateKey, err := wgcommands.CreateNewPrivateKey()

	if err != nil {
		return &protocol.ApplyWorkspaceEditParams{}, err
	}

	section := d.Config.FindSectionByLine(args.Line)
	property := d.Config.FindPropertyByLine(args.Line)

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
						Range: protocol.Range{
							Start: protocol.Position{
								Line:      property.End.Line,
								Character: property.End.Character,
							},
							End: protocol.Position{
								Line:      property.End.Line,
								Character: property.End.Character,
							},
						},
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

func (args CodeActionGeneratePresharedKeyArgs) RunCommand(d *wireguard.WGDocument) (*protocol.ApplyWorkspaceEditParams, error) {
	presharedKey, err := wgcommands.CreatePresharedKey()

	if err != nil {
		return &protocol.ApplyWorkspaceEditParams{}, err
	}

	section := d.Config.FindSectionByLine(args.Line)
	property := d.Config.FindPropertyByLine(args.Line)

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
						Range: protocol.Range{
							Start: protocol.Position{
								Line:      property.End.Line,
								Character: property.End.Character,
							},
							End: protocol.Position{
								Line:      property.End.Line,
								Character: property.End.Character,
							},
						},
					},
				},
			},
		},
	}, nil
}

