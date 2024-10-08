// Code generated from Match.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type MatchLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var MatchLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func matchlexerLexerInit() {
	staticData := &MatchLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "','", "", "", "", "", "", "", "", "", "", "", "", "", "", "'\"'",
	}
	staticData.SymbolicNames = []string{
		"", "COMMA", "ALL", "CANONICAL", "FINAL", "EXEC", "LOCALNETWORK", "HOST",
		"ORIGINALHOST", "TAGGED", "USER", "LOCALUSER", "STRING", "WHITESPACE",
		"QUOTED_STRING", "QUOTE",
	}
	staticData.RuleNames = []string{
		"COMMA", "ALL", "CANONICAL", "FINAL", "EXEC", "LOCALNETWORK", "HOST",
		"ORIGINALHOST", "TAGGED", "USER", "LOCALUSER", "STRING", "WHITESPACE",
		"QUOTED_STRING", "QUOTE",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 15, 141, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 1, 0,
		1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 4, 11, 113,
		8, 11, 11, 11, 12, 11, 114, 1, 12, 4, 12, 118, 8, 12, 11, 12, 12, 12, 119,
		1, 13, 1, 13, 3, 13, 124, 8, 13, 1, 13, 1, 13, 1, 13, 5, 13, 129, 8, 13,
		10, 13, 12, 13, 132, 9, 13, 1, 13, 3, 13, 135, 8, 13, 1, 13, 3, 13, 138,
		8, 13, 1, 14, 1, 14, 0, 0, 15, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13,
		7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 1, 0,
		20, 2, 0, 65, 65, 97, 97, 2, 0, 76, 76, 108, 108, 2, 0, 67, 67, 99, 99,
		2, 0, 78, 78, 110, 110, 2, 0, 79, 79, 111, 111, 2, 0, 73, 73, 105, 105,
		2, 0, 70, 70, 102, 102, 2, 0, 69, 69, 101, 101, 2, 0, 88, 88, 120, 120,
		2, 0, 84, 84, 116, 116, 2, 0, 87, 87, 119, 119, 2, 0, 82, 82, 114, 114,
		2, 0, 75, 75, 107, 107, 2, 0, 72, 72, 104, 104, 2, 0, 83, 83, 115, 115,
		2, 0, 71, 71, 103, 103, 2, 0, 68, 68, 100, 100, 2, 0, 85, 85, 117, 117,
		5, 0, 9, 10, 13, 13, 32, 32, 35, 35, 44, 44, 2, 0, 9, 9, 32, 32, 146, 0,
		1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0,
		9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0,
		0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0,
		0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 1, 31, 1, 0,
		0, 0, 3, 33, 1, 0, 0, 0, 5, 37, 1, 0, 0, 0, 7, 47, 1, 0, 0, 0, 9, 53, 1,
		0, 0, 0, 11, 58, 1, 0, 0, 0, 13, 71, 1, 0, 0, 0, 15, 76, 1, 0, 0, 0, 17,
		89, 1, 0, 0, 0, 19, 96, 1, 0, 0, 0, 21, 101, 1, 0, 0, 0, 23, 112, 1, 0,
		0, 0, 25, 117, 1, 0, 0, 0, 27, 121, 1, 0, 0, 0, 29, 139, 1, 0, 0, 0, 31,
		32, 5, 44, 0, 0, 32, 2, 1, 0, 0, 0, 33, 34, 7, 0, 0, 0, 34, 35, 7, 1, 0,
		0, 35, 36, 7, 1, 0, 0, 36, 4, 1, 0, 0, 0, 37, 38, 7, 2, 0, 0, 38, 39, 7,
		0, 0, 0, 39, 40, 7, 3, 0, 0, 40, 41, 7, 4, 0, 0, 41, 42, 7, 3, 0, 0, 42,
		43, 7, 5, 0, 0, 43, 44, 7, 2, 0, 0, 44, 45, 7, 0, 0, 0, 45, 46, 7, 1, 0,
		0, 46, 6, 1, 0, 0, 0, 47, 48, 7, 6, 0, 0, 48, 49, 7, 5, 0, 0, 49, 50, 7,
		3, 0, 0, 50, 51, 7, 0, 0, 0, 51, 52, 7, 1, 0, 0, 52, 8, 1, 0, 0, 0, 53,
		54, 7, 7, 0, 0, 54, 55, 7, 8, 0, 0, 55, 56, 7, 7, 0, 0, 56, 57, 7, 2, 0,
		0, 57, 10, 1, 0, 0, 0, 58, 59, 7, 1, 0, 0, 59, 60, 7, 4, 0, 0, 60, 61,
		7, 2, 0, 0, 61, 62, 7, 0, 0, 0, 62, 63, 7, 1, 0, 0, 63, 64, 7, 3, 0, 0,
		64, 65, 7, 7, 0, 0, 65, 66, 7, 9, 0, 0, 66, 67, 7, 10, 0, 0, 67, 68, 7,
		4, 0, 0, 68, 69, 7, 11, 0, 0, 69, 70, 7, 12, 0, 0, 70, 12, 1, 0, 0, 0,
		71, 72, 7, 13, 0, 0, 72, 73, 7, 4, 0, 0, 73, 74, 7, 14, 0, 0, 74, 75, 7,
		9, 0, 0, 75, 14, 1, 0, 0, 0, 76, 77, 7, 4, 0, 0, 77, 78, 7, 11, 0, 0, 78,
		79, 7, 5, 0, 0, 79, 80, 7, 15, 0, 0, 80, 81, 7, 5, 0, 0, 81, 82, 7, 3,
		0, 0, 82, 83, 7, 0, 0, 0, 83, 84, 7, 1, 0, 0, 84, 85, 7, 13, 0, 0, 85,
		86, 7, 4, 0, 0, 86, 87, 7, 14, 0, 0, 87, 88, 7, 9, 0, 0, 88, 16, 1, 0,
		0, 0, 89, 90, 7, 9, 0, 0, 90, 91, 7, 0, 0, 0, 91, 92, 7, 15, 0, 0, 92,
		93, 7, 15, 0, 0, 93, 94, 7, 7, 0, 0, 94, 95, 7, 16, 0, 0, 95, 18, 1, 0,
		0, 0, 96, 97, 7, 17, 0, 0, 97, 98, 7, 14, 0, 0, 98, 99, 7, 7, 0, 0, 99,
		100, 7, 11, 0, 0, 100, 20, 1, 0, 0, 0, 101, 102, 7, 1, 0, 0, 102, 103,
		7, 4, 0, 0, 103, 104, 7, 2, 0, 0, 104, 105, 7, 0, 0, 0, 105, 106, 7, 1,
		0, 0, 106, 107, 7, 17, 0, 0, 107, 108, 7, 14, 0, 0, 108, 109, 7, 7, 0,
		0, 109, 110, 7, 11, 0, 0, 110, 22, 1, 0, 0, 0, 111, 113, 8, 18, 0, 0, 112,
		111, 1, 0, 0, 0, 113, 114, 1, 0, 0, 0, 114, 112, 1, 0, 0, 0, 114, 115,
		1, 0, 0, 0, 115, 24, 1, 0, 0, 0, 116, 118, 7, 19, 0, 0, 117, 116, 1, 0,
		0, 0, 118, 119, 1, 0, 0, 0, 119, 117, 1, 0, 0, 0, 119, 120, 1, 0, 0, 0,
		120, 26, 1, 0, 0, 0, 121, 123, 3, 29, 14, 0, 122, 124, 3, 25, 12, 0, 123,
		122, 1, 0, 0, 0, 123, 124, 1, 0, 0, 0, 124, 130, 1, 0, 0, 0, 125, 126,
		3, 23, 11, 0, 126, 127, 3, 25, 12, 0, 127, 129, 1, 0, 0, 0, 128, 125, 1,
		0, 0, 0, 129, 132, 1, 0, 0, 0, 130, 128, 1, 0, 0, 0, 130, 131, 1, 0, 0,
		0, 131, 134, 1, 0, 0, 0, 132, 130, 1, 0, 0, 0, 133, 135, 3, 23, 11, 0,
		134, 133, 1, 0, 0, 0, 134, 135, 1, 0, 0, 0, 135, 137, 1, 0, 0, 0, 136,
		138, 3, 29, 14, 0, 137, 136, 1, 0, 0, 0, 137, 138, 1, 0, 0, 0, 138, 28,
		1, 0, 0, 0, 139, 140, 5, 34, 0, 0, 140, 30, 1, 0, 0, 0, 7, 0, 114, 119,
		123, 130, 134, 137, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// MatchLexerInit initializes any static state used to implement MatchLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewMatchLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func MatchLexerInit() {
	staticData := &MatchLexerLexerStaticData
	staticData.once.Do(matchlexerLexerInit)
}

// NewMatchLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewMatchLexer(input antlr.CharStream) *MatchLexer {
	MatchLexerInit()
	l := new(MatchLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &MatchLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Match.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// MatchLexer tokens.
const (
	MatchLexerCOMMA         = 1
	MatchLexerALL           = 2
	MatchLexerCANONICAL     = 3
	MatchLexerFINAL         = 4
	MatchLexerEXEC          = 5
	MatchLexerLOCALNETWORK  = 6
	MatchLexerHOST          = 7
	MatchLexerORIGINALHOST  = 8
	MatchLexerTAGGED        = 9
	MatchLexerUSER          = 10
	MatchLexerLOCALUSER     = 11
	MatchLexerSTRING        = 12
	MatchLexerWHITESPACE    = 13
	MatchLexerQUOTED_STRING = 14
	MatchLexerQUOTE         = 15
)