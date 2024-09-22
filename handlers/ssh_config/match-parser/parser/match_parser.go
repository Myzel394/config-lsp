// Code generated from Match.g4 by ANTLR 4.13.0. DO NOT EDIT.

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
		"", "','",
	}
	staticData.SymbolicNames = []string{
		"", "COMMA", "STRING", "WHITESPACE", "QUOTED_STRING",
	}
	staticData.RuleNames = []string{
		"root", "matchEntry", "separator", "criteria", "values", "value", "string",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 4, 56, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7, 4,
		2, 5, 7, 5, 2, 6, 7, 6, 1, 0, 3, 0, 16, 8, 0, 1, 0, 1, 0, 3, 0, 20, 8,
		0, 5, 0, 22, 8, 0, 10, 0, 12, 0, 25, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 3, 1,
		31, 8, 1, 1, 1, 3, 1, 34, 8, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 3, 4, 41,
		8, 4, 1, 4, 1, 4, 3, 4, 45, 8, 4, 5, 4, 47, 8, 4, 10, 4, 12, 4, 50, 9,
		4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 0, 0, 7, 0, 2, 4, 6, 8, 10, 12, 0, 1,
		2, 0, 2, 2, 4, 4, 56, 0, 15, 1, 0, 0, 0, 2, 28, 1, 0, 0, 0, 4, 35, 1, 0,
		0, 0, 6, 37, 1, 0, 0, 0, 8, 40, 1, 0, 0, 0, 10, 51, 1, 0, 0, 0, 12, 53,
		1, 0, 0, 0, 14, 16, 3, 2, 1, 0, 15, 14, 1, 0, 0, 0, 15, 16, 1, 0, 0, 0,
		16, 23, 1, 0, 0, 0, 17, 19, 5, 3, 0, 0, 18, 20, 3, 2, 1, 0, 19, 18, 1,
		0, 0, 0, 19, 20, 1, 0, 0, 0, 20, 22, 1, 0, 0, 0, 21, 17, 1, 0, 0, 0, 22,
		25, 1, 0, 0, 0, 23, 21, 1, 0, 0, 0, 23, 24, 1, 0, 0, 0, 24, 26, 1, 0, 0,
		0, 25, 23, 1, 0, 0, 0, 26, 27, 5, 0, 0, 1, 27, 1, 1, 0, 0, 0, 28, 30, 3,
		6, 3, 0, 29, 31, 3, 4, 2, 0, 30, 29, 1, 0, 0, 0, 30, 31, 1, 0, 0, 0, 31,
		33, 1, 0, 0, 0, 32, 34, 3, 8, 4, 0, 33, 32, 1, 0, 0, 0, 33, 34, 1, 0, 0,
		0, 34, 3, 1, 0, 0, 0, 35, 36, 5, 3, 0, 0, 36, 5, 1, 0, 0, 0, 37, 38, 3,
		12, 6, 0, 38, 7, 1, 0, 0, 0, 39, 41, 3, 10, 5, 0, 40, 39, 1, 0, 0, 0, 40,
		41, 1, 0, 0, 0, 41, 48, 1, 0, 0, 0, 42, 44, 5, 1, 0, 0, 43, 45, 3, 10,
		5, 0, 44, 43, 1, 0, 0, 0, 44, 45, 1, 0, 0, 0, 45, 47, 1, 0, 0, 0, 46, 42,
		1, 0, 0, 0, 47, 50, 1, 0, 0, 0, 48, 46, 1, 0, 0, 0, 48, 49, 1, 0, 0, 0,
		49, 9, 1, 0, 0, 0, 50, 48, 1, 0, 0, 0, 51, 52, 3, 12, 6, 0, 52, 11, 1,
		0, 0, 0, 53, 54, 7, 0, 0, 0, 54, 13, 1, 0, 0, 0, 8, 15, 19, 23, 30, 33,
		40, 44, 48,
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
	MatchParserSTRING        = 2
	MatchParserWHITESPACE    = 3
	MatchParserQUOTED_STRING = 4
)

// MatchParser rules.
const (
	MatchParserRULE_root       = 0
	MatchParserRULE_matchEntry = 1
	MatchParserRULE_separator  = 2
	MatchParserRULE_criteria   = 3
	MatchParserRULE_values     = 4
	MatchParserRULE_value      = 5
	MatchParserRULE_string     = 6
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
	p.SetState(15)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == MatchParserSTRING || _la == MatchParserQUOTED_STRING {
		{
			p.SetState(14)
			p.MatchEntry()
		}

	}
	p.SetState(23)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MatchParserWHITESPACE {
		{
			p.SetState(17)
			p.Match(MatchParserWHITESPACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(19)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MatchParserSTRING || _la == MatchParserQUOTED_STRING {
			{
				p.SetState(18)
				p.MatchEntry()
			}

		}

		p.SetState(25)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(26)
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
	Criteria() ICriteriaContext
	Separator() ISeparatorContext
	Values() IValuesContext

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

func (s *MatchEntryContext) Criteria() ICriteriaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICriteriaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICriteriaContext)
}

func (s *MatchEntryContext) Separator() ISeparatorContext {
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

func (s *MatchEntryContext) Values() IValuesContext {
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
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(28)
		p.Criteria()
	}
	p.SetState(30)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(29)
			p.Separator()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(33)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(32)
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
	p.EnterRule(localctx, 4, MatchParserRULE_separator)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(35)
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

// ICriteriaContext is an interface to support dynamic dispatch.
type ICriteriaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	String_() IStringContext

	// IsCriteriaContext differentiates from other interfaces.
	IsCriteriaContext()
}

type CriteriaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCriteriaContext() *CriteriaContext {
	var p = new(CriteriaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MatchParserRULE_criteria
	return p
}

func InitEmptyCriteriaContext(p *CriteriaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MatchParserRULE_criteria
}

func (*CriteriaContext) IsCriteriaContext() {}

func NewCriteriaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CriteriaContext {
	var p = new(CriteriaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MatchParserRULE_criteria

	return p
}

func (s *CriteriaContext) GetParser() antlr.Parser { return s.parser }

func (s *CriteriaContext) String_() IStringContext {
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

func (s *CriteriaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CriteriaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CriteriaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MatchListener); ok {
		listenerT.EnterCriteria(s)
	}
}

func (s *CriteriaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MatchListener); ok {
		listenerT.ExitCriteria(s)
	}
}

func (p *MatchParser) Criteria() (localctx ICriteriaContext) {
	localctx = NewCriteriaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, MatchParserRULE_criteria)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(37)
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
	p.EnterRule(localctx, 8, MatchParserRULE_values)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(40)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == MatchParserSTRING || _la == MatchParserQUOTED_STRING {
		{
			p.SetState(39)
			p.Value()
		}

	}
	p.SetState(48)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MatchParserCOMMA {
		{
			p.SetState(42)
			p.Match(MatchParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(44)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MatchParserSTRING || _la == MatchParserQUOTED_STRING {
			{
				p.SetState(43)
				p.Value()
			}

		}

		p.SetState(50)
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
	p.EnterRule(localctx, 10, MatchParserRULE_value)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(51)
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
	p.EnterRule(localctx, 12, MatchParserRULE_string)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(53)
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
