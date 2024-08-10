package fstab

import (
	"config-lsp/common"
	docvalues "config-lsp/doc-values"
	fstabdocumentation "config-lsp/handlers/fstab/documentation"
	"fmt"
	"regexp"
	"slices"
	"strings"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

var commentPattern = regexp.MustCompile(`^\s*#`)
var ignoreLinePattern = regexp.MustCompile(`^\s*$`)
var whitespacePattern = regexp.MustCompile(`\S+`)

type MalformedLineError struct{}

func (e MalformedLineError) Error() string {
	return "Malformed line"
}

type Field struct {
	Value string
	Start uint32
	End   uint32
}

func (f Field) String() string {
	return fmt.Sprintf("%v [%v-%v]", f.Value, f.Start, f.End)
}

func (f *Field) CreateRange(fieldLine uint32) protocol.Range {
	return protocol.Range{
		Start: protocol.Position{
			Line:      fieldLine,
			Character: f.Start,
		},
		End: protocol.Position{
			Line:      fieldLine,
			Character: f.End,
		},
	}
}

type FstabField string

const (
	FstabFieldSpec           FstabField = "spec"
	FstabFieldMountPoint     FstabField = "mountpoint"
	FstabFieldFileSystemType FstabField = "filesystemtype"
	FstabFieldOptions        FstabField = "options"
	FstabFieldFreq           FstabField = "freq"
	FstabFieldPass           FstabField = "pass"
)

type FstabFields struct {
	Spec           *Field
	MountPoint     *Field
	FilesystemType *Field
	Options        *Field
	Freq           *Field
	Pass           *Field
}

func (f FstabFields) String() string {
	return fmt.Sprintf(
		"Spec: %s, MountPoint: %s, FilesystemType: %s, Options: %s, Freq: %s, Pass: %s",
		f.Spec,
		f.MountPoint,
		f.FilesystemType,
		f.Options,
		f.Freq,
		f.Pass,
	)
}

type FstabLine struct {
	Line   uint32
	Fields FstabFields
}

func (e *FstabLine) CheckIsValid() []protocol.Diagnostic {
	diagnostics := make([]protocol.Diagnostic, 0)

	if e.Fields.Spec != nil {
		errors := fstabdocumentation.SpecField.CheckIsValid(e.Fields.Spec.Value)

		if len(errors) > 0 {
			diagnostics = append(
				diagnostics,
				docvalues.InvalidValuesToErrorDiagnostics(e.Line, e.Fields.Spec.Start, errors)...,
			)
		}
	}

	if e.Fields.MountPoint != nil {
		errors := fstabdocumentation.MountPointField.CheckIsValid(e.Fields.MountPoint.Value)

		if len(errors) > 0 {
			diagnostics = append(
				diagnostics,
				docvalues.InvalidValuesToErrorDiagnostics(e.Line, e.Fields.MountPoint.Start, errors)...,
			)
		}
	}

	var fileSystemType string = ""

	if e.Fields.FilesystemType != nil {
		errors := fstabdocumentation.FileSystemTypeField.CheckIsValid(e.Fields.FilesystemType.Value)

		if len(errors) > 0 {
			diagnostics = append(
				diagnostics,
				docvalues.InvalidValuesToErrorDiagnostics(e.Line, e.Fields.FilesystemType.Start, errors)...,
			)
		} else {
			fileSystemType = e.Fields.FilesystemType.Value
		}
	}

	if e.Fields.Options != nil && fileSystemType != "" {
		var optionsField docvalues.Value

		if foundField, found := fstabdocumentation.MountOptionsMapField[fileSystemType]; found {
			optionsField = foundField
		} else {
			optionsField = fstabdocumentation.DefaultMountOptionsField
		}

		errors := optionsField.CheckIsValid(e.Fields.Options.Value)

		if len(errors) > 0 {
			diagnostics = append(
				diagnostics,
				docvalues.InvalidValuesToErrorDiagnostics(e.Line, e.Fields.Options.Start, errors)...,
			)
		}
	}

	if e.Fields.Freq != nil {
		errors := fstabdocumentation.FreqField.CheckIsValid(e.Fields.Freq.Value)

		if len(errors) > 0 {
			diagnostics = append(
				diagnostics,
				docvalues.InvalidValuesToErrorDiagnostics(e.Line, e.Fields.Freq.Start, errors)...,
			)
		}
	}

	if e.Fields.Pass != nil {
		errors := fstabdocumentation.PassField.CheckIsValid(e.Fields.Pass.Value)

		if len(errors) > 0 {
			diagnostics = append(
				diagnostics,
				docvalues.InvalidValuesToErrorDiagnostics(e.Line, e.Fields.Pass.Start, errors)...,
			)
		}
	}

	return diagnostics
}

func (e FstabLine) GetFieldAtPosition(cursor uint32) FstabField {
	if e.Fields.Spec == nil || (cursor >= e.Fields.Spec.Start && cursor <= e.Fields.Spec.End) {
		return FstabFieldSpec
	}

	if e.Fields.MountPoint == nil || (cursor >= e.Fields.MountPoint.Start && cursor <= e.Fields.MountPoint.End) {
		return FstabFieldMountPoint
	}

	if e.Fields.FilesystemType == nil || (cursor >= e.Fields.FilesystemType.Start && cursor <= e.Fields.FilesystemType.End) {
		return FstabFieldFileSystemType
	}

	if e.Fields.Options == nil || (cursor >= e.Fields.Options.Start && cursor <= e.Fields.Options.End) {
		return FstabFieldOptions
	}

	println(fmt.Sprintf("cursor: %v, freq: %v", cursor, e.Fields.Freq))
	if e.Fields.Freq == nil || (cursor >= e.Fields.Freq.Start && cursor <= e.Fields.Freq.End) {
		return FstabFieldFreq
	}

	return FstabFieldPass
}

type FstabEntryType string

const (
	FstabEntryTypeLine    FstabEntryType = "line"
	FstabEntryTypeComment FstabEntryType = "comment"
)

type FstabEntry struct {
	Type FstabEntryType
	Line FstabLine
}

type FstabParser struct {
	entries []FstabEntry
}

func (p *FstabParser) AddLine(line string, lineNumber int) error {
	fields := whitespacePattern.FindAllStringIndex(line, -1)

	if len(fields) == 0 {
		return MalformedLineError{}
	}

	var spec *Field
	var mountPoint *Field
	var filesystemType *Field
	var options *Field
	var freq *Field
	var pass *Field

	switch len(fields) {
	case 6:
		field := fields[5]
		start := uint32(field[0])
		end := uint32(field[1])

		pass = &Field{
			Value: line[start:end],
			Start: start,
			End:   end,
		}
		fallthrough
	case 5:
		field := fields[4]
		start := uint32(field[0])
		end := uint32(field[1])

		freq = &Field{
			Value: line[start:end],
			Start: start,
			End:   end,
		}
		fallthrough
	case 4:
		field := fields[3]
		start := uint32(field[0])
		end := uint32(field[1])

		options = &Field{
			Value: line[start:end],
			Start: start,
			End:   end,
		}
		fallthrough
	case 3:
		field := fields[2]
		start := uint32(field[0])
		end := uint32(field[1])

		filesystemType = &Field{
			Value: line[start:end],
			Start: start,
			End:   end,
		}
		fallthrough
	case 2:
		field := fields[1]
		start := uint32(field[0])
		end := uint32(field[1])

		mountPoint = &Field{
			Value: line[start:end],
			Start: start,
			End:   end,
		}
		fallthrough
	case 1:
		field := fields[0]
		start := uint32(field[0])
		end := uint32(field[1])

		spec = &Field{
			Value: line[start:end],
			Start: start,
			End:   end,
		}
	}

	entry := FstabEntry{
		Type: FstabEntryTypeLine,
		Line: FstabLine{
			Line: uint32(lineNumber),
			Fields: FstabFields{
				Spec:           spec,
				MountPoint:     mountPoint,
				FilesystemType: filesystemType,
				Options:        options,
				Freq:           freq,
				Pass:           pass,
			},
		},
	}
	p.entries = append(p.entries, entry)

	return nil
}

func (p *FstabParser) AddCommentLine(line string, lineNumber int) {
	entry := FstabLine{
		Line: uint32(lineNumber),
	}
	p.entries = append(p.entries, FstabEntry{
		Type: FstabEntryTypeComment,
		Line: entry,
	})
}

func (p *FstabParser) ParseFromContent(content string) []common.ParseError {
	errors := []common.ParseError{}
	lines := strings.Split(content, "\n")

	for index, line := range lines {
		if ignoreLinePattern.MatchString(line) {
			continue
		}

		if commentPattern.MatchString(line) {
			p.AddCommentLine(line, index)
			continue
		}

		err := p.AddLine(line, index)

		if err != nil {
			errors = append(errors, common.ParseError{
				Line: uint32(index),
				Err:  err,
			})
		}
	}

	return errors
}

func (p *FstabParser) GetEntry(line uint32) (*FstabEntry, bool) {
	index, found := slices.BinarySearchFunc(p.entries, line, func(entry FstabEntry, line uint32) int {
		if entry.Line.Line < line {
			return -1
		}

		if entry.Line.Line > line {
			return 1
		}

		return 0
	})

	if !found {
		return nil, false
	}

	return &p.entries[index], true
}

func (p *FstabParser) Clear() {
	p.entries = []FstabEntry{}
}

func (p *FstabParser) AnalyzeValues() []protocol.Diagnostic {
	diagnostics := []protocol.Diagnostic{}

	for _, entry := range p.entries {
		switch entry.Type {
		case FstabEntryTypeLine:
			newDiagnostics := entry.Line.CheckIsValid()

			if len(newDiagnostics) > 0 {
				diagnostics = append(diagnostics, newDiagnostics...)
			}
		case FstabEntryTypeComment:
			// Do nothing
		}
	}

	return diagnostics
}
