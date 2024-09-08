// Code generated from Hosts.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Hosts

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

type HostsParser struct {
	*antlr.BaseParser
}

var HostsParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func hostsParserInit() {
	staticData := &HostsParserStaticData
	staticData.LiteralNames = []string{
		"", "", "'/'", "'.'", "':'", "'#'",
	}
	staticData.SymbolicNames = []string{
		"", "COMMENTLINE", "SLASH", "DOT", "COLON", "HASHTAG", "SEPARATOR",
		"NEWLINE", "STRING",
	}
	staticData.RuleNames = []string{
		"lineStatement", "entry", "aliases", "alias", "hostname", "domain",
		"ipAddress", "ipv4Address", "ipv6Address", "ipRange", "ipRangeBits",
		"ipPort", "comment", "leadingComment",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 8, 131, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 1, 0, 3, 0, 30, 8, 0, 1,
		0, 1, 0, 3, 0, 34, 8, 0, 1, 0, 3, 0, 37, 8, 0, 1, 0, 1, 0, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 3, 1, 46, 8, 1, 1, 2, 1, 2, 3, 2, 50, 8, 2, 4, 2, 52,
		8, 2, 11, 2, 12, 2, 53, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 4, 5, 61, 8, 5, 11,
		5, 12, 5, 62, 1, 5, 1, 5, 5, 5, 67, 8, 5, 10, 5, 12, 5, 70, 9, 5, 5, 5,
		72, 8, 5, 10, 5, 12, 5, 75, 9, 5, 1, 6, 1, 6, 3, 6, 79, 8, 6, 1, 7, 1,
		7, 4, 7, 83, 8, 7, 11, 7, 12, 7, 84, 1, 7, 1, 7, 3, 7, 89, 8, 7, 1, 7,
		3, 7, 92, 8, 7, 3, 7, 94, 8, 7, 1, 8, 1, 8, 4, 8, 98, 8, 8, 11, 8, 12,
		8, 99, 1, 8, 1, 8, 3, 8, 104, 8, 8, 1, 8, 1, 8, 1, 8, 3, 8, 109, 8, 8,
		1, 8, 3, 8, 112, 8, 8, 1, 8, 3, 8, 115, 8, 8, 3, 8, 117, 8, 8, 1, 9, 1,
		9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13,
		1, 13, 0, 0, 14, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 0,
		0, 136, 0, 29, 1, 0, 0, 0, 2, 40, 1, 0, 0, 0, 4, 51, 1, 0, 0, 0, 6, 55,
		1, 0, 0, 0, 8, 57, 1, 0, 0, 0, 10, 60, 1, 0, 0, 0, 12, 78, 1, 0, 0, 0,
		14, 82, 1, 0, 0, 0, 16, 108, 1, 0, 0, 0, 18, 118, 1, 0, 0, 0, 20, 121,
		1, 0, 0, 0, 22, 123, 1, 0, 0, 0, 24, 126, 1, 0, 0, 0, 26, 128, 1, 0, 0,
		0, 28, 30, 5, 6, 0, 0, 29, 28, 1, 0, 0, 0, 29, 30, 1, 0, 0, 0, 30, 31,
		1, 0, 0, 0, 31, 33, 3, 2, 1, 0, 32, 34, 5, 6, 0, 0, 33, 32, 1, 0, 0, 0,
		33, 34, 1, 0, 0, 0, 34, 36, 1, 0, 0, 0, 35, 37, 3, 26, 13, 0, 36, 35, 1,
		0, 0, 0, 36, 37, 1, 0, 0, 0, 37, 38, 1, 0, 0, 0, 38, 39, 5, 0, 0, 1, 39,
		1, 1, 0, 0, 0, 40, 41, 3, 12, 6, 0, 41, 42, 5, 6, 0, 0, 42, 45, 3, 8, 4,
		0, 43, 44, 5, 6, 0, 0, 44, 46, 3, 4, 2, 0, 45, 43, 1, 0, 0, 0, 45, 46,
		1, 0, 0, 0, 46, 3, 1, 0, 0, 0, 47, 49, 3, 6, 3, 0, 48, 50, 5, 6, 0, 0,
		49, 48, 1, 0, 0, 0, 49, 50, 1, 0, 0, 0, 50, 52, 1, 0, 0, 0, 51, 47, 1,
		0, 0, 0, 52, 53, 1, 0, 0, 0, 53, 51, 1, 0, 0, 0, 53, 54, 1, 0, 0, 0, 54,
		5, 1, 0, 0, 0, 55, 56, 3, 10, 5, 0, 56, 7, 1, 0, 0, 0, 57, 58, 3, 10, 5,
		0, 58, 9, 1, 0, 0, 0, 59, 61, 5, 8, 0, 0, 60, 59, 1, 0, 0, 0, 61, 62, 1,
		0, 0, 0, 62, 60, 1, 0, 0, 0, 62, 63, 1, 0, 0, 0, 63, 73, 1, 0, 0, 0, 64,
		68, 5, 3, 0, 0, 65, 67, 5, 8, 0, 0, 66, 65, 1, 0, 0, 0, 67, 70, 1, 0, 0,
		0, 68, 66, 1, 0, 0, 0, 68, 69, 1, 0, 0, 0, 69, 72, 1, 0, 0, 0, 70, 68,
		1, 0, 0, 0, 71, 64, 1, 0, 0, 0, 72, 75, 1, 0, 0, 0, 73, 71, 1, 0, 0, 0,
		73, 74, 1, 0, 0, 0, 74, 11, 1, 0, 0, 0, 75, 73, 1, 0, 0, 0, 76, 79, 3,
		14, 7, 0, 77, 79, 3, 16, 8, 0, 78, 76, 1, 0, 0, 0, 78, 77, 1, 0, 0, 0,
		79, 13, 1, 0, 0, 0, 80, 81, 5, 8, 0, 0, 81, 83, 5, 3, 0, 0, 82, 80, 1,
		0, 0, 0, 83, 84, 1, 0, 0, 0, 84, 82, 1, 0, 0, 0, 84, 85, 1, 0, 0, 0, 85,
		86, 1, 0, 0, 0, 86, 93, 5, 8, 0, 0, 87, 89, 3, 18, 9, 0, 88, 87, 1, 0,
		0, 0, 88, 89, 1, 0, 0, 0, 89, 94, 1, 0, 0, 0, 90, 92, 3, 22, 11, 0, 91,
		90, 1, 0, 0, 0, 91, 92, 1, 0, 0, 0, 92, 94, 1, 0, 0, 0, 93, 88, 1, 0, 0,
		0, 93, 91, 1, 0, 0, 0, 94, 15, 1, 0, 0, 0, 95, 96, 5, 8, 0, 0, 96, 98,
		5, 4, 0, 0, 97, 95, 1, 0, 0, 0, 98, 99, 1, 0, 0, 0, 99, 97, 1, 0, 0, 0,
		99, 100, 1, 0, 0, 0, 100, 101, 1, 0, 0, 0, 101, 109, 5, 8, 0, 0, 102, 104,
		5, 8, 0, 0, 103, 102, 1, 0, 0, 0, 103, 104, 1, 0, 0, 0, 104, 105, 1, 0,
		0, 0, 105, 106, 5, 4, 0, 0, 106, 107, 5, 4, 0, 0, 107, 109, 5, 8, 0, 0,
		108, 97, 1, 0, 0, 0, 108, 103, 1, 0, 0, 0, 109, 116, 1, 0, 0, 0, 110, 112,
		3, 18, 9, 0, 111, 110, 1, 0, 0, 0, 111, 112, 1, 0, 0, 0, 112, 117, 1, 0,
		0, 0, 113, 115, 3, 22, 11, 0, 114, 113, 1, 0, 0, 0, 114, 115, 1, 0, 0,
		0, 115, 117, 1, 0, 0, 0, 116, 111, 1, 0, 0, 0, 116, 114, 1, 0, 0, 0, 117,
		17, 1, 0, 0, 0, 118, 119, 5, 2, 0, 0, 119, 120, 3, 20, 10, 0, 120, 19,
		1, 0, 0, 0, 121, 122, 5, 8, 0, 0, 122, 21, 1, 0, 0, 0, 123, 124, 5, 4,
		0, 0, 124, 125, 5, 8, 0, 0, 125, 23, 1, 0, 0, 0, 126, 127, 5, 1, 0, 0,
		127, 25, 1, 0, 0, 0, 128, 129, 5, 1, 0, 0, 129, 27, 1, 0, 0, 0, 20, 29,
		33, 36, 45, 49, 53, 62, 68, 73, 78, 84, 88, 91, 93, 99, 103, 108, 111,
		114, 116,
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

// HostsParserInit initializes any static state used to implement HostsParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewHostsParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func HostsParserInit() {
	staticData := &HostsParserStaticData
	staticData.once.Do(hostsParserInit)
}

// NewHostsParser produces a new parser instance for the optional input antlr.TokenStream.
func NewHostsParser(input antlr.TokenStream) *HostsParser {
	HostsParserInit()
	this := new(HostsParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &HostsParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Hosts.g4"

	return this
}

// HostsParser tokens.
const (
	HostsParserEOF         = antlr.TokenEOF
	HostsParserCOMMENTLINE = 1
	HostsParserSLASH       = 2
	HostsParserDOT         = 3
	HostsParserCOLON       = 4
	HostsParserHASHTAG     = 5
	HostsParserSEPARATOR   = 6
	HostsParserNEWLINE     = 7
	HostsParserSTRING      = 8
)

// HostsParser rules.
const (
	HostsParserRULE_lineStatement  = 0
	HostsParserRULE_entry          = 1
	HostsParserRULE_aliases        = 2
	HostsParserRULE_alias          = 3
	HostsParserRULE_hostname       = 4
	HostsParserRULE_domain         = 5
	HostsParserRULE_ipAddress      = 6
	HostsParserRULE_ipv4Address    = 7
	HostsParserRULE_ipv6Address    = 8
	HostsParserRULE_ipRange        = 9
	HostsParserRULE_ipRangeBits    = 10
	HostsParserRULE_ipPort         = 11
	HostsParserRULE_comment        = 12
	HostsParserRULE_leadingComment = 13
)

// ILineStatementContext is an interface to support dynamic dispatch.
type ILineStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Entry() IEntryContext
	EOF() antlr.TerminalNode
	AllSEPARATOR() []antlr.TerminalNode
	SEPARATOR(i int) antlr.TerminalNode
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
	p.RuleIndex = HostsParserRULE_lineStatement
	return p
}

func InitEmptyLineStatementContext(p *LineStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HostsParserRULE_lineStatement
}

func (*LineStatementContext) IsLineStatementContext() {}

func NewLineStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LineStatementContext {
	var p = new(LineStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = HostsParserRULE_lineStatement

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
	return s.GetToken(HostsParserEOF, 0)
}

func (s *LineStatementContext) AllSEPARATOR() []antlr.TerminalNode {
	return s.GetTokens(HostsParserSEPARATOR)
}

func (s *LineStatementContext) SEPARATOR(i int) antlr.TerminalNode {
	return s.GetToken(HostsParserSEPARATOR, i)
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
	if listenerT, ok := listener.(HostsListener); ok {
		listenerT.EnterLineStatement(s)
	}
}

func (s *LineStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HostsListener); ok {
		listenerT.ExitLineStatement(s)
	}
}

func (p *HostsParser) LineStatement() (localctx ILineStatementContext) {
	localctx = NewLineStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, HostsParserRULE_lineStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(29)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == HostsParserSEPARATOR {
		{
			p.SetState(28)
			p.Match(HostsParserSEPARATOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(31)
		p.Entry()
	}
	p.SetState(33)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == HostsParserSEPARATOR {
		{
			p.SetState(32)
			p.Match(HostsParserSEPARATOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(36)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == HostsParserCOMMENTLINE {
		{
			p.SetState(35)
			p.LeadingComment()
		}

	}
	{
		p.SetState(38)
		p.Match(HostsParserEOF)
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
	IpAddress() IIpAddressContext
	AllSEPARATOR() []antlr.TerminalNode
	SEPARATOR(i int) antlr.TerminalNode
	Hostname() IHostnameContext
	Aliases() IAliasesContext

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
	p.RuleIndex = HostsParserRULE_entry
	return p
}

func InitEmptyEntryContext(p *EntryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HostsParserRULE_entry
}

func (*EntryContext) IsEntryContext() {}

func NewEntryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EntryContext {
	var p = new(EntryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = HostsParserRULE_entry

	return p
}

func (s *EntryContext) GetParser() antlr.Parser { return s.parser }

func (s *EntryContext) IpAddress() IIpAddressContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIpAddressContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIpAddressContext)
}

func (s *EntryContext) AllSEPARATOR() []antlr.TerminalNode {
	return s.GetTokens(HostsParserSEPARATOR)
}

func (s *EntryContext) SEPARATOR(i int) antlr.TerminalNode {
	return s.GetToken(HostsParserSEPARATOR, i)
}

func (s *EntryContext) Hostname() IHostnameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHostnameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHostnameContext)
}

func (s *EntryContext) Aliases() IAliasesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAliasesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAliasesContext)
}

func (s *EntryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EntryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EntryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HostsListener); ok {
		listenerT.EnterEntry(s)
	}
}

func (s *EntryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HostsListener); ok {
		listenerT.ExitEntry(s)
	}
}

