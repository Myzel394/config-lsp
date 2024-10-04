// Code generated from Config.g4 by ANTLR 4.13.0. DO NOT EDIT.

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

type ConfigLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var ConfigLexerLexerStaticData struct {
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

func configlexerLexerInit() {
	staticData := &ConfigLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'#'",
	}
	staticData.SymbolicNames = []string{
		"", "HASH", "WHITESPACE", "STRING", "NEWLINE", "QUOTED_STRING",
	}
	staticData.RuleNames = []string{
		"HASH", "WHITESPACE", "STRING", "NEWLINE", "QUOTED_STRING",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 5, 46, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 1, 0, 1, 0, 1, 1, 4, 1, 15, 8, 1, 11, 1, 12, 1, 16, 1, 2, 4, 2,
		20, 8, 2, 11, 2, 12, 2, 21, 1, 3, 3, 3, 25, 8, 3, 1, 3, 1, 3, 1, 4, 1,
		4, 3, 4, 31, 8, 4, 1, 4, 1, 4, 1, 4, 5, 4, 36, 8, 4, 10, 4, 12, 4, 39,
		9, 4, 1, 4, 3, 4, 42, 8, 4, 1, 4, 3, 4, 45, 8, 4, 0, 0, 5, 1, 1, 3, 2,
		5, 3, 7, 4, 9, 5, 1, 0, 2, 2, 0, 9, 9, 32, 32, 4, 0, 9, 10, 13, 13, 32,
		32, 34, 35, 52, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0,
		7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 1, 11, 1, 0, 0, 0, 3, 14, 1, 0, 0, 0,
		5, 19, 1, 0, 0, 0, 7, 24, 1, 0, 0, 0, 9, 28, 1, 0, 0, 0, 11, 12, 5, 35,
		0, 0, 12, 2, 1, 0, 0, 0, 13, 15, 7, 0, 0, 0, 14, 13, 1, 0, 0, 0, 15, 16,
		1, 0, 0, 0, 16, 14, 1, 0, 0, 0, 16, 17, 1, 0, 0, 0, 17, 4, 1, 0, 0, 0,
		18, 20, 8, 1, 0, 0, 19, 18, 1, 0, 0, 0, 20, 21, 1, 0, 0, 0, 21, 19, 1,
		0, 0, 0, 21, 22, 1, 0, 0, 0, 22, 6, 1, 0, 0, 0, 23, 25, 5, 13, 0, 0, 24,
		23, 1, 0, 0, 0, 24, 25, 1, 0, 0, 0, 25, 26, 1, 0, 0, 0, 26, 27, 5, 10,
		0, 0, 27, 8, 1, 0, 0, 0, 28, 30, 5, 34, 0, 0, 29, 31, 3, 3, 1, 0, 30, 29,
		1, 0, 0, 0, 30, 31, 1, 0, 0, 0, 31, 37, 1, 0, 0, 0, 32, 33, 3, 5, 2, 0,
		33, 34, 3, 3, 1, 0, 34, 36, 1, 0, 0, 0, 35, 32, 1, 0, 0, 0, 36, 39, 1,
		0, 0, 0, 37, 35, 1, 0, 0, 0, 37, 38, 1, 0, 0, 0, 38, 41, 1, 0, 0, 0, 39,
		37, 1, 0, 0, 0, 40, 42, 3, 5, 2, 0, 41, 40, 1, 0, 0, 0, 41, 42, 1, 0, 0,
		0, 42, 44, 1, 0, 0, 0, 43, 45, 5, 34, 0, 0, 44, 43, 1, 0, 0, 0, 44, 45,
		1, 0, 0, 0, 45, 10, 1, 0, 0, 0, 8, 0, 16, 21, 24, 30, 37, 41, 44, 0,
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

// ConfigLexerInit initializes any static state used to implement ConfigLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewConfigLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func ConfigLexerInit() {
	staticData := &ConfigLexerLexerStaticData
	staticData.once.Do(configlexerLexerInit)
}

// NewConfigLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewConfigLexer(input antlr.CharStream) *ConfigLexer {
	ConfigLexerInit()
	l := new(ConfigLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &ConfigLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Config.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ConfigLexer tokens.
const (
	ConfigLexerHASH          = 1
	ConfigLexerWHITESPACE    = 2
	ConfigLexerSTRING        = 3
	ConfigLexerNEWLINE       = 4
	ConfigLexerQUOTED_STRING = 5
)
