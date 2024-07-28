package common

import (
	"regexp"
	"strings"
)

type SimpleConfigPosition struct {
	Line int
}

type SimpleConfigLine struct {
	Value    string
	Position SimpleConfigPosition
}

func (l SimpleConfigLine) IsCursorAtRootOption(cursor int) bool {
	if cursor <= len(l.Value) {
		return true
	}

	return false
}

type SimpleConfigOptions struct {
	Separator        string
	IgnorePattern    regexp.Regexp
	AvailableOptions *map[string]Option
}

type SimpleConfigParser struct {
	Lines   map[string]SimpleConfigLine
	Options SimpleConfigOptions
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
			Option:      option,
			FoundOnLine: uint32(lineNumber),
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
		p.AddLine(option+p.Options.Separator+value, len(p.Lines))
	}
}

func (p *SimpleConfigParser) GetOption(option string) (SimpleConfigLine, error) {
	if _, exists := p.Lines[option]; exists {
		return p.Lines[option], nil
	}

	return SimpleConfigLine{
			Value: "",
			Position: SimpleConfigPosition{
				Line: 0,
			},
		},
		OptionUnknownError{
			Option: option,
		}
}

func (p *SimpleConfigParser) ParseFromFile(content string) []ParserError {
	lines := strings.Split(content, "\n")
	errors := make([]ParserError, 0)

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
func (p *SimpleConfigParser) FindByLineNumber(lineNumber int) (string, SimpleConfigLine, error) {
	for option, line := range p.Lines {
		if line.Position.Line == lineNumber {
			return option, line, nil
		}
	}

	return "", SimpleConfigLine{Value: "", Position: SimpleConfigPosition{Line: 0}}, LineNotFoundError{}
}
