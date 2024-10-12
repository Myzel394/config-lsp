// Code generated from Config.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Config

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

type ConfigParser struct {
	*antlr.BaseParser
}

var ConfigParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func configParserInit() {
	staticData := &ConfigParserStaticData
	staticData.LiteralNames = []string{
		"", "'='", "'#'", "';'",
	}
	staticData.SymbolicNames = []string{
		"", "EQUAL", "HASH", "SEMICOLON", "WHITESPACE", "STRING", "QUOTED_STRING",
	}
	staticData.RuleNames = []string{
		"lineStatement", "entry", "leadingComment", "key", "separator", "value",
		"string", "commentSymbol",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 6, 76, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7, 4,
		2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 1, 0, 3, 0, 18, 8, 0, 1, 0, 1, 0, 3,
		0, 22, 8, 0, 1, 0, 3, 0, 25, 8, 0, 1, 0, 1, 0, 1, 1, 3, 1, 30, 8, 1, 1,
		1, 3, 1, 33, 8, 1, 1, 1, 3, 1, 36, 8, 1, 1, 1, 3, 1, 39, 8, 1, 1, 1, 3,
		1, 42, 8, 1, 1, 2, 1, 2, 3, 2, 46, 8, 2, 1, 2, 1, 2, 3, 2, 50, 8, 2, 4,
		2, 52, 8, 2, 11, 2, 12, 2, 53, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6,
		3, 6, 63, 8, 6, 1, 6, 1, 6, 5, 6, 67, 8, 6, 10, 6, 12, 6, 70, 9, 6, 1,
		6, 1, 6, 1, 7, 1, 7, 1, 7, 0, 0, 8, 0, 2, 4, 6, 8, 10, 12, 14, 0, 2, 1,
		0, 5, 6, 1, 0, 2, 3, 80, 0, 17, 1, 0, 0, 0, 2, 29, 1, 0, 0, 0, 4, 43, 1,
		0, 0, 0, 6, 55, 1, 0, 0, 0, 8, 57, 1, 0, 0, 0, 10, 59, 1, 0, 0, 0, 12,
		62, 1, 0, 0, 0, 14, 73, 1, 0, 0, 0, 16, 18, 5, 4, 0, 0, 17, 16, 1, 0, 0,
		0, 17, 18, 1, 0, 0, 0, 18, 19, 1, 0, 0, 0, 19, 21, 3, 2, 1, 0, 20, 22,
		5, 4, 0, 0, 21, 20, 1, 0, 0, 0, 21, 22, 1, 0, 0, 0, 22, 24, 1, 0, 0, 0,
		23, 25, 3, 4, 2, 0, 24, 23, 1, 0, 0, 0, 24, 25, 1, 0, 0, 0, 25, 26, 1,
		0, 0, 0, 26, 27, 5, 0, 0, 1, 27, 1, 1, 0, 0, 0, 28, 30, 3, 6, 3, 0, 29,
		28, 1, 0, 0, 0, 29, 30, 1, 0, 0, 0, 30, 32, 1, 0, 0, 0, 31, 33, 5, 4, 0,
		0, 32, 31, 1, 0, 0, 0, 32, 33, 1, 0, 0, 0, 33, 35, 1, 0, 0, 0, 34, 36,
		3, 8, 4, 0, 35, 34, 1, 0, 0, 0, 35, 36, 1, 0, 0, 0, 36, 38, 1, 0, 0, 0,
		37, 39, 5, 4, 0, 0, 38, 37, 1, 0, 0, 0, 38, 39, 1, 0, 0, 0, 39, 41, 1,
		0, 0, 0, 40, 42, 3, 10, 5, 0, 41, 40, 1, 0, 0, 0, 41, 42, 1, 0, 0, 0, 42,
		3, 1, 0, 0, 0, 43, 45, 3, 14, 7, 0, 44, 46, 5, 4, 0, 0, 45, 44, 1, 0, 0,
		0, 45, 46, 1, 0, 0, 0, 46, 51, 1, 0, 0, 0, 47, 49, 3, 12, 6, 0, 48, 50,
		5, 4, 0, 0, 49, 48, 1, 0, 0, 0, 49, 50, 1, 0, 0, 0, 50, 52, 1, 0, 0, 0,
		51, 47, 1, 0, 0, 0, 52, 53, 1, 0, 0, 0, 53, 51, 1, 0, 0, 0, 53, 54, 1,
		0, 0, 0, 54, 5, 1, 0, 0, 0, 55, 56, 3, 12, 6, 0, 56, 7, 1, 0, 0, 0, 57,
		58, 5, 1, 0, 0, 58, 9, 1, 0, 0, 0, 59, 60, 3, 12, 6, 0, 60, 11, 1, 0, 0,
		0, 61, 63, 5, 4, 0, 0, 62, 61, 1, 0, 0, 0, 62, 63, 1, 0, 0, 0, 63, 68,
		1, 0, 0, 0, 64, 65, 7, 0, 0, 0, 65, 67, 5, 4, 0, 0, 66, 64, 1, 0, 0, 0,
		67, 70, 1, 0, 0, 0, 68, 66, 1, 0, 0, 0, 68, 69, 1, 0, 0, 0, 69, 71, 1,
		0, 0, 0, 70, 68, 1, 0, 0, 0, 71, 72, 7, 0, 0, 0, 72, 13, 1, 0, 0, 0, 73,
		74, 7, 1, 0, 0, 74, 15, 1, 0, 0, 0, 13, 17, 21, 24, 29, 32, 35, 38, 41,
		45, 49, 53, 62, 68,
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

// ConfigParserInit initializes any static state used to implement ConfigParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewConfigParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ConfigParserInit() {
	staticData := &ConfigParserStaticData
	staticData.once.Do(configParserInit)
}

// NewConfigParser produces a new parser instance for the optional input antlr.TokenStream.
func NewConfigParser(input antlr.TokenStream) *ConfigParser {
	ConfigParserInit()
	this := new(ConfigParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &ConfigParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Config.g4"

	return this
}

// ConfigParser tokens.
const (
	ConfigParserEOF           = antlr.TokenEOF
	ConfigParserEQUAL         = 1
	ConfigParserHASH          = 2
	ConfigParserSEMICOLON     = 3
	ConfigParserWHITESPACE    = 4
	ConfigParserSTRING        = 5
	ConfigParserQUOTED_STRING = 6
)

// ConfigParser rules.
const (
	ConfigParserRULE_lineStatement  = 0
	ConfigParserRULE_entry          = 1
	ConfigParserRULE_leadingComment = 2
	ConfigParserRULE_key            = 3
	ConfigParserRULE_separator      = 4
	ConfigParserRULE_value          = 5
	ConfigParserRULE_string         = 6
	ConfigParserRULE_commentSymbol  = 7
)

// ILineStatementContext is an interface to support dynamic dispatch.
type ILineStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Entry() IEntryContext
	EOF() antlr.TerminalNode
	AllWHITESPACE() []antlr.TerminalNode
	WHITESPACE(i int) antlr.TerminalNode
	LeadingComment() ILeadingCommentContext

	// IsLineStatementContext differentiates from other interfaces.
	IsLineStatementContext()
}

type LineStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLineStatementContext() *LineStatementContext {
	var p = new(LineStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ConfigParserRULE_lineStatement
	return p
}

func InitEmptyLineStatementContext(p *LineStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ConfigParserRULE_lineStatement
}

func (*LineStatementContext) IsLineStatementContext() {}

func NewLineStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LineStatementContext {
	var p = new(LineStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConfigParserRULE_lineStatement

	return p
}

func (s *LineStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *LineStatementContext) Entry() IEntryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEntryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEntryContext)
}

func (s *LineStatementContext) EOF() antlr.TerminalNode {
	return s.GetToken(ConfigParserEOF, 0)
}

func (s *LineStatementContext) AllWHITESPACE() []antlr.TerminalNode {
	return s.GetTokens(ConfigParserWHITESPACE)
}

func (s *LineStatementContext) WHITESPACE(i int) antlr.TerminalNode {
	return s.GetToken(ConfigParserWHITESPACE, i)
}

func (s *LineStatementContext) LeadingComment() ILeadingCommentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILeadingCommentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILeadingCommentContext)
}

