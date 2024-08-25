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
		"NEWLINE", "DIGITS", "OCTETS", "DOMAIN",
	}
	staticData.RuleNames = []string{
		"lineStatement", "entry", "aliases", "alias", "hostname", "domain",
		"ipAddress", "ipv4Address", "singleIPv4Address", "ipv6Address", "singleIPv6Address",
		"ipv4Digit", "ipv6Octet", "ipRange", "ipRangeBits", "comment", "leadingComment",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 10, 110, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 1, 0, 3, 0, 36, 8, 0, 1, 0, 1, 0, 3, 0, 40, 8, 0, 1, 0, 3,
		0, 43, 8, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 52, 8, 1,
		1, 2, 1, 2, 3, 2, 56, 8, 2, 4, 2, 58, 8, 2, 11, 2, 12, 2, 59, 1, 3, 1,
		3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 3, 6, 70, 8, 6, 1, 7, 1, 7, 3, 7,
		74, 8, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 3,
		9, 86, 8, 9, 1, 10, 1, 10, 1, 10, 4, 10, 91, 8, 10, 11, 10, 12, 10, 92,
		1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 14, 1,
		14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 0, 0, 17, 0, 2, 4, 6, 8, 10, 12,
		14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 0, 0, 102, 0, 35, 1, 0, 0, 0, 2,
		46, 1, 0, 0, 0, 4, 57, 1, 0, 0, 0, 6, 61, 1, 0, 0, 0, 8, 63, 1, 0, 0, 0,
		10, 65, 1, 0, 0, 0, 12, 69, 1, 0, 0, 0, 14, 71, 1, 0, 0, 0, 16, 75, 1,
		0, 0, 0, 18, 83, 1, 0, 0, 0, 20, 90, 1, 0, 0, 0, 22, 96, 1, 0, 0, 0, 24,
		98, 1, 0, 0, 0, 26, 100, 1, 0, 0, 0, 28, 103, 1, 0, 0, 0, 30, 105, 1, 0,
		0, 0, 32, 107, 1, 0, 0, 0, 34, 36, 5, 6, 0, 0, 35, 34, 1, 0, 0, 0, 35,
		36, 1, 0, 0, 0, 36, 37, 1, 0, 0, 0, 37, 39, 3, 2, 1, 0, 38, 40, 5, 6, 0,
		0, 39, 38, 1, 0, 0, 0, 39, 40, 1, 0, 0, 0, 40, 42, 1, 0, 0, 0, 41, 43,
		3, 32, 16, 0, 42, 41, 1, 0, 0, 0, 42, 43, 1, 0, 0, 0, 43, 44, 1, 0, 0,
		0, 44, 45, 5, 0, 0, 1, 45, 1, 1, 0, 0, 0, 46, 47, 3, 12, 6, 0, 47, 48,
		5, 6, 0, 0, 48, 51, 3, 8, 4, 0, 49, 50, 5, 6, 0, 0, 50, 52, 3, 4, 2, 0,
		51, 49, 1, 0, 0, 0, 51, 52, 1, 0, 0, 0, 52, 3, 1, 0, 0, 0, 53, 55, 3, 6,
		3, 0, 54, 56, 5, 6, 0, 0, 55, 54, 1, 0, 0, 0, 55, 56, 1, 0, 0, 0, 56, 58,
		1, 0, 0, 0, 57, 53, 1, 0, 0, 0, 58, 59, 1, 0, 0, 0, 59, 57, 1, 0, 0, 0,
		59, 60, 1, 0, 0, 0, 60, 5, 1, 0, 0, 0, 61, 62, 5, 10, 0, 0, 62, 7, 1, 0,
		0, 0, 63, 64, 3, 10, 5, 0, 64, 9, 1, 0, 0, 0, 65, 66, 5, 10, 0, 0, 66,
		11, 1, 0, 0, 0, 67, 70, 3, 14, 7, 0, 68, 70, 3, 18, 9, 0, 69, 67, 1, 0,
		0, 0, 69, 68, 1, 0, 0, 0, 70, 13, 1, 0, 0, 0, 71, 73, 3, 16, 8, 0, 72,
		74, 3, 26, 13, 0, 73, 72, 1, 0, 0, 0, 73, 74, 1, 0, 0, 0, 74, 15, 1, 0,
		0, 0, 75, 76, 3, 22, 11, 0, 76, 77, 5, 3, 0, 0, 77, 78, 3, 22, 11, 0, 78,
		79, 5, 3, 0, 0, 79, 80, 3, 22, 11, 0, 80, 81, 5, 3, 0, 0, 81, 82, 3, 22,
		11, 0, 82, 17, 1, 0, 0, 0, 83, 85, 3, 20, 10, 0, 84, 86, 3, 26, 13, 0,
		85, 84, 1, 0, 0, 0, 85, 86, 1, 0, 0, 0, 86, 19, 1, 0, 0, 0, 87, 88, 3,
		24, 12, 0, 88, 89, 5, 4, 0, 0, 89, 91, 1, 0, 0, 0, 90, 87, 1, 0, 0, 0,
		91, 92, 1, 0, 0, 0, 92, 90, 1, 0, 0, 0, 92, 93, 1, 0, 0, 0, 93, 94, 1,
		0, 0, 0, 94, 95, 3, 24, 12, 0, 95, 21, 1, 0, 0, 0, 96, 97, 5, 8, 0, 0,
		97, 23, 1, 0, 0, 0, 98, 99, 5, 9, 0, 0, 99, 25, 1, 0, 0, 0, 100, 101, 5,
		2, 0, 0, 101, 102, 3, 28, 14, 0, 102, 27, 1, 0, 0, 0, 103, 104, 5, 8, 0,
		0, 104, 29, 1, 0, 0, 0, 105, 106, 5, 1, 0, 0, 106, 31, 1, 0, 0, 0, 107,
		108, 5, 1, 0, 0, 108, 33, 1, 0, 0, 0, 10, 35, 39, 42, 51, 55, 59, 69, 73,
		85, 92,
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
	HostsParserDIGITS      = 8
	HostsParserOCTETS      = 9
	HostsParserDOMAIN      = 10
)

// HostsParser rules.
const (
	HostsParserRULE_lineStatement     = 0
	HostsParserRULE_entry             = 1
	HostsParserRULE_aliases           = 2
	HostsParserRULE_alias             = 3
	HostsParserRULE_hostname          = 4
	HostsParserRULE_domain            = 5
	HostsParserRULE_ipAddress         = 6
	HostsParserRULE_ipv4Address       = 7
	HostsParserRULE_singleIPv4Address = 8
	HostsParserRULE_ipv6Address       = 9
	HostsParserRULE_singleIPv6Address = 10
	HostsParserRULE_ipv4Digit         = 11
	HostsParserRULE_ipv6Octet         = 12
	HostsParserRULE_ipRange           = 13
	HostsParserRULE_ipRangeBits       = 14
	HostsParserRULE_comment           = 15
	HostsParserRULE_leadingComment    = 16
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
	p.SetState(35)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == HostsParserSEPARATOR {
		{
			p.SetState(34)
			p.Match(HostsParserSEPARATOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(37)
		p.Entry()
	}
	p.SetState(39)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == HostsParserSEPARATOR {
		{
			p.SetState(38)
			p.Match(HostsParserSEPARATOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(42)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == HostsParserCOMMENTLINE {
		{
			p.SetState(41)
			p.LeadingComment()
		}

	}
	{
		p.SetState(44)
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
		p.SetState(46)
		p.IpAddress()
	}
	{
		p.SetState(47)
		p.Match(HostsParserSEPARATOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(48)
		p.Hostname()
	}
	p.SetState(51)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(49)
			p.Match(HostsParserSEPARATOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(50)
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
	p.SetState(57)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == HostsParserDOMAIN {
		{
			p.SetState(53)
			p.Alias()
		}
		p.SetState(55)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(54)
				p.Match(HostsParserSEPARATOR)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

		p.SetState(59)
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
	DOMAIN() antlr.TerminalNode

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

func (s *AliasContext) DOMAIN() antlr.TerminalNode {
	return s.GetToken(HostsParserDOMAIN, 0)
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
		p.SetState(61)
		p.Match(HostsParserDOMAIN)
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
		p.SetState(63)
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
	DOMAIN() antlr.TerminalNode

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

func (s *DomainContext) DOMAIN() antlr.TerminalNode {
	return s.GetToken(HostsParserDOMAIN, 0)
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
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(65)
		p.Match(HostsParserDOMAIN)
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
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case HostsParserDIGITS:
		{
			p.SetState(67)
			p.Ipv4Address()
		}

	case HostsParserOCTETS:
		{
			p.SetState(68)
			p.Ipv6Address()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	SingleIPv4Address() ISingleIPv4AddressContext
	IpRange() IIpRangeContext

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

func (s *Ipv4AddressContext) SingleIPv4Address() ISingleIPv4AddressContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleIPv4AddressContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISingleIPv4AddressContext)
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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(71)
		p.SingleIPv4Address()
	}
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == HostsParserSLASH {
		{
			p.SetState(72)
			p.IpRange()
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

// ISingleIPv4AddressContext is an interface to support dynamic dispatch.
type ISingleIPv4AddressContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIpv4Digit() []IIpv4DigitContext
	Ipv4Digit(i int) IIpv4DigitContext
	AllDOT() []antlr.TerminalNode
	DOT(i int) antlr.TerminalNode

	// IsSingleIPv4AddressContext differentiates from other interfaces.
	IsSingleIPv4AddressContext()
}

type SingleIPv4AddressContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingleIPv4AddressContext() *SingleIPv4AddressContext {
	var p = new(SingleIPv4AddressContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HostsParserRULE_singleIPv4Address
	return p
}

func InitEmptySingleIPv4AddressContext(p *SingleIPv4AddressContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HostsParserRULE_singleIPv4Address
}

func (*SingleIPv4AddressContext) IsSingleIPv4AddressContext() {}

func NewSingleIPv4AddressContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SingleIPv4AddressContext {
	var p = new(SingleIPv4AddressContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = HostsParserRULE_singleIPv4Address

	return p
}

func (s *SingleIPv4AddressContext) GetParser() antlr.Parser { return s.parser }

func (s *SingleIPv4AddressContext) AllIpv4Digit() []IIpv4DigitContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIpv4DigitContext); ok {
			len++
		}
	}

	tst := make([]IIpv4DigitContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIpv4DigitContext); ok {
			tst[i] = t.(IIpv4DigitContext)
			i++
		}
	}

	return tst
}

func (s *SingleIPv4AddressContext) Ipv4Digit(i int) IIpv4DigitContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIpv4DigitContext); ok {
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

	return t.(IIpv4DigitContext)
}

func (s *SingleIPv4AddressContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(HostsParserDOT)
}

func (s *SingleIPv4AddressContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(HostsParserDOT, i)
}

func (s *SingleIPv4AddressContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleIPv4AddressContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SingleIPv4AddressContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HostsListener); ok {
		listenerT.EnterSingleIPv4Address(s)
	}
}

func (s *SingleIPv4AddressContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HostsListener); ok {
		listenerT.ExitSingleIPv4Address(s)
	}
}

func (p *HostsParser) SingleIPv4Address() (localctx ISingleIPv4AddressContext) {
	localctx = NewSingleIPv4AddressContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, HostsParserRULE_singleIPv4Address)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(75)
		p.Ipv4Digit()
	}
	{
		p.SetState(76)
		p.Match(HostsParserDOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(77)
		p.Ipv4Digit()
	}
	{
		p.SetState(78)
		p.Match(HostsParserDOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(79)
		p.Ipv4Digit()
	}
	{
		p.SetState(80)
		p.Match(HostsParserDOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(81)
		p.Ipv4Digit()
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
	SingleIPv6Address() ISingleIPv6AddressContext
	IpRange() IIpRangeContext

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

func (s *Ipv6AddressContext) SingleIPv6Address() ISingleIPv6AddressContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleIPv6AddressContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISingleIPv6AddressContext)
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
	p.EnterRule(localctx, 18, HostsParserRULE_ipv6Address)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(83)
		p.SingleIPv6Address()
	}
	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == HostsParserSLASH {
		{
			p.SetState(84)
			p.IpRange()
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

// ISingleIPv6AddressContext is an interface to support dynamic dispatch.
type ISingleIPv6AddressContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIpv6Octet() []IIpv6OctetContext
	Ipv6Octet(i int) IIpv6OctetContext
	AllCOLON() []antlr.TerminalNode
	COLON(i int) antlr.TerminalNode

	// IsSingleIPv6AddressContext differentiates from other interfaces.
	IsSingleIPv6AddressContext()
}

type SingleIPv6AddressContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingleIPv6AddressContext() *SingleIPv6AddressContext {
	var p = new(SingleIPv6AddressContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HostsParserRULE_singleIPv6Address
	return p
}

func InitEmptySingleIPv6AddressContext(p *SingleIPv6AddressContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HostsParserRULE_singleIPv6Address
}

func (*SingleIPv6AddressContext) IsSingleIPv6AddressContext() {}

func NewSingleIPv6AddressContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SingleIPv6AddressContext {
	var p = new(SingleIPv6AddressContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = HostsParserRULE_singleIPv6Address

	return p
}

func (s *SingleIPv6AddressContext) GetParser() antlr.Parser { return s.parser }

func (s *SingleIPv6AddressContext) AllIpv6Octet() []IIpv6OctetContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIpv6OctetContext); ok {
			len++
		}
	}

	tst := make([]IIpv6OctetContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIpv6OctetContext); ok {
			tst[i] = t.(IIpv6OctetContext)
			i++
		}
	}

	return tst
}

func (s *SingleIPv6AddressContext) Ipv6Octet(i int) IIpv6OctetContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIpv6OctetContext); ok {
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

	return t.(IIpv6OctetContext)
}

func (s *SingleIPv6AddressContext) AllCOLON() []antlr.TerminalNode {
	return s.GetTokens(HostsParserCOLON)
}

func (s *SingleIPv6AddressContext) COLON(i int) antlr.TerminalNode {
	return s.GetToken(HostsParserCOLON, i)
}

func (s *SingleIPv6AddressContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleIPv6AddressContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SingleIPv6AddressContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HostsListener); ok {
		listenerT.EnterSingleIPv6Address(s)
	}
}

func (s *SingleIPv6AddressContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HostsListener); ok {
		listenerT.ExitSingleIPv6Address(s)
	}
}

func (p *HostsParser) SingleIPv6Address() (localctx ISingleIPv6AddressContext) {
	localctx = NewSingleIPv6AddressContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, HostsParserRULE_singleIPv6Address)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(90)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(87)
				p.Ipv6Octet()
			}
			{
				p.SetState(88)
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

		p.SetState(92)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	{
		p.SetState(94)
		p.Ipv6Octet()
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

// IIpv4DigitContext is an interface to support dynamic dispatch.
type IIpv4DigitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DIGITS() antlr.TerminalNode

	// IsIpv4DigitContext differentiates from other interfaces.
	IsIpv4DigitContext()
}

type Ipv4DigitContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIpv4DigitContext() *Ipv4DigitContext {
	var p = new(Ipv4DigitContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HostsParserRULE_ipv4Digit
	return p
}

func InitEmptyIpv4DigitContext(p *Ipv4DigitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HostsParserRULE_ipv4Digit
}

func (*Ipv4DigitContext) IsIpv4DigitContext() {}

func NewIpv4DigitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ipv4DigitContext {
	var p = new(Ipv4DigitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = HostsParserRULE_ipv4Digit

	return p
}

func (s *Ipv4DigitContext) GetParser() antlr.Parser { return s.parser }

func (s *Ipv4DigitContext) DIGITS() antlr.TerminalNode {
	return s.GetToken(HostsParserDIGITS, 0)
}

func (s *Ipv4DigitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ipv4DigitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ipv4DigitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HostsListener); ok {
		listenerT.EnterIpv4Digit(s)
	}
}

func (s *Ipv4DigitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HostsListener); ok {
		listenerT.ExitIpv4Digit(s)
	}
}

func (p *HostsParser) Ipv4Digit() (localctx IIpv4DigitContext) {
	localctx = NewIpv4DigitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, HostsParserRULE_ipv4Digit)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(96)
		p.Match(HostsParserDIGITS)
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

// IIpv6OctetContext is an interface to support dynamic dispatch.
type IIpv6OctetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OCTETS() antlr.TerminalNode

	// IsIpv6OctetContext differentiates from other interfaces.
	IsIpv6OctetContext()
}

type Ipv6OctetContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIpv6OctetContext() *Ipv6OctetContext {
	var p = new(Ipv6OctetContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HostsParserRULE_ipv6Octet
	return p
}

func InitEmptyIpv6OctetContext(p *Ipv6OctetContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = HostsParserRULE_ipv6Octet
}

func (*Ipv6OctetContext) IsIpv6OctetContext() {}

func NewIpv6OctetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ipv6OctetContext {
	var p = new(Ipv6OctetContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = HostsParserRULE_ipv6Octet

	return p
}

func (s *Ipv6OctetContext) GetParser() antlr.Parser { return s.parser }

func (s *Ipv6OctetContext) OCTETS() antlr.TerminalNode {
	return s.GetToken(HostsParserOCTETS, 0)
}

func (s *Ipv6OctetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ipv6OctetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ipv6OctetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HostsListener); ok {
		listenerT.EnterIpv6Octet(s)
	}
}

func (s *Ipv6OctetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HostsListener); ok {
		listenerT.ExitIpv6Octet(s)
	}
}

func (p *HostsParser) Ipv6Octet() (localctx IIpv6OctetContext) {
	localctx = NewIpv6OctetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, HostsParserRULE_ipv6Octet)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(98)
		p.Match(HostsParserOCTETS)
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
	p.EnterRule(localctx, 26, HostsParserRULE_ipRange)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(100)
		p.Match(HostsParserSLASH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(101)
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
	DIGITS() antlr.TerminalNode

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

func (s *IpRangeBitsContext) DIGITS() antlr.TerminalNode {
	return s.GetToken(HostsParserDIGITS, 0)
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
	p.EnterRule(localctx, 28, HostsParserRULE_ipRangeBits)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(103)
		p.Match(HostsParserDIGITS)
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
	p.EnterRule(localctx, 30, HostsParserRULE_comment)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(105)
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
	p.EnterRule(localctx, 32, HostsParserRULE_leadingComment)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(107)
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
