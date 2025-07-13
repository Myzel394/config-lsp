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

		return fields.SpecField.FetchCompletions(
			value,
			cursor,
		), nil
	case ast.FstabFieldMountPoint:
		value, cursor := getFieldSafely(entry.Fields.MountPoint, cursor)

		return fields.MountPointField.FetchCompletions(
			value,
			cursor,
		), nil
	case ast.FstabFieldFileSystemType:
		value, cursor := getFieldSafely(entry.Fields.FilesystemType, cursor)

		return fields.FileSystemTypeField.FetchCompletions(
			value,
			cursor,
		), nil
	case ast.FstabFieldOptions:
		line, cursor := getFieldSafely(entry.Fields.Options, cursor)
		fileSystemType := entry.Fields.FilesystemType.Value.Value
		completions := make([]protocol.CompletionItem, 0, 50)

		optionsValue := entry.FetchMountOptionsField(false)

		if optionsValue != nil {
			for _, completion := range optionsValue.FetchCompletions(line, cursor) {
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
		}

		// Add defaults
		completions = append(completions, fields.DefaultMountOptionsField.FetchCompletions(line, cursor)...)

		return completions, nil
	case ast.FstabFieldFreq:
		value, cursor := getFieldSafely(entry.Fields.Freq, cursor)

		return fields.FreqField.FetchCompletions(
			value,
			cursor,
		), nil
	case ast.FstabFieldFsck:
		value, cursor := getFieldSafely(entry.Fields.Fsck, cursor)

		if entry.Fields.FilesystemType != nil &&
			utils.KeyExists(fields.FsckOneDisabledFilesystems, strings.ToLower(entry.Fields.FilesystemType.Value.Value)) {
			return fields.FsckFieldWhenDisabledFilesystems.FetchCompletions(
				value,
				cursor,
			), nil
		} else {
			return fields.FsckField.FetchCompletions(
				value,
				cursor,
			), nil
		}
	}

	return nil, nil
}

// Safely get value and new cursor position
// If field is nil, return empty string and 0
func getFieldSafely(field *ast.FstabField, cursor common.CursorPosition) (string, common.CursorPosition) {
	if field == nil {
		return "", 0
	}

	if field.Value.Value == "" {
		return "", 0
	}

	if uint32(cursor) < field.Start.Character {
		return "", 0
	}

	return field.Value.Raw, cursor.ShiftHorizontal(-int(field.Start.Character))
}
