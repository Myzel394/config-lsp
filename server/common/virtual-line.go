package common

import (
	"config-lsp/utils"
)

type VirtualLine struct {
	// This is the true location of the text
	// This ranges from the start of the first line and character
	// to the end of the last line and character
	LocationRange

	Parts []VirtualLinePart
}

func (l VirtualLine) GetText() string {
	text := ""

	for _, part := range l.Parts {
		text += part.Text
	}

	return text
}

// GetSubset Get a subset of the virtual line starting from `start` and ending at `end`
func (l VirtualLine) GetSubset(start uint32, end uint32) VirtualLine {
	parts := make([]VirtualLinePart, 0, 5)
	currentIndex := uint32(0)

	for _, part := range l.Parts {
		partStart := currentIndex
		partEnd := currentIndex + uint32(len(part.Text))

		if partEnd < start {
			continue
		}

		if start <= partEnd {
			var rangeStart uint32
			var rangeEnd uint32

			if start >= partStart {
				rangeStart = start - partStart
			} else {
				rangeStart = 0
			}

			if end <= partEnd {
				rangeEnd = end - partStart
			} else {
				rangeEnd = partEnd
			}

			parts = append(parts, VirtualLinePart{
				LocationRange: LocationRange{
					Start: Location{
						Line:      part.Start.Line,
						Character: part.Start.Character + rangeStart,
					},
					End: Location{
						Line:      part.Start.Line,
						Character: part.Start.Character + rangeEnd,
					},
				},
				Text: part.Text[rangeStart:rangeEnd],
			})
		}

		currentIndex = partEnd

		if currentIndex >= end {
			break
		}
	}

	return VirtualLine{
		LocationRange: LocationRange{
			Start: parts[0].Start,
			End:   parts[len(parts)-1].End,
		},
		Parts: parts,
	}
}

// ConvertRangeToTextRange Convert a given start and end range to a text range
// start and end are the start and end ranges of the virtual line
// This will return the start and end ranges of the actual text lines so that they
// match to the text of the virtual line
// The `start` and `end` are expected to be within the range of the virtual line
func (l VirtualLine) ConvertRangeToTextRange(start uint32, end uint32) []LocationRange {
	virtualLine := l.GetSubset(start, end)

	ranges := make([]LocationRange, 0, 5)

	for _, part := range virtualLine.Parts {
		ranges = append(ranges, part.LocationRange)
	}

	return ranges
}

func (l VirtualLine) AsTrimmed() VirtualLine {
	if len(l.Parts) == 0 {
		// There's nothing that could be trimmed, so we can also just
		// return the original line, as it's identical
		return l
	}

	parts := make([]VirtualLinePart, len(l.Parts))

	for index, part := range l.Parts {
		parts[index] = part.AsTrimmed()
	}

	return VirtualLine{
		LocationRange: LocationRange{
			Start: parts[0].Start,
			End:   parts[len(parts)-1].End,
		},
		Parts: parts,
	}
}

type VirtualLinePart struct {
	// This is the true location of the text
	LocationRange

	Text string
}

func (p VirtualLinePart) AsTrimmed() VirtualLinePart {
	firstNonWhitespace := utils.FindFirstNonMatch(p.Text, UnicodeWhitespace, 0)

	if firstNonWhitespace == -1 {
		// Empty line
		return p
	}

	lastNonWhitespace := utils.FindLastNonMatch(p.Text, UnicodeWhitespace, len(p.Text)-1)

	if lastNonWhitespace == -1 {
		lastNonWhitespace = len(p.Text) - 1
	}

	return VirtualLinePart{
		LocationRange: LocationRange{
			Start: Location{
				Line:      p.Start.Line,
				Character: p.Start.Character + uint32(firstNonWhitespace),
			},
			End: Location{
				Line:      p.Start.Line,
				Character: p.Start.Character + uint32(lastNonWhitespace) + 1,
			},
		},
		Text: p.Text[firstNonWhitespace : lastNonWhitespace+1],
	}
}

func SplitIntoVirtualLines(input string) []VirtualLine {
	stringLines := utils.SplitIntoVirtualLines(input)

	lines := make([]VirtualLine, 0, len(stringLines))

	for rawLineNumber, line := range stringLines {
		parts := make([]VirtualLinePart, 0)

		for virtualLineNumber, part := range line {
			if part == "" {
				continue
			}

			lineNumber := uint32(rawLineNumber) + uint32(virtualLineNumber)

			parts = append(parts, VirtualLinePart{
				LocationRange: LocationRange{
					Start: Location{
						Line:      lineNumber,
						Character: 0,
					},
					End: Location{
						Line:      lineNumber,
						Character: uint32(len(part)),
					},
				},
				Text: part,
			})
		}

		if len(parts) == 0 {
			continue
		}

		lines = append(lines, VirtualLine{
			LocationRange: LocationRange{
				Start: parts[0].Start,
				End:   parts[len(parts)-1].End,
			},
			Parts: parts,
		})
	}

	return lines
}
