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
		"", "'='", "'#'", "';'",
	}
	staticData.SymbolicNames = []string{
		"", "EQUAL", "HASH", "SEMICOLON", "WHITESPACE", "STRING", "QUOTED_STRING",
	}
	staticData.RuleNames = []string{
		"EQUAL", "HASH", "SEMICOLON", "WHITESPACE", "STRING", "QUOTED_STRING",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 6, 47, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 4, 3, 21,
		8, 3, 11, 3, 12, 3, 22, 1, 4, 4, 4, 26, 8, 4, 11, 4, 12, 4, 27, 1, 5, 1,
		5, 3, 5, 32, 8, 5, 1, 5, 1, 5, 1, 5, 5, 5, 37, 8, 5, 10, 5, 12, 5, 40,
		9, 5, 1, 5, 3, 5, 43, 8, 5, 1, 5, 3, 5, 46, 8, 5, 0, 0, 6, 1, 1, 3, 2,
		5, 3, 7, 4, 9, 5, 11, 6, 1, 0, 2, 2, 0, 9, 9, 32, 32, 4, 0, 9, 10, 13,
		13, 32, 32, 34, 34, 52, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0,
		0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 1, 13, 1,
		0, 0, 0, 3, 15, 1, 0, 0, 0, 5, 17, 1, 0, 0, 0, 7, 20, 1, 0, 0, 0, 9, 25,
		1, 0, 0, 0, 11, 29, 1, 0, 0, 0, 13, 14, 5, 61, 0, 0, 14, 2, 1, 0, 0, 0,
		15, 16, 5, 35, 0, 0, 16, 4, 1, 0, 0, 0, 17, 18, 5, 59, 0, 0, 18, 6, 1,
		0, 0, 0, 19, 21, 7, 0, 0, 0, 20, 19, 1, 0, 0, 0, 21, 22, 1, 0, 0, 0, 22,
		20, 1, 0, 0, 0, 22, 23, 1, 0, 0, 0, 23, 8, 1, 0, 0, 0, 24, 26, 8, 1, 0,
		0, 25, 24, 1, 0, 0, 0, 26, 27, 1, 0, 0, 0, 27, 25, 1, 0, 0, 0, 27, 28,
		1, 0, 0, 0, 28, 10, 1, 0, 0, 0, 29, 31, 5, 34, 0, 0, 30, 32, 3, 7, 3, 0,
		31, 30, 1, 0, 0, 0, 31, 32, 1, 0, 0, 0, 32, 38, 1, 0, 0, 0, 33, 34, 3,
		9, 4, 0, 34, 35, 3, 7, 3, 0, 35, 37, 1, 0, 0, 0, 36, 33, 1, 0, 0, 0, 37,
		40, 1, 0, 0, 0, 38, 36, 1, 0, 0, 0, 38, 39, 1, 0, 0, 0, 39, 42, 1, 0, 0,
		0, 40, 38, 1, 0, 0, 0, 41, 43, 3, 9, 4, 0, 42, 41, 1, 0, 0, 0, 42, 43,
		1, 0, 0, 0, 43, 45, 1, 0, 0, 0, 44, 46, 5, 34, 0, 0, 45, 44, 1, 0, 0, 0,
		45, 46, 1, 0, 0, 0, 46, 12, 1, 0, 0, 0, 7, 0, 22, 27, 31, 38, 42, 45, 0,
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
	ConfigLexerEQUAL         = 1
	ConfigLexerHASH          = 2
	ConfigLexerSEMICOLON     = 3
	ConfigLexerWHITESPACE    = 4
	ConfigLexerSTRING        = 5
	ConfigLexerQUOTED_STRING = 6
)
