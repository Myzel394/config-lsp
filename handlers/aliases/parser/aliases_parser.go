// Code generated from Aliases.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Aliases

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

type AliasesParser struct {
	*antlr.BaseParser
}

var AliasesParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func aliasesParserInit() {
	staticData := &AliasesParserStaticData
	staticData.LiteralNames = []string{
		"", "", "'@'", "", "'|'", "':'", "','", "'#'", "'/'",
	}
	staticData.SymbolicNames = []string{
		"", "SEPARATOR", "AT", "INCLUDE", "VERTLINE", "COLON", "COMMA", "NUMBER_SIGN",
		"SLASH", "STRING",
	}
	staticData.RuleNames = []string{
		"lineStatement", "entry", "separator", "key", "values", "value", "user",
		"file", "command", "include", "comment", "email", "error", "errorStatus",
		"errorCode", "errorMessage",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 9, 128, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		1, 0, 1, 0, 3, 0, 35, 8, 0, 1, 0, 3, 0, 38, 8, 0, 1, 0, 1, 0, 1, 1, 3,
		1, 43, 8, 1, 1, 1, 1, 1, 3, 1, 47, 8, 1, 1, 1, 1, 1, 3, 1, 51, 8, 1, 1,
		1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 5, 4, 63, 8, 4,
		10, 4, 12, 4, 66, 9, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5,
		75, 8, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 5, 7, 82, 8, 7, 10, 7, 12, 7, 85,
		9, 7, 1, 7, 3, 7, 88, 8, 7, 1, 8, 1, 8, 3, 8, 92, 8, 8, 1, 9, 1, 9, 1,
		9, 1, 9, 3, 9, 98, 8, 9, 1, 10, 1, 10, 3, 10, 102, 8, 10, 1, 10, 4, 10,
		105, 8, 10, 11, 10, 12, 10, 106, 1, 10, 3, 10, 110, 8, 10, 1, 11, 1, 11,
		1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1,
		14, 1, 14, 1, 15, 1, 15, 1, 15, 0, 0, 16, 0, 2, 4, 6, 8, 10, 12, 14, 16,
		18, 20, 22, 24, 26, 28, 30, 0, 0, 128, 0, 32, 1, 0, 0, 0, 2, 42, 1, 0,
		0, 0, 4, 54, 1, 0, 0, 0, 6, 56, 1, 0, 0, 0, 8, 64, 1, 0, 0, 0, 10, 74,
		1, 0, 0, 0, 12, 76, 1, 0, 0, 0, 14, 78, 1, 0, 0, 0, 16, 89, 1, 0, 0, 0,
		18, 93, 1, 0, 0, 0, 20, 99, 1, 0, 0, 0, 22, 111, 1, 0, 0, 0, 24, 115, 1,
		0, 0, 0, 26, 121, 1, 0, 0, 0, 28, 123, 1, 0, 0, 0, 30, 125, 1, 0, 0, 0,
		32, 34, 3, 2, 1, 0, 33, 35, 5, 1, 0, 0, 34, 33, 1, 0, 0, 0, 34, 35, 1,
		0, 0, 0, 35, 37, 1, 0, 0, 0, 36, 38, 3, 20, 10, 0, 37, 36, 1, 0, 0, 0,
		37, 38, 1, 0, 0, 0, 38, 39, 1, 0, 0, 0, 39, 40, 5, 0, 0, 1, 40, 1, 1, 0,
		0, 0, 41, 43, 5, 1, 0, 0, 42, 41, 1, 0, 0, 0, 42, 43, 1, 0, 0, 0, 43, 44,
		1, 0, 0, 0, 44, 46, 3, 6, 3, 0, 45, 47, 5, 1, 0, 0, 46, 45, 1, 0, 0, 0,
		46, 47, 1, 0, 0, 0, 47, 48, 1, 0, 0, 0, 48, 50, 3, 4, 2, 0, 49, 51, 5,
		1, 0, 0, 50, 49, 1, 0, 0, 0, 50, 51, 1, 0, 0, 0, 51, 52, 1, 0, 0, 0, 52,
		53, 3, 8, 4, 0, 53, 3, 1, 0, 0, 0, 54, 55, 5, 5, 0, 0, 55, 5, 1, 0, 0,
		0, 56, 57, 5, 9, 0, 0, 57, 7, 1, 0, 0, 0, 58, 59, 3, 10, 5, 0, 59, 60,
		5, 6, 0, 0, 60, 61, 5, 1, 0, 0, 61, 63, 1, 0, 0, 0, 62, 58, 1, 0, 0, 0,
		63, 66, 1, 0, 0, 0, 64, 62, 1, 0, 0, 0, 64, 65, 1, 0, 0, 0, 65, 67, 1,
		0, 0, 0, 66, 64, 1, 0, 0, 0, 67, 68, 3, 10, 5, 0, 68, 9, 1, 0, 0, 0, 69,
		75, 3, 12, 6, 0, 70, 75, 3, 14, 7, 0, 71, 75, 3, 16, 8, 0, 72, 75, 3, 18,
		9, 0, 73, 75, 3, 22, 11, 0, 74, 69, 1, 0, 0, 0, 74, 70, 1, 0, 0, 0, 74,
		71, 1, 0, 0, 0, 74, 72, 1, 0, 0, 0, 74, 73, 1, 0, 0, 0, 75, 11, 1, 0, 0,
		0, 76, 77, 5, 9, 0, 0, 77, 13, 1, 0, 0, 0, 78, 83, 5, 8, 0, 0, 79, 80,
		5, 9, 0, 0, 80, 82, 5, 8, 0, 0, 81, 79, 1, 0, 0, 0, 82, 85, 1, 0, 0, 0,
		83, 81, 1, 0, 0, 0, 83, 84, 1, 0, 0, 0, 84, 87, 1, 0, 0, 0, 85, 83, 1,
		0, 0, 0, 86, 88, 5, 9, 0, 0, 87, 86, 1, 0, 0, 0, 87, 88, 1, 0, 0, 0, 88,
		15, 1, 0, 0, 0, 89, 91, 5, 4, 0, 0, 90, 92, 5, 9, 0, 0, 91, 90, 1, 0, 0,
		0, 91, 92, 1, 0, 0, 0, 92, 17, 1, 0, 0, 0, 93, 94, 5, 5, 0, 0, 94, 95,
		5, 3, 0, 0, 95, 97, 5, 5, 0, 0, 96, 98, 3, 14, 7, 0, 97, 96, 1, 0, 0, 0,
		97, 98, 1, 0, 0, 0, 98, 19, 1, 0, 0, 0, 99, 104, 5, 7, 0, 0, 100, 102,
		5, 1, 0, 0, 101, 100, 1, 0, 0, 0, 101, 102, 1, 0, 0, 0, 102, 103, 1, 0,
		0, 0, 103, 105, 5, 9, 0, 0, 104, 101, 1, 0, 0, 0, 105, 106, 1, 0, 0, 0,
		106, 104, 1, 0, 0, 0, 106, 107, 1, 0, 0, 0, 107, 109, 1, 0, 0, 0, 108,
		110, 5, 1, 0, 0, 109, 108, 1, 0, 0, 0, 109, 110, 1, 0, 0, 0, 110, 21, 1,
		0, 0, 0, 111, 112, 5, 9, 0, 0, 112, 113, 5, 2, 0, 0, 113, 114, 5, 9, 0,
		0, 114, 23, 1, 0, 0, 0, 115, 116, 3, 26, 13, 0, 116, 117, 5, 5, 0, 0, 117,
		118, 3, 28, 14, 0, 118, 119, 5, 1, 0, 0, 119, 120, 3, 30, 15, 0, 120, 25,
		1, 0, 0, 0, 121, 122, 5, 9, 0, 0, 122, 27, 1, 0, 0, 0, 123, 124, 5, 9,
		0, 0, 124, 29, 1, 0, 0, 0, 125, 126, 5, 9, 0, 0, 126, 31, 1, 0, 0, 0, 14,
		34, 37, 42, 46, 50, 64, 74, 83, 87, 91, 97, 101, 106, 109,
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

// AliasesParserInit initializes any static state used to implement AliasesParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewAliasesParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func AliasesParserInit() {
	staticData := &AliasesParserStaticData
	staticData.once.Do(aliasesParserInit)
}

// NewAliasesParser produces a new parser instance for the optional input antlr.TokenStream.
func NewAliasesParser(input antlr.TokenStream) *AliasesParser {
	AliasesParserInit()
	this := new(AliasesParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &AliasesParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Aliases.g4"

	return this
}

// AliasesParser tokens.
const (
	AliasesParserEOF         = antlr.TokenEOF
	AliasesParserSEPARATOR   = 1
	AliasesParserAT          = 2
	AliasesParserINCLUDE     = 3
	AliasesParserVERTLINE    = 4
	AliasesParserCOLON       = 5
	AliasesParserCOMMA       = 6
	AliasesParserNUMBER_SIGN = 7
	AliasesParserSLASH       = 8
	AliasesParserSTRING      = 9
)

// AliasesParser rules.
const (
	AliasesParserRULE_lineStatement = 0
	AliasesParserRULE_entry         = 1
	AliasesParserRULE_separator     = 2
	AliasesParserRULE_key           = 3
	AliasesParserRULE_values        = 4
	AliasesParserRULE_value         = 5
	AliasesParserRULE_user          = 6
	AliasesParserRULE_file          = 7
	AliasesParserRULE_command       = 8
	AliasesParserRULE_include       = 9
	AliasesParserRULE_comment       = 10
	AliasesParserRULE_email         = 11
	AliasesParserRULE_error         = 12
	AliasesParserRULE_errorStatus   = 13
	AliasesParserRULE_errorCode     = 14
	AliasesParserRULE_errorMessage  = 15
)

// ILineStatementContext is an interface to support dynamic dispatch.
type ILineStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Entry() IEntryContext
	EOF() antlr.TerminalNode
	SEPARATOR() antlr.TerminalNode
	Comment() ICommentContext

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
	p.RuleIndex = AliasesParserRULE_lineStatement
	return p
}

func InitEmptyLineStatementContext(p *LineStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AliasesParserRULE_lineStatement
}

func (*LineStatementContext) IsLineStatementContext() {}

func NewLineStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LineStatementContext {
	var p = new(LineStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AliasesParserRULE_lineStatement

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
	return s.GetToken(AliasesParserEOF, 0)
}

func (s *LineStatementContext) SEPARATOR() antlr.TerminalNode {
	return s.GetToken(AliasesParserSEPARATOR, 0)
}

func (s *LineStatementContext) Comment() ICommentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICommentContext)
}

