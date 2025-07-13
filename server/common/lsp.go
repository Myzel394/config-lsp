package common

import protocol "github.com/tliron/glsp/protocol_3_16"

// LSPCharacterAsCursorPosition:
// @deprecated
func CursorToCharacterIndex(cursor uint32) uint32 {
	return max(1, cursor) - 1
}

func DeprecatedImprovedCursorToIndex(
	c CursorPosition,
	line string,
	offset uint32,
) uint32 {
	if len(line) == 0 {
		return 0
	}

	return min(uint32(len(line)-1), uint32(c)-offset)
}

var SeverityError = protocol.DiagnosticSeverityError
var SeverityWarning = protocol.DiagnosticSeverityWarning
var SeverityInformation = protocol.DiagnosticSeverityInformation
var SeverityHint = protocol.DiagnosticSeverityHint