func (p *HostsParser) Entry() (localctx IEntryContext) {
	localctx = NewEntryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, HostsParserRULE_entry)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(40)
		p.IpAddress()
	}
	{
		p.SetState(41)
		p.Match(HostsParserSEPARATOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(42)
		p.Hostname()
	}
	p.SetState(45)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(43)
			p.Match(HostsParserSEPARATOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(44)
			p.Aliases()
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

// IAliasesContext is an interface to support dynamic dispatch.
type IAliasesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAlias() []IAliasContext
	Alias(i int) IAliasContext
	AllSEPARATOR() []antlr.TerminalNode
	SEPARATOR(i int) antlr.TerminalNode

	// IsAliasesContext differentiates from other interfaces.
	IsAliasesContext()
}

type AliasesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAliasesContext() *AliasesContext {
	var p = new(AliasesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HostsParserRULE_aliases
	return p
}

func InitEmptyAliasesContext(p *AliasesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HostsParserRULE_aliases
}

func (*AliasesContext) IsAliasesContext() {}

func NewAliasesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AliasesContext {
	var p = new(AliasesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = HostsParserRULE_aliases

	return p
}

func (s *AliasesContext) GetParser() antlr.Parser { return s.parser }

func (s *AliasesContext) AllAlias() []IAliasContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAliasContext); ok {
			len++
		}
	}

	tst := make([]IAliasContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAliasContext); ok {
			tst[i] = t.(IAliasContext)
			i++
		}
	}

	return tst
}

