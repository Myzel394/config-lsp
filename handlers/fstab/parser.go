package fstab

import (
	"config-lsp/common"
	docvalues "config-lsp/doc-values"
	"fmt"
	"regexp"
	"slices"
	"strings"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

var ignoreLinePattern = regexp.MustCompile(`^\s*(#|$)`)
var whitespacePattern = regexp.MustCompile(`\s+`)

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
	return f.Value
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

type FstabFields struct {
	Spec           *Field
	MountPoint     *Field
	FilesystemType *Field
	Options        *Field
	Freq           *Field
	Pass           *Field
}

func (f FstabFields) String() string {
	return fmt.Sprintf("Spec: %s, MountPoint: %s, FilesystemType: %s, Options: %s, Freq: %s, Pass: %s", f.Spec, f.MountPoint, f.FilesystemType, f.Options, f.Freq, f.Pass)
}

type FstabEntry struct {
	Line   uint32
	Fields FstabFields
}

func (e *FstabEntry) CheckIsValid() []protocol.Diagnostic {
	println(fmt.Sprintf("Checking entry at line %d; fields: %v", e.Line, e.Fields))
	diagnostics := make([]protocol.Diagnostic, 0)
	severity := protocol.DiagnosticSeverityError

	if e.Fields.Spec != nil {
		err := specField.CheckIsValid(e.Fields.Spec.Value)

		if err != nil {
			diagnostics = append(diagnostics, protocol.Diagnostic{
				Range:    e.Fields.Spec.CreateRange(e.Line),
				Message:  err.Error(),
				Severity: &severity,
			})
		}
	}

	if e.Fields.MountPoint != nil {
		err := mountPointField.CheckIsValid(e.Fields.MountPoint.Value)

		if err != nil {
			diagnostics = append(diagnostics, protocol.Diagnostic{
				Range:    e.Fields.Spec.CreateRange(e.Line),
				Message:  err.Error(),
				Severity: &severity,
			})
		}
	}

	var fileSystemType string = ""

	if e.Fields.FilesystemType != nil {
		err := fileSystemTypeField.CheckIsValid(e.Fields.FilesystemType.Value)

		if err != nil {
			diagnostics = append(diagnostics, protocol.Diagnostic{
				Range:    e.Fields.FilesystemType.CreateRange(e.Line),
				Message:  err.Error(),
				Severity: &severity,
			})
		} else {
			fileSystemType = e.Fields.FilesystemType.Value
		}
	}

	if e.Fields.Options != nil && fileSystemType != "" {
		var optionsField docvalues.Value

		if foundField, found := mountOptionsMapField[fileSystemType]; found {
			optionsField = foundField
		} else {
			optionsField = defaultMountOptionsField
		}

		println(fmt.Sprintf("Checking options for %s", fileSystemType))

		err := optionsField.CheckIsValid(e.Fields.Options.Value)

		if err != nil {
			diagnostics = append(diagnostics, protocol.Diagnostic{
				Range:    e.Fields.Options.CreateRange(e.Line),
				Message:  err.Error(),
				Severity: &severity,
			})
		}
	}

	return diagnostics
}

type FstabParser struct {
	entries []FstabEntry
}

func (p *FstabParser) AddLine(line string, lineNumber int) error {
	fields := whitespacePattern.Split(line, -1)

	if len(fields) == 0 {
		return MalformedLineError{}
	}

	var spec Field
	var mountPoint Field
	var filesystemType Field
	var options Field
	var freq Field
	var pass Field

	switch len(fields) {
	case 6:
		value := fields[5]
		start := uint32(strings.Index(line, value))
		pass = Field{
			Value: fields[5],
			Start: start,
			End:   start + uint32(len(value)),
		}
		fallthrough
	case 5:
		value := fields[4]
		start := uint32(strings.Index(line, value))

		freq = Field{
			Value: value,
			Start: start,
			End:   start + uint32(len(value)),
		}
		fallthrough
	case 4:
		value := fields[3]
		start := uint32(strings.Index(line, value))

		options = Field{
			Value: value,
			Start: start,
			End:   start + uint32(len(value)),
		}
		fallthrough
	case 3:
		value := fields[2]
		start := uint32(strings.Index(line, value))

		filesystemType = Field{
			Value: value,
			Start: start,
			End:   start + uint32(len(value)),
		}
		fallthrough
	case 2:
		value := fields[1]
		start := uint32(strings.Index(line, value))

		mountPoint = Field{
			Value: value,
			Start: start,
			End:   start + uint32(len(value)),
		}
		fallthrough
	case 1:
		value := fields[0]
		start := uint32(strings.Index(line, value))

		spec = Field{
			Value: value,
			Start: start,
			End:   start + uint32(len(value)),
		}
	}

	entry := FstabEntry{
		Line: uint32(lineNumber),
		Fields: FstabFields{
			Spec:           &spec,
			MountPoint:     &mountPoint,
			FilesystemType: &filesystemType,
			Options:        &options,
			Freq:           &freq,
			Pass:           &pass,
		},
	}
	p.entries = append(p.entries, entry)

	return nil
}

func (p *FstabParser) ParseFromContent(content string) []common.ParseError {
	errors := []common.ParseError{}
	lines := strings.Split(content, "\n")

	for index, line := range lines {
		if ignoreLinePattern.MatchString(line) {
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

func (p *FstabParser) GetEntry(line uint32) (FstabEntry, bool) {
	index, found := slices.BinarySearchFunc(p.entries, line, func(entry FstabEntry, line uint32) int {
		if entry.Line < line {
			return -1
		}

		if entry.Line > line {
			return 1
		}

		return 0
	})

	if !found {
		return FstabEntry{}, false
	}

	return p.entries[index], true
}

func (p *FstabParser) Clear() {
	p.entries = []FstabEntry{}
}

func (p *FstabParser) AnalyzeValues() []protocol.Diagnostic {
	diagnostics := []protocol.Diagnostic{}

	for _, entry := range p.entries {
		newDiagnostics := entry.CheckIsValid()

		if len(newDiagnostics) > 0 {
			diagnostics = append(diagnostics, newDiagnostics...)
		}
	}

	return diagnostics
}
