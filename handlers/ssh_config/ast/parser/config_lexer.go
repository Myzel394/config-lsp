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
		"", "HASH", "WHITESPACE", "STRING", "NEWLINE",
	}
	staticData.RuleNames = []string{
		"HASH", "WHITESPACE", "STRING", "NEWLINE",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 4, 26, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 1,
		0, 1, 0, 1, 1, 4, 1, 13, 8, 1, 11, 1, 12, 1, 14, 1, 2, 4, 2, 18, 8, 2,
		11, 2, 12, 2, 19, 1, 3, 3, 3, 23, 8, 3, 1, 3, 1, 3, 0, 0, 4, 1, 1, 3, 2,
		5, 3, 7, 4, 1, 0, 2, 2, 0, 9, 9, 32, 32, 4, 0, 9, 10, 13, 13, 32, 32, 35,
		35, 28, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1,
		0, 0, 0, 1, 9, 1, 0, 0, 0, 3, 12, 1, 0, 0, 0, 5, 17, 1, 0, 0, 0, 7, 22,
		1, 0, 0, 0, 9, 10, 5, 35, 0, 0, 10, 2, 1, 0, 0, 0, 11, 13, 7, 0, 0, 0,
		12, 11, 1, 0, 0, 0, 13, 14, 1, 0, 0, 0, 14, 12, 1, 0, 0, 0, 14, 15, 1,
		0, 0, 0, 15, 4, 1, 0, 0, 0, 16, 18, 8, 1, 0, 0, 17, 16, 1, 0, 0, 0, 18,
		19, 1, 0, 0, 0, 19, 17, 1, 0, 0, 0, 19, 20, 1, 0, 0, 0, 20, 6, 1, 0, 0,
		0, 21, 23, 5, 13, 0, 0, 22, 21, 1, 0, 0, 0, 22, 23, 1, 0, 0, 0, 23, 24,
		1, 0, 0, 0, 24, 25, 5, 10, 0, 0, 25, 8, 1, 0, 0, 0, 4, 0, 14, 19, 22, 0,
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
	ConfigLexerWHITESPACE = 2
	ConfigLexerSTRING     = 3
	ConfigLexerNEWLINE    = 4
)
