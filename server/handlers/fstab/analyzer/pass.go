package analyzer

import (
	"config-lsp/common"
	"config-lsp/handlers/fstab/ast"
	"fmt"
	"strings"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func analyzePassFields(ctx *analyzerContext) {
	it := ctx.document.Config.Entries.Iterator()

	var rootEntry *ast.FstabEntry

	for it.Next() {
		entry := it.Value().(*ast.FstabEntry)

		if entry.Fields != nil && entry.Fields.Pass != nil && entry.Fields.Pass.Value.Value == "1" {
			fileSystem := strings.ToLower(entry.Fields.FilesystemType.Value.Value)

			if fileSystem == "btrfs" || fileSystem == "xfs" {
				// From https://wiki.archlinux.org/title/Fstab

				ctx.diagnostics = append(ctx.diagnostics, protocol.Diagnostic{
					Range:    entry.Fields.Pass.ToLSPRange(),
					Message:  "If the root file system is btrfs or XFS, the fsck order should be set to 0 instead of 1. See fsck.btrfs(8) and fsck.xfs(8).",
					Severity: &common.SeverityWarning,
				})

				continue
			}

			if entry.Fields.Pass.Value.Value == "1" {
				if rootEntry != nil {
					ctx.diagnostics = append(ctx.diagnostics, protocol.Diagnostic{
						Range:    entry.Fields.Pass.ToLSPRange(),
						Message:  fmt.Sprintf("Only the root file system should have a fsck of 1. Other file systems should have a fsck of 2 or 0. The root file system is already using a fsck=1 on line %d", rootEntry.Fields.Start.Line),
						Severity: &common.SeverityWarning,
					})
				} else {
					rootEntry = entry
				}
			}
		}
	}
}
