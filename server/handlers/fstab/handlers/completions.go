package handlers

import (
	"config-lsp/common"
	"config-lsp/handlers/fstab/ast"
	"config-lsp/handlers/fstab/fields"
	"config-lsp/utils"
	"fmt"
	"strings"

	"github.com/tliron/glsp/protocol_3_16"
)

func GetCompletion(
	entry *ast.FstabEntry,
	cursor common.CursorPosition,
) ([]protocol.CompletionItem, error) {
	targetField := entry.GetFieldAtPosition(cursor)

	switch targetField {
	case ast.FstabFieldSpec:
		value, cursor := getFieldSafely(entry.Fields.Spec, cursor)

		return fields.SpecField.DeprecatedFetchCompletions(
			value,
			cursor,
		), nil
	case ast.FstabFieldMountPoint:
		value, cursor := getFieldSafely(entry.Fields.MountPoint, cursor)

		return fields.MountPointField.DeprecatedFetchCompletions(
			value,
			cursor,
		), nil
	case ast.FstabFieldFileSystemType:
		value, cursor := getFieldSafely(entry.Fields.FilesystemType, cursor)

		return fields.FileSystemTypeField.DeprecatedFetchCompletions(
			value,
			cursor,
		), nil
	case ast.FstabFieldOptions:
		line, cursor := getFieldSafely(entry.Fields.Options, cursor)
		fileSystemType := entry.Fields.FilesystemType.Value.Value
		completions := make([]protocol.CompletionItem, 0, 50)

		optionsValue := entry.FetchMountOptionsField(false)

		if optionsValue == nil {
			optionsValue = fields.DefaultMountOptionsField
		}

		for _, completion := range optionsValue.DeprecatedFetchCompletions(line, cursor) {
			var documentation string

			switch completion.Documentation.(type) {
			case string:
				documentation = completion.Documentation.(string)
			case *string:
				documentation = *completion.Documentation.(*string)
			}

			completion.Documentation = protocol.MarkupContent{
				Kind:  protocol.MarkupKindMarkdown,
				Value: documentation + "\n\n" + fmt.Sprintf("From: _%s_", fileSystemType),
			}
			completions = append(completions, completion)
		}

		return completions, nil
	case ast.FstabFieldFreq:
		value, cursor := getFieldSafely(entry.Fields.Freq, cursor)

		return fields.FreqField.DeprecatedFetchCompletions(
			value,
			cursor,
		), nil
	case ast.FstabFieldFsck:
		value, cursor := getFieldSafely(entry.Fields.Fsck, cursor)

		if entry.Fields.FilesystemType != nil &&
			utils.KeyExists(fields.FsckOneDisabledFilesystems, strings.ToLower(entry.Fields.FilesystemType.Value.Value)) {
			return fields.FsckFieldWhenDisabledFilesystems.DeprecatedFetchCompletions(
				value,
				cursor,
			), nil
		} else {
			return fields.FsckField.DeprecatedFetchCompletions(
				value,
				cursor,
			), nil
		}
	}

	return nil, nil
}

// Safely get value and new cursor position
// If field is nil, return empty string and 0
func getFieldSafely(field *ast.FstabField, cursor common.CursorPosition) (string, uint32) {
	if field == nil {
		return "", 0
	}

	if field.Value.Value == "" {
		return "", 0
	}

	if uint32(cursor) < field.Start.Character {
		return "", 0
	}

	return field.Value.Raw, common.CursorToCharacterIndex(uint32(cursor) - field.Start.Character)
}
