package handlers

import (
	"config-lsp/common"
	"config-lsp/handlers/hosts/ast"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func GetEntrySignatureHelp(
	entry *ast.HostsEntry,
	cursor common.CursorPosition,
) *protocol.SignatureHelp {
	var index uint32

	if entry == nil || entry.IPAddress == nil || entry.IPAddress.Location.ContainsPosition(cursor) {
		index = 0
	} else if entry.Hostname == nil || entry.Hostname.Location.ContainsPosition(cursor) {
		index = 1
	} else {
		index = 2
	}

	signature := uint32(0)

	return &protocol.SignatureHelp{
		ActiveSignature: &signature,
		Signatures: []protocol.SignatureInformation{
			{
				Label:           "<ip address> <hostname> [<alias>...]",
				ActiveParameter: &index,
				Parameters: []protocol.ParameterInformation{
					{
						Label: []uint32{
							0,
							uint32(len("<ip address>")),
						},
						Documentation: "The ip address to forward to",
					},
					{
						Label: []uint32{
							uint32(len("<ip address>")),
							uint32(len("<ip address> ") + len("<hostname>")),
						},
						Documentation: "The hostname to forward to",
					},
					{
						Label: []uint32{
							uint32(len("<ip address> ") + len("<hostname>")),
							uint32(len("<ip address> ") + len("<hostname> ") + len("[<alias>...]")),
						},
						Documentation: "An optional list of aliases that can also forward",
					},
				},
			},
		},
	}
}
