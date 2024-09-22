// Code generated from handlers/sshd_config/Config.g4 by ANTLR 4.13.0. DO NOT EDIT.

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
		"", "'#'",
	}
	staticData.SymbolicNames = []string{
		"", "HASH", "WHITESPACE", "STRING", "NEWLINE", "QUOTED_STRING",
	}
	staticData.RuleNames = []string{
		"lineStatement", "entry", "separator", "key", "value", "leadingComment",
		"string",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 5, 71, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7, 4,
		2, 5, 7, 5, 2, 6, 7, 6, 1, 0, 1, 0, 1, 0, 3, 0, 18, 8, 0, 3, 0, 20, 8,
		0, 1, 0, 1, 0, 1, 1, 3, 1, 25, 8, 1, 1, 1, 3, 1, 28, 8, 1, 1, 1, 3, 1,
		31, 8, 1, 1, 1, 3, 1, 34, 8, 1, 1, 1, 3, 1, 37, 8, 1, 1, 2, 1, 2, 1, 3,
		1, 3, 1, 4, 1, 4, 1, 4, 5, 4, 46, 8, 4, 10, 4, 12, 4, 49, 9, 4, 1, 4, 3,
		4, 52, 8, 4, 1, 4, 3, 4, 55, 8, 4, 1, 5, 1, 5, 3, 5, 59, 8, 5, 1, 5, 1,
		5, 3, 5, 63, 8, 5, 4, 5, 65, 8, 5, 11, 5, 12, 5, 66, 1, 6, 1, 6, 1, 6,
		0, 0, 7, 0, 2, 4, 6, 8, 10, 12, 0, 1, 2, 0, 3, 3, 5, 5, 77, 0, 19, 1, 0,
		0, 0, 2, 24, 1, 0, 0, 0, 4, 38, 1, 0, 0, 0, 6, 40, 1, 0, 0, 0, 8, 47, 1,
		0, 0, 0, 10, 56, 1, 0, 0, 0, 12, 68, 1, 0, 0, 0, 14, 20, 3, 2, 1, 0, 15,
		20, 3, 10, 5, 0, 16, 18, 5, 2, 0, 0, 17, 16, 1, 0, 0, 0, 17, 18, 1, 0,
		0, 0, 18, 20, 1, 0, 0, 0, 19, 14, 1, 0, 0, 0, 19, 15, 1, 0, 0, 0, 19, 17,
		1, 0, 0, 0, 20, 21, 1, 0, 0, 0, 21, 22, 5, 0, 0, 1, 22, 1, 1, 0, 0, 0,
		23, 25, 5, 2, 0, 0, 24, 23, 1, 0, 0, 0, 24, 25, 1, 0, 0, 0, 25, 27, 1,
		0, 0, 0, 26, 28, 3, 6, 3, 0, 27, 26, 1, 0, 0, 0, 27, 28, 1, 0, 0, 0, 28,
		30, 1, 0, 0, 0, 29, 31, 3, 4, 2, 0, 30, 29, 1, 0, 0, 0, 30, 31, 1, 0, 0,
		0, 31, 33, 1, 0, 0, 0, 32, 34, 3, 8, 4, 0, 33, 32, 1, 0, 0, 0, 33, 34,
		1, 0, 0, 0, 34, 36, 1, 0, 0, 0, 35, 37, 3, 10, 5, 0, 36, 35, 1, 0, 0, 0,
		36, 37, 1, 0, 0, 0, 37, 3, 1, 0, 0, 0, 38, 39, 5, 2, 0, 0, 39, 5, 1, 0,
		0, 0, 40, 41, 3, 12, 6, 0, 41, 7, 1, 0, 0, 0, 42, 43, 3, 12, 6, 0, 43,
		44, 5, 2, 0, 0, 44, 46, 1, 0, 0, 0, 45, 42, 1, 0, 0, 0, 46, 49, 1, 0, 0,
		0, 47, 45, 1, 0, 0, 0, 47, 48, 1, 0, 0, 0, 48, 51, 1, 0, 0, 0, 49, 47,
		1, 0, 0, 0, 50, 52, 3, 12, 6, 0, 51, 50, 1, 0, 0, 0, 51, 52, 1, 0, 0, 0,
		52, 54, 1, 0, 0, 0, 53, 55, 5, 2, 0, 0, 54, 53, 1, 0, 0, 0, 54, 55, 1,
		0, 0, 0, 55, 9, 1, 0, 0, 0, 56, 58, 5, 1, 0, 0, 57, 59, 5, 2, 0, 0, 58,
		57, 1, 0, 0, 0, 58, 59, 1, 0, 0, 0, 59, 64, 1, 0, 0, 0, 60, 62, 3, 12,
		6, 0, 61, 63, 5, 2, 0, 0, 62, 61, 1, 0, 0, 0, 62, 63, 1, 0, 0, 0, 63, 65,
		1, 0, 0, 0, 64, 60, 1, 0, 0, 0, 65, 66, 1, 0, 0, 0, 66, 64, 1, 0, 0, 0,
		66, 67, 1, 0, 0, 0, 67, 11, 1, 0, 0, 0, 68, 69, 7, 0, 0, 0, 69, 13, 1,
		0, 0, 0, 13, 17, 19, 24, 27, 30, 33, 36, 47, 51, 54, 58, 62, 66,
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
	ConfigParserHASH          = 1
	ConfigParserWHITESPACE    = 2
	ConfigParserSTRING        = 3
	ConfigParserNEWLINE       = 4
	ConfigParserQUOTED_STRING = 5
)

