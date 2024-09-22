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
		"", "','",
	}
	staticData.SymbolicNames = []string{
		"", "COMMA", "STRING", "WHITESPACE", "QUOTED_STRING",
	}
	staticData.RuleNames = []string{
		"COMMA", "STRING", "WHITESPACE", "QUOTED_STRING",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 4, 39, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 1,
		0, 1, 0, 1, 1, 4, 1, 13, 8, 1, 11, 1, 12, 1, 14, 1, 2, 4, 2, 18, 8, 2,
		11, 2, 12, 2, 19, 1, 3, 1, 3, 3, 3, 24, 8, 3, 1, 3, 1, 3, 1, 3, 5, 3, 29,
		8, 3, 10, 3, 12, 3, 32, 9, 3, 1, 3, 3, 3, 35, 8, 3, 1, 3, 3, 3, 38, 8,
		3, 0, 0, 4, 1, 1, 3, 2, 5, 3, 7, 4, 1, 0, 2, 5, 0, 9, 10, 13, 13, 32, 32,
		35, 35, 44, 44, 2, 0, 9, 9, 32, 32, 44, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0,
		0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 1, 9, 1, 0, 0, 0, 3, 12, 1, 0, 0,
		0, 5, 17, 1, 0, 0, 0, 7, 21, 1, 0, 0, 0, 9, 10, 5, 44, 0, 0, 10, 2, 1,
		0, 0, 0, 11, 13, 8, 0, 0, 0, 12, 11, 1, 0, 0, 0, 13, 14, 1, 0, 0, 0, 14,
		12, 1, 0, 0, 0, 14, 15, 1, 0, 0, 0, 15, 4, 1, 0, 0, 0, 16, 18, 7, 1, 0,
		0, 17, 16, 1, 0, 0, 0, 18, 19, 1, 0, 0, 0, 19, 17, 1, 0, 0, 0, 19, 20,
		1, 0, 0, 0, 20, 6, 1, 0, 0, 0, 21, 23, 5, 34, 0, 0, 22, 24, 3, 5, 2, 0,
		23, 22, 1, 0, 0, 0, 23, 24, 1, 0, 0, 0, 24, 30, 1, 0, 0, 0, 25, 26, 3,
		3, 1, 0, 26, 27, 3, 5, 2, 0, 27, 29, 1, 0, 0, 0, 28, 25, 1, 0, 0, 0, 29,
		32, 1, 0, 0, 0, 30, 28, 1, 0, 0, 0, 30, 31, 1, 0, 0, 0, 31, 34, 1, 0, 0,
		0, 32, 30, 1, 0, 0, 0, 33, 35, 3, 3, 1, 0, 34, 33, 1, 0, 0, 0, 34, 35,
		1, 0, 0, 0, 35, 37, 1, 0, 0, 0, 36, 38, 5, 34, 0, 0, 37, 36, 1, 0, 0, 0,
		37, 38, 1, 0, 0, 0, 38, 8, 1, 0, 0, 0, 7, 0, 14, 19, 23, 30, 34, 37, 0,
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
	MatchLexerSTRING        = 2
	MatchLexerWHITESPACE    = 3
	MatchLexerQUOTED_STRING = 4
)
