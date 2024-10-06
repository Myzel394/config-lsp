package handlers

import (
	"config-lsp/common/formatting"
	sshconfig "config-lsp/handlers/ssh_config"
	"config-lsp/handlers/ssh_config/ast"
	"fmt"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

type codeActionName string

const (
	CodeActionAddToUnknown codeActionName = "addToUnknown"
)

type codeAction interface {
	RunCommand(*sshconfig.SSHDocument) (*protocol.ApplyWorkspaceEditParams, error)
}

type codeActionArgs interface{}

type codeActionAddToUnknownArgs struct {
	URI protocol.DocumentUri
	// Where the option is defined
	OptionLine uint32
	// Where the block is defined, if nil, option is globally defined
	BlockLine *uint32
}

func CodeActionAddToUnknownArgsFromArguments(arguments map[string]interface{}) codeActionAddToUnknownArgs {
	var blockLine *uint32

	if arguments["BlockLine"] != nil {
		blockLineValue := uint32(arguments["BlockLine"].(float64))
		blockLine = &blockLineValue
	}

	return codeActionAddToUnknownArgs{
		URI:        arguments["URI"].(protocol.DocumentUri),
		OptionLine: uint32(arguments["OptionLine"].(float64)),
		BlockLine:  blockLine,
	}
}

var addToUnknownOptionTemplate = formatting.FormatTemplate("/!'%s/!'")

func (args codeActionAddToUnknownArgs) RunCommand(d *sshconfig.SSHDocument) (*protocol.ApplyWorkspaceEditParams, error) {
	var option *ast.SSHOption
	var block ast.SSHBlock

	// Either this or `insertionLine` must be set
	// `ignoreUnknownOption` is used if an `IgnoreUnknown` option is set already
	// `insertionLine` is used if no `IgnoreUnknown` option is set
	var ignoreUnknownOption *ast.SSHOption
	var insertionLine uint32

	if args.BlockLine == nil {
		// Global
		rawOption, found := d.Config.Options.Get(args.OptionLine)

		if !found {
			return nil, fmt.Errorf("No option found at line %d", args.OptionLine)
		}

		option = rawOption.(*ast.SSHOption)

		if ignoreOption, found := d.Indexes.IgnoredOptions[nil]; found {
			ignoreUnknownOption = ignoreOption.OptionValue
		} else {
			insertionLine = 0
		}
	} else {
		// Block
		rawBlock, found := d.Config.Options.Get(*args.BlockLine)

		if !found {
			return nil, fmt.Errorf("No block found at line %d", *args.BlockLine)
		}

		block = rawBlock.(ast.SSHBlock)

		rawOption, found := block.GetOptions().Get(args.OptionLine)

		if !found {
			return nil, fmt.Errorf("No option found at line %d", args.OptionLine)
		}

		option = rawOption.(*ast.SSHOption)

		if ignoreOption, found := d.Indexes.IgnoredOptions[block]; found {
			ignoreUnknownOption = ignoreOption.OptionValue
		} else {
			insertionLine = block.GetEntryOption().Start.Line + 1
		}
	}

	rawOptionName := option.Key.Value.Raw
	optionName := addToUnknownOptionTemplate.Format(formatting.DefaultFormattingOptions, rawOptionName)
	label := fmt.Sprintf("Add %s to unknown options", option.Key.Key)

	// We got everything, let's build the edit!
	if ignoreUnknownOption == nil {
		// Insert a completely new IgnoreUnknown option
		if block == nil {
			// Global
			return &protocol.ApplyWorkspaceEditParams{
				Label: &label,
				Edit: protocol.WorkspaceEdit{
					Changes: map[protocol.DocumentUri][]protocol.TextEdit{
						args.URI: {
							{
								Range: protocol.Range{
									Start: protocol.Position{
										Line:      insertionLine,
										Character: 0,
									},
									End: protocol.Position{
										Line:      insertionLine,
										Character: 0,
									},
								},
								NewText: fmt.Sprintf("IgnoreUnknown %s\n", optionName),
							},
						},
					},
				},
			}, nil
		} else {
			// Block
			return &protocol.ApplyWorkspaceEditParams{
				Label: &label,
				Edit: protocol.WorkspaceEdit{
					Changes: map[protocol.DocumentUri][]protocol.TextEdit{
						args.URI: {
							{
								Range: protocol.Range{
									Start: protocol.Position{
										Line:      insertionLine,
										Character: 0,
									},
									End: protocol.Position{
										Line:      insertionLine,
										Character: 0,
									},
								},
								NewText: fmt.Sprintf("    IgnoreUnknown %s\n", optionName),
							},
						},
					},
				},
			}, nil
		}
	} else {
		// Append to the existing IgnoreUnknown option
		return &protocol.ApplyWorkspaceEditParams{
			Label: &label,
			Edit: protocol.WorkspaceEdit{
				Changes: map[protocol.DocumentUri][]protocol.TextEdit{
					args.URI: {
						{
							Range: protocol.Range{
								Start: protocol.Position{
									Line:      ignoreUnknownOption.Start.Line,
									Character: ignoreUnknownOption.End.Character,
								},
								End: protocol.Position{
									Line:      ignoreUnknownOption.Start.Line,
									Character: ignoreUnknownOption.End.Character,
								},
							},
							NewText: fmt.Sprintf(" %s", optionName),
						},
					},
				},
			},
		}, nil
	}
}
