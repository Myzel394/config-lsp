package fstab

import (
	docvalues "config-lsp/doc-values"
	fstabdocumentation "config-lsp/handlers/fstab/documentation"
	"strings"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentHover(context *glsp.Context, params *protocol.HoverParams) (*protocol.Hover, error) {
	cursor := params.Position.Character

	parser := documentParserMap[params.TextDocument.URI]

	entry, found := parser.GetEntry(params.Position.Line)

	// Empty line
	if !found {
		return nil, nil
	}

	// Comment line
	if entry.Type == FstabEntryTypeComment {
		return nil, nil
	}

	line := entry.Line
	targetField := line.GetFieldAtPosition(cursor)

	switch targetField {
	case FstabFieldSpec:
		return &SpecHoverField, nil
	case FstabFieldMountPoint:
		return &MountPointHoverField, nil
	case FstabFieldFileSystemType:
		return &FileSystemTypeField, nil
	case FstabFieldOptions:
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
	case FstabFieldFreq:
		return &FreqHoverField, nil
	case FstabFieldPass:
		return &PassHoverField, nil
	}

	return nil, nil
}