func (s *AliasesContext) Alias(i int) IAliasContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAliasContext); ok {
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

	return t.(IAliasContext)
}

func (s *AliasesContext) AllSEPARATOR() []antlr.TerminalNode {
	return s.GetTokens(HostsParserSEPARATOR)
}

func (s *AliasesContext) SEPARATOR(i int) antlr.TerminalNode {
	return s.GetToken(HostsParserSEPARATOR, i)
}

func (s *AliasesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AliasesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AliasesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HostsListener); ok {
		listenerT.EnterAliases(s)
	}
}

func (s *AliasesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HostsListener); ok {
		listenerT.ExitAliases(s)
	}
}

func (p *HostsParser) Aliases() (localctx IAliasesContext) {
	localctx = NewAliasesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, HostsParserRULE_aliases)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == HostsParserSTRING {
		{
			p.SetState(47)
			p.Alias()
		}
		p.SetState(49)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(48)
				p.Match(HostsParserSEPARATOR)
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

// IAliasContext is an interface to support dynamic dispatch.
type IAliasContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Domain() IDomainContext

	// IsAliasContext differentiates from other interfaces.
	IsAliasContext()
}

type AliasContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAliasContext() *AliasContext {
	var p = new(AliasContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HostsParserRULE_alias
	return p
}

func InitEmptyAliasContext(p *AliasContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HostsParserRULE_alias
}

func (*AliasContext) IsAliasContext() {}

func NewAliasContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AliasContext {
	var p = new(AliasContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = HostsParserRULE_alias

	return p
}

func (s *AliasContext) GetParser() antlr.Parser { return s.parser }

func (s *AliasContext) Domain() IDomainContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDomainContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDomainContext)
}

