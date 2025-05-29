// Code generated from Match.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Match

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type MatchParser struct {
	*antlr.BaseParser
}

var MatchParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func matchParserInit() {
	staticData := &MatchParserStaticData
	staticData.LiteralNames = []string{
		"", "','", "", "", "", "", "", "", "", "", "", "", "", "", "", "'\"'",
	}
	staticData.SymbolicNames = []string{
		"", "COMMA", "ALL", "CANONICAL", "FINAL", "EXEC", "LOCALNETWORK", "HOST",
		"ORIGINALHOST", "TAGGED", "USER", "LOCALUSER", "STRING", "WHITESPACE",
		"QUOTED_STRING", "QUOTE",
	}
	staticData.RuleNames = []string{
		"root", "matchEntry", "entrySingle", "entryWithValue", "separator",
		"values", "value", "criteriaSingle", "criteriaWithValue", "string",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 15, 80, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 1, 0, 3,
		0, 22, 8, 0, 1, 0, 1, 0, 3, 0, 26, 8, 0, 5, 0, 28, 8, 0, 10, 0, 12, 0,
		31, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 3, 1, 37, 8, 1, 1, 2, 1, 2, 1, 3, 1,
		3, 3, 3, 43, 8, 3, 1, 3, 3, 3, 46, 8, 3, 1, 4, 1, 4, 1, 5, 3, 5, 51, 8,
		5, 1, 5, 1, 5, 3, 5, 55, 8, 5, 5, 5, 57, 8, 5, 10, 5, 12, 5, 60, 9, 5,
		1, 6, 1, 6, 1, 7, 3, 7, 65, 8, 7, 1, 7, 1, 7, 3, 7, 69, 8, 7, 1, 8, 3,
		8, 72, 8, 8, 1, 8, 1, 8, 3, 8, 76, 8, 8, 1, 9, 1, 9, 1, 9, 0, 0, 10, 0,
		2, 4, 6, 8, 10, 12, 14, 16, 18, 0, 3, 1, 0, 2, 4, 1, 0, 5, 11, 2, 0, 12,
		12, 14, 14, 82, 0, 21, 1, 0, 0, 0, 2, 36, 1, 0, 0, 0, 4, 38, 1, 0, 0, 0,
		6, 40, 1, 0, 0, 0, 8, 47, 1, 0, 0, 0, 10, 50, 1, 0, 0, 0, 12, 61, 1, 0,
		0, 0, 14, 64, 1, 0, 0, 0, 16, 71, 1, 0, 0, 0, 18, 77, 1, 0, 0, 0, 20, 22,
		3, 2, 1, 0, 21, 20, 1, 0, 0, 0, 21, 22, 1, 0, 0, 0, 22, 29, 1, 0, 0, 0,
		23, 25, 5, 13, 0, 0, 24, 26, 3, 2, 1, 0, 25, 24, 1, 0, 0, 0, 25, 26, 1,
		0, 0, 0, 26, 28, 1, 0, 0, 0, 27, 23, 1, 0, 0, 0, 28, 31, 1, 0, 0, 0, 29,
		27, 1, 0, 0, 0, 29, 30, 1, 0, 0, 0, 30, 32, 1, 0, 0, 0, 31, 29, 1, 0, 0,
		0, 32, 33, 5, 0, 0, 1, 33, 1, 1, 0, 0, 0, 34, 37, 3, 4, 2, 0, 35, 37, 3,
		6, 3, 0, 36, 34, 1, 0, 0, 0, 36, 35, 1, 0, 0, 0, 37, 3, 1, 0, 0, 0, 38,
		39, 3, 14, 7, 0, 39, 5, 1, 0, 0, 0, 40, 42, 3, 16, 8, 0, 41, 43, 3, 8,
		4, 0, 42, 41, 1, 0, 0, 0, 42, 43, 1, 0, 0, 0, 43, 45, 1, 0, 0, 0, 44, 46,
		3, 10, 5, 0, 45, 44, 1, 0, 0, 0, 45, 46, 1, 0, 0, 0, 46, 7, 1, 0, 0, 0,
		47, 48, 5, 13, 0, 0, 48, 9, 1, 0, 0, 0, 49, 51, 3, 12, 6, 0, 50, 49, 1,
		0, 0, 0, 50, 51, 1, 0, 0, 0, 51, 58, 1, 0, 0, 0, 52, 54, 5, 1, 0, 0, 53,
		55, 3, 12, 6, 0, 54, 53, 1, 0, 0, 0, 54, 55, 1, 0, 0, 0, 55, 57, 1, 0,
		0, 0, 56, 52, 1, 0, 0, 0, 57, 60, 1, 0, 0, 0, 58, 56, 1, 0, 0, 0, 58, 59,
		1, 0, 0, 0, 59, 11, 1, 0, 0, 0, 60, 58, 1, 0, 0, 0, 61, 62, 3, 18, 9, 0,
		62, 13, 1, 0, 0, 0, 63, 65, 5, 15, 0, 0, 64, 63, 1, 0, 0, 0, 64, 65, 1,
		0, 0, 0, 65, 66, 1, 0, 0, 0, 66, 68, 7, 0, 0, 0, 67, 69, 5, 15, 0, 0, 68,
		67, 1, 0, 0, 0, 68, 69, 1, 0, 0, 0, 69, 15, 1, 0, 0, 0, 70, 72, 5, 15,
		0, 0, 71, 70, 1, 0, 0, 0, 71, 72, 1, 0, 0, 0, 72, 73, 1, 0, 0, 0, 73, 75,
		7, 1, 0, 0, 74, 76, 5, 15, 0, 0, 75, 74, 1, 0, 0, 0, 75, 76, 1, 0, 0, 0,
		76, 17, 1, 0, 0, 0, 77, 78, 7, 2, 0, 0, 78, 19, 1, 0, 0, 0, 13, 21, 25,
		29, 36, 42, 45, 50, 54, 58, 64, 68, 71, 75,
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

// MatchParserInit initializes any static state used to implement MatchParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewMatchParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func MatchParserInit() {
	staticData := &MatchParserStaticData
	staticData.once.Do(matchParserInit)
}

// NewMatchParser produces a new parser instance for the optional input antlr.TokenStream.
func NewMatchParser(input antlr.TokenStream) *MatchParser {
	MatchParserInit()
	this := new(MatchParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &MatchParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Match.g4"

	return this
}

// MatchParser tokens.
const (
	MatchParserEOF           = antlr.TokenEOF
	MatchParserCOMMA         = 1
	MatchParserALL           = 2
	MatchParserCANONICAL     = 3
	MatchParserFINAL         = 4
	MatchParserEXEC          = 5
	MatchParserLOCALNETWORK  = 6
	MatchParserHOST          = 7
	MatchParserORIGINALHOST  = 8
	MatchParserTAGGED        = 9
	MatchParserUSER          = 10
	MatchParserLOCALUSER     = 11
	MatchParserSTRING        = 12
	MatchParserWHITESPACE    = 13
	MatchParserQUOTED_STRING = 14
	MatchParserQUOTE         = 15
)

// MatchParser rules.
const (
	MatchParserRULE_root              = 0
	MatchParserRULE_matchEntry        = 1
	MatchParserRULE_entrySingle       = 2
	MatchParserRULE_entryWithValue    = 3
	MatchParserRULE_separator         = 4
	MatchParserRULE_values            = 5
	MatchParserRULE_value             = 6
	MatchParserRULE_criteriaSingle    = 7
	MatchParserRULE_criteriaWithValue = 8
	MatchParserRULE_string            = 9
)

// IRootContext is an interface to support dynamic dispatch.
type IRootContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllMatchEntry() []IMatchEntryContext
	MatchEntry(i int) IMatchEntryContext
	AllWHITESPACE() []antlr.TerminalNode
	WHITESPACE(i int) antlr.TerminalNode

	// IsRootContext differentiates from other interfaces.
	IsRootContext()
}

type RootContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRootContext() *RootContext {
	var p = new(RootContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MatchParserRULE_root
	return p
}

func InitEmptyRootContext(p *RootContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MatchParserRULE_root
}

func (*RootContext) IsRootContext() {}

func NewRootContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RootContext {
	var p = new(RootContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MatchParserRULE_root

	return p
}

func (s *RootContext) GetParser() antlr.Parser { return s.parser }

func (s *RootContext) EOF() antlr.TerminalNode {
	return s.GetToken(MatchParserEOF, 0)
}

func (s *RootContext) AllMatchEntry() []IMatchEntryContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMatchEntryContext); ok {
			len++
		}
	}

	tst := make([]IMatchEntryContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMatchEntryContext); ok {
			tst[i] = t.(IMatchEntryContext)
			i++
		}
	}

	return tst
}

