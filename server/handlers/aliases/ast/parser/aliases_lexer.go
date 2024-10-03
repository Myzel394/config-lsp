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
		"", "", "", "", "'@'", "", "'|'", "':'", "','", "'#'", "'/'",
	}
	staticData.SymbolicNames = []string{
		"", "DIGITS", "ERROR", "SEPARATOR", "AT", "INCLUDE", "VERTLINE", "COLON",
		"COMMA", "NUMBER_SIGN", "SLASH", "STRING",
	}
	staticData.RuleNames = []string{
		"DIGITS", "ERROR", "SEPARATOR", "AT", "INCLUDE", "VERTLINE", "COLON",
		"COMMA", "NUMBER_SIGN", "SLASH", "STRING",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 11, 64, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 1, 0, 4, 0, 25, 8, 0, 11, 0, 12, 0, 26, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 2, 4, 2, 36, 8, 2, 11, 2, 12, 2, 37, 1, 3, 1, 3, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7,
		1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 4, 10, 61, 8, 10, 11, 10, 12, 10,
		62, 0, 0, 11, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9,
		19, 10, 21, 11, 1, 0, 3, 1, 0, 48, 57, 2, 0, 9, 9, 32, 32, 9, 0, 9, 10,
		13, 13, 32, 32, 35, 35, 44, 44, 47, 47, 58, 58, 64, 64, 124, 124, 66, 0,
		1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0,
		9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0,
		0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 1, 24, 1, 0, 0,
		0, 3, 28, 1, 0, 0, 0, 5, 35, 1, 0, 0, 0, 7, 39, 1, 0, 0, 0, 9, 41, 1, 0,
		0, 0, 11, 49, 1, 0, 0, 0, 13, 51, 1, 0, 0, 0, 15, 53, 1, 0, 0, 0, 17, 55,
		1, 0, 0, 0, 19, 57, 1, 0, 0, 0, 21, 60, 1, 0, 0, 0, 23, 25, 7, 0, 0, 0,
		24, 23, 1, 0, 0, 0, 25, 26, 1, 0, 0, 0, 26, 24, 1, 0, 0, 0, 26, 27, 1,
		0, 0, 0, 27, 2, 1, 0, 0, 0, 28, 29, 5, 101, 0, 0, 29, 30, 5, 114, 0, 0,
		30, 31, 5, 114, 0, 0, 31, 32, 5, 111, 0, 0, 32, 33, 5, 114, 0, 0, 33, 4,
		1, 0, 0, 0, 34, 36, 7, 1, 0, 0, 35, 34, 1, 0, 0, 0, 36, 37, 1, 0, 0, 0,
		37, 35, 1, 0, 0, 0, 37, 38, 1, 0, 0, 0, 38, 6, 1, 0, 0, 0, 39, 40, 5, 64,
		0, 0, 40, 8, 1, 0, 0, 0, 41, 42, 5, 105, 0, 0, 42, 43, 5, 110, 0, 0, 43,
		44, 5, 99, 0, 0, 44, 45, 5, 108, 0, 0, 45, 46, 5, 117, 0, 0, 46, 47, 5,
		100, 0, 0, 47, 48, 5, 101, 0, 0, 48, 10, 1, 0, 0, 0, 49, 50, 5, 124, 0,
		0, 50, 12, 1, 0, 0, 0, 51, 52, 5, 58, 0, 0, 52, 14, 1, 0, 0, 0, 53, 54,
		5, 44, 0, 0, 54, 16, 1, 0, 0, 0, 55, 56, 5, 35, 0, 0, 56, 18, 1, 0, 0,
		0, 57, 58, 5, 47, 0, 0, 58, 20, 1, 0, 0, 0, 59, 61, 8, 2, 0, 0, 60, 59,
		1, 0, 0, 0, 61, 62, 1, 0, 0, 0, 62, 60, 1, 0, 0, 0, 62, 63, 1, 0, 0, 0,
		63, 22, 1, 0, 0, 0, 4, 0, 26, 37, 62, 0,
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
	AliasesLexerDIGITS      = 1
	AliasesLexerERROR       = 2
	AliasesLexerSEPARATOR   = 3
	AliasesLexerAT          = 4
	AliasesLexerINCLUDE     = 5
	AliasesLexerVERTLINE    = 6
	AliasesLexerCOLON       = 7
	AliasesLexerCOMMA       = 8
	AliasesLexerNUMBER_SIGN = 9
	AliasesLexerSLASH       = 10
	AliasesLexerSTRING      = 11
)