func (s *LineStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LineStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LineStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfigListener); ok {
		listenerT.EnterLineStatement(s)
	}
}

func (s *LineStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfigListener); ok {
		listenerT.ExitLineStatement(s)
	}
}

func (p *ConfigParser) LineStatement() (localctx ILineStatementContext) {
	localctx = NewLineStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ConfigParserRULE_lineStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(17)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(16)
			p.Match(ConfigParserWHITESPACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(19)
		p.Entry()
	}
	p.SetState(21)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ConfigParserWHITESPACE {
		{
			p.SetState(20)
			p.Match(ConfigParserWHITESPACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(24)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ConfigParserHASH || _la == ConfigParserSEMICOLON {
		{
			p.SetState(23)
			p.LeadingComment()
		}

	}
	{
		p.SetState(26)
		p.Match(ConfigParserEOF)
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

// IEntryContext is an interface to support dynamic dispatch.
type IEntryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Key() IKeyContext
	AllWHITESPACE() []antlr.TerminalNode
	WHITESPACE(i int) antlr.TerminalNode
	Separator() ISeparatorContext
	Value() IValueContext

	// IsEntryContext differentiates from other interfaces.
	IsEntryContext()
}

type EntryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEntryContext() *EntryContext {
	var p = new(EntryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ConfigParserRULE_entry
	return p
}

func InitEmptyEntryContext(p *EntryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ConfigParserRULE_entry
}

func (*EntryContext) IsEntryContext() {}

func NewEntryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EntryContext {
	var p = new(EntryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConfigParserRULE_entry

	return p
}

func (s *EntryContext) GetParser() antlr.Parser { return s.parser }

func (s *EntryContext) Key() IKeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKeyContext)
}

func (s *EntryContext) AllWHITESPACE() []antlr.TerminalNode {
	return s.GetTokens(ConfigParserWHITESPACE)
}

func (s *EntryContext) WHITESPACE(i int) antlr.TerminalNode {
	return s.GetToken(ConfigParserWHITESPACE, i)
}

func (s *EntryContext) Separator() ISeparatorContext {
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

func (s *EntryContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *EntryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EntryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EntryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfigListener); ok {
		listenerT.EnterEntry(s)
	}
}

func (s *EntryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfigListener); ok {
		listenerT.ExitEntry(s)
	}
}

func (p *ConfigParser) Entry() (localctx IEntryContext) {
	localctx = NewEntryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ConfigParserRULE_entry)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(29)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(28)
			p.Key()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(32)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(31)
			p.Match(ConfigParserWHITESPACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(35)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ConfigParserEQUAL {
		{
			p.SetState(34)
			p.Separator()
		}

	}
	p.SetState(38)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(37)
			p.Match(ConfigParserWHITESPACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(41)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(40)
			p.Value()
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

// ILeadingCommentContext is an interface to support dynamic dispatch.
type ILeadingCommentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CommentSymbol() ICommentSymbolContext
	AllWHITESPACE() []antlr.TerminalNode
	WHITESPACE(i int) antlr.TerminalNode
	AllString_() []IStringContext
	String_(i int) IStringContext

	// IsLeadingCommentContext differentiates from other interfaces.
	IsLeadingCommentContext()
}

type LeadingCommentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLeadingCommentContext() *LeadingCommentContext {
	var p = new(LeadingCommentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ConfigParserRULE_leadingComment
	return p
}

func InitEmptyLeadingCommentContext(p *LeadingCommentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ConfigParserRULE_leadingComment
}

func (*LeadingCommentContext) IsLeadingCommentContext() {}

func NewLeadingCommentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LeadingCommentContext {
	var p = new(LeadingCommentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConfigParserRULE_leadingComment

	return p
}

func (s *LeadingCommentContext) GetParser() antlr.Parser { return s.parser }

func (s *LeadingCommentContext) CommentSymbol() ICommentSymbolContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommentSymbolContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICommentSymbolContext)
}

func (s *LeadingCommentContext) AllWHITESPACE() []antlr.TerminalNode {
	return s.GetTokens(ConfigParserWHITESPACE)
}

func (s *LeadingCommentContext) WHITESPACE(i int) antlr.TerminalNode {
	return s.GetToken(ConfigParserWHITESPACE, i)
}

func (s *LeadingCommentContext) AllString_() []IStringContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStringContext); ok {
			len++
		}
	}

	tst := make([]IStringContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStringContext); ok {
			tst[i] = t.(IStringContext)
			i++
		}
	}

	return tst
}

