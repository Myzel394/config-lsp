package common

import protocol "github.com/tliron/glsp/protocol_3_16"

type Location struct {
	Line      uint32
	Character uint32
}

type LocationRange struct {
	Start Location
	End   Location
}

func (l LocationRange) ToLSPRange() protocol.Range {
	return protocol.Range{
		Start: protocol.Position{
			Line:      l.Start.Line,
			Character: l.Start.Character,
		},
		End: protocol.Position{
			Line:      l.End.Line,
			Character: l.End.Character,
		},
	}
}

func (l *LocationRange) ChangeBothLines(newLine uint32) {
	l.Start.Line = newLine
	l.End.Line = newLine
}

func CreateFullLineRange(line uint32) LocationRange {
	return LocationRange{
		Start: Location{
			Line:      line,
			Character: 0,
		},
		End: Location{
			Line:      line,
			Character: 999999,
		},
	}
}

func CreateSingleCharRange(line uint32, character uint32) LocationRange {
	return LocationRange{
		Start: Location{
			Line:      line,
			Character: character,
		},
		End: Location{
			Line:      line,
			Character: character,
		},
	}
}