func (s *LineStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LineStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LineStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AliasesListener); ok {
		listenerT.EnterLineStatement(s)
	}
}

func (s *LineStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AliasesListener); ok {
		listenerT.ExitLineStatement(s)
	}
}

func (p *AliasesParser) LineStatement() (localctx ILineStatementContext) {
	localctx = NewLineStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, AliasesParserRULE_lineStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(32)
		p.Entry()
	}
	p.SetState(34)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == AliasesParserSEPARATOR {
		{
			p.SetState(33)
			p.Match(AliasesParserSEPARATOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(37)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == AliasesParserNUMBER_SIGN {
		{
			p.SetState(36)
			p.Comment()
		}

	}
	{
		p.SetState(39)
		p.Match(AliasesParserEOF)
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
	Separator() ISeparatorContext
	Values() IValuesContext
	AllSEPARATOR() []antlr.TerminalNode
	SEPARATOR(i int) antlr.TerminalNode

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
	p.RuleIndex = AliasesParserRULE_entry
	return p
}

func InitEmptyEntryContext(p *EntryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AliasesParserRULE_entry
}

func (*EntryContext) IsEntryContext() {}

func NewEntryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EntryContext {
	var p = new(EntryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AliasesParserRULE_entry

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

func (s *EntryContext) Values() IValuesContext {
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

func (s *EntryContext) AllSEPARATOR() []antlr.TerminalNode {
	return s.GetTokens(AliasesParserSEPARATOR)
}

func (s *EntryContext) SEPARATOR(i int) antlr.TerminalNode {
	return s.GetToken(AliasesParserSEPARATOR, i)
}

func (s *EntryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EntryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EntryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AliasesListener); ok {
		listenerT.EnterEntry(s)
	}
}

func (s *EntryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AliasesListener); ok {
		listenerT.ExitEntry(s)
	}
}

func (p *AliasesParser) Entry() (localctx IEntryContext) {
	localctx = NewEntryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, AliasesParserRULE_entry)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(42)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == AliasesParserSEPARATOR {
		{
			p.SetState(41)
			p.Match(AliasesParserSEPARATOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(44)
		p.Key()
	}
	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == AliasesParserSEPARATOR {
		{
			p.SetState(45)
			p.Match(AliasesParserSEPARATOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(48)
		p.Separator()
	}
	p.SetState(50)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == AliasesParserSEPARATOR {
		{
			p.SetState(49)
			p.Match(AliasesParserSEPARATOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(52)
		p.Values()
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
	COLON() antlr.TerminalNode

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
	p.RuleIndex = AliasesParserRULE_separator
	return p
}

func InitEmptySeparatorContext(p *SeparatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AliasesParserRULE_separator
}

func (*SeparatorContext) IsSeparatorContext() {}

func NewSeparatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SeparatorContext {
	var p = new(SeparatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AliasesParserRULE_separator

	return p
}

func (s *SeparatorContext) GetParser() antlr.Parser { return s.parser }

func (s *SeparatorContext) COLON() antlr.TerminalNode {
	return s.GetToken(AliasesParserCOLON, 0)
}

func (s *SeparatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SeparatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SeparatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AliasesListener); ok {
		listenerT.EnterSeparator(s)
	}
}

func (s *SeparatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AliasesListener); ok {
		listenerT.ExitSeparator(s)
	}
}

func (p *AliasesParser) Separator() (localctx ISeparatorContext) {
	localctx = NewSeparatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, AliasesParserRULE_separator)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(54)
		p.Match(AliasesParserCOLON)
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
	STRING() antlr.TerminalNode

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
	p.RuleIndex = AliasesParserRULE_key
	return p
}

func InitEmptyKeyContext(p *KeyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AliasesParserRULE_key
}

func (*KeyContext) IsKeyContext() {}

func NewKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyContext {
	var p = new(KeyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AliasesParserRULE_key

	return p
}

func (s *KeyContext) GetParser() antlr.Parser { return s.parser }

func (s *KeyContext) STRING() antlr.TerminalNode {
	return s.GetToken(AliasesParserSTRING, 0)
}

func (s *KeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AliasesListener); ok {
		listenerT.EnterKey(s)
	}
}

func (s *KeyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AliasesListener); ok {
		listenerT.ExitKey(s)
	}
}

func (p *AliasesParser) Key() (localctx IKeyContext) {
	localctx = NewKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, AliasesParserRULE_key)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(56)
		p.Match(AliasesParserSTRING)
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
	AllSEPARATOR() []antlr.TerminalNode
	SEPARATOR(i int) antlr.TerminalNode

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
	p.RuleIndex = AliasesParserRULE_values
	return p
}

func InitEmptyValuesContext(p *ValuesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AliasesParserRULE_values
}

func (*ValuesContext) IsValuesContext() {}

func NewValuesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValuesContext {
	var p = new(ValuesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AliasesParserRULE_values

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
	return s.GetTokens(AliasesParserCOMMA)
}

func (s *ValuesContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(AliasesParserCOMMA, i)
}

func (s *ValuesContext) AllSEPARATOR() []antlr.TerminalNode {
	return s.GetTokens(AliasesParserSEPARATOR)
}

func (s *ValuesContext) SEPARATOR(i int) antlr.TerminalNode {
	return s.GetToken(AliasesParserSEPARATOR, i)
}

func (s *ValuesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValuesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValuesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AliasesListener); ok {
		listenerT.EnterValues(s)
	}
}

func (s *ValuesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AliasesListener); ok {
		listenerT.ExitValues(s)
	}
}

func (p *AliasesParser) Values() (localctx IValuesContext) {
	localctx = NewValuesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, AliasesParserRULE_values)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(58)
				p.Value()
			}
			{
				p.SetState(59)
				p.Match(AliasesParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(60)
				p.Match(AliasesParserSEPARATOR)
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
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	{
		p.SetState(67)
		p.Value()
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
	User() IUserContext
	File() IFileContext
	Command() ICommandContext
	Include() IIncludeContext
	Email() IEmailContext

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
	p.RuleIndex = AliasesParserRULE_value
	return p
}

func InitEmptyValueContext(p *ValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AliasesParserRULE_value
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AliasesParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) User() IUserContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUserContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUserContext)
}

func (s *ValueContext) File() IFileContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFileContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFileContext)
}

func (s *ValueContext) Command() ICommandContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommandContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICommandContext)
}

