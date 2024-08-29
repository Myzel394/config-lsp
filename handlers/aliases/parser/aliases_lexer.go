// Code generated from Aliases.g4 by ANTLR 4.13.0. DO NOT EDIT.

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

type AliasesLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var AliasesLexerLexerStaticData struct {
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

func aliaseslexerLexerInit() {
	staticData := &AliasesLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "", "'@'", "", "'|'", "':'", "','", "'#'", "'/'",
	}
	staticData.SymbolicNames = []string{
		"", "SEPARATOR", "AT", "INCLUDE", "VERTLINE", "COLON", "COMMA", "NUMBER_SIGN",
		"SLASH", "STRING",
	}
	staticData.RuleNames = []string{
		"SEPARATOR", "AT", "INCLUDE", "VERTLINE", "COLON", "COMMA", "NUMBER_SIGN",
		"SLASH", "STRING",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 9, 49, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 1, 0, 4, 0, 21,
		8, 0, 11, 0, 12, 0, 22, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7,
		1, 8, 4, 8, 46, 8, 8, 11, 8, 12, 8, 47, 0, 0, 9, 1, 1, 3, 2, 5, 3, 7, 4,
		9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 1, 0, 2, 2, 0, 9, 9, 32, 32, 9, 0, 9,
		10, 13, 13, 32, 32, 35, 35, 44, 44, 47, 47, 58, 58, 64, 64, 124, 124, 50,
		0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0,
		0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0,
		0, 0, 17, 1, 0, 0, 0, 1, 20, 1, 0, 0, 0, 3, 24, 1, 0, 0, 0, 5, 26, 1, 0,
		0, 0, 7, 34, 1, 0, 0, 0, 9, 36, 1, 0, 0, 0, 11, 38, 1, 0, 0, 0, 13, 40,
		1, 0, 0, 0, 15, 42, 1, 0, 0, 0, 17, 45, 1, 0, 0, 0, 19, 21, 7, 0, 0, 0,
		20, 19, 1, 0, 0, 0, 21, 22, 1, 0, 0, 0, 22, 20, 1, 0, 0, 0, 22, 23, 1,
		0, 0, 0, 23, 2, 1, 0, 0, 0, 24, 25, 5, 64, 0, 0, 25, 4, 1, 0, 0, 0, 26,
		27, 5, 105, 0, 0, 27, 28, 5, 110, 0, 0, 28, 29, 5, 99, 0, 0, 29, 30, 5,
		108, 0, 0, 30, 31, 5, 117, 0, 0, 31, 32, 5, 100, 0, 0, 32, 33, 5, 101,
		0, 0, 33, 6, 1, 0, 0, 0, 34, 35, 5, 124, 0, 0, 35, 8, 1, 0, 0, 0, 36, 37,
		5, 58, 0, 0, 37, 10, 1, 0, 0, 0, 38, 39, 5, 44, 0, 0, 39, 12, 1, 0, 0,
		0, 40, 41, 5, 35, 0, 0, 41, 14, 1, 0, 0, 0, 42, 43, 5, 47, 0, 0, 43, 16,
		1, 0, 0, 0, 44, 46, 8, 1, 0, 0, 45, 44, 1, 0, 0, 0, 46, 47, 1, 0, 0, 0,
		47, 45, 1, 0, 0, 0, 47, 48, 1, 0, 0, 0, 48, 18, 1, 0, 0, 0, 3, 0, 22, 47,
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

// AliasesLexerInit initializes any static state used to implement AliasesLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewAliasesLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func AliasesLexerInit() {
	staticData := &AliasesLexerLexerStaticData
	staticData.once.Do(aliaseslexerLexerInit)
}

// NewAliasesLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewAliasesLexer(input antlr.CharStream) *AliasesLexer {
	AliasesLexerInit()
	l := new(AliasesLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &AliasesLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Aliases.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// AliasesLexer tokens.
const (
	AliasesLexerSEPARATOR   = 1
	AliasesLexerAT          = 2
	AliasesLexerINCLUDE     = 3
	AliasesLexerVERTLINE    = 4
	AliasesLexerCOLON       = 5
	AliasesLexerCOMMA       = 6
	AliasesLexerNUMBER_SIGN = 7
	AliasesLexerSLASH       = 8
	AliasesLexerSTRING      = 9
)
