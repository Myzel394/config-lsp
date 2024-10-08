package analyzer

import (
	"config-lsp/common"
	"config-lsp/handlers/fstab/ast"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func analyzeFieldAreFilled(
	ctx *analyzerContext,
) {
	it := ctx.document.Config.Entries.Iterator()
	for it.Next() {
		entry := it.Value().(*ast.FstabEntry)

		if entry.Fields.Spec == nil || entry.Fields.Spec.Value.Value == "" {
			ctx.diagnostics = append(ctx.diagnostics, protocol.Diagnostic{
				Range: protocol.Range{
					Start: protocol.Position{
						Line:      entry.Fields.Start.Line,
						Character: 0,
					},
					End: protocol.Position{
						Line:      entry.Fields.Start.Line,
						Character: 0,
					},
				},
				Message:  "The spec field is missing",
				Severity: &common.SeverityError,
			})

			continue
		}

		if entry.Fields.MountPoint == nil || entry.Fields.MountPoint.Value.Value == "" {
			ctx.diagnostics = append(ctx.diagnostics, protocol.Diagnostic{
				Range: protocol.Range{
					Start: protocol.Position{
						Line:      entry.Fields.Start.Line,
						Character: entry.Fields.Spec.End.Character,
					},
					End: protocol.Position{
						Line:      entry.Fields.Start.Line,
						Character: entry.Fields.Spec.End.Character,
					},
				},
				Message:  "The mount point field is missing",
				Severity: &common.SeverityError,
			})

			continue
		}

		if entry.Fields.FilesystemType == nil || entry.Fields.FilesystemType.Value.Value == "" {
			ctx.diagnostics = append(ctx.diagnostics, protocol.Diagnostic{
				Range: protocol.Range{
					Start: protocol.Position{
						Line:      entry.Fields.Start.Line,
						Character: entry.Fields.MountPoint.End.Character,
					},
					End: protocol.Position{
						Line:      entry.Fields.Start.Line,
						Character: entry.Fields.MountPoint.End.Character,
					},
				},
				Message:  "The file system type field is missing",
				Severity: &common.SeverityError,
			})

			continue
		}

		if entry.Fields.Options == nil || entry.Fields.Options.Value.Value == "" {
			ctx.diagnostics = append(ctx.diagnostics, protocol.Diagnostic{
				Range: protocol.Range{
					Start: protocol.Position{
						Line:      entry.Fields.Start.Line,
						Character: entry.Fields.FilesystemType.End.Character,
					},
					End: protocol.Position{
						Line:      entry.Fields.Start.Line,
						Character: entry.Fields.FilesystemType.End.Character,
					},
				},
				Message:  "The options field is missing",
				Severity: &common.SeverityError,
			})

			continue
		}

		if entry.Fields.Freq == nil || entry.Fields.Freq.Value.Value == "" {
			ctx.diagnostics = append(ctx.diagnostics, protocol.Diagnostic{
				Range: protocol.Range{
					Start: protocol.Position{
						Line:      entry.Fields.Start.Line,
						Character: entry.Fields.Options.End.Character,
					},
					End: protocol.Position{
						Line:      entry.Fields.Start.Line,
						Character: entry.Fields.Options.End.Character,
					},
				},
				Message:  "The freq field is missing",
				Severity: &common.SeverityError,
			})

			continue
		}

		if entry.Fields.Pass == nil || entry.Fields.Pass.Value.Value == "" {
			ctx.diagnostics = append(ctx.diagnostics, protocol.Diagnostic{
				Range: protocol.Range{
					Start: protocol.Position{
						Line:      entry.Fields.Start.Line,
						Character: entry.Fields.Freq.End.Character,
					},
					End: protocol.Position{
						Line:      entry.Fields.Start.Line,
						Character: entry.Fields.Freq.End.Character,
					},
				},
				Message:  "The pass field is missing",
				Severity: &common.SeverityError,
			})

			continue
		}
	}
}
