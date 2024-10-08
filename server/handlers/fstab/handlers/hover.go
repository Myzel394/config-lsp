package handlers

import (
	"config-lsp/common"
	"config-lsp/doc-values"
	"config-lsp/handlers/fstab/ast"
	"config-lsp/handlers/fstab/fields"
	"strings"

	"github.com/tliron/glsp/protocol_3_16"
)

func GetHoverInfo(
	line uint32,
	index common.IndexPosition,
	entry *ast.FstabEntry,
) (*protocol.Hover, error) {
	targetField := entry.GetFieldAtPosition(index)

	switch targetField {
	case ast.FstabFieldSpec:
		return &SpecHoverField, nil
	case ast.FstabFieldMountPoint:
		return &MountPointHoverField, nil
	case ast.FstabFieldFileSystemType:
		return &FileSystemTypeField, nil
	case ast.FstabFieldOptions:
		fileSystemType := entry.Fields.FilesystemType.Value.Value
		var optionsField docvalues.DeprecatedValue

		if foundField, found := fields.MountOptionsMapField[fileSystemType]; found {
			optionsField = foundField
		} else {
			optionsField = fields.DefaultMountOptionsField
		}

		relativeCursor := uint32(index) - entry.Fields.Options.Start.Character
		fieldInfo := optionsField.DeprecatedFetchHoverInfo(entry.Fields.Options.Value.Value, relativeCursor)

		hover := protocol.Hover{
			Contents: protocol.MarkupContent{
				Kind:  protocol.MarkupKindMarkdown,
				Value: strings.Join(fieldInfo, "\n"),
			},
		}

		return &hover, nil
	case ast.FstabFieldFreq:
		return &FreqHoverField, nil
	case ast.FstabFieldPass:
		return &PassHoverField, nil
	}

	return nil, nil
}
