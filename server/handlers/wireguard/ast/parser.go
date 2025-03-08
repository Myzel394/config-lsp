package ast

import (
	"config-lsp/common"
	"config-lsp/utils"
	"fmt"
	"regexp"
	"strings"
)

func NewWGConfig() *WGConfig {
	config := &WGConfig{}
	config.Clear()

	return config
}

func (c *WGConfig) Clear() {
	c.Sections = make([]*WGSection, 0, 2)
	c.CommentLines = make(map[uint32]struct{})
}

var commentPattern = regexp.MustCompile(`^\s*([;#])`)
var emptyPattern = regexp.MustCompile(`^\s*$`)
var headerPattern = regexp.MustCompile(`^\s*\[(\w+)?]?`)
var linePattern = regexp.MustCompile(`^\s*(?P<key>.+?)\s*(?P<separator>=)\s*(?P<value>\S.*?)?\s*(?:[;#].*)?\s*$`)

func (c *WGConfig) Parse(input string) []common.LSPError {
	errors := make([]common.LSPError, 0)
	lines := utils.SplitIntoLines(input)

	var currentSection *WGSection

	for rawLineNumber, line := range lines {
		lineNumber := uint32(rawLineNumber)

		if emptyPattern.MatchString(line) {
			// Set end of last section
			if currentSection != nil {
				currentSection.End.Line = lineNumber
				currentSection.End.Character = 0
			}
			continue
		}

		if commentPattern.MatchString(line) {
			c.CommentLines[lineNumber] = struct{}{}
			// Set end of last section
			if currentSection != nil {
				currentSection.End.Line = lineNumber
				currentSection.End.Character = uint32(len(line))
			}
			continue
		}

		if headerPattern.MatchString(line) {
			name := headerPattern.FindStringSubmatch(line)[1]

			currentSection = &WGSection{
				LocationRange: common.LocationRange{
					Start: common.Location{
						Line:      lineNumber,
						Character: 0,
					},
					End: common.Location{
						Line:      lineNumber,
						Character: uint32(len(line)) + 1,
					},
				},
				Header: WGHeader{
					LocationRange: common.LocationRange{
						Start: common.Location{
							Line:      lineNumber,
							Character: 0,
						},
						End: common.Location{
							Line:      lineNumber,
							Character: uint32(len(line)) + 1,
						},
					},
					Name: name,
				},
				Properties: make(map[uint32]*WGProperty),
			}

			c.Sections = append(c.Sections, currentSection)

			continue
		}

		// Else property

		// Set end of last section
		if currentSection != nil {
			currentSection.End.Line = lineNumber
			currentSection.End.Character = uint32(len(line))
		}

		if currentSection == nil {
			// Root properties are not allowed
			errors = append(errors, common.LSPError{
				Range: common.LocationRange{
					Start: common.Location{
						Line:      lineNumber,
						Character: 0,
					},
					End: common.Location{
						Line:      lineNumber,
						Character: uint32(len(line)),
					},
				},
				Err: fmt.Errorf("A header is missing before a property. This property has no header above it."),
			})

			continue
		}

		if !strings.Contains(line, "=") {
			// Incomplete property
			indexes := utils.GetTrimIndex(line)

			currentSection.Properties[lineNumber] = &WGProperty{
				Key: WGPropertyKey{
					LocationRange: common.LocationRange{
						Start: common.Location{
							Line:      lineNumber,
							Character: uint32(indexes[0]),
						},
						End: common.Location{
							Line:      lineNumber,
							Character: uint32(indexes[1]),
						},
					},
					Name: line[indexes[0]:indexes[1]],
				},
			}
		} else {
			// Fully written out property

			indexes := linePattern.FindStringSubmatchIndex(line)

			if indexes == nil || len(indexes) == 0 {
				// Error
				errors = append(errors, common.LSPError{
					Range: common.LocationRange{
						Start: common.Location{
							Line:      lineNumber,
							Character: 0,
						},
						End: common.Location{
							Line:      lineNumber,
							Character: uint32(len(line)),
						},
					},
					Err: fmt.Errorf("This property seems to be malformed"),
				})

				continue
			}

			// Construct key
			keyStart := uint32(indexes[2])
			keyEnd := uint32(indexes[3])
			key := WGPropertyKey{
				LocationRange: common.LocationRange{
					Start: common.Location{
						Line:      lineNumber,
						Character: keyStart,
					},
					End: common.Location{
						Line:      lineNumber,
						Character: keyEnd,
					},
				},
				Name: line[keyStart:keyEnd],
			}

			// Construct separator
			separatorStart := uint32(indexes[4])
			separatorEnd := uint32(indexes[5])
			separator := WGPropertySeparator{
				LocationRange: common.LocationRange{
					Start: common.Location{
						Line:      lineNumber,
						Character: separatorStart,
					},
					End: common.Location{
						Line:      lineNumber,
						Character: separatorEnd,
					},
				},
			}

			// Construct value
			var value *WGPropertyValue
			propertyEnd := uint32(len(line))

			if indexes[6] != -1 && indexes[7] != -1 {
				// value exists
				valueStart := uint32(indexes[6])
				valueEnd := uint32(indexes[7])
				propertyEnd = valueEnd

				value = &WGPropertyValue{
					LocationRange: common.LocationRange{
						Start: common.Location{
							Line:      lineNumber,
							Character: valueStart,
						},
						End: common.Location{
							Line:      lineNumber,
							Character: valueEnd,
						},
					},
					Value: line[valueStart:valueEnd],
				}
			}

			// And lastly, add the property
			currentSection.Properties[lineNumber] = &WGProperty{
				LocationRange: common.LocationRange{
					Start: common.Location{
						Line:      lineNumber,
						Character: keyStart,
					},
					End: common.Location{
						Line:      lineNumber,
						Character: propertyEnd,
					},
				},
				RawValue:  line,
				Key:       key,
				Separator: &separator,
				Value:     value,
			}
		}
	}

	return errors
}
