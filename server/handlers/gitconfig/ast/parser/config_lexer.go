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
		"", "'='", "'#'",
	}
	staticData.SymbolicNames = []string{
		"", "EQUAL", "HASH", "WHITESPACE", "STRING", "QUOTED_STRING",
	}
	staticData.RuleNames = []string{
		"EQUAL", "HASH", "WHITESPACE", "STRING", "QUOTED_STRING",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 5, 43, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 4, 2, 17, 8, 2, 11, 2, 12, 2, 18,
		1, 3, 4, 3, 22, 8, 3, 11, 3, 12, 3, 23, 1, 4, 1, 4, 3, 4, 28, 8, 4, 1,
		4, 1, 4, 1, 4, 5, 4, 33, 8, 4, 10, 4, 12, 4, 36, 9, 4, 1, 4, 3, 4, 39,
		8, 4, 1, 4, 3, 4, 42, 8, 4, 0, 0, 5, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 1, 0,
		2, 2, 0, 9, 9, 32, 32, 4, 0, 9, 10, 13, 13, 32, 32, 34, 35, 48, 0, 1, 1,
		0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1,
		0, 0, 0, 1, 11, 1, 0, 0, 0, 3, 13, 1, 0, 0, 0, 5, 16, 1, 0, 0, 0, 7, 21,
		1, 0, 0, 0, 9, 25, 1, 0, 0, 0, 11, 12, 5, 61, 0, 0, 12, 2, 1, 0, 0, 0,
		13, 14, 5, 35, 0, 0, 14, 4, 1, 0, 0, 0, 15, 17, 7, 0, 0, 0, 16, 15, 1,
		0, 0, 0, 17, 18, 1, 0, 0, 0, 18, 16, 1, 0, 0, 0, 18, 19, 1, 0, 0, 0, 19,
		6, 1, 0, 0, 0, 20, 22, 8, 1, 0, 0, 21, 20, 1, 0, 0, 0, 22, 23, 1, 0, 0,
		0, 23, 21, 1, 0, 0, 0, 23, 24, 1, 0, 0, 0, 24, 8, 1, 0, 0, 0, 25, 27, 5,
		34, 0, 0, 26, 28, 3, 5, 2, 0, 27, 26, 1, 0, 0, 0, 27, 28, 1, 0, 0, 0, 28,
		34, 1, 0, 0, 0, 29, 30, 3, 7, 3, 0, 30, 31, 3, 5, 2, 0, 31, 33, 1, 0, 0,
		0, 32, 29, 1, 0, 0, 0, 33, 36, 1, 0, 0, 0, 34, 32, 1, 0, 0, 0, 34, 35,
		1, 0, 0, 0, 35, 38, 1, 0, 0, 0, 36, 34, 1, 0, 0, 0, 37, 39, 3, 7, 3, 0,
		38, 37, 1, 0, 0, 0, 38, 39, 1, 0, 0, 0, 39, 41, 1, 0, 0, 0, 40, 42, 5,
		34, 0, 0, 41, 40, 1, 0, 0, 0, 41, 42, 1, 0, 0, 0, 42, 10, 1, 0, 0, 0, 7,
		0, 18, 23, 27, 34, 38, 41, 0,
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
	ConfigLexerWHITESPACE    = 3
	ConfigLexerSTRING        = 4
	ConfigLexerQUOTED_STRING = 5
)