func (s *AliasContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AliasContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AliasContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HostsListener); ok {
		listenerT.EnterAlias(s)
	}
}

func (s *AliasContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HostsListener); ok {
		listenerT.ExitAlias(s)
	}
}

func (p *HostsParser) Alias() (localctx IAliasContext) {
	localctx = NewAliasContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, HostsParserRULE_alias)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(55)
		p.Domain()
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

// IHostnameContext is an interface to support dynamic dispatch.
type IHostnameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Domain() IDomainContext

	// IsHostnameContext differentiates from other interfaces.
	IsHostnameContext()
}

type HostnameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHostnameContext() *HostnameContext {
	var p = new(HostnameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HostsParserRULE_hostname
	return p
}

func InitEmptyHostnameContext(p *HostnameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HostsParserRULE_hostname
}

func (*HostnameContext) IsHostnameContext() {}

func NewHostnameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HostnameContext {
	var p = new(HostnameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = HostsParserRULE_hostname

	return p
}

func (s *HostnameContext) GetParser() antlr.Parser { return s.parser }

func (s *HostnameContext) Domain() IDomainContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDomainContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDomainContext)
}

func (s *HostnameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HostnameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HostnameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HostsListener); ok {
		listenerT.EnterHostname(s)
	}
}

