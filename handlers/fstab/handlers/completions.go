package handlers

import (
	"config-lsp/doc-values"
	"config-lsp/handlers/fstab/documentation"
	"config-lsp/handlers/fstab/parser"
	"github.com/tliron/glsp/protocol_3_16"
)

func GetCompletion(
	line parser.FstabLine,
	cursor uint32,
) ([]protocol.CompletionItem, error) {
	targetField := line.GetFieldAtPosition(cursor)

	switch targetField {
	case parser.FstabFieldSpec:
		value, cursor := GetFieldSafely(line.Fields.Spec, cursor)

		return fstabdocumentation.SpecField.DeprecatedFetchCompletions(
			value,
			cursor,
		), nil
	case parser.FstabFieldMountPoint:
		value, cursor := GetFieldSafely(line.Fields.MountPoint, cursor)

		return fstabdocumentation.MountPointField.DeprecatedFetchCompletions(
			value,
			cursor,
		), nil
	case parser.FstabFieldFileSystemType:
		value, cursor := GetFieldSafely(line.Fields.FilesystemType, cursor)

		return fstabdocumentation.FileSystemTypeField.DeprecatedFetchCompletions(
			value,
			cursor,
		), nil
	case parser.FstabFieldOptions:
		fileSystemType := line.Fields.FilesystemType.Value

		var optionsField docvalues.DeprecatedValue

		if foundField, found := fstabdocumentation.MountOptionsMapField[fileSystemType]; found {
			optionsField = foundField
		} else {
			optionsField = fstabdocumentation.DefaultMountOptionsField
		}

		value, cursor := GetFieldSafely(line.Fields.Options, cursor)

		completions := optionsField.DeprecatedFetchCompletions(
			value,
			cursor,
		)

		return completions, nil
	case parser.FstabFieldFreq:
		value, cursor := GetFieldSafely(line.Fields.Freq, cursor)

		return fstabdocumentation.FreqField.DeprecatedFetchCompletions(
			value,
			cursor,
		), nil
	case parser.FstabFieldPass:
		value, cursor := GetFieldSafely(line.Fields.Pass, cursor)

		return fstabdocumentation.PassField.DeprecatedFetchCompletions(
			value,
			cursor,
		), nil
	}

	return nil, nil
}

// Safely get value and new cursor position
// If field is nil, return empty string and 0
func GetFieldSafely(field *parser.Field, character uint32) (string, uint32) {
	if field == nil {
		return "", 0
	}

	if field.Value == "" {
		return "", 0
	}

	return field.Value, character - field.Start
}