func (s *RootContext) MatchEntry(i int) IMatchEntryContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMatchEntryContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMatchEntryContext)
}

func (s *RootContext) AllWHITESPACE() []antlr.TerminalNode {
	return s.GetTokens(MatchParserWHITESPACE)
}

func (s *RootContext) WHITESPACE(i int) antlr.TerminalNode {
	return s.GetToken(MatchParserWHITESPACE, i)
}

func (s *RootContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RootContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RootContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MatchListener); ok {
		listenerT.EnterRoot(s)
	}
}

func (s *RootContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MatchListener); ok {
		listenerT.ExitRoot(s)
	}
}

func (p *MatchParser) Root() (localctx IRootContext) {
	localctx = NewRootContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MatchParserRULE_root)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(21)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&36860) != 0 {
		{
			p.SetState(20)
			p.MatchEntry()
		}

	}
	p.SetState(29)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MatchParserWHITESPACE {
		{
			p.SetState(23)
			p.Match(MatchParserWHITESPACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(25)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&36860) != 0 {
			{
				p.SetState(24)
				p.MatchEntry()
			}

		}

		p.SetState(31)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(32)
		p.Match(MatchParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMatchEntryContext is an interface to support dynamic dispatch.
type IMatchEntryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EntrySingle() IEntrySingleContext
	EntryWithValue() IEntryWithValueContext

	// IsMatchEntryContext differentiates from other interfaces.
	IsMatchEntryContext()
}

type MatchEntryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMatchEntryContext() *MatchEntryContext {
	var p = new(MatchEntryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MatchParserRULE_matchEntry
	return p
}

func InitEmptyMatchEntryContext(p *MatchEntryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MatchParserRULE_matchEntry
}

func (*MatchEntryContext) IsMatchEntryContext() {}

func NewMatchEntryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MatchEntryContext {
	var p = new(MatchEntryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MatchParserRULE_matchEntry

	return p
}

func (s *MatchEntryContext) GetParser() antlr.Parser { return s.parser }

func (s *MatchEntryContext) EntrySingle() IEntrySingleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEntrySingleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEntrySingleContext)
}

func (s *MatchEntryContext) EntryWithValue() IEntryWithValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEntryWithValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEntryWithValueContext)
}