func (s *HostnameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HostsListener); ok {
		listenerT.ExitHostname(s)
	}
}

func (p *HostsParser) Hostname() (localctx IHostnameContext) {
	localctx = NewHostnameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, HostsParserRULE_hostname)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(57)
		p.Domain()
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

// IDomainContext is an interface to support dynamic dispatch.
type IDomainContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSTRING() []antlr.TerminalNode
	STRING(i int) antlr.TerminalNode
	AllDOT() []antlr.TerminalNode
	DOT(i int) antlr.TerminalNode

	// IsDomainContext differentiates from other interfaces.
	IsDomainContext()
}

type DomainContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDomainContext() *DomainContext {
	var p = new(DomainContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HostsParserRULE_domain
	return p
}

func InitEmptyDomainContext(p *DomainContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HostsParserRULE_domain
}

func (*DomainContext) IsDomainContext() {}

func NewDomainContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DomainContext {
	var p = new(DomainContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = HostsParserRULE_domain

	return p
}

func (s *DomainContext) GetParser() antlr.Parser { return s.parser }

func (s *DomainContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(HostsParserSTRING)
}

func (s *DomainContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(HostsParserSTRING, i)
}

func (s *DomainContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(HostsParserDOT)
}

func (s *DomainContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(HostsParserDOT, i)
}

func (s *DomainContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DomainContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DomainContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HostsListener); ok {
		listenerT.EnterDomain(s)
	}
}

func (s *DomainContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HostsListener); ok {
		listenerT.ExitDomain(s)
	}
}

func (p *HostsParser) Domain() (localctx IDomainContext) {
	localctx = NewDomainContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, HostsParserRULE_domain)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(60)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(59)
				p.Match(HostsParserSTRING)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(62)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == HostsParserDOT {
		{
			p.SetState(64)
			p.Match(HostsParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(68)
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
					p.SetState(65)
					p.Match(HostsParserSTRING)
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
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

		p.SetState(75)
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

// IIpAddressContext is an interface to support dynamic dispatch.
type IIpAddressContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Ipv4Address() IIpv4AddressContext
	Ipv6Address() IIpv6AddressContext

	// IsIpAddressContext differentiates from other interfaces.
	IsIpAddressContext()
}

type IpAddressContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIpAddressContext() *IpAddressContext {
	var p = new(IpAddressContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HostsParserRULE_ipAddress
	return p
}

func InitEmptyIpAddressContext(p *IpAddressContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HostsParserRULE_ipAddress
}

func (*IpAddressContext) IsIpAddressContext() {}

func NewIpAddressContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IpAddressContext {
	var p = new(IpAddressContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = HostsParserRULE_ipAddress

	return p
}

func (s *IpAddressContext) GetParser() antlr.Parser { return s.parser }

func (s *IpAddressContext) Ipv4Address() IIpv4AddressContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIpv4AddressContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIpv4AddressContext)
}

func (s *IpAddressContext) Ipv6Address() IIpv6AddressContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIpv6AddressContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIpv6AddressContext)
}

func (s *IpAddressContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IpAddressContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IpAddressContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HostsListener); ok {
		listenerT.EnterIpAddress(s)
	}
}

