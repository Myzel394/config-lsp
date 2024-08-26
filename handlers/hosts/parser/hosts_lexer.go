// Code generated from Hosts.g4 by ANTLR 4.13.0. DO NOT EDIT.

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

type HostsLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var HostsLexerLexerStaticData struct {
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

func hostslexerLexerInit() {
	staticData := &HostsLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "", "'/'", "'.'", "':'", "'#'",
	}
	staticData.SymbolicNames = []string{
		"", "COMMENTLINE", "SLASH", "DOT", "COLON", "HASHTAG", "SEPARATOR",
		"NEWLINE", "DIGITS", "OCTETS", "DOMAIN",
	}
	staticData.RuleNames = []string{
		"COMMENTLINE", "SLASH", "DOT", "COLON", "HASHTAG", "SEPARATOR", "NEWLINE",
		"DIGITS", "DIGIT", "OCTETS", "OCTET", "DOMAIN", "STRING",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 10, 83, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 1, 0, 1, 0, 4, 0, 30, 8, 0, 11,
		0, 12, 0, 31, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 4,
		5, 43, 8, 5, 11, 5, 12, 5, 44, 1, 6, 3, 6, 48, 8, 6, 1, 6, 1, 6, 1, 7,
		4, 7, 53, 8, 7, 11, 7, 12, 7, 54, 1, 8, 1, 8, 1, 9, 4, 9, 60, 8, 9, 11,
		9, 12, 9, 61, 1, 10, 1, 10, 1, 11, 4, 11, 67, 8, 11, 11, 11, 12, 11, 68,
		1, 11, 1, 11, 4, 11, 73, 8, 11, 11, 11, 12, 11, 74, 5, 11, 77, 8, 11, 10,
		11, 12, 11, 80, 9, 11, 1, 12, 1, 12, 0, 0, 13, 1, 1, 3, 2, 5, 3, 7, 4,
		9, 5, 11, 6, 13, 7, 15, 8, 17, 0, 19, 9, 21, 0, 23, 10, 25, 0, 1, 0, 8,
		2, 0, 10, 10, 13, 13, 2, 0, 9, 9, 32, 32, 1, 0, 13, 13, 1, 0, 10, 10, 1,
		0, 48, 57, 3, 0, 48, 57, 65, 70, 97, 102, 2, 0, 65, 90, 97, 122, 5, 0,
		9, 10, 13, 13, 32, 32, 35, 35, 46, 46, 87, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0,
		0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0,
		0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 23, 1,
		0, 0, 0, 1, 27, 1, 0, 0, 0, 3, 33, 1, 0, 0, 0, 5, 35, 1, 0, 0, 0, 7, 37,
		1, 0, 0, 0, 9, 39, 1, 0, 0, 0, 11, 42, 1, 0, 0, 0, 13, 47, 1, 0, 0, 0,
		15, 52, 1, 0, 0, 0, 17, 56, 1, 0, 0, 0, 19, 59, 1, 0, 0, 0, 21, 63, 1,
		0, 0, 0, 23, 66, 1, 0, 0, 0, 25, 81, 1, 0, 0, 0, 27, 29, 3, 9, 4, 0, 28,
		30, 8, 0, 0, 0, 29, 28, 1, 0, 0, 0, 30, 31, 1, 0, 0, 0, 31, 29, 1, 0, 0,
		0, 31, 32, 1, 0, 0, 0, 32, 2, 1, 0, 0, 0, 33, 34, 5, 47, 0, 0, 34, 4, 1,
		0, 0, 0, 35, 36, 5, 46, 0, 0, 36, 6, 1, 0, 0, 0, 37, 38, 5, 58, 0, 0, 38,
		8, 1, 0, 0, 0, 39, 40, 5, 35, 0, 0, 40, 10, 1, 0, 0, 0, 41, 43, 7, 1, 0,
		0, 42, 41, 1, 0, 0, 0, 43, 44, 1, 0, 0, 0, 44, 42, 1, 0, 0, 0, 44, 45,
		1, 0, 0, 0, 45, 12, 1, 0, 0, 0, 46, 48, 7, 2, 0, 0, 47, 46, 1, 0, 0, 0,
		47, 48, 1, 0, 0, 0, 48, 49, 1, 0, 0, 0, 49, 50, 7, 3, 0, 0, 50, 14, 1,
		0, 0, 0, 51, 53, 3, 17, 8, 0, 52, 51, 1, 0, 0, 0, 53, 54, 1, 0, 0, 0, 54,
		52, 1, 0, 0, 0, 54, 55, 1, 0, 0, 0, 55, 16, 1, 0, 0, 0, 56, 57, 7, 4, 0,
		0, 57, 18, 1, 0, 0, 0, 58, 60, 3, 21, 10, 0, 59, 58, 1, 0, 0, 0, 60, 61,
		1, 0, 0, 0, 61, 59, 1, 0, 0, 0, 61, 62, 1, 0, 0, 0, 62, 20, 1, 0, 0, 0,
		63, 64, 7, 5, 0, 0, 64, 22, 1, 0, 0, 0, 65, 67, 3, 25, 12, 0, 66, 65, 1,
		0, 0, 0, 67, 68, 1, 0, 0, 0, 68, 66, 1, 0, 0, 0, 68, 69, 1, 0, 0, 0, 69,
		78, 1, 0, 0, 0, 70, 72, 3, 5, 2, 0, 71, 73, 7, 6, 0, 0, 72, 71, 1, 0, 0,
		0, 73, 74, 1, 0, 0, 0, 74, 72, 1, 0, 0, 0, 74, 75, 1, 0, 0, 0, 75, 77,
		1, 0, 0, 0, 76, 70, 1, 0, 0, 0, 77, 80, 1, 0, 0, 0, 78, 76, 1, 0, 0, 0,
		78, 79, 1, 0, 0, 0, 79, 24, 1, 0, 0, 0, 80, 78, 1, 0, 0, 0, 81, 82, 8,
		7, 0, 0, 82, 26, 1, 0, 0, 0, 9, 0, 31, 44, 47, 54, 61, 68, 74, 78, 0,
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

// HostsLexerInit initializes any static state used to implement HostsLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewHostsLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func HostsLexerInit() {
	staticData := &HostsLexerLexerStaticData
	staticData.once.Do(hostslexerLexerInit)
}

// NewHostsLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewHostsLexer(input antlr.CharStream) *HostsLexer {
	HostsLexerInit()
	l := new(HostsLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &HostsLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Hosts.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// HostsLexer tokens.
const (
	HostsLexerCOMMENTLINE = 1
	HostsLexerSLASH       = 2
	HostsLexerDOT         = 3
	HostsLexerCOLON       = 4
	HostsLexerHASHTAG     = 5
	HostsLexerSEPARATOR   = 6
	HostsLexerNEWLINE     = 7
	HostsLexerDIGITS      = 8
	HostsLexerOCTETS      = 9
	HostsLexerDOMAIN      = 10
)