func (s *MatchEntryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MatchEntryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MatchEntryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MatchListener); ok {
		listenerT.EnterMatchEntry(s)
	}
}

func (s *MatchEntryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MatchListener); ok {
		listenerT.ExitMatchEntry(s)
	}
}

func (p *MatchParser) MatchEntry() (localctx IMatchEntryContext) {
	localctx = NewMatchEntryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MatchParserRULE_matchEntry)
	p.SetState(36)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(34)
			p.EntrySingle()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(35)
			p.EntryWithValue()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEntrySingleContext is an interface to support dynamic dispatch.
type IEntrySingleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CriteriaSingle() ICriteriaSingleContext

	// IsEntrySingleContext differentiates from other interfaces.
	IsEntrySingleContext()
}

type EntrySingleContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEntrySingleContext() *EntrySingleContext {
	var p = new(EntrySingleContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MatchParserRULE_entrySingle
	return p
}

func InitEmptyEntrySingleContext(p *EntrySingleContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MatchParserRULE_entrySingle
}

func (*EntrySingleContext) IsEntrySingleContext() {}

func NewEntrySingleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EntrySingleContext {
	var p = new(EntrySingleContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MatchParserRULE_entrySingle

	return p
}

func (s *EntrySingleContext) GetParser() antlr.Parser { return s.parser }

func (s *EntrySingleContext) CriteriaSingle() ICriteriaSingleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICriteriaSingleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICriteriaSingleContext)
}

func (s *EntrySingleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EntrySingleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EntrySingleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MatchListener); ok {
		listenerT.EnterEntrySingle(s)
	}
}

func (s *EntrySingleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MatchListener); ok {
		listenerT.ExitEntrySingle(s)
	}
}

