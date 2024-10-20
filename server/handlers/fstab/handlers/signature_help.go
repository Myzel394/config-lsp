package handlers

import (
	"config-lsp/common"
	"config-lsp/handlers/fstab/ast"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func GetEntrySignatureHelp(
	entry *ast.FstabEntry,
	cursor common.CursorPosition,
) *protocol.SignatureHelp {
	var index uint32

	if entry == nil || entry.Fields.Spec == nil || entry.Fields.Spec.ContainsPosition(cursor) {
		index = 0
	} else if entry.Fields.MountPoint == nil && entry.Fields.MountPoint.ContainsPosition(cursor) {
		index = 1
	} else if entry.Fields.FilesystemType == nil && entry.Fields.FilesystemType.ContainsPosition(cursor) {
		index = 2
	} else if entry.Fields.Options == nil || entry.Fields.Options.ContainsPosition(cursor) {
		index = 3
	} else if entry.Fields.Freq == nil || entry.Fields.Freq.ContainsPosition(cursor) {
		index = 4
	} else {
		index = 5
	}

	signature := uint32(0)

	return &protocol.SignatureHelp{
		ActiveSignature: &signature,
		Signatures: []protocol.SignatureInformation{
			{
				Label:           "<spec> <mount point> <file system type> <options> <freq> <pass>",
				ActiveParameter: &index,
				Parameters: []protocol.ParameterInformation{
					{
						Label: []uint32{
							0,
							uint32(len("<spec>")),
						},
						Documentation: "The device or remote filesystem to mount",
					},
					{
						Label: []uint32{
							uint32(len("<spec>")),
							uint32(len("<spec> ") + len("<mount point>")),
						},
						Documentation: "The directory to mount the device or remote filesystem",
					},
					{
						Label: []uint32{
							uint32(len("<spec> <mount point>")),
							uint32(len("<spec> <mount point> ") + len("<file system type>")),
						},
						Documentation: "The type of filesystem",
					},
					{
						Label: []uint32{
							uint32(len("<spec> <mount point> <file system type>")),
							uint32(len("<spec> <mount point> <file system type> ") + len("<options>")),
						},
						Documentation: "Mount options",
					},
					{
						Label: []uint32{
							uint32(len("<spec> <mount point> <file system type> <options>")),
							uint32(len("<spec> <mount point> <file system type> <options> ") + len("<freq>")),
						},
						Documentation: "Used by dump(8) to determine which filesystems need to be dumped",
					},
					{
						Label: []uint32{
							uint32(len("<spec> <mount point> <file system type> <options> <freq>")),
							uint32(len("<spec> <mount point> <file system type> <options> <freq> ") + len("<pass>")),
						},
						Documentation: "Used by fsck(8) to determine the order in which filesystem checks are done at boot time",
					},
				},
			},
		},
	}
}
