package analyzer

import (
	"config-lsp/common"
	docvalues "config-lsp/doc-values"
	"config-lsp/handlers/fstab/ast"
	"config-lsp/handlers/fstab/fields"
	"errors"
	"regexp"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func analyzeValuesAreValid(
	ctx *analyzerContext,
) {
	it := ctx.document.Config.Entries.Iterator()

	for it.Next() {
		entry := it.Value().(*ast.FstabEntry)

		checkField(ctx, entry.Fields.Spec, fields.SpecField)
		checkField(ctx, entry.Fields.MountPoint, fields.MountPointField)
		checkField(ctx, entry.Fields.FilesystemType, fields.FileSystemTypeField)

		analyzeSpecField(ctx, entry.Fields.Spec)

		if entry.Fields.Options != nil {
			checkMountOptions(ctx, entry)
		}

		if entry.Fields.Freq != nil {
			checkField(ctx, entry.Fields.Freq, fields.FreqField)
		}

		if entry.Fields.Fsck != nil {
			checkField(ctx, entry.Fields.Fsck, fields.FsckField)
		}
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

func checkMountOptions(
	ctx *analyzerContext,
	entry *ast.FstabEntry,
) {
	mountOptions := entry.FetchMountOptionsField(true)

	if mountOptions == nil {
		return
	}

	field := entry.Fields.Options
	invalidValues := mountOptions.DeprecatedCheckIsValid(field.Value.Value)

	isZFS := entry.Fields.FilesystemType.Value.Value == "zfs"

	for _, invalidValue := range invalidValues {
		// Edge case for ZFS:
		// ZFS allows "User Properties", basically arbitrary keys to be set
		// See https://openzfs.github.io/openzfs-docs/man/master/7/zfsprops.7.html#User_Properties
		if isZFS {
			notInEnumError, isNotInEnumError := (invalidValue.Err).(any).(docvalues.ValueEnumValNotInEnumsError)

			if isNotInEnumError {
				// Edge case confirmed, just validate if the user property follows the rules
				invalidValue = checkZFSUserProperty(notInEnumError.ProvidedValue)

				if invalidValue == nil {
					continue
				}
			}
		}

		err := docvalues.LSPErrorFromInvalidValue(field.Start.Line, *invalidValue).ShiftCharacter(field.Start.Character)

		ctx.diagnostics = append(ctx.diagnostics, protocol.Diagnostic{
			Range:    err.Range.ToLSPRange(),
			Message:  err.Err.Error(),
			Severity: &common.SeverityError,
		})
	}
}

var validZFSUserProperties = regexp.MustCompile(`^(?<key>[a-z0-9:._][a-z0-9:._-]*?:[a-z0-9:._-]+)(?:=(?<value>.+))?$`)

func checkZFSUserProperty(
	propertyName string,
) *docvalues.InvalidValue {
	match := validZFSUserProperties.FindStringSubmatch(propertyName)

	if match == nil {
		return &docvalues.InvalidValue{
			Err:   errors.New(`Invalid ZFS user property format. User property names must contain a colon (":") character to distinguish them from native properties. They may contain lowercase letters, numbers, and the following punctuation characters: colon (":"), dash ("-"), period ("."), and underscore ("_")`),
			Start: 0,
			End:   uint32(len(propertyName)),
		}
	}

	key := match[1]
	// value := match[2]

	if len(key) > 256 {
		return &docvalues.InvalidValue{
			Err:   errors.New("ZFS user property key is too long, maximum length is 256 characters"),
			Start: 0,
			End:   uint32(len(key)),
		}
	}

	// TODO: Value check can be added later

	return nil
}