func (s *IpAddressContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HostsListener); ok {
		listenerT.ExitIpAddress(s)
	}
}

func (p *HostsParser) IpAddress() (localctx IIpAddressContext) {
	localctx = NewIpAddressContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, HostsParserRULE_ipAddress)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(78)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(76)
			p.Ipv4Address()
		}

	case 2:
		{
			p.SetState(77)
			p.Ipv6Address()
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

// IIpv4AddressContext is an interface to support dynamic dispatch.
type IIpv4AddressContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSTRING() []antlr.TerminalNode
	STRING(i int) antlr.TerminalNode
	AllDOT() []antlr.TerminalNode
	DOT(i int) antlr.TerminalNode
	IpRange() IIpRangeContext
	IpPort() IIpPortContext

	// IsIpv4AddressContext differentiates from other interfaces.
	IsIpv4AddressContext()
}

type Ipv4AddressContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIpv4AddressContext() *Ipv4AddressContext {
	var p = new(Ipv4AddressContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HostsParserRULE_ipv4Address
	return p
}

func InitEmptyIpv4AddressContext(p *Ipv4AddressContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HostsParserRULE_ipv4Address
}

func (*Ipv4AddressContext) IsIpv4AddressContext() {}

func NewIpv4AddressContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ipv4AddressContext {
	var p = new(Ipv4AddressContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = HostsParserRULE_ipv4Address

	return p
}

func (s *Ipv4AddressContext) GetParser() antlr.Parser { return s.parser }

func (s *Ipv4AddressContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(HostsParserSTRING)
}

func (s *Ipv4AddressContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(HostsParserSTRING, i)
}

func (s *Ipv4AddressContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(HostsParserDOT)
}

func (s *Ipv4AddressContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(HostsParserDOT, i)
}

func (s *Ipv4AddressContext) IpRange() IIpRangeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIpRangeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIpRangeContext)
}

func (s *Ipv4AddressContext) IpPort() IIpPortContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIpPortContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIpPortContext)
}

func (s *Ipv4AddressContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ipv4AddressContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ipv4AddressContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HostsListener); ok {
		listenerT.EnterIpv4Address(s)
	}
}

func (s *Ipv4AddressContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HostsListener); ok {
		listenerT.ExitIpv4Address(s)
	}
}

func (p *HostsParser) Ipv4Address() (localctx IIpv4AddressContext) {
	localctx = NewIpv4AddressContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, HostsParserRULE_ipv4Address)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(80)
				p.Match(HostsParserSTRING)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(81)
				p.Match(HostsParserDOT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(84)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	{
		p.SetState(86)
		p.Match(HostsParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.SetState(88)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == HostsParserSLASH {
			{
				p.SetState(87)
				p.IpRange()
			}

		}

	case 2:
		p.SetState(91)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == HostsParserCOLON {
			{
				p.SetState(90)
				p.IpPort()
			}

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

// IIpv6AddressContext is an interface to support dynamic dispatch.
type IIpv6AddressContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSTRING() []antlr.TerminalNode
	STRING(i int) antlr.TerminalNode
	AllCOLON() []antlr.TerminalNode
	COLON(i int) antlr.TerminalNode
	IpRange() IIpRangeContext
	IpPort() IIpPortContext

	// IsIpv6AddressContext differentiates from other interfaces.
	IsIpv6AddressContext()
}

type Ipv6AddressContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIpv6AddressContext() *Ipv6AddressContext {
	var p = new(Ipv6AddressContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HostsParserRULE_ipv6Address
	return p
}

func InitEmptyIpv6AddressContext(p *Ipv6AddressContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HostsParserRULE_ipv6Address
}

func (*Ipv6AddressContext) IsIpv6AddressContext() {}

func NewIpv6AddressContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ipv6AddressContext {
	var p = new(Ipv6AddressContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = HostsParserRULE_ipv6Address

	return p
}

func (s *Ipv6AddressContext) GetParser() antlr.Parser { return s.parser }

func (s *Ipv6AddressContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(HostsParserSTRING)
}

func (s *Ipv6AddressContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(HostsParserSTRING, i)
}

func (s *Ipv6AddressContext) AllCOLON() []antlr.TerminalNode {
	return s.GetTokens(HostsParserCOLON)
}

func (s *Ipv6AddressContext) COLON(i int) antlr.TerminalNode {
	return s.GetToken(HostsParserCOLON, i)
}

func (s *Ipv6AddressContext) IpRange() IIpRangeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIpRangeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIpRangeContext)
}

func (s *Ipv6AddressContext) IpPort() IIpPortContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIpPortContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIpPortContext)
}

func (s *Ipv6AddressContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ipv6AddressContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ipv6AddressContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HostsListener); ok {
		listenerT.EnterIpv6Address(s)
	}
}

