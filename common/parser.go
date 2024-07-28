package common

import (
	"fmt"
	"regexp"
	"strings"
)

type SimpleConfigPosition struct {
	Line int
}

type SimpleConfigLine struct {
	Value string
	Position SimpleConfigPosition
}

func (l SimpleConfigLine) IsCursorAtRootOption(cursor int) bool {
	if cursor <= len(l.Value) {
		return true
	}

	return false
}

type SimpleConfigOptions struct {
	Separator string
	IgnorePattern regexp.Regexp
	AvailableOptions *map[string]Option
}

type SimpleConfigParser struct {
	Lines map[string]SimpleConfigLine
	Options SimpleConfigOptions
}

type OptionAlreadyExistsError struct {
	Option string
}
func (e OptionAlreadyExistsError) Error() string {
	return fmt.Sprintf("Option %s already exists", e.Option)
}
type OptionUnknownError struct {
	Option string
}
func (e OptionUnknownError) Error() string {
	return fmt.Sprintf("Option '%s' does not exist", e.Option)
}
type MalformedLineError struct {
	Line string
}
func (e MalformedLineError) Error() string {
	return fmt.Sprintf("Malformed line: %s", e.Line)
}
type LineNotFoundError struct {}
func (e LineNotFoundError) Error() string {
	return "Line not found"
}

func (p *SimpleConfigParser) AddLine(line string, lineNumber int) error {
	parts := strings.SplitN(line, p.Options.Separator, 2)

	if len(parts) == 0 {
		return MalformedLineError{
			Line: line,
		}
	}

	option := parts[0]

	if _, exists := (*p.Options.AvailableOptions)[option]; !exists {
		return OptionUnknownError{
			Option: option,
		}
	}

	value := ""

	if len(parts) > 1 {
		value = parts[1]
	}

	if _, exists := p.Lines[option]; exists {
		return OptionAlreadyExistsError{
			Option: option,
		}
	}

	p.Lines[option] = SimpleConfigLine{
		Value: value,
		Position: SimpleConfigPosition{
			Line: lineNumber,
		},
	}

	return nil
	
}

func (p *SimpleConfigParser) ReplaceOption(option string, value string) {
	p.Lines[option] = SimpleConfigLine{
		Value: value,
		Position: SimpleConfigPosition{
			Line: p.Lines[option].Position.Line,
		},
	}
}

func (p *SimpleConfigParser) RemoveOption(option string) {
	delete(p.Lines, option)
}

func (p *SimpleConfigParser) UpsertOption(option string, value string) {
	if _, exists := p.Lines[option]; exists {
		p.ReplaceOption(option, value)
	} else {
		p.AddLine(option + p.Options.Separator + value, len(p.Lines))
	}
}

func (p *SimpleConfigParser) ParseFromFile(content string) []error {
	lines := strings.Split(content, "\n")
	errors := make([]error, 0)

	for index, line := range lines {
		if p.Options.IgnorePattern.MatchString(line) {
			continue
		}

		err := p.AddLine(line, index)

		if err != nil {
			errors = append(errors, err)
		}
	}

	return errors
}

func (p *SimpleConfigParser) Clear() {
	clear(p.Lines)
}

// TODO: Use better approach: Store an extra array of lines in order; with references to the SimpleConfigLine
func (p SimpleConfigParser) FindByLineNumber(lineNumber int) (string, SimpleConfigLine, error) {
	for option, line := range p.Lines {
		if line.Position.Line == lineNumber {
			return option, line, nil
		}
	}

	return "", SimpleConfigLine{Value: "", Position: SimpleConfigPosition{Line: 0}}, LineNotFoundError{}
}

