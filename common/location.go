package common

type Location struct {
	Line      uint32
	Character uint32
}

type LocationRange struct {
	Start Location
	End   Location
}

func (l *LocationRange) ChangeBothLines(newLine uint32) {
	l.Start.Line = newLine
	l.End.Line = newLine
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
