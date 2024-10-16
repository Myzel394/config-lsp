package handlers

import (
	"config-lsp/common"
	"config-lsp/handlers/ssh_config/ast"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func GetOptionSignatureHelp(
	option *ast.SSHOption,
	cursor common.CursorPosition,
) *protocol.SignatureHelp {
	var index uint32

	if option == nil || option.Key == nil || (option.OptionValue == nil || option.Key.ContainsPosition(cursor)) {
		index = 0
	} else {
		index = 1
	}

	signature := uint32(0)
	return &protocol.SignatureHelp{
		ActiveSignature: &signature,
		Signatures: []protocol.SignatureInformation{
			{
				Label:           "<option> <value>",
				ActiveParameter: &index,
				Parameters: []protocol.ParameterInformation{
					{
						Label: []uint32{
							0,
							uint32(len("<option>") + 1),
						},
						Documentation: "The option name",
					},
					{
						Label: []uint32{
							uint32(len("<option>")),
							uint32(len("<option>") + len("<value>") + 1),
						},
						Documentation: "The value for the option",
					},
				},
			},
		},
	}
}

func GetMatchSignatureHelp(
	match *ast.SSHMatchBlock,
	cursor common.CursorPosition,
) *protocol.SignatureHelp {
	var index uint32

	if match.MatchOption.Key.ContainsPosition(cursor) {
		index = 0
	} else {
		entry := match.MatchValue.GetEntryAtPosition(cursor)

		if entry == nil || entry.Criteria.ContainsPosition(cursor) {
			index = 1
		} else {
			index = 2
		}
	}

	signature := uint32(0)
	return &protocol.SignatureHelp{
		ActiveSignature: &signature,
		Signatures: []protocol.SignatureInformation{
			{
				Label:           "Match <criteria> <values>",
				ActiveParameter: &index,
				Parameters: []protocol.ParameterInformation{
					{
						Label: []uint32{
							0,
							uint32(len("Match") + 1),
						},
						Documentation: "The \"Match\" keyword",
					},
					{
						Label: []uint32{
							uint32(len("Match ")),
							uint32(len("Match ") + len("<criteria>")),
						},
						Documentation: "The criteria name",
					},
					{
						Label: []uint32{
							uint32(len("Host <criteria> ")),
							uint32(len("Host <criteria> ") + len("<values>") + 1),
						},
						Documentation: "Values for the criteria",
					},
				},
			},
		},
	}
}

func GetHostSignatureHelp(
	host *ast.SSHHostBlock,
	cursor common.CursorPosition,
) *protocol.SignatureHelp {
	var index uint32

	if host.HostOption.Key.ContainsPosition(cursor) {
		index = 0
	} else if len(host.HostValue.Hosts) >= 1 && host.HostValue.Hosts[0].ContainsPosition(cursor) {
		index = 1
	} else {
		index = 2
	}

	signature := uint32(0)
	return &protocol.SignatureHelp{
		ActiveSignature: &signature,
		Signatures: []protocol.SignatureInformation{
			{
				Label:           "Host <host1> [<host2> ...]",
				ActiveParameter: &index,
				Parameters: []protocol.ParameterInformation{
					{
						Label: []uint32{
							0,
							uint32(len("Host") + 1),
						},
						Documentation: "The \"Host\" keyword",
					},
					{
						Label: []uint32{
							uint32(len("Host ")),
							uint32(len("Host ") + len("<host1>") + 1),
						},
						Documentation: "A host that should match",
					},
					{
						Label: []uint32{
							uint32(len("Host <host1> ")),
							uint32(len("Host <host1> ") + len("[<host2> ...]") + 1),
						},
						Documentation: "Additional (optional) hosts that should match",
					},
				},
			},
		},
	}
}
