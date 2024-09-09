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
		"", "HASH", "MATCH", "WHITESPACE", "STRING", "NEWLINE",
	}
	staticData.RuleNames = []string{
		"HASH", "MATCH", "WHITESPACE", "STRING", "NEWLINE",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 5, 34, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 4, 2, 21,
		8, 2, 11, 2, 12, 2, 22, 1, 3, 4, 3, 26, 8, 3, 11, 3, 12, 3, 27, 1, 4, 3,
		4, 31, 8, 4, 1, 4, 1, 4, 0, 0, 5, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 1, 0, 7,
		2, 0, 77, 77, 109, 109, 2, 0, 65, 65, 97, 97, 2, 0, 84, 84, 116, 116, 2,
		0, 67, 67, 99, 99, 2, 0, 72, 72, 104, 104, 2, 0, 9, 9, 32, 32, 4, 0, 9,
		10, 13, 13, 32, 32, 35, 35, 36, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0,
		5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 1, 11, 1, 0, 0, 0, 3,
		13, 1, 0, 0, 0, 5, 20, 1, 0, 0, 0, 7, 25, 1, 0, 0, 0, 9, 30, 1, 0, 0, 0,
		11, 12, 5, 35, 0, 0, 12, 2, 1, 0, 0, 0, 13, 14, 7, 0, 0, 0, 14, 15, 7,
		1, 0, 0, 15, 16, 7, 2, 0, 0, 16, 17, 7, 3, 0, 0, 17, 18, 7, 4, 0, 0, 18,
		4, 1, 0, 0, 0, 19, 21, 7, 5, 0, 0, 20, 19, 1, 0, 0, 0, 21, 22, 1, 0, 0,
		0, 22, 20, 1, 0, 0, 0, 22, 23, 1, 0, 0, 0, 23, 6, 1, 0, 0, 0, 24, 26, 8,
		6, 0, 0, 25, 24, 1, 0, 0, 0, 26, 27, 1, 0, 0, 0, 27, 25, 1, 0, 0, 0, 27,
		28, 1, 0, 0, 0, 28, 8, 1, 0, 0, 0, 29, 31, 5, 13, 0, 0, 30, 29, 1, 0, 0,
		0, 30, 31, 1, 0, 0, 0, 31, 32, 1, 0, 0, 0, 32, 33, 5, 10, 0, 0, 33, 10,
		1, 0, 0, 0, 4, 0, 22, 27, 30, 0,
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
	ConfigLexerHASH       = 1
	ConfigLexerMATCH      = 2
	ConfigLexerWHITESPACE = 3
	ConfigLexerSTRING     = 4
	ConfigLexerNEWLINE    = 5
)
