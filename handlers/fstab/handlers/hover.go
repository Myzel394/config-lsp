package handlers

import (
	"config-lsp/doc-values"
	"config-lsp/handlers/fstab/documentation"
	"config-lsp/handlers/fstab/parser"
	"github.com/tliron/glsp/protocol_3_16"
	"strings"
)

func GetHoverInfo(entry *parser.FstabEntry, cursor uint32) (*protocol.Hover, error) {
	line := entry.Line
	targetField := line.GetFieldAtPosition(cursor)

	switch targetField {
	case parser.FstabFieldSpec:
		return &SpecHoverField, nil
	case parser.FstabFieldMountPoint:
		return &MountPointHoverField, nil
	case parser.FstabFieldFileSystemType:
		return &FileSystemTypeField, nil
	case parser.FstabFieldOptions:
		fileSystemType := line.Fields.FilesystemType.Value
		var optionsField docvalues.Value

		if foundField, found := fstabdocumentation.MountOptionsMapField[fileSystemType]; found {
			optionsField = foundField
		} else {
			optionsField = fstabdocumentation.DefaultMountOptionsField
		}

		relativeCursor := cursor - line.Fields.Options.Start
		fieldInfo := optionsField.FetchHoverInfo(line.Fields.Options.Value, relativeCursor)

		hover := protocol.Hover{
			Contents: protocol.MarkupContent{
				Kind:  protocol.MarkupKindMarkdown,
				Value: strings.Join(fieldInfo, "\n"),
			},
		}

		return &hover, nil
	case parser.FstabFieldFreq:
		return &FreqHoverField, nil
	case parser.FstabFieldPass:
		return &PassHoverField, nil
	}

	return nil, nil
}