func (p *MatchParser) EntrySingle() (localctx IEntrySingleContext) {
	localctx = NewEntrySingleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, MatchParserRULE_entrySingle)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(38)
		p.CriteriaSingle()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEntryWithValueContext is an interface to support dynamic dispatch.
type IEntryWithValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CriteriaWithValue() ICriteriaWithValueContext
	Separator() ISeparatorContext
	Values() IValuesContext

	// IsEntryWithValueContext differentiates from other interfaces.
	IsEntryWithValueContext()
}

type EntryWithValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEntryWithValueContext() *EntryWithValueContext {
	var p = new(EntryWithValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MatchParserRULE_entryWithValue
	return p
}

func InitEmptyEntryWithValueContext(p *EntryWithValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MatchParserRULE_entryWithValue
}

func (*EntryWithValueContext) IsEntryWithValueContext() {}

func NewEntryWithValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EntryWithValueContext {
	var p = new(EntryWithValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MatchParserRULE_entryWithValue

	return p
}

func (s *EntryWithValueContext) GetParser() antlr.Parser { return s.parser }

func (s *EntryWithValueContext) CriteriaWithValue() ICriteriaWithValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICriteriaWithValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICriteriaWithValueContext)
}

func (s *EntryWithValueContext) Separator() ISeparatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISeparatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISeparatorContext)
}

func (s *EntryWithValueContext) Values() IValuesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValuesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValuesContext)
}

func (s *EntryWithValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EntryWithValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EntryWithValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MatchListener); ok {
		listenerT.EnterEntryWithValue(s)
	}
}

func (s *EntryWithValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MatchListener); ok {
		listenerT.ExitEntryWithValue(s)
	}
}

func (p *MatchParser) EntryWithValue() (localctx IEntryWithValueContext) {
	localctx = NewEntryWithValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, MatchParserRULE_entryWithValue)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(40)
		p.CriteriaWithValue()
	}
	p.SetState(42)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(41)
			p.Separator()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(45)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(44)
			p.Values()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISeparatorContext is an interface to support dynamic dispatch.
type ISeparatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	WHITESPACE() antlr.TerminalNode

	// IsSeparatorContext differentiates from other interfaces.
	IsSeparatorContext()
}

type SeparatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySeparatorContext() *SeparatorContext {
	var p = new(SeparatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MatchParserRULE_separator
	return p
}

func InitEmptySeparatorContext(p *SeparatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MatchParserRULE_separator
}

func (*SeparatorContext) IsSeparatorContext() {}

func NewSeparatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SeparatorContext {
	var p = new(SeparatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MatchParserRULE_separator

	return p
}

func (s *SeparatorContext) GetParser() antlr.Parser { return s.parser }

func (s *SeparatorContext) WHITESPACE() antlr.TerminalNode {
	return s.GetToken(MatchParserWHITESPACE, 0)
}

func (s *SeparatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SeparatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SeparatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MatchListener); ok {
		listenerT.EnterSeparator(s)
	}
}

func (s *SeparatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MatchListener); ok {
		listenerT.ExitSeparator(s)
	}
}

func (p *MatchParser) Separator() (localctx ISeparatorContext) {
	localctx = NewSeparatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, MatchParserRULE_separator)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(47)
		p.Match(MatchParserWHITESPACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IValuesContext is an interface to support dynamic dispatch.
type IValuesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllValue() []IValueContext
	Value(i int) IValueContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsValuesContext differentiates from other interfaces.
	IsValuesContext()
}

type ValuesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValuesContext() *ValuesContext {
	var p = new(ValuesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MatchParserRULE_values
	return p
}

func InitEmptyValuesContext(p *ValuesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MatchParserRULE_values
}

func (*ValuesContext) IsValuesContext() {}

func NewValuesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValuesContext {
	var p = new(ValuesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MatchParserRULE_values

	return p
}

func (s *ValuesContext) GetParser() antlr.Parser { return s.parser }

func (s *ValuesContext) AllValue() []IValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IValueContext); ok {
			len++
		}
	}

	tst := make([]IValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IValueContext); ok {
			tst[i] = t.(IValueContext)
			i++
		}
	}

	return tst
}

func (s *ValuesContext) Value(i int) IValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ValuesContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MatchParserCOMMA)
}

