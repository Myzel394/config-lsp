package ast

import (
	"config-lsp/common"
	commonparser "config-lsp/common/parser"
	"config-lsp/utils"
	"regexp"

	"github.com/emirpasic/gods/maps/treemap"

	gods "github.com/emirpasic/gods/utils"
)

func NewFstabConfig() *FstabConfig {
	config := &FstabConfig{}
	config.Clear()

	return config
}

func (c *FstabConfig) Clear() {
	c.Entries = treemap.NewWith(gods.UInt32Comparator)
	c.CommentLines = map[uint32]struct{}{}
}

var commentPattern = regexp.MustCompile(`^\s*#`)
var emptyPattern = regexp.MustCompile(`^\s*$`)
var whitespacePattern = regexp.MustCompile(`\S+`)

func (c *FstabConfig) Parse(input string) []common.LSPError {
	errors := make([]common.LSPError, 0)
	lines := utils.SplitIntoLines(input)

	for rawLineNumber, line := range lines {
		lineNumber := uint32(rawLineNumber)

		if emptyPattern.MatchString(line) {
			continue
		}

		if commentPattern.MatchString(line) {
			c.CommentLines[lineNumber] = struct{}{}
			continue
		}

		errors = append(
			errors,
			c.parseStatement(lineNumber, line)...,
		)
	}

	return errors
}

func (c *FstabConfig) parseStatement(
	line uint32,
	input string,
) []common.LSPError {
	fields := whitespacePattern.FindAllStringIndex(input, -1)

	if len(fields) == 0 {
		return []common.LSPError{
			{
				Range: common.LocationRange{
					Start: common.Location{
						Line:      line,
						Character: 0,
					},
					End: common.Location{
						Line:      line,
						Character: 0,
					},
				},
			},
		}
	}

	var spec *FstabField
	var mountPoint *FstabField
	var filesystemType *FstabField
	var options *FstabField
	var freq *FstabField
	var pass *FstabField

	switch len(fields) {
	case 6:
		pass = parseField(line, input, fields[5])
		fallthrough
	case 5:
		freq = parseField(line, input, fields[4])

		if pass == nil && fields[4][1] < len(input) {
			pass = createPartialField(line, input, fields[4][1], len(input))
		}

		fallthrough
	case 4:
		options = parseField(line, input, fields[3])

		if freq == nil && fields[3][1] < len(input) {
			freq = createPartialField(line, input, fields[3][1], len(input))
		}

		fallthrough
	case 3:
		filesystemType = parseField(line, input, fields[2])

		if options == nil && fields[2][1] < len(input) {
			options = createPartialField(line, input, fields[2][1], len(input))
		}

		fallthrough
	case 2:
		mountPoint = parseField(line, input, fields[1])

		if filesystemType == nil && fields[1][1] < len(input) {
			filesystemType = createPartialField(line, input, fields[1][1], len(input))
		}

		fallthrough
	case 1:
		spec = parseField(line, input, fields[0])

		if mountPoint == nil && fields[0][1] < len(input) {
			mountPoint = createPartialField(line, input, fields[0][1], len(input))
		}
	}

	fstabLine := &FstabEntry{
		Fields: FstabFields{
			LocationRange: common.LocationRange{
				Start: common.Location{
					Line:      line,
					Character: 0,
				},
				End: common.Location{
					Line:      line,
					Character: uint32(len(input)),
				},
			},
			Spec:           spec,
			MountPoint:     mountPoint,
			FilesystemType: filesystemType,
			Options:        options,
			Freq:           freq,
			Pass:           pass,
		},
	}

	c.Entries.Put(line, fstabLine)

	return nil
}

func parseField(
	line uint32,
	input string,
	field []int,
) *FstabField {
	start := uint32(field[0])
	end := uint32(field[1])
	value := input[start:end]

	return &FstabField{
		LocationRange: common.LocationRange{
			Start: common.Location{
				Line:      line,
				Character: start,
			},
			End: common.Location{
				Line:      line,
				Character: end,
			},
		},
		Value: commonparser.ParseRawString(value, commonparser.ParseFeatures{
			ParseEscapedCharacters: true,
			ParseDoubleQuotes:      true,
			Replacements: &map[string]string{
				`\\040`: " ",
			},
		}),
	}
}

func createPartialField(
	line uint32,
	input string,
	start int,
	end int,
) *FstabField {
	return nil
	return &FstabField{
		LocationRange: common.LocationRange{
			Start: common.Location{
				Line:      line,
				Character: uint32(start),
			},
			End: common.Location{
				Line:      line,
				Character: uint32(end),
			},
		},
		Value: commonparser.ParseRawString(input[end:end], commonparser.ParseFeatures{
			ParseEscapedCharacters: true,
			ParseDoubleQuotes:      true,
			Replacements:           &map[string]string{},
		}),
	}
}
