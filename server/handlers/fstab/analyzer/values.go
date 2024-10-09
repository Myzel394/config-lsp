package analyzer

import (
	"config-lsp/common"
	docvalues "config-lsp/doc-values"
	"config-lsp/handlers/fstab/ast"
	"config-lsp/handlers/fstab/fields"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func analyzeValuesAreValid(
	ctx *analyzerContext,
) {
	it := ctx.document.Config.Entries.Iterator()

	for it.Next() {
		entry := it.Value().(*ast.FstabEntry)
		mountOptions := entry.FetchMountOptionsField(true)

		checkField(ctx, entry.Fields.Spec, fields.SpecField)
		checkField(ctx, entry.Fields.MountPoint, fields.MountPointField)
		checkField(ctx, entry.Fields.FilesystemType, fields.FileSystemTypeField)

		if mountOptions != nil {
			checkField(ctx, entry.Fields.Options, mountOptions)
		}

		checkField(ctx, entry.Fields.Freq, fields.FreqField)
		checkField(ctx, entry.Fields.Pass, fields.PassField)
	}
}

func checkField(
	ctx *analyzerContext,
	field *ast.FstabField,
	docOption docvalues.DeprecatedValue,
) {
	invalidValues := docOption.DeprecatedCheckIsValid(field.Value.Value)

	for _, invalidValue := range invalidValues {
		err := docvalues.LSPErrorFromInvalidValue(field.Start.Line, *invalidValue).ShiftCharacter(field.Start.Character)

		ctx.diagnostics = append(ctx.diagnostics, protocol.Diagnostic{
			Range:    err.Range.ToLSPRange(),
			Message:  err.Err.Error(),
			Severity: &common.SeverityError,
		})
	}
}
