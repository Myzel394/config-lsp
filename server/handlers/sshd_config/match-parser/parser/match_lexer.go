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
		"", "COMMA", "STRING", "WHITESPACE",
	}
	staticData.RuleNames = []string{
		"COMMA", "STRING", "WHITESPACE",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 3, 19, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 1, 0, 1, 0, 1,
		1, 4, 1, 11, 8, 1, 11, 1, 12, 1, 12, 1, 2, 4, 2, 16, 8, 2, 11, 2, 12, 2,
		17, 0, 0, 3, 1, 1, 3, 2, 5, 3, 1, 0, 2, 5, 0, 9, 10, 13, 13, 32, 32, 35,
		35, 44, 44, 2, 0, 9, 9, 32, 32, 20, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0,
		0, 5, 1, 0, 0, 0, 1, 7, 1, 0, 0, 0, 3, 10, 1, 0, 0, 0, 5, 15, 1, 0, 0,
		0, 7, 8, 5, 44, 0, 0, 8, 2, 1, 0, 0, 0, 9, 11, 8, 0, 0, 0, 10, 9, 1, 0,
		0, 0, 11, 12, 1, 0, 0, 0, 12, 10, 1, 0, 0, 0, 12, 13, 1, 0, 0, 0, 13, 4,
		1, 0, 0, 0, 14, 16, 7, 1, 0, 0, 15, 14, 1, 0, 0, 0, 16, 17, 1, 0, 0, 0,
		17, 15, 1, 0, 0, 0, 17, 18, 1, 0, 0, 0, 18, 6, 1, 0, 0, 0, 3, 0, 12, 17,
		0,
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
	MatchLexerCOMMA      = 1
	MatchLexerSTRING     = 2
	MatchLexerWHITESPACE = 3
)
