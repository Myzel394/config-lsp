package handlers

import (
	"config-lsp/handlers/hosts/ast"
	"config-lsp/utils"
	"fmt"
	"strings"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

type CodeActionName string

const (
	CodeActionInlineAliases CodeActionName = "inlineAliases"
)

type CodeAction interface {
	RunCommand(ast.HostsParser) (*protocol.ApplyWorkspaceEditParams, error)
}

type CodeActionArgs interface{}

type CodeActionInlineAliasesArgs struct {
	URI      protocol.DocumentUri
	FromLine uint32
	ToLine   uint32
}

func CodeActionInlineAliasesArgsFromArguments(arguments map[string]any) CodeActionInlineAliasesArgs {
	return CodeActionInlineAliasesArgs{
		URI:      arguments["URI"].(protocol.DocumentUri),
		FromLine: uint32(arguments["FromLine"].(float64)),
		ToLine:   uint32(arguments["ToLine"].(float64)),
	}
}

func (args CodeActionInlineAliasesArgs) RunCommand(hostsParser ast.HostsParser) (*protocol.ApplyWorkspaceEditParams, error) {
	fromEntry := hostsParser.Tree.Entries[args.FromLine]
	toEntry := hostsParser.Tree.Entries[args.ToLine]

	if fromEntry == nil || toEntry == nil {
		// Weird
		return nil, nil
	}

	var insertCharacter uint32

	if toEntry.Aliases != nil {
		insertCharacter = toEntry.Aliases[len(toEntry.Aliases)-1].Location.End.Character
	} else {
		insertCharacter = toEntry.Hostname.Location.End.Character
	}

	hostnames := append(
		[]string{
			fromEntry.Hostname.Value,
		},
		utils.Map(
			fromEntry.Aliases,
			func(alias *ast.HostsHostname) string {
				return alias.Value
			},
		)...,
	)

	label := fmt.Sprintf("Inline aliases from %d to %d", args.FromLine, args.ToLine)
	return &protocol.ApplyWorkspaceEditParams{
		Label: &label,
		Edit: protocol.WorkspaceEdit{
			Changes: map[protocol.DocumentUri][]protocol.TextEdit{
				args.URI: {
					// Delete old line
					{
						Range: protocol.Range{
							Start: protocol.Position{
								Line:      args.FromLine,
								Character: 0,
							},
							End: protocol.Position{
								Line:      args.FromLine + 1,
								Character: 0,
							},
						},
						NewText: "",
					},
					// Insert aliases
					{
						Range: protocol.Range{
							Start: protocol.Position{
								Line:      args.ToLine,
								Character: insertCharacter,
							},
							End: protocol.Position{
								Line:      args.ToLine,
								Character: insertCharacter,
							},
						},
						NewText: " " + strings.Join(hostnames, " "),
					},
				},
			},
		},
	}, nil
}