func (s *ValuesContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MatchParserCOMMA, i)
}

func (s *ValuesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValuesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValuesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MatchListener); ok {
		listenerT.EnterValues(s)
	}
}

func (s *ValuesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MatchListener); ok {
		listenerT.ExitValues(s)
	}
}

func (p *MatchParser) Values() (localctx IValuesContext) {
	localctx = NewValuesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, MatchParserRULE_values)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(50)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == MatchParserSTRING || _la == MatchParserQUOTED_STRING {
		{
			p.SetState(49)
			p.Value()
		}

	}
	p.SetState(58)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MatchParserCOMMA {
		{
			p.SetState(52)
			p.Match(MatchParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(54)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MatchParserSTRING || _la == MatchParserQUOTED_STRING {
			{
				p.SetState(53)
				p.Value()
			}

		}

		p.SetState(60)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	String_() IStringContext

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MatchParserRULE_value
	return p
}

func InitEmptyValueContext(p *ValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MatchParserRULE_value
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MatchParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) String_() IStringContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringContext)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MatchListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MatchListener); ok {
		listenerT.ExitValue(s)
	}
}

func (p *MatchParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, MatchParserRULE_value)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(61)
		p.String_()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICriteriaSingleContext is an interface to support dynamic dispatch.
type ICriteriaSingleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ALL() antlr.TerminalNode
	CANONICAL() antlr.TerminalNode
	FINAL() antlr.TerminalNode
	AllQUOTE() []antlr.TerminalNode
	QUOTE(i int) antlr.TerminalNode

	// IsCriteriaSingleContext differentiates from other interfaces.
	IsCriteriaSingleContext()
}

type CriteriaSingleContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCriteriaSingleContext() *CriteriaSingleContext {
	var p = new(CriteriaSingleContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MatchParserRULE_criteriaSingle
	return p
}

func InitEmptyCriteriaSingleContext(p *CriteriaSingleContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MatchParserRULE_criteriaSingle
}

func (*CriteriaSingleContext) IsCriteriaSingleContext() {}

func NewCriteriaSingleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CriteriaSingleContext {
	var p = new(CriteriaSingleContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MatchParserRULE_criteriaSingle

	return p
}

func (s *CriteriaSingleContext) GetParser() antlr.Parser { return s.parser }

func (s *CriteriaSingleContext) ALL() antlr.TerminalNode {
	return s.GetToken(MatchParserALL, 0)
}

func (s *CriteriaSingleContext) CANONICAL() antlr.TerminalNode {
	return s.GetToken(MatchParserCANONICAL, 0)
}

func (s *CriteriaSingleContext) FINAL() antlr.TerminalNode {
	return s.GetToken(MatchParserFINAL, 0)
}

func (s *CriteriaSingleContext) AllQUOTE() []antlr.TerminalNode {
	return s.GetTokens(MatchParserQUOTE)
}

func (s *CriteriaSingleContext) QUOTE(i int) antlr.TerminalNode {
	return s.GetToken(MatchParserQUOTE, i)
}

func (s *CriteriaSingleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CriteriaSingleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CriteriaSingleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MatchListener); ok {
		listenerT.EnterCriteriaSingle(s)
	}
}

func (s *CriteriaSingleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MatchListener); ok {
		listenerT.ExitCriteriaSingle(s)
	}
}

