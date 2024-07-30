package common

import (
	docvalues "config-lsp/doc-values"
	"regexp"
	"strings"
)

type SimpleConfigPosition struct {
	Line uint32
}

type SimpleConfigLine struct {
	Value    string
	Position SimpleConfigPosition
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

func (p *SimpleConfigParser) AddLine(line string, lineNumber uint32) (string, error) {
	parts := strings.SplitN(line, p.Options.Separator, 2)

	if len(parts) == 0 {
		return "", docvalues.MalformedLineError{}
	}

	option := parts[0]

	if _, exists := (*p.Options.AvailableOptions)[option]; !exists {
		return option, docvalues.OptionUnknownError{}
	}

	value := ""

	if len(parts) > 1 {
		value = parts[1]
	}

	if _, exists := p.Lines[option]; exists {
		return option, docvalues.OptionAlreadyExistsError{
			AlreadyLine: p.Lines[option].Position.Line,
		}
	}

	p.Lines[option] = SimpleConfigLine{
		Value: value,
		Position: SimpleConfigPosition{
			Line: lineNumber,
		},
	}

	return option, nil

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
		p.AddLine(option+p.Options.Separator+value, uint32(len(p.Lines)))
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
		docvalues.OptionUnknownError{}
}

func (p *SimpleConfigParser) ParseFromFile(content string) []docvalues.OptionError {
	lines := strings.Split(content, "\n")
	errors := make([]docvalues.OptionError, 0)

	for index, line := range lines {
		if p.Options.IgnorePattern.MatchString(line) {
			continue
		}

		option, err := p.AddLine(line, uint32(index))

		if err != nil {
			errors = append(errors, docvalues.OptionError{
				Line:           uint32(index),
				ProvidedOption: option,
				DocError:       err,
			})
		}
	}

	return errors
}

func (p *SimpleConfigParser) Clear() {
	clear(p.Lines)
}

// TODO: Use better approach: Store an extra array of lines in order; with references to the SimpleConfigLine
func (p *SimpleConfigParser) FindByLineNumber(lineNumber uint32) (string, SimpleConfigLine, error) {
	for option, line := range p.Lines {
		if line.Position.Line == lineNumber {
			return option, line, nil
		}
	}

	return "", SimpleConfigLine{Value: "", Position: SimpleConfigPosition{Line: 0}}, docvalues.LineNotFoundError{}
}