func (s *LeadingCommentContext) String_(i int) IStringContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringContext); ok {
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

	return t.(IStringContext)
}

func (s *LeadingCommentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LeadingCommentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LeadingCommentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfigListener); ok {
		listenerT.EnterLeadingComment(s)
	}
}

func (s *LeadingCommentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfigListener); ok {
		listenerT.ExitLeadingComment(s)
	}
}

func (p *ConfigParser) LeadingComment() (localctx ILeadingCommentContext) {
	localctx = NewLeadingCommentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ConfigParserRULE_leadingComment)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(43)
		p.CommentSymbol()
	}
	p.SetState(45)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(44)
			p.Match(ConfigParserWHITESPACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&112) != 0) {
		{
			p.SetState(47)
			p.String_()
		}
		p.SetState(49)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(48)
				p.Match(ConfigParserWHITESPACE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

		p.SetState(53)
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

// IKeyContext is an interface to support dynamic dispatch.
type IKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	String_() IStringContext

	// IsKeyContext differentiates from other interfaces.
	IsKeyContext()
}

type KeyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeyContext() *KeyContext {
	var p = new(KeyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ConfigParserRULE_key
	return p
}

func InitEmptyKeyContext(p *KeyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ConfigParserRULE_key
}

func (*KeyContext) IsKeyContext() {}

func NewKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyContext {
	var p = new(KeyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConfigParserRULE_key

	return p
}

func (s *KeyContext) GetParser() antlr.Parser { return s.parser }

func (s *KeyContext) String_() IStringContext {
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

func (s *KeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfigListener); ok {
		listenerT.EnterKey(s)
	}
}

func (s *KeyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfigListener); ok {
		listenerT.ExitKey(s)
	}
}

func (p *ConfigParser) Key() (localctx IKeyContext) {
	localctx = NewKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ConfigParserRULE_key)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(55)
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

// ISeparatorContext is an interface to support dynamic dispatch.
type ISeparatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EQUAL() antlr.TerminalNode

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
	p.RuleIndex = ConfigParserRULE_separator
	return p
}

func InitEmptySeparatorContext(p *SeparatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ConfigParserRULE_separator
}

func (*SeparatorContext) IsSeparatorContext() {}

func NewSeparatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SeparatorContext {
	var p = new(SeparatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConfigParserRULE_separator

	return p
}

func (s *SeparatorContext) GetParser() antlr.Parser { return s.parser }

func (s *SeparatorContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(ConfigParserEQUAL, 0)
}

func (s *SeparatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SeparatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SeparatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfigListener); ok {
		listenerT.EnterSeparator(s)
	}
}

func (s *SeparatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfigListener); ok {
		listenerT.ExitSeparator(s)
	}
}