func (s *Ipv6AddressContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HostsListener); ok {
		listenerT.ExitIpv6Address(s)
	}
}

func (p *HostsParser) Ipv6Address() (localctx IIpv6AddressContext) {
	localctx = NewIpv6AddressContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, HostsParserRULE_ipv6Address)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(108)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.SetState(97)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(95)
					p.Match(HostsParserSTRING)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(96)
					p.Match(HostsParserCOLON)
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
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		{
			p.SetState(101)
			p.Match(HostsParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.SetState(103)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == HostsParserSTRING {
			{
				p.SetState(102)
				p.Match(HostsParserSTRING)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(105)
			p.Match(HostsParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(106)
			p.Match(HostsParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(107)
			p.Match(HostsParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.SetState(116)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		p.SetState(111)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == HostsParserSLASH {
			{
				p.SetState(110)
				p.IpRange()
			}

		}

	case 2:
		p.SetState(114)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == HostsParserCOLON {
			{
				p.SetState(113)
				p.IpPort()
			}

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

// IIpRangeContext is an interface to support dynamic dispatch.
type IIpRangeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SLASH() antlr.TerminalNode
	IpRangeBits() IIpRangeBitsContext

	// IsIpRangeContext differentiates from other interfaces.
	IsIpRangeContext()
}

type IpRangeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIpRangeContext() *IpRangeContext {
	var p = new(IpRangeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HostsParserRULE_ipRange
	return p
}

func InitEmptyIpRangeContext(p *IpRangeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HostsParserRULE_ipRange
}

func (*IpRangeContext) IsIpRangeContext() {}

func NewIpRangeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IpRangeContext {
	var p = new(IpRangeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = HostsParserRULE_ipRange

	return p
}

func (s *IpRangeContext) GetParser() antlr.Parser { return s.parser }

func (s *IpRangeContext) SLASH() antlr.TerminalNode {
	return s.GetToken(HostsParserSLASH, 0)
}

func (s *IpRangeContext) IpRangeBits() IIpRangeBitsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIpRangeBitsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIpRangeBitsContext)
}

func (s *IpRangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IpRangeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IpRangeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HostsListener); ok {
		listenerT.EnterIpRange(s)
	}
}

func (s *IpRangeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HostsListener); ok {
		listenerT.ExitIpRange(s)
	}
}

func (p *HostsParser) IpRange() (localctx IIpRangeContext) {
	localctx = NewIpRangeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, HostsParserRULE_ipRange)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(118)
		p.Match(HostsParserSLASH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(119)
		p.IpRangeBits()
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

// IIpRangeBitsContext is an interface to support dynamic dispatch.
type IIpRangeBitsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode

	// IsIpRangeBitsContext differentiates from other interfaces.
	IsIpRangeBitsContext()
}

type IpRangeBitsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIpRangeBitsContext() *IpRangeBitsContext {
	var p = new(IpRangeBitsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HostsParserRULE_ipRangeBits
	return p
}

func InitEmptyIpRangeBitsContext(p *IpRangeBitsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HostsParserRULE_ipRangeBits
}

func (*IpRangeBitsContext) IsIpRangeBitsContext() {}

func NewIpRangeBitsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IpRangeBitsContext {
	var p = new(IpRangeBitsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = HostsParserRULE_ipRangeBits

	return p
}

func (s *IpRangeBitsContext) GetParser() antlr.Parser { return s.parser }

func (s *IpRangeBitsContext) STRING() antlr.TerminalNode {
	return s.GetToken(HostsParserSTRING, 0)
}

func (s *IpRangeBitsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IpRangeBitsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IpRangeBitsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HostsListener); ok {
		listenerT.EnterIpRangeBits(s)
	}
}

func (s *IpRangeBitsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HostsListener); ok {
		listenerT.ExitIpRangeBits(s)
	}
}

func (p *HostsParser) IpRangeBits() (localctx IIpRangeBitsContext) {
	localctx = NewIpRangeBitsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, HostsParserRULE_ipRangeBits)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(121)
		p.Match(HostsParserSTRING)
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

// IIpPortContext is an interface to support dynamic dispatch.
type IIpPortContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	COLON() antlr.TerminalNode
	STRING() antlr.TerminalNode

	// IsIpPortContext differentiates from other interfaces.
	IsIpPortContext()
}

type IpPortContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIpPortContext() *IpPortContext {
	var p = new(IpPortContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HostsParserRULE_ipPort
	return p
}

func InitEmptyIpPortContext(p *IpPortContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HostsParserRULE_ipPort
}

func (*IpPortContext) IsIpPortContext() {}

func NewIpPortContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IpPortContext {
	var p = new(IpPortContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = HostsParserRULE_ipPort

	return p
}

func (s *IpPortContext) GetParser() antlr.Parser { return s.parser }

func (s *IpPortContext) COLON() antlr.TerminalNode {
	return s.GetToken(HostsParserCOLON, 0)
}

func (s *IpPortContext) STRING() antlr.TerminalNode {
	return s.GetToken(HostsParserSTRING, 0)
}

func (s *IpPortContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IpPortContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IpPortContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HostsListener); ok {
		listenerT.EnterIpPort(s)
	}
}

func (s *IpPortContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HostsListener); ok {
		listenerT.ExitIpPort(s)
	}
}

func (p *HostsParser) IpPort() (localctx IIpPortContext) {
	localctx = NewIpPortContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, HostsParserRULE_ipPort)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(123)
		p.Match(HostsParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(124)
		p.Match(HostsParserSTRING)
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

// ICommentContext is an interface to support dynamic dispatch.
type ICommentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	COMMENTLINE() antlr.TerminalNode

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
	p.RuleIndex = HostsParserRULE_comment
	return p
}

func InitEmptyCommentContext(p *CommentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HostsParserRULE_comment
}

func (*CommentContext) IsCommentContext() {}

func NewCommentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommentContext {
	var p = new(CommentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = HostsParserRULE_comment

	return p
}

func (s *CommentContext) GetParser() antlr.Parser { return s.parser }

func (s *CommentContext) COMMENTLINE() antlr.TerminalNode {
	return s.GetToken(HostsParserCOMMENTLINE, 0)
}

func (s *CommentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HostsListener); ok {
		listenerT.EnterComment(s)
	}
}

func (s *CommentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HostsListener); ok {
		listenerT.ExitComment(s)
	}
}

func (p *HostsParser) Comment() (localctx ICommentContext) {
	localctx = NewCommentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, HostsParserRULE_comment)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(126)
		p.Match(HostsParserCOMMENTLINE)
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

// ILeadingCommentContext is an interface to support dynamic dispatch.
type ILeadingCommentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	COMMENTLINE() antlr.TerminalNode

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
	p.RuleIndex = HostsParserRULE_leadingComment
	return p
}

func InitEmptyLeadingCommentContext(p *LeadingCommentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HostsParserRULE_leadingComment
}

func (*LeadingCommentContext) IsLeadingCommentContext() {}

func NewLeadingCommentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LeadingCommentContext {
	var p = new(LeadingCommentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = HostsParserRULE_leadingComment

	return p
}

func (s *LeadingCommentContext) GetParser() antlr.Parser { return s.parser }

func (s *LeadingCommentContext) COMMENTLINE() antlr.TerminalNode {
	return s.GetToken(HostsParserCOMMENTLINE, 0)
}

func (s *LeadingCommentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LeadingCommentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LeadingCommentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HostsListener); ok {
		listenerT.EnterLeadingComment(s)
	}
}

func (s *LeadingCommentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HostsListener); ok {
		listenerT.ExitLeadingComment(s)
	}
}

func (p *HostsParser) LeadingComment() (localctx ILeadingCommentContext) {
	localctx = NewLeadingCommentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, HostsParserRULE_leadingComment)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(128)
		p.Match(HostsParserCOMMENTLINE)
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
