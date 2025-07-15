package ini

import (
	"config-lsp/common"
	"config-lsp/utils"
	"fmt"
	"regexp"
	"strings"

	"github.com/emirpasic/gods/maps/treemap"
	gods "github.com/emirpasic/gods/utils"
)

// NewConfig creates a new empty INI configuration
func NewConfig() *Config {
	config := &Config{}
	config.Clear()

	return config
}

// Clear resets the configuration
func (c *Config) Clear() {
	c.Sections = make([]*Section, 0, 2)
	c.CommentLines = make(map[uint32]struct{})
}

var commentPattern = regexp.MustCompile(`^\s*([;#])`)
var emptyPattern = regexp.MustCompile(`^\s*$`)
var headerPattern = regexp.MustCompile(`^\s*\[(\w+)?]?`)
var linePattern = regexp.MustCompile(`^\s*(?P<key>.+?)\s*(?P<separator>=)\s*(?P<value>\S.*?)?\s*(?: [;#].*)?\s*$`)

// Parse parses an INI string and returns any errors encountered
func (c *Config) Parse(input string) []common.LSPError {
	errors := make([]common.LSPError, 0)
	lines := utils.SplitIntoLines(input)

	var currentSection *Section

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

			currentSection = &Section{
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
				Header: &Header{
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
				Properties: treemap.NewWith(gods.UInt32Comparator),
			}

			c.Sections = append(c.Sections, currentSection)

			continue
		}

		///// Else property

		if currentSection == nil {
			if c.XParseConfig.AllowRootProperties {
				currentSection = &Section{
					LocationRange: common.LocationRange{
						Start: common.Location{
							Line:      lineNumber,
							Character: 0,
						},
						End: common.Location{
							Line:      lineNumber,
							Character: uint32(len(line)),
						},
					},
					Header:     nil, // No header for empty Sections
					Properties: treemap.NewWith(gods.UInt32Comparator),
				}

				c.Sections = append(c.Sections, currentSection)
			} else {
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
		}

		// Set end of last section
		currentSection.End.Line = lineNumber
		currentSection.End.Character = uint32(len(line))

		if !strings.Contains(line, "=") {
			// Incomplete property

			indexes := utils.GetTrimIndex(line)

			newProperty := &Property{
				Key: PropertyKey{
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

			currentSection.Properties.Put(lineNumber, newProperty)
		} else {
			// Fully written out property

			indexes := linePattern.FindStringSubmatchIndex(line)

			if len(indexes) == 0 {
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
			key := PropertyKey{
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
			separator := PropertySeparator{
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
			var value *PropertyValue
			propertyEnd := uint32(len(line))

			if indexes[6] != -1 && indexes[7] != -1 {
				// value exists
				valueStart := uint32(indexes[6])
				valueEnd := uint32(indexes[7])
				propertyEnd = valueEnd

				value = &PropertyValue{
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
			newProperty := &Property{
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
			currentSection.Properties.Put(lineNumber, newProperty)
		}
	}

	// Edge case, empty file
	if len(c.Sections) == 0 && c.XParseConfig.AllowRootProperties {
		// If the file is empty, we create a root section
		c.Sections = append(c.Sections, &Section{
			LocationRange: common.LocationRange{
				Start: common.Location{
					Line:      0,
					Character: 0,
				},
				End: common.Location{
					Line:      uint32(len(lines)),
					Character: 0,
				},
			},
			Header:     nil, // No header for empty Sections
			Properties: treemap.NewWith(gods.UInt32Comparator),
		})
	}

	return errors
}