func (s *ValueContext) Include() IIncludeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIncludeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIncludeContext)
}

func (s *ValueContext) Email() IEmailContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEmailContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEmailContext)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AliasesListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AliasesListener); ok {
		listenerT.ExitValue(s)
	}
}

func (p *AliasesParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, AliasesParserRULE_value)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(69)
			p.User()
		}

	case 2:
		{
			p.SetState(70)
			p.File()
		}

	case 3:
		{
			p.SetState(71)
			p.Command()
		}

	case 4:
		{
			p.SetState(72)
			p.Include()
		}

	case 5:
		{
			p.SetState(73)
			p.Email()
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

// IUserContext is an interface to support dynamic dispatch.
type IUserContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode

	// IsUserContext differentiates from other interfaces.
	IsUserContext()
}

type UserContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUserContext() *UserContext {
	var p = new(UserContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AliasesParserRULE_user
	return p
}

func InitEmptyUserContext(p *UserContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AliasesParserRULE_user
}

func (*UserContext) IsUserContext() {}

func NewUserContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UserContext {
	var p = new(UserContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AliasesParserRULE_user

	return p
}

func (s *UserContext) GetParser() antlr.Parser { return s.parser }

func (s *UserContext) STRING() antlr.TerminalNode {
	return s.GetToken(AliasesParserSTRING, 0)
}

func (s *UserContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UserContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UserContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AliasesListener); ok {
		listenerT.EnterUser(s)
	}
}

func (s *UserContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AliasesListener); ok {
		listenerT.ExitUser(s)
	}
}

