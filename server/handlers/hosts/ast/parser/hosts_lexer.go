// Code generated from Hosts.g4 by ANTLR 4.13.2. DO NOT EDIT.

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
		"NEWLINE", "STRING",
	}
	staticData.RuleNames = []string{
		"COMMENTLINE", "SLASH", "DOT", "COLON", "HASHTAG", "SEPARATOR", "NEWLINE",
		"STRING",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 8, 46, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 1, 0, 1, 0, 4, 0, 20, 8, 0,
		11, 0, 12, 0, 21, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5,
		4, 5, 33, 8, 5, 11, 5, 12, 5, 34, 1, 6, 3, 6, 38, 8, 6, 1, 6, 1, 6, 1,
		7, 4, 7, 43, 8, 7, 11, 7, 12, 7, 44, 0, 0, 8, 1, 1, 3, 2, 5, 3, 7, 4, 9,
		5, 11, 6, 13, 7, 15, 8, 1, 0, 5, 2, 0, 10, 10, 13, 13, 2, 0, 9, 9, 32,
		32, 1, 0, 13, 13, 1, 0, 10, 10, 5, 0, 45, 45, 48, 57, 65, 90, 95, 95, 97,
		122, 49, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1,
		0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15,
		1, 0, 0, 0, 1, 17, 1, 0, 0, 0, 3, 23, 1, 0, 0, 0, 5, 25, 1, 0, 0, 0, 7,
		27, 1, 0, 0, 0, 9, 29, 1, 0, 0, 0, 11, 32, 1, 0, 0, 0, 13, 37, 1, 0, 0,
		0, 15, 42, 1, 0, 0, 0, 17, 19, 3, 9, 4, 0, 18, 20, 8, 0, 0, 0, 19, 18,
		1, 0, 0, 0, 20, 21, 1, 0, 0, 0, 21, 19, 1, 0, 0, 0, 21, 22, 1, 0, 0, 0,
		22, 2, 1, 0, 0, 0, 23, 24, 5, 47, 0, 0, 24, 4, 1, 0, 0, 0, 25, 26, 5, 46,
		0, 0, 26, 6, 1, 0, 0, 0, 27, 28, 5, 58, 0, 0, 28, 8, 1, 0, 0, 0, 29, 30,
		5, 35, 0, 0, 30, 10, 1, 0, 0, 0, 31, 33, 7, 1, 0, 0, 32, 31, 1, 0, 0, 0,
		33, 34, 1, 0, 0, 0, 34, 32, 1, 0, 0, 0, 34, 35, 1, 0, 0, 0, 35, 12, 1,
		0, 0, 0, 36, 38, 7, 2, 0, 0, 37, 36, 1, 0, 0, 0, 37, 38, 1, 0, 0, 0, 38,
		39, 1, 0, 0, 0, 39, 40, 7, 3, 0, 0, 40, 14, 1, 0, 0, 0, 41, 43, 7, 4, 0,
		0, 42, 41, 1, 0, 0, 0, 43, 44, 1, 0, 0, 0, 44, 42, 1, 0, 0, 0, 44, 45,
		1, 0, 0, 0, 45, 16, 1, 0, 0, 0, 5, 0, 21, 34, 37, 44, 0,
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
	HostsLexerSTRING      = 8
)
