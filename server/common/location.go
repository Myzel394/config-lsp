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

func (l Location) GetRelativeIndexPosition(i IndexPosition) IndexPosition {
	return i - IndexPosition(l.Character)
}

func (l Location) ToLSPPosition() protocol.Position {
	return protocol.Position{
		Line:      l.Line,
		Character: l.Character,
	}
}

// LocationRange: Represents a range of characters in a document
// Locations are zero-based, start-inclusive and end-exclusive
// This approach is preferred over using an index-based range, because
// it allows to check very easily for cursor positions, as well as for
// index-based ranges.
type LocationRange struct {
	Start Location
	End   Location
}

func (l LocationRange) ContainsPosition(p Position) bool {
	return l.IsPositionAfterStart(p) && l.IsPositionBeforeEnd(p)
}

// Check if the given position is after the start of the range
// It's like: Position >= Start
// This checks inclusively
func (l LocationRange) IsPositionAfterStart(p Position) bool {
	return p.getValue() >= l.Start.Character
}

func (l LocationRange) IsPositionBeforeStart(p Position) bool {
	return p.getValue() < l.Start.Character
}

// Check if the given position is before the end of the range
// It's like: Position <= End
// This checks inclusively
func (l LocationRange) IsPositionBeforeEnd(p Position) bool {
	switch p.(type) {
	case CursorPosition:
		return p.getValue() <= l.End.Character
	case IndexPosition:
		return p.getValue() < l.End.Character
	}

	return false
}

func (l LocationRange) IsPositionAfterEnd(p Position) bool {
	switch p.(type) {
	case CursorPosition:
		return p.getValue() > l.End.Character
	case IndexPosition:
		return p.getValue() >= l.End.Character
	}

	return false
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

// This will include the next line, so the end will be at the start of the next line
func (l LocationRange) IncludeNextLine() LocationRange {
	return LocationRange{
		Start: l.Start,
		End: Location{
			Line:      l.End.Line + 1,
			Character: 0,
		},
	}
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

func (l LocationRange) ChangeBothLines(newLine uint32) LocationRange {
	return LocationRange{
		Start: Location{
			Line:      newLine,
			Character: l.Start.Character,
		},
		End: Location{
			Line:      newLine,
			Character: l.End.Character,
		},
	}
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
			Character: end + 1,
		},
	}
}

type Position interface {
	getValue() uint32
}

// Use this type if you want to use a cursor based position
// A cursor based position is a position that represents a cursor
// Given the example:
// "PermitRootLogin yes"
// Taking a look at the first character "P" - the index is 0.
// However, the cursor can either be at:
//
//	"|P" - 0 or
//	"P|" - 1
//
// This is used for example for textDocument/completion or textDocument/signature
type CursorPosition uint32

func (c CursorPosition) getValue() uint32 {
	return uint32(c)
}

func (c CursorPosition) ShiftHorizontal(offset int) CursorPosition {
	return CursorPosition(uint32(int(c) + offset))
}

func LSPCharacterAsCursorPosition(character uint32) CursorPosition {
	return CursorPosition(character)
}

func (c CursorPosition) IsBeforeIndexPosition(i IndexPosition) bool {
	// |H[e]llo
	return uint32(c) < uint32(i)
}

func (c CursorPosition) IsAfterIndexPosition(i IndexPosition) bool {
	// H[e]|llo
	return uint32(c) > uint32(i)+1
}

// Get the byte that is before the cursor position
// This expects that the cursor is not out of bounds
func (c CursorPosition) GetCharacterBefore(value string) byte {
	if c.getValue() == 0 {
		return value[0]
	} else {
		return value[max(0, c.getValue()-1)]
	}
}

// Get the byte that is after the cursor position
// This expects that the cursor is not out of bounds
func (c CursorPosition) GetCharacterAfter(value string) byte {
	if c.getValue() >= uint32(len(value)) {
		return value[len(value)-1]
	} else {
		return value[c.getValue()]
	}
}

func (c CursorPosition) IsAtEdge(value string) bool {
	// If the cursor is at the start or end of the value
	return c.getValue() == 0 || c.getValue() >= uint32(len(value))
}

// Use this type if you want to use an index based position
type IndexPosition uint32

func (i IndexPosition) getValue() uint32 {
	return uint32(i)
}

func LSPCharacterAsIndexPosition(character uint32) IndexPosition {
	return IndexPosition(character)
}
