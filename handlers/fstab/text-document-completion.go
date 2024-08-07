package fstab

import (
	docvalues "config-lsp/doc-values"
	fstabdocumentation "config-lsp/handlers/fstab/documentation"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentCompletion(context *glsp.Context, params *protocol.CompletionParams) (any, error) {
	parser := documentParserMap[params.TextDocument.URI]

	entry, found := parser.GetEntry(params.Position.Line)

	if !found {
		// Empty line, return spec completions
		return fstabdocumentation.SpecField.FetchCompletions(
			"",
			params.Position.Character,
		), nil
	}

	cursor := params.Position.Character
	targetField := entry.GetFieldAtPosition(cursor - 1)

	switch targetField {
	case FstabFieldSpec:
		value, cursor := GetFieldSafely(entry.Fields.Spec, cursor)

		return fstabdocumentation.SpecField.FetchCompletions(
			value,
			cursor,
		), nil
	case FstabFieldMountPoint:
		value, cursor := GetFieldSafely(entry.Fields.MountPoint, cursor)

		return fstabdocumentation.MountPointField.FetchCompletions(
			value,
			cursor,
		), nil
	case FstabFieldFileSystemType:
		value, cursor := GetFieldSafely(entry.Fields.FilesystemType, cursor)

		return fstabdocumentation.FileSystemTypeField.FetchCompletions(
			value,
			cursor,
		), nil
	case FstabFieldOptions:
		fileSystemType := entry.Fields.FilesystemType.Value

		var optionsField docvalues.Value

		if foundField, found := fstabdocumentation.MountOptionsMapField[fileSystemType]; found {
			optionsField = foundField
		} else {
			optionsField = fstabdocumentation.DefaultMountOptionsField
		}

		value, cursor := GetFieldSafely(entry.Fields.Options, cursor)

		completions := optionsField.FetchCompletions(
			value,
			cursor,
		)

		return completions, nil
	case FstabFieldFreq:
		value, cursor := GetFieldSafely(entry.Fields.Freq, cursor)

		return fstabdocumentation.FreqField.FetchCompletions(
			value,
			cursor,
		), nil
	case FstabFieldPass:
		value, cursor := GetFieldSafely(entry.Fields.Pass, cursor)

		return fstabdocumentation.PassField.FetchCompletions(
			value,
			cursor,
		), nil
	}

	return nil, nil
}

// Safely get value and new cursor position
// If field is nil, return empty string and 0
func GetFieldSafely(field *Field, character uint32) (string, uint32) {
	if field == nil {
		return "", 0
	}

	if field.Value == "" {
		return "", 0
	}

	return field.Value, character - field.Start
}