func (p *AliasesParser) User() (localctx IUserContext) {
	localctx = NewUserContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, AliasesParserRULE_user)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(76)
		p.Match(AliasesParserSTRING)
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

// IFileContext is an interface to support dynamic dispatch.
type IFileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSLASH() []antlr.TerminalNode
	SLASH(i int) antlr.TerminalNode
	AllSTRING() []antlr.TerminalNode
	STRING(i int) antlr.TerminalNode

	// IsFileContext differentiates from other interfaces.
	IsFileContext()
}

type FileContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFileContext() *FileContext {
	var p = new(FileContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AliasesParserRULE_file
	return p
}

func InitEmptyFileContext(p *FileContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AliasesParserRULE_file
}

func (*FileContext) IsFileContext() {}

func NewFileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FileContext {
	var p = new(FileContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AliasesParserRULE_file

	return p
}

func (s *FileContext) GetParser() antlr.Parser { return s.parser }

func (s *FileContext) AllSLASH() []antlr.TerminalNode {
	return s.GetTokens(AliasesParserSLASH)
}

func (s *FileContext) SLASH(i int) antlr.TerminalNode {
	return s.GetToken(AliasesParserSLASH, i)
}

func (s *FileContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(AliasesParserSTRING)
}

func (s *FileContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(AliasesParserSTRING, i)
}

func (s *FileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AliasesListener); ok {
		listenerT.EnterFile(s)
	}
}

func (s *FileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AliasesListener); ok {
		listenerT.ExitFile(s)
	}
}

