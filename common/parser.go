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
	Value     string
	Separator string
	Position  SimpleConfigPosition
}

// Get the character positions of [Option End, Separator End, Value End]
func (l SimpleConfigLine) GetCharacterPositions(optionName string) [3]int {
	return [3]int{len(optionName), len(optionName + l.Separator), len(optionName + l.Separator + l.Value)}
}

type SimpleConfigOptions struct {
	Separator     regexp.Regexp
	IgnorePattern regexp.Regexp
	// This is the separator that will be used when adding a new line
	IdealSeparator   string
	AvailableOptions *map[string]Option
}

type SimpleConfigParser struct {
	Lines   map[string]SimpleConfigLine
	Options SimpleConfigOptions
}

func (p *SimpleConfigParser) AddLine(line string, lineNumber uint32) (string, error) {
	var option string
	var separator string
	var value string

	re := p.Options.Separator
	matches := re.FindStringSubmatch(line)

	if len(matches) == 0 {
		return "", docvalues.MalformedLineError{}
	}

	optionIndex := re.SubexpIndex("OptionName")

	if optionIndex == -1 {
		return "", docvalues.MalformedLineError{}
	}

	option = matches[optionIndex]

	if _, exists := (*p.Options.AvailableOptions)[option]; !exists {
		return option, docvalues.OptionUnknownError{}
	}

	separatorIndex := re.SubexpIndex("Separator")

	if separatorIndex == -1 {
		return option, docvalues.MalformedLineError{}
	}

	valueIndex := re.SubexpIndex("Value")

	if valueIndex == -1 {
		return option, docvalues.MalformedLineError{}
	}

	value = matches[valueIndex]

	if _, exists := p.Lines[option]; exists {
		return option, docvalues.OptionAlreadyExistsError{
			AlreadyLine: p.Lines[option].Position.Line,
		}
	}

	p.Lines[option] = SimpleConfigLine{
		Value:     value,
		Separator: separator,
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

func (p *SimpleConfigParser) GetOption(option string) (SimpleConfigLine, error) {
	if _, exists := p.Lines[option]; exists {
		return p.Lines[option], nil
	}

	return SimpleConfigLine{}, docvalues.OptionUnknownError{}
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
