package common

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

type Location struct {
	Line      uint32
	Character uint32
}

type LocationRange struct {
	Start Location
	End   Location
}

func (l LocationRange) ShiftHorizontal(offset uint32) LocationRange {
	return LocationRange{
		Start: Location{
			Line:      l.Start.Line,
			Character: l.Start.Character + offset,
		},
		End: Location{
			Line:      l.End.Line,
			Character: l.End.Character + offset,
		},
	}
}

func (l LocationRange) String() string {
	if l.Start.Line == l.End.Line {
		return fmt.Sprintf("%d:%d-%d", l.Start.Line, l.Start.Character, l.End.Character)
	}

	return fmt.Sprintf("%d:%d-%d:%d", l.Start.Line, l.Start.Character, l.End.Line, l.End.Character)
}

var GlobalLocationRange = LocationRange{
	Start: Location{
		Line:      0,
		Character: 0,
	},
	End: Location{
		Line:      0,
		Character: 0,
	},
}

func (l LocationRange) ToLSPRange() protocol.Range {
	return protocol.Range{
		Start: protocol.Position{
			Line:      l.Start.Line,
			Character: l.Start.Character,
		},
		End: protocol.Position{
			Line:      l.End.Line,
			Character: l.End.Character + 1,
		},
	}
}

func (l *LocationRange) ChangeBothLines(newLine uint32) {
	l.Start.Line = newLine
	l.End.Line = newLine
}

func (l LocationRange) ContainsCursor(line uint32, character uint32) bool {
	return line == l.Start.Line && character >= l.Start.Character && character <= l.End.Character
}

func (l LocationRange) ContainsCursorByCharacter(character uint32) bool {
	return character >= l.Start.Character && character <= l.End.Character
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

func CharacterRangeFromCtx(
	ctx antlr.BaseParserRuleContext,
) LocationRange {
	line := uint32(ctx.GetStart().GetLine())
	start := uint32(ctx.GetStart().GetStart())
	end := uint32(ctx.GetStop().GetStop())

	return LocationRange{
		Start: Location{
			Line:      line,
			Character: start,
		},
		End: Location{
			Line:      line,
			Character: end,
		},
	}
}