func (p *MatchParser) CriteriaSingle() (localctx ICriteriaSingleContext) {
	localctx = NewCriteriaSingleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, MatchParserRULE_criteriaSingle)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == MatchParserQUOTE {
		{
			p.SetState(63)
			p.Match(MatchParserQUOTE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(66)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&28) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(68)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == MatchParserQUOTE {
		{
			p.SetState(67)
			p.Match(MatchParserQUOTE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICriteriaWithValueContext is an interface to support dynamic dispatch.
type ICriteriaWithValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EXEC() antlr.TerminalNode
	LOCALNETWORK() antlr.TerminalNode
	HOST() antlr.TerminalNode
	ORIGINALHOST() antlr.TerminalNode
	TAGGED() antlr.TerminalNode
	USER() antlr.TerminalNode
	LOCALUSER() antlr.TerminalNode
	AllQUOTE() []antlr.TerminalNode
	QUOTE(i int) antlr.TerminalNode

	// IsCriteriaWithValueContext differentiates from other interfaces.
	IsCriteriaWithValueContext()
}

type CriteriaWithValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCriteriaWithValueContext() *CriteriaWithValueContext {
	var p = new(CriteriaWithValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MatchParserRULE_criteriaWithValue
	return p
}

func InitEmptyCriteriaWithValueContext(p *CriteriaWithValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MatchParserRULE_criteriaWithValue
}

func (*CriteriaWithValueContext) IsCriteriaWithValueContext() {}

func NewCriteriaWithValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CriteriaWithValueContext {
	var p = new(CriteriaWithValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MatchParserRULE_criteriaWithValue

	return p
}

func (s *CriteriaWithValueContext) GetParser() antlr.Parser { return s.parser }

func (s *CriteriaWithValueContext) EXEC() antlr.TerminalNode {
	return s.GetToken(MatchParserEXEC, 0)
}

func (s *CriteriaWithValueContext) LOCALNETWORK() antlr.TerminalNode {
	return s.GetToken(MatchParserLOCALNETWORK, 0)
}

func (s *CriteriaWithValueContext) HOST() antlr.TerminalNode {
	return s.GetToken(MatchParserHOST, 0)
}

func (s *CriteriaWithValueContext) ORIGINALHOST() antlr.TerminalNode {
	return s.GetToken(MatchParserORIGINALHOST, 0)
}

func (s *CriteriaWithValueContext) TAGGED() antlr.TerminalNode {
	return s.GetToken(MatchParserTAGGED, 0)
}

func (s *CriteriaWithValueContext) USER() antlr.TerminalNode {
	return s.GetToken(MatchParserUSER, 0)
}

func (s *CriteriaWithValueContext) LOCALUSER() antlr.TerminalNode {
	return s.GetToken(MatchParserLOCALUSER, 0)
}

func (s *CriteriaWithValueContext) AllQUOTE() []antlr.TerminalNode {
	return s.GetTokens(MatchParserQUOTE)
}

func (s *CriteriaWithValueContext) QUOTE(i int) antlr.TerminalNode {
	return s.GetToken(MatchParserQUOTE, i)
}

func (s *CriteriaWithValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CriteriaWithValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CriteriaWithValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MatchListener); ok {
		listenerT.EnterCriteriaWithValue(s)
	}
}

func (s *CriteriaWithValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MatchListener); ok {
		listenerT.ExitCriteriaWithValue(s)
	}
}

func (p *MatchParser) CriteriaWithValue() (localctx ICriteriaWithValueContext) {
	localctx = NewCriteriaWithValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, MatchParserRULE_criteriaWithValue)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == MatchParserQUOTE {
		{
			p.SetState(70)
			p.Match(MatchParserQUOTE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(73)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4064) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == MatchParserQUOTE {
		{
			p.SetState(74)
			p.Match(MatchParserQUOTE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStringContext is an interface to support dynamic dispatch.
type IStringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	QUOTED_STRING() antlr.TerminalNode
	STRING() antlr.TerminalNode

	// IsStringContext differentiates from other interfaces.
	IsStringContext()
}

type StringContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringContext() *StringContext {
	var p = new(StringContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MatchParserRULE_string
	return p
}

func InitEmptyStringContext(p *StringContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MatchParserRULE_string
}

func (*StringContext) IsStringContext() {}

func NewStringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringContext {
	var p = new(StringContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MatchParserRULE_string

	return p
}

func (s *StringContext) GetParser() antlr.Parser { return s.parser }

func (s *StringContext) QUOTED_STRING() antlr.TerminalNode {
	return s.GetToken(MatchParserQUOTED_STRING, 0)
}

func (s *StringContext) STRING() antlr.TerminalNode {
	return s.GetToken(MatchParserSTRING, 0)
}

func (s *StringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MatchListener); ok {
		listenerT.EnterString(s)
	}
}

func (s *StringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MatchListener); ok {
		listenerT.ExitString(s)
	}
}

func (p *MatchParser) String_() (localctx IStringContext) {
	localctx = NewStringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, MatchParserRULE_string)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(77)
		_la = p.GetTokenStream().LA(1)

		if !(_la == MatchParserSTRING || _la == MatchParserQUOTED_STRING) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
