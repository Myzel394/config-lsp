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
		"", "", "", "", "'@'", "", "'|'", "':'", "','", "'#'", "'/'",
	}
	staticData.SymbolicNames = []string{
		"", "DIGITS", "ERROR", "SEPARATOR", "AT", "INCLUDE", "VERTLINE", "COLON",
		"COMMA", "NUMBER_SIGN", "SLASH", "STRING",
	}
	staticData.RuleNames = []string{
		"lineStatement", "entry", "separator", "key", "values", "value", "user",
		"file", "command", "include", "comment", "email", "error", "errorCode",
		"errorMessage",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 11, 124, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 1, 0, 1, 0,
		1, 0, 1, 1, 3, 1, 35, 8, 1, 1, 1, 1, 1, 3, 1, 39, 8, 1, 1, 1, 1, 1, 3,
		1, 43, 8, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4,
		5, 4, 55, 8, 4, 10, 4, 12, 4, 58, 9, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 3, 5, 68, 8, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 5, 7, 75,
		8, 7, 10, 7, 12, 7, 78, 9, 7, 1, 7, 3, 7, 81, 8, 7, 1, 8, 1, 8, 3, 8, 85,
		8, 8, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 91, 8, 9, 1, 10, 1, 10, 3, 10, 95,
		8, 10, 1, 10, 4, 10, 98, 8, 10, 11, 10, 12, 10, 99, 1, 10, 3, 10, 103,
		8, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 3, 12, 112, 8,
		12, 1, 12, 3, 12, 115, 8, 12, 1, 12, 3, 12, 118, 8, 12, 1, 13, 1, 13, 1,
		14, 1, 14, 1, 14, 0, 0, 15, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22,
		24, 26, 28, 0, 0, 127, 0, 30, 1, 0, 0, 0, 2, 34, 1, 0, 0, 0, 4, 46, 1,
		0, 0, 0, 6, 48, 1, 0, 0, 0, 8, 56, 1, 0, 0, 0, 10, 67, 1, 0, 0, 0, 12,
		69, 1, 0, 0, 0, 14, 71, 1, 0, 0, 0, 16, 82, 1, 0, 0, 0, 18, 86, 1, 0, 0,
		0, 20, 92, 1, 0, 0, 0, 22, 104, 1, 0, 0, 0, 24, 108, 1, 0, 0, 0, 26, 119,
		1, 0, 0, 0, 28, 121, 1, 0, 0, 0, 30, 31, 3, 2, 1, 0, 31, 32, 5, 0, 0, 1,
		32, 1, 1, 0, 0, 0, 33, 35, 5, 3, 0, 0, 34, 33, 1, 0, 0, 0, 34, 35, 1, 0,
		0, 0, 35, 36, 1, 0, 0, 0, 36, 38, 3, 6, 3, 0, 37, 39, 5, 3, 0, 0, 38, 37,
		1, 0, 0, 0, 38, 39, 1, 0, 0, 0, 39, 40, 1, 0, 0, 0, 40, 42, 3, 4, 2, 0,
		41, 43, 5, 3, 0, 0, 42, 41, 1, 0, 0, 0, 42, 43, 1, 0, 0, 0, 43, 44, 1,
		0, 0, 0, 44, 45, 3, 8, 4, 0, 45, 3, 1, 0, 0, 0, 46, 47, 5, 7, 0, 0, 47,
		5, 1, 0, 0, 0, 48, 49, 5, 11, 0, 0, 49, 7, 1, 0, 0, 0, 50, 51, 3, 10, 5,
		0, 51, 52, 5, 8, 0, 0, 52, 53, 5, 3, 0, 0, 53, 55, 1, 0, 0, 0, 54, 50,
		1, 0, 0, 0, 55, 58, 1, 0, 0, 0, 56, 54, 1, 0, 0, 0, 56, 57, 1, 0, 0, 0,
		57, 59, 1, 0, 0, 0, 58, 56, 1, 0, 0, 0, 59, 60, 3, 10, 5, 0, 60, 9, 1,
		0, 0, 0, 61, 68, 3, 12, 6, 0, 62, 68, 3, 14, 7, 0, 63, 68, 3, 16, 8, 0,
		64, 68, 3, 18, 9, 0, 65, 68, 3, 22, 11, 0, 66, 68, 3, 24, 12, 0, 67, 61,
		1, 0, 0, 0, 67, 62, 1, 0, 0, 0, 67, 63, 1, 0, 0, 0, 67, 64, 1, 0, 0, 0,
		67, 65, 1, 0, 0, 0, 67, 66, 1, 0, 0, 0, 68, 11, 1, 0, 0, 0, 69, 70, 5,
		11, 0, 0, 70, 13, 1, 0, 0, 0, 71, 76, 5, 10, 0, 0, 72, 73, 5, 11, 0, 0,
		73, 75, 5, 10, 0, 0, 74, 72, 1, 0, 0, 0, 75, 78, 1, 0, 0, 0, 76, 74, 1,
		0, 0, 0, 76, 77, 1, 0, 0, 0, 77, 80, 1, 0, 0, 0, 78, 76, 1, 0, 0, 0, 79,
		81, 5, 11, 0, 0, 80, 79, 1, 0, 0, 0, 80, 81, 1, 0, 0, 0, 81, 15, 1, 0,
		0, 0, 82, 84, 5, 6, 0, 0, 83, 85, 5, 11, 0, 0, 84, 83, 1, 0, 0, 0, 84,
		85, 1, 0, 0, 0, 85, 17, 1, 0, 0, 0, 86, 87, 5, 7, 0, 0, 87, 88, 5, 5, 0,
		0, 88, 90, 5, 7, 0, 0, 89, 91, 3, 14, 7, 0, 90, 89, 1, 0, 0, 0, 90, 91,
		1, 0, 0, 0, 91, 19, 1, 0, 0, 0, 92, 97, 5, 9, 0, 0, 93, 95, 5, 3, 0, 0,
		94, 93, 1, 0, 0, 0, 94, 95, 1, 0, 0, 0, 95, 96, 1, 0, 0, 0, 96, 98, 5,
		11, 0, 0, 97, 94, 1, 0, 0, 0, 98, 99, 1, 0, 0, 0, 99, 97, 1, 0, 0, 0, 99,
		100, 1, 0, 0, 0, 100, 102, 1, 0, 0, 0, 101, 103, 5, 3, 0, 0, 102, 101,
		1, 0, 0, 0, 102, 103, 1, 0, 0, 0, 103, 21, 1, 0, 0, 0, 104, 105, 5, 11,
		0, 0, 105, 106, 5, 4, 0, 0, 106, 107, 5, 11, 0, 0, 107, 23, 1, 0, 0, 0,
		108, 109, 5, 2, 0, 0, 109, 111, 5, 7, 0, 0, 110, 112, 3, 26, 13, 0, 111,
		110, 1, 0, 0, 0, 111, 112, 1, 0, 0, 0, 112, 114, 1, 0, 0, 0, 113, 115,
		5, 3, 0, 0, 114, 113, 1, 0, 0, 0, 114, 115, 1, 0, 0, 0, 115, 117, 1, 0,
		0, 0, 116, 118, 3, 28, 14, 0, 117, 116, 1, 0, 0, 0, 117, 118, 1, 0, 0,
		0, 118, 25, 1, 0, 0, 0, 119, 120, 5, 1, 0, 0, 120, 27, 1, 0, 0, 0, 121,
		122, 5, 11, 0, 0, 122, 29, 1, 0, 0, 0, 15, 34, 38, 42, 56, 67, 76, 80,
		84, 90, 94, 99, 102, 111, 114, 117,
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
	AliasesParserDIGITS      = 1
	AliasesParserERROR       = 2
	AliasesParserSEPARATOR   = 3
	AliasesParserAT          = 4
	AliasesParserINCLUDE     = 5
	AliasesParserVERTLINE    = 6
	AliasesParserCOLON       = 7
	AliasesParserCOMMA       = 8
	AliasesParserNUMBER_SIGN = 9
	AliasesParserSLASH       = 10
	AliasesParserSTRING      = 11
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
	AliasesParserRULE_errorCode     = 13
	AliasesParserRULE_errorMessage  = 14
)