func (p *ConfigParser) Separator() (localctx ISeparatorContext) {
	localctx = NewSeparatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ConfigParserRULE_separator)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(57)
		p.Match(ConfigParserEQUAL)
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
	p.RuleIndex = ConfigParserRULE_value
	return p
}

func InitEmptyValueContext(p *ValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ConfigParserRULE_value
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConfigParserRULE_value

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
	if listenerT, ok := listener.(ConfigListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfigListener); ok {
		listenerT.ExitValue(s)
	}
}

func (p *ConfigParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ConfigParserRULE_value)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(59)
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
	AllSTRING() []antlr.TerminalNode
	STRING(i int) antlr.TerminalNode
	AllQUOTED_STRING() []antlr.TerminalNode
	QUOTED_STRING(i int) antlr.TerminalNode
	AllWHITESPACE() []antlr.TerminalNode
	WHITESPACE(i int) antlr.TerminalNode

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
	p.RuleIndex = ConfigParserRULE_string
	return p
}

func InitEmptyStringContext(p *StringContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ConfigParserRULE_string
}

func (*StringContext) IsStringContext() {}

func NewStringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringContext {
	var p = new(StringContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConfigParserRULE_string

	return p
}

func (s *StringContext) GetParser() antlr.Parser { return s.parser }

func (s *StringContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(ConfigParserSTRING)
}

func (s *StringContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(ConfigParserSTRING, i)
}

func (s *StringContext) AllQUOTED_STRING() []antlr.TerminalNode {
	return s.GetTokens(ConfigParserQUOTED_STRING)
}

func (s *StringContext) QUOTED_STRING(i int) antlr.TerminalNode {
	return s.GetToken(ConfigParserQUOTED_STRING, i)
}

func (s *StringContext) AllWHITESPACE() []antlr.TerminalNode {
	return s.GetTokens(ConfigParserWHITESPACE)
}

func (s *StringContext) WHITESPACE(i int) antlr.TerminalNode {
	return s.GetToken(ConfigParserWHITESPACE, i)
}

func (s *StringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfigListener); ok {
		listenerT.EnterString(s)
	}
}