func (p *AliasesParser) File() (localctx IFileContext) {
	localctx = NewFileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, AliasesParserRULE_file)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(78)
		p.Match(AliasesParserSLASH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(83)
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
				p.SetState(79)
				p.Match(AliasesParserSTRING)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(80)
				p.Match(AliasesParserSLASH)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(85)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == AliasesParserSTRING {
		{
			p.SetState(86)
			p.Match(AliasesParserSTRING)
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

// ICommandContext is an interface to support dynamic dispatch.
type ICommandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VERTLINE() antlr.TerminalNode
	STRING() antlr.TerminalNode

	// IsCommandContext differentiates from other interfaces.
	IsCommandContext()
}

type CommandContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommandContext() *CommandContext {
	var p = new(CommandContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AliasesParserRULE_command
	return p
}

func InitEmptyCommandContext(p *CommandContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AliasesParserRULE_command
}

func (*CommandContext) IsCommandContext() {}

func NewCommandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommandContext {
	var p = new(CommandContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AliasesParserRULE_command

	return p
}

func (s *CommandContext) GetParser() antlr.Parser { return s.parser }

func (s *CommandContext) VERTLINE() antlr.TerminalNode {
	return s.GetToken(AliasesParserVERTLINE, 0)
}

func (s *CommandContext) STRING() antlr.TerminalNode {
	return s.GetToken(AliasesParserSTRING, 0)
}

func (s *CommandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AliasesListener); ok {
		listenerT.EnterCommand(s)
	}
}

func (s *CommandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AliasesListener); ok {
		listenerT.ExitCommand(s)
	}
}

func (p *AliasesParser) Command() (localctx ICommandContext) {
	localctx = NewCommandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, AliasesParserRULE_command)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(89)
		p.Match(AliasesParserVERTLINE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(91)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == AliasesParserSTRING {
		{
			p.SetState(90)
			p.Match(AliasesParserSTRING)
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

// IIncludeContext is an interface to support dynamic dispatch.
type IIncludeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllCOLON() []antlr.TerminalNode
	COLON(i int) antlr.TerminalNode
	INCLUDE() antlr.TerminalNode
	File() IFileContext

	// IsIncludeContext differentiates from other interfaces.
	IsIncludeContext()
}

type IncludeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIncludeContext() *IncludeContext {
	var p = new(IncludeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AliasesParserRULE_include
	return p
}

func InitEmptyIncludeContext(p *IncludeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AliasesParserRULE_include
}

func (*IncludeContext) IsIncludeContext() {}

func NewIncludeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IncludeContext {
	var p = new(IncludeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AliasesParserRULE_include

	return p
}

func (s *IncludeContext) GetParser() antlr.Parser { return s.parser }

func (s *IncludeContext) AllCOLON() []antlr.TerminalNode {
	return s.GetTokens(AliasesParserCOLON)
}

func (s *IncludeContext) COLON(i int) antlr.TerminalNode {
	return s.GetToken(AliasesParserCOLON, i)
}

func (s *IncludeContext) INCLUDE() antlr.TerminalNode {
	return s.GetToken(AliasesParserINCLUDE, 0)
}

func (s *IncludeContext) File() IFileContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFileContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFileContext)
}

func (s *IncludeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IncludeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IncludeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AliasesListener); ok {
		listenerT.EnterInclude(s)
	}
}

func (s *IncludeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AliasesListener); ok {
		listenerT.ExitInclude(s)
	}
}

func (p *AliasesParser) Include() (localctx IIncludeContext) {
	localctx = NewIncludeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, AliasesParserRULE_include)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(93)
		p.Match(AliasesParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(94)
		p.Match(AliasesParserINCLUDE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(95)
		p.Match(AliasesParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(97)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == AliasesParserSLASH {
		{
			p.SetState(96)
			p.File()
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

// ICommentContext is an interface to support dynamic dispatch.
type ICommentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMBER_SIGN() antlr.TerminalNode
	AllSTRING() []antlr.TerminalNode
	STRING(i int) antlr.TerminalNode
	AllSEPARATOR() []antlr.TerminalNode
	SEPARATOR(i int) antlr.TerminalNode

	// IsCommentContext differentiates from other interfaces.
	IsCommentContext()
}

type CommentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommentContext() *CommentContext {
	var p = new(CommentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AliasesParserRULE_comment
	return p
}

func InitEmptyCommentContext(p *CommentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AliasesParserRULE_comment
}

func (*CommentContext) IsCommentContext() {}

func NewCommentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommentContext {
	var p = new(CommentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AliasesParserRULE_comment

	return p
}

func (s *CommentContext) GetParser() antlr.Parser { return s.parser }

func (s *CommentContext) NUMBER_SIGN() antlr.TerminalNode {
	return s.GetToken(AliasesParserNUMBER_SIGN, 0)
}

func (s *CommentContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(AliasesParserSTRING)
}

func (s *CommentContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(AliasesParserSTRING, i)
}

func (s *CommentContext) AllSEPARATOR() []antlr.TerminalNode {
	return s.GetTokens(AliasesParserSEPARATOR)
}

func (s *CommentContext) SEPARATOR(i int) antlr.TerminalNode {
	return s.GetToken(AliasesParserSEPARATOR, i)
}

func (s *CommentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AliasesListener); ok {
		listenerT.EnterComment(s)
	}
}

func (s *CommentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AliasesListener); ok {
		listenerT.ExitComment(s)
	}
}

func (p *AliasesParser) Comment() (localctx ICommentContext) {
	localctx = NewCommentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, AliasesParserRULE_comment)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(99)
		p.Match(AliasesParserNUMBER_SIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(104)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			p.SetState(101)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == AliasesParserSEPARATOR {
				{
					p.SetState(100)
					p.Match(AliasesParserSEPARATOR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			{
				p.SetState(103)
				p.Match(AliasesParserSTRING)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(106)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(109)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == AliasesParserSEPARATOR {
		{
			p.SetState(108)
			p.Match(AliasesParserSEPARATOR)
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

// IEmailContext is an interface to support dynamic dispatch.
type IEmailContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSTRING() []antlr.TerminalNode
	STRING(i int) antlr.TerminalNode
	AT() antlr.TerminalNode

	// IsEmailContext differentiates from other interfaces.
	IsEmailContext()
}

type EmailContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEmailContext() *EmailContext {
	var p = new(EmailContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AliasesParserRULE_email
	return p
}

func InitEmptyEmailContext(p *EmailContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AliasesParserRULE_email
}

func (*EmailContext) IsEmailContext() {}

func NewEmailContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EmailContext {
	var p = new(EmailContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AliasesParserRULE_email

	return p
}

func (s *EmailContext) GetParser() antlr.Parser { return s.parser }

func (s *EmailContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(AliasesParserSTRING)
}

func (s *EmailContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(AliasesParserSTRING, i)
}

func (s *EmailContext) AT() antlr.TerminalNode {
	return s.GetToken(AliasesParserAT, 0)
}

func (s *EmailContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EmailContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EmailContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AliasesListener); ok {
		listenerT.EnterEmail(s)
	}
}

func (s *EmailContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AliasesListener); ok {
		listenerT.ExitEmail(s)
	}
}

func (p *AliasesParser) Email() (localctx IEmailContext) {
	localctx = NewEmailContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, AliasesParserRULE_email)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(111)
		p.Match(AliasesParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(112)
		p.Match(AliasesParserAT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(113)
		p.Match(AliasesParserSTRING)
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

// IErrorContext is an interface to support dynamic dispatch.
type IErrorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ErrorStatus() IErrorStatusContext
	COLON() antlr.TerminalNode
	ErrorCode() IErrorCodeContext
	SEPARATOR() antlr.TerminalNode
	ErrorMessage() IErrorMessageContext

	// IsErrorContext differentiates from other interfaces.
	IsErrorContext()
}

type ErrorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyErrorContext() *ErrorContext {
	var p = new(ErrorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AliasesParserRULE_error
	return p
}

func InitEmptyErrorContext(p *ErrorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AliasesParserRULE_error
}

func (*ErrorContext) IsErrorContext() {}

func NewErrorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ErrorContext {
	var p = new(ErrorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AliasesParserRULE_error

	return p
}

func (s *ErrorContext) GetParser() antlr.Parser { return s.parser }

func (s *ErrorContext) ErrorStatus() IErrorStatusContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IErrorStatusContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IErrorStatusContext)
}

func (s *ErrorContext) COLON() antlr.TerminalNode {
	return s.GetToken(AliasesParserCOLON, 0)
}

func (s *ErrorContext) ErrorCode() IErrorCodeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IErrorCodeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IErrorCodeContext)
}

func (s *ErrorContext) SEPARATOR() antlr.TerminalNode {
	return s.GetToken(AliasesParserSEPARATOR, 0)
}

func (s *ErrorContext) ErrorMessage() IErrorMessageContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IErrorMessageContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IErrorMessageContext)
}

func (s *ErrorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ErrorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ErrorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AliasesListener); ok {
		listenerT.EnterError(s)
	}
}

func (s *ErrorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AliasesListener); ok {
		listenerT.ExitError(s)
	}
}

func (p *AliasesParser) Error_() (localctx IErrorContext) {
	localctx = NewErrorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, AliasesParserRULE_error)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(115)
		p.ErrorStatus()
	}
	{
		p.SetState(116)
		p.Match(AliasesParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(117)
		p.ErrorCode()
	}
	{
		p.SetState(118)
		p.Match(AliasesParserSEPARATOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(119)
		p.ErrorMessage()
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

// IErrorStatusContext is an interface to support dynamic dispatch.
type IErrorStatusContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode

	// IsErrorStatusContext differentiates from other interfaces.
	IsErrorStatusContext()
}

type ErrorStatusContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyErrorStatusContext() *ErrorStatusContext {
	var p = new(ErrorStatusContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AliasesParserRULE_errorStatus
	return p
}

func InitEmptyErrorStatusContext(p *ErrorStatusContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AliasesParserRULE_errorStatus
}

func (*ErrorStatusContext) IsErrorStatusContext() {}

func NewErrorStatusContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ErrorStatusContext {
	var p = new(ErrorStatusContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AliasesParserRULE_errorStatus

	return p
}

func (s *ErrorStatusContext) GetParser() antlr.Parser { return s.parser }

func (s *ErrorStatusContext) STRING() antlr.TerminalNode {
	return s.GetToken(AliasesParserSTRING, 0)
}

func (s *ErrorStatusContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ErrorStatusContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ErrorStatusContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AliasesListener); ok {
		listenerT.EnterErrorStatus(s)
	}
}

func (s *ErrorStatusContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AliasesListener); ok {
		listenerT.ExitErrorStatus(s)
	}
}

func (p *AliasesParser) ErrorStatus() (localctx IErrorStatusContext) {
	localctx = NewErrorStatusContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, AliasesParserRULE_errorStatus)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(121)
		p.Match(AliasesParserSTRING)
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

// IErrorCodeContext is an interface to support dynamic dispatch.
type IErrorCodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode

	// IsErrorCodeContext differentiates from other interfaces.
	IsErrorCodeContext()
}

type ErrorCodeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyErrorCodeContext() *ErrorCodeContext {
	var p = new(ErrorCodeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AliasesParserRULE_errorCode
	return p
}

func InitEmptyErrorCodeContext(p *ErrorCodeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AliasesParserRULE_errorCode
}

func (*ErrorCodeContext) IsErrorCodeContext() {}

func NewErrorCodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ErrorCodeContext {
	var p = new(ErrorCodeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AliasesParserRULE_errorCode

	return p
}

func (s *ErrorCodeContext) GetParser() antlr.Parser { return s.parser }

func (s *ErrorCodeContext) STRING() antlr.TerminalNode {
	return s.GetToken(AliasesParserSTRING, 0)
}

func (s *ErrorCodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ErrorCodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ErrorCodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AliasesListener); ok {
		listenerT.EnterErrorCode(s)
	}
}

func (s *ErrorCodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AliasesListener); ok {
		listenerT.ExitErrorCode(s)
	}
}

func (p *AliasesParser) ErrorCode() (localctx IErrorCodeContext) {
	localctx = NewErrorCodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, AliasesParserRULE_errorCode)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(123)
		p.Match(AliasesParserSTRING)
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

// IErrorMessageContext is an interface to support dynamic dispatch.
type IErrorMessageContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode

	// IsErrorMessageContext differentiates from other interfaces.
	IsErrorMessageContext()
}

type ErrorMessageContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyErrorMessageContext() *ErrorMessageContext {
	var p = new(ErrorMessageContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AliasesParserRULE_errorMessage
	return p
}

func InitEmptyErrorMessageContext(p *ErrorMessageContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = AliasesParserRULE_errorMessage
}

func (*ErrorMessageContext) IsErrorMessageContext() {}

func NewErrorMessageContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ErrorMessageContext {
	var p = new(ErrorMessageContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = AliasesParserRULE_errorMessage

	return p
}

func (s *ErrorMessageContext) GetParser() antlr.Parser { return s.parser }

func (s *ErrorMessageContext) STRING() antlr.TerminalNode {
	return s.GetToken(AliasesParserSTRING, 0)
}

func (s *ErrorMessageContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ErrorMessageContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ErrorMessageContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AliasesListener); ok {
		listenerT.EnterErrorMessage(s)
	}
}

func (s *ErrorMessageContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AliasesListener); ok {
		listenerT.ExitErrorMessage(s)
	}
}

func (p *AliasesParser) ErrorMessage() (localctx IErrorMessageContext) {
	localctx = NewErrorMessageContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, AliasesParserRULE_errorMessage)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(125)
		p.Match(AliasesParserSTRING)
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