// ILineStatementContext is an interface to support dynamic dispatch.
type ILineStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Entry() IEntryContext
	EOF() antlr.TerminalNode

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
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(30)
		p.Entry()
	}
	{
		p.SetState(31)
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
	{
		p.SetState(36)
		p.Key()
	}
	p.SetState(38)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == AliasesParserSEPARATOR {
		{
			p.SetState(37)
			p.Match(AliasesParserSEPARATOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(40)
		p.Separator()
	}
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
		p.SetState(46)
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
		p.SetState(48)
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
	p.SetState(56)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(50)
				p.Value()
			}
			{
				p.SetState(51)
				p.Match(AliasesParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(52)
				p.Match(AliasesParserSEPARATOR)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(58)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	{
		p.SetState(59)
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
	Error_() IErrorContext

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

func (s *ValueContext) Error_() IErrorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IErrorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IErrorContext)
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
	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(61)
			p.User()
		}

	case 2:
		{
			p.SetState(62)
			p.File()
		}

	case 3:
		{
			p.SetState(63)
			p.Command()
		}

	case 4:
		{
			p.SetState(64)
			p.Include()
		}

	case 5:
		{
			p.SetState(65)
			p.Email()
		}

	case 6:
		{
			p.SetState(66)
			p.Error_()
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
		p.SetState(69)
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
		p.SetState(71)
		p.Match(AliasesParserSLASH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(76)
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
				p.SetState(72)
				p.Match(AliasesParserSTRING)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(73)
				p.Match(AliasesParserSLASH)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(78)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(80)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == AliasesParserSTRING {
		{
			p.SetState(79)
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
		p.SetState(82)
		p.Match(AliasesParserVERTLINE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == AliasesParserSTRING {
		{
			p.SetState(83)
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
		p.SetState(86)
		p.Match(AliasesParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(87)
		p.Match(AliasesParserINCLUDE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(88)
		p.Match(AliasesParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(90)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == AliasesParserSLASH {
		{
			p.SetState(89)
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
		p.SetState(92)
		p.Match(AliasesParserNUMBER_SIGN)
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
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			p.SetState(94)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == AliasesParserSEPARATOR {
				{
					p.SetState(93)
					p.Match(AliasesParserSEPARATOR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			{
				p.SetState(96)
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

		p.SetState(99)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == AliasesParserSEPARATOR {
		{
			p.SetState(101)
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
		p.SetState(104)
		p.Match(AliasesParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(105)
		p.Match(AliasesParserAT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(106)
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
	ERROR() antlr.TerminalNode
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

func (s *ErrorContext) ERROR() antlr.TerminalNode {
	return s.GetToken(AliasesParserERROR, 0)
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
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(108)
		p.Match(AliasesParserERROR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(109)
		p.Match(AliasesParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(111)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == AliasesParserDIGITS {
		{
			p.SetState(110)
			p.ErrorCode()
		}

	}
	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == AliasesParserSEPARATOR {
		{
			p.SetState(113)
			p.Match(AliasesParserSEPARATOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == AliasesParserSTRING {
		{
			p.SetState(116)
			p.ErrorMessage()
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
	DIGITS() antlr.TerminalNode

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

func (s *ErrorCodeContext) DIGITS() antlr.TerminalNode {
	return s.GetToken(AliasesParserDIGITS, 0)
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
	p.EnterRule(localctx, 26, AliasesParserRULE_errorCode)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(119)
		p.Match(AliasesParserDIGITS)
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
	p.EnterRule(localctx, 28, AliasesParserRULE_errorMessage)
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
