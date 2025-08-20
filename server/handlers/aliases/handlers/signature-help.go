package handlers

import (
	"config-lsp/common"
	"config-lsp/handlers/aliases/ast"
	"strings"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func GetRootSignatureHelp(
	activeParameter uint32,
) *protocol.SignatureHelp {
	index := uint32(0)
	return &protocol.SignatureHelp{
		ActiveSignature: &index,
		Signatures: []protocol.SignatureInformation{
			{
				Label:           "<alias>: <value1>, <value2>, ...",
				ActiveParameter: &activeParameter,
				Parameters: []protocol.ParameterInformation{
					{
						Label: []uint32{
							0,
							uint32(len("<alias>") + 1),
						},
						Documentation: "The alias to define",
					},
					{
						Label: []uint32{
							uint32(len("<alias>: ")),
							uint32(len("<alias>: ") + len("<value1>")),
						},
						Documentation: "A value to associate with the alias",
					},
				},
			},
		},
	}
}

func GetAllValuesSignatureHelp() *protocol.SignatureHelp {
	index := uint32(0)
	return &protocol.SignatureHelp{
		Signatures: []protocol.SignatureInformation{
			{
				Label:           "<user>",
				ActiveParameter: &index,
				Parameters: []protocol.ParameterInformation{
					{
						Label: []uint32{
							0,
							uint32(len("<user>")),
						},
					},
				},
			},
			{
				Label:           "<user>@<host>",
				ActiveParameter: &index,
				Parameters: []protocol.ParameterInformation{
					{
						Label: []uint32{
							0,
							uint32(len("<user>")),
						},
					},
					{
						Label: []uint32{
							uint32(len("<user>@")),
							uint32(len("<user>@<host>")),
						},
					},
				},
			},
			{
				Label:           "<file>",
				ActiveParameter: &index,
				Parameters: []protocol.ParameterInformation{
					{
						Label: []uint32{
							0,
							uint32(len("<file>")),
						},
					},
				},
			},
			{
				Label:           ":include:<file>",
				ActiveParameter: &index,
				Parameters: []protocol.ParameterInformation{
					{
						Label: []uint32{
							0,
							uint32(len(":include:")),
						},
					},
				},
			},
			{
				Label:           "|<command>",
				ActiveParameter: &index,
				Parameters: []protocol.ParameterInformation{
					{
						Label: []uint32{
							0,
							1,
						},
					},
				},
			},
			{
				Label:           "error:<code> <message>",
				ActiveParameter: &index,
				Parameters: []protocol.ParameterInformation{
					{
						Label: []uint32{
							0,
							uint32(len("error")),
						},
					},
				},
			},
		},
	}
}

func GetValueSignatureHelp(
	cursor common.CursorPosition,
	value ast.AliasValueInterface,
) *protocol.SignatureHelp {
	switch value.(type) {
	case ast.AliasValueUser:
		index := uint32(0)
		return &protocol.SignatureHelp{
			Signatures: []protocol.SignatureInformation{
				{
					Label:           "<user>",
					ActiveParameter: &index,
					Parameters: []protocol.ParameterInformation{
						{
							Label: []uint32{
								0,
								uint32(len("<user>")),
							},
						},
					},
				},
				{
					Label:           "<user>@<host>",
					ActiveParameter: &index,
					Parameters: []protocol.ParameterInformation{
						{
							Label: []uint32{
								0,
								uint32(len("<user>")),
							},
						},
						{
							Label: []uint32{
								uint32(len("<user>@")),
								uint32(len("<user>@") + len("<host>")),
							},
						},
					},
				},
			},
		}
	case ast.AliasValueEmail:
		indexPosition := common.LSPCharacterAsIndexPosition(uint32(strings.Index(value.GetAliasValue().Value, "@")))
		isBeforeAtSymbol := cursor.IsBeforeIndexPosition(indexPosition)

		var index uint32

		if isBeforeAtSymbol {
			index = 0
		} else {
			index = 1
		}

		return &protocol.SignatureHelp{
			Signatures: []protocol.SignatureInformation{
				{
					Label:           "<user>@<host>",
					ActiveParameter: &index,
					Parameters: []protocol.ParameterInformation{
						{
							Label: []uint32{
								0,
								uint32(len("<user>")),
							},
						},
						{
							Label: []uint32{
								uint32(len("<user>@")),
								uint32(len("<user>@") + len("<host>")),
							},
						},
					},
				},
			},
		}
	case ast.AliasValueFile:
		index := uint32(0)
		return &protocol.SignatureHelp{
			Signatures: []protocol.SignatureInformation{
				{
					Label:           "<file>",
					ActiveParameter: &index,
					Parameters: []protocol.ParameterInformation{
						{
							Label: []uint32{
								0,
								uint32(len("<file>")),
							},
						},
					},
				},
			},
		}
	case ast.AliasValueInclude:
		index := uint32(0)
		return &protocol.SignatureHelp{
			Signatures: []protocol.SignatureInformation{
				{
					Label:           "include:<file>",
					ActiveParameter: &index,
					Parameters: []protocol.ParameterInformation{
						{
							Label: []uint32{
								uint32(len("include:")),
								uint32(len("include:<file>")),
							},
						},
					},
				},
			},
		}
	case ast.AliasValueCommand:
		var index uint32

		if cursor == 0 {
			index = 0
		} else {
			index = 1
		}

		return &protocol.SignatureHelp{
			Signatures: []protocol.SignatureInformation{
				{
					Label:           "|<command>",
					ActiveParameter: &index,
					Parameters: []protocol.ParameterInformation{
						{
							Label: []uint32{
								0,
								1,
							},
						},
						{
							Label: []uint32{
								1,
								uint32(1 + len("<command>")),
							},
						},
					},
				},
			},
		}
	case ast.AliasValueError:
		errorValue := value.(ast.AliasValueError)
		var index uint32

		if errorValue.Code == nil || errorValue.Code.Location.IsPositionBeforeEnd(cursor) {
			index = 1
		} else {
			index = 2
		}

		return &protocol.SignatureHelp{
			Signatures: []protocol.SignatureInformation{
				{
					Label:           "error:<code> <message>",
					ActiveParameter: &index,
					Parameters: []protocol.ParameterInformation{
						{
							Label: []uint32{
								0,
								uint32(len("error:")),
							},
						},
						{
							Label: []uint32{
								uint32(len("error:")),
								uint32(len("error:<code>")),
							},
						},
						{
							Label: []uint32{
								uint32(len("error:<code> ")),
								uint32(len("error:<code> <message>")),
							},
						},
					},
				},
			},
		}
	}

	return nil
}