func (s *StringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfigListener); ok {
		listenerT.ExitString(s)
	}
}

func (p *ConfigParser) String_() (localctx IStringContext) {
	localctx = NewStringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ConfigParserRULE_string)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(62)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ConfigParserWHITESPACE {
		{
			p.SetState(61)
			p.Match(ConfigParserWHITESPACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(68)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(64)
				_la = p.GetTokenStream().LA(1)

				if !(_la == ConfigParserSTRING || _la == ConfigParserQUOTED_STRING) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(65)
				p.Match(ConfigParserWHITESPACE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(70)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	{
		p.SetState(71)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ConfigParserSTRING || _la == ConfigParserQUOTED_STRING) {
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

// ICommentSymbolContext is an interface to support dynamic dispatch.
type ICommentSymbolContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	HASH() antlr.TerminalNode
	SEMICOLON() antlr.TerminalNode

	// IsCommentSymbolContext differentiates from other interfaces.
	IsCommentSymbolContext()
}

type CommentSymbolContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommentSymbolContext() *CommentSymbolContext {
	var p = new(CommentSymbolContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ConfigParserRULE_commentSymbol
	return p
}

func InitEmptyCommentSymbolContext(p *CommentSymbolContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ConfigParserRULE_commentSymbol
}

func (*CommentSymbolContext) IsCommentSymbolContext() {}

func NewCommentSymbolContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommentSymbolContext {
	var p = new(CommentSymbolContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConfigParserRULE_commentSymbol

	return p
}

func (s *CommentSymbolContext) GetParser() antlr.Parser { return s.parser }

func (s *CommentSymbolContext) HASH() antlr.TerminalNode {
	return s.GetToken(ConfigParserHASH, 0)
}

func (s *CommentSymbolContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(ConfigParserSEMICOLON, 0)
}

func (s *CommentSymbolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommentSymbolContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommentSymbolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfigListener); ok {
		listenerT.EnterCommentSymbol(s)
	}
}

func (s *CommentSymbolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfigListener); ok {
		listenerT.ExitCommentSymbol(s)
	}
}

func (p *ConfigParser) CommentSymbol() (localctx ICommentSymbolContext) {
	localctx = NewCommentSymbolContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ConfigParserRULE_commentSymbol)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(73)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ConfigParserHASH || _la == ConfigParserSEMICOLON) {
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