// ConfigParser rules.
const (
	ConfigParserRULE_lineStatement  = 0
	ConfigParserRULE_entry          = 1
	ConfigParserRULE_separator      = 2
	ConfigParserRULE_key            = 3
	ConfigParserRULE_value          = 4
	ConfigParserRULE_leadingComment = 5
	ConfigParserRULE_string         = 6
)

// ILineStatementContext is an interface to support dynamic dispatch.
type ILineStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	Entry() IEntryContext
	LeadingComment() ILeadingCommentContext
	WHITESPACE() antlr.TerminalNode

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

func (s *LineStatementContext) EOF() antlr.TerminalNode {
	return s.GetToken(ConfigParserEOF, 0)
}

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

func (s *LineStatementContext) WHITESPACE() antlr.TerminalNode {
	return s.GetToken(ConfigParserWHITESPACE, 0)
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
	p.SetState(19)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(14)
			p.Entry()
		}

	case 2:
		{
			p.SetState(15)
			p.LeadingComment()
		}

	case 3:
		p.SetState(17)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ConfigParserWHITESPACE {
			{
				p.SetState(16)
				p.Match(ConfigParserWHITESPACE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	{
		p.SetState(21)
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
	WHITESPACE() antlr.TerminalNode
	Key() IKeyContext
	Separator() ISeparatorContext
	Value() IValueContext
	LeadingComment() ILeadingCommentContext

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

func (s *EntryContext) WHITESPACE() antlr.TerminalNode {
	return s.GetToken(ConfigParserWHITESPACE, 0)
}

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

func (s *EntryContext) LeadingComment() ILeadingCommentContext {
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
	p.SetState(24)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(23)
			p.Match(ConfigParserWHITESPACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(27)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(26)
			p.Key()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(30)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(29)
			p.Separator()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(33)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(32)
			p.Value()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(36)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ConfigParserHASH {
		{
			p.SetState(35)
			p.LeadingComment()
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

func (s *SeparatorContext) WHITESPACE() antlr.TerminalNode {
	return s.GetToken(ConfigParserWHITESPACE, 0)
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
	p.EnterRule(localctx, 4, ConfigParserRULE_separator)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(38)
		p.Match(ConfigParserWHITESPACE)
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
		p.SetState(40)
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

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllString_() []IStringContext
	String_(i int) IStringContext
	AllWHITESPACE() []antlr.TerminalNode
	WHITESPACE(i int) antlr.TerminalNode

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

func (s *ValueContext) AllString_() []IStringContext {
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

func (s *ValueContext) String_(i int) IStringContext {
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

func (s *ValueContext) AllWHITESPACE() []antlr.TerminalNode {
	return s.GetTokens(ConfigParserWHITESPACE)
}

func (s *ValueContext) WHITESPACE(i int) antlr.TerminalNode {
	return s.GetToken(ConfigParserWHITESPACE, i)
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
	p.EnterRule(localctx, 8, ConfigParserRULE_value)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(47)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(42)
				p.String_()
			}
			{
				p.SetState(43)
				p.Match(ConfigParserWHITESPACE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(49)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ConfigParserSTRING || _la == ConfigParserQUOTED_STRING {
		{
			p.SetState(50)
			p.String_()
		}

	}
	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ConfigParserWHITESPACE {
		{
			p.SetState(53)
			p.Match(ConfigParserWHITESPACE)
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

// ILeadingCommentContext is an interface to support dynamic dispatch.
type ILeadingCommentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	HASH() antlr.TerminalNode
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

func (s *LeadingCommentContext) HASH() antlr.TerminalNode {
	return s.GetToken(ConfigParserHASH, 0)
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
	p.EnterRule(localctx, 10, ConfigParserRULE_leadingComment)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(56)
		p.Match(ConfigParserHASH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(58)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ConfigParserWHITESPACE {
		{
			p.SetState(57)
			p.Match(ConfigParserWHITESPACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ConfigParserSTRING || _la == ConfigParserQUOTED_STRING {
		{
			p.SetState(60)
			p.String_()
		}
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

		p.SetState(66)
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

func (s *StringContext) QUOTED_STRING() antlr.TerminalNode {
	return s.GetToken(ConfigParserQUOTED_STRING, 0)
}

func (s *StringContext) STRING() antlr.TerminalNode {
	return s.GetToken(ConfigParserSTRING, 0)
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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(68)
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
