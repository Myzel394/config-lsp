// Code generated from Fstab.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Fstab

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

type FstabParser struct {
	*antlr.BaseParser
}

var FstabParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func fstabParserInit() {
	staticData := &FstabParserStaticData
	staticData.LiteralNames = []string{
		"", "", "", "'#'",
	}
	staticData.SymbolicNames = []string{
		"", "DIGITS", "WHITESPACE", "HASH", "STRING", "QUOTED_STRING", "ADFS",
		"AFFS", "BTRFS", "EXFAT",
	}
	staticData.RuleNames = []string{
		"entry", "spec", "mountPoint", "fileSystem", "mountOptions", "freq",
		"pass",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 9, 68, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7, 4,
		2, 5, 7, 5, 2, 6, 7, 6, 1, 0, 3, 0, 16, 8, 0, 1, 0, 3, 0, 19, 8, 0, 1,
		0, 3, 0, 22, 8, 0, 1, 0, 3, 0, 25, 8, 0, 1, 0, 3, 0, 28, 8, 0, 1, 0, 3,
		0, 31, 8, 0, 1, 0, 3, 0, 34, 8, 0, 1, 0, 3, 0, 37, 8, 0, 1, 0, 3, 0, 40,
		8, 0, 1, 0, 3, 0, 43, 8, 0, 1, 0, 3, 0, 46, 8, 0, 1, 0, 3, 0, 49, 8, 0,
		1, 0, 3, 0, 52, 8, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1,
		4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 0, 0, 7, 0, 2, 4, 6, 8, 10, 12,
		0, 2, 1, 0, 4, 5, 1, 0, 4, 9, 73, 0, 15, 1, 0, 0, 0, 2, 55, 1, 0, 0, 0,
		4, 57, 1, 0, 0, 0, 6, 59, 1, 0, 0, 0, 8, 61, 1, 0, 0, 0, 10, 63, 1, 0,
		0, 0, 12, 65, 1, 0, 0, 0, 14, 16, 5, 2, 0, 0, 15, 14, 1, 0, 0, 0, 15, 16,
		1, 0, 0, 0, 16, 18, 1, 0, 0, 0, 17, 19, 3, 2, 1, 0, 18, 17, 1, 0, 0, 0,
		18, 19, 1, 0, 0, 0, 19, 21, 1, 0, 0, 0, 20, 22, 5, 2, 0, 0, 21, 20, 1,
		0, 0, 0, 21, 22, 1, 0, 0, 0, 22, 24, 1, 0, 0, 0, 23, 25, 3, 4, 2, 0, 24,
		23, 1, 0, 0, 0, 24, 25, 1, 0, 0, 0, 25, 27, 1, 0, 0, 0, 26, 28, 5, 2, 0,
		0, 27, 26, 1, 0, 0, 0, 27, 28, 1, 0, 0, 0, 28, 30, 1, 0, 0, 0, 29, 31,
		3, 6, 3, 0, 30, 29, 1, 0, 0, 0, 30, 31, 1, 0, 0, 0, 31, 33, 1, 0, 0, 0,
		32, 34, 5, 2, 0, 0, 33, 32, 1, 0, 0, 0, 33, 34, 1, 0, 0, 0, 34, 36, 1,
		0, 0, 0, 35, 37, 3, 8, 4, 0, 36, 35, 1, 0, 0, 0, 36, 37, 1, 0, 0, 0, 37,
		39, 1, 0, 0, 0, 38, 40, 5, 2, 0, 0, 39, 38, 1, 0, 0, 0, 39, 40, 1, 0, 0,
		0, 40, 42, 1, 0, 0, 0, 41, 43, 3, 10, 5, 0, 42, 41, 1, 0, 0, 0, 42, 43,
		1, 0, 0, 0, 43, 45, 1, 0, 0, 0, 44, 46, 5, 2, 0, 0, 45, 44, 1, 0, 0, 0,
		45, 46, 1, 0, 0, 0, 46, 48, 1, 0, 0, 0, 47, 49, 3, 12, 6, 0, 48, 47, 1,
		0, 0, 0, 48, 49, 1, 0, 0, 0, 49, 51, 1, 0, 0, 0, 50, 52, 5, 2, 0, 0, 51,
		50, 1, 0, 0, 0, 51, 52, 1, 0, 0, 0, 52, 53, 1, 0, 0, 0, 53, 54, 5, 0, 0,
		1, 54, 1, 1, 0, 0, 0, 55, 56, 7, 0, 0, 0, 56, 3, 1, 0, 0, 0, 57, 58, 7,
		0, 0, 0, 58, 5, 1, 0, 0, 0, 59, 60, 7, 1, 0, 0, 60, 7, 1, 0, 0, 0, 61,
		62, 7, 0, 0, 0, 62, 9, 1, 0, 0, 0, 63, 64, 5, 1, 0, 0, 64, 11, 1, 0, 0,
		0, 65, 66, 5, 1, 0, 0, 66, 13, 1, 0, 0, 0, 13, 15, 18, 21, 24, 27, 30,
		33, 36, 39, 42, 45, 48, 51,
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

// FstabParserInit initializes any static state used to implement FstabParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewFstabParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func FstabParserInit() {
	staticData := &FstabParserStaticData
	staticData.once.Do(fstabParserInit)
}

// NewFstabParser produces a new parser instance for the optional input antlr.TokenStream.
func NewFstabParser(input antlr.TokenStream) *FstabParser {
	FstabParserInit()
	this := new(FstabParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &FstabParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Fstab.g4"

	return this
}

// FstabParser tokens.
const (
	FstabParserEOF           = antlr.TokenEOF
	FstabParserDIGITS        = 1
	FstabParserWHITESPACE    = 2
	FstabParserHASH          = 3
	FstabParserSTRING        = 4
	FstabParserQUOTED_STRING = 5
	FstabParserADFS          = 6
	FstabParserAFFS          = 7
	FstabParserBTRFS         = 8
	FstabParserEXFAT         = 9
)

// FstabParser rules.
const (
	FstabParserRULE_entry        = 0
	FstabParserRULE_spec         = 1
	FstabParserRULE_mountPoint   = 2
	FstabParserRULE_fileSystem   = 3
	FstabParserRULE_mountOptions = 4
	FstabParserRULE_freq         = 5
	FstabParserRULE_pass         = 6
)

// IEntryContext is an interface to support dynamic dispatch.
type IEntryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllWHITESPACE() []antlr.TerminalNode
	WHITESPACE(i int) antlr.TerminalNode
	Spec() ISpecContext
	MountPoint() IMountPointContext
	FileSystem() IFileSystemContext
	MountOptions() IMountOptionsContext
	Freq() IFreqContext
	Pass() IPassContext

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
	p.RuleIndex = FstabParserRULE_entry
	return p
}

func InitEmptyEntryContext(p *EntryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FstabParserRULE_entry
}

func (*EntryContext) IsEntryContext() {}

func NewEntryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EntryContext {
	var p = new(EntryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FstabParserRULE_entry

	return p
}

func (s *EntryContext) GetParser() antlr.Parser { return s.parser }

func (s *EntryContext) EOF() antlr.TerminalNode {
	return s.GetToken(FstabParserEOF, 0)
}

func (s *EntryContext) AllWHITESPACE() []antlr.TerminalNode {
	return s.GetTokens(FstabParserWHITESPACE)
}

func (s *EntryContext) WHITESPACE(i int) antlr.TerminalNode {
	return s.GetToken(FstabParserWHITESPACE, i)
}

func (s *EntryContext) Spec() ISpecContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISpecContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISpecContext)
}

func (s *EntryContext) MountPoint() IMountPointContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMountPointContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMountPointContext)
}

func (s *EntryContext) FileSystem() IFileSystemContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFileSystemContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFileSystemContext)
}

func (s *EntryContext) MountOptions() IMountOptionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMountOptionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMountOptionsContext)
}

func (s *EntryContext) Freq() IFreqContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFreqContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFreqContext)
}

func (s *EntryContext) Pass() IPassContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPassContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPassContext)
}

func (s *EntryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EntryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EntryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FstabListener); ok {
		listenerT.EnterEntry(s)
	}
}

func (s *EntryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FstabListener); ok {
		listenerT.ExitEntry(s)
	}
}

func (p *FstabParser) Entry() (localctx IEntryContext) {
	localctx = NewEntryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, FstabParserRULE_entry)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(15)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(14)
			p.Match(FstabParserWHITESPACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(18)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(17)
			p.Spec()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(21)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(20)
			p.Match(FstabParserWHITESPACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(24)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(23)
			p.MountPoint()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(27)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(26)
			p.Match(FstabParserWHITESPACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(30)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(29)
			p.FileSystem()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(33)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(32)
			p.Match(FstabParserWHITESPACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

	if _la == FstabParserSTRING || _la == FstabParserQUOTED_STRING {
		{
			p.SetState(35)
			p.MountOptions()
		}

	}
	p.SetState(39)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(38)
			p.Match(FstabParserWHITESPACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(42)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(41)
			p.Freq()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(45)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(44)
			p.Match(FstabParserWHITESPACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(48)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FstabParserDIGITS {
		{
			p.SetState(47)
			p.Pass()
		}

	}
	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FstabParserWHITESPACE {
		{
			p.SetState(50)
			p.Match(FstabParserWHITESPACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(53)
		p.Match(FstabParserEOF)
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

// ISpecContext is an interface to support dynamic dispatch.
type ISpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	QUOTED_STRING() antlr.TerminalNode
	STRING() antlr.TerminalNode

	// IsSpecContext differentiates from other interfaces.
	IsSpecContext()
}

type SpecContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySpecContext() *SpecContext {
	var p = new(SpecContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FstabParserRULE_spec
	return p
}

func InitEmptySpecContext(p *SpecContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FstabParserRULE_spec
}

func (*SpecContext) IsSpecContext() {}

func NewSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SpecContext {
	var p = new(SpecContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FstabParserRULE_spec

	return p
}

func (s *SpecContext) GetParser() antlr.Parser { return s.parser }

func (s *SpecContext) QUOTED_STRING() antlr.TerminalNode {
	return s.GetToken(FstabParserQUOTED_STRING, 0)
}

func (s *SpecContext) STRING() antlr.TerminalNode {
	return s.GetToken(FstabParserSTRING, 0)
}

func (s *SpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FstabListener); ok {
		listenerT.EnterSpec(s)
	}
}

func (s *SpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FstabListener); ok {
		listenerT.ExitSpec(s)
	}
}

func (p *FstabParser) Spec() (localctx ISpecContext) {
	localctx = NewSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, FstabParserRULE_spec)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(55)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FstabParserSTRING || _la == FstabParserQUOTED_STRING) {
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

// IMountPointContext is an interface to support dynamic dispatch.
type IMountPointContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	QUOTED_STRING() antlr.TerminalNode
	STRING() antlr.TerminalNode

	// IsMountPointContext differentiates from other interfaces.
	IsMountPointContext()
}

type MountPointContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMountPointContext() *MountPointContext {
	var p = new(MountPointContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FstabParserRULE_mountPoint
	return p
}

func InitEmptyMountPointContext(p *MountPointContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FstabParserRULE_mountPoint
}

func (*MountPointContext) IsMountPointContext() {}

func NewMountPointContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MountPointContext {
	var p = new(MountPointContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FstabParserRULE_mountPoint

	return p
}

func (s *MountPointContext) GetParser() antlr.Parser { return s.parser }

func (s *MountPointContext) QUOTED_STRING() antlr.TerminalNode {
	return s.GetToken(FstabParserQUOTED_STRING, 0)
}

func (s *MountPointContext) STRING() antlr.TerminalNode {
	return s.GetToken(FstabParserSTRING, 0)
}

func (s *MountPointContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MountPointContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MountPointContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FstabListener); ok {
		listenerT.EnterMountPoint(s)
	}
}

func (s *MountPointContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FstabListener); ok {
		listenerT.ExitMountPoint(s)
	}
}

func (p *FstabParser) MountPoint() (localctx IMountPointContext) {
	localctx = NewMountPointContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, FstabParserRULE_mountPoint)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(57)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FstabParserSTRING || _la == FstabParserQUOTED_STRING) {
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

// IFileSystemContext is an interface to support dynamic dispatch.
type IFileSystemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ADFS() antlr.TerminalNode
	AFFS() antlr.TerminalNode
	BTRFS() antlr.TerminalNode
	EXFAT() antlr.TerminalNode
	STRING() antlr.TerminalNode
	QUOTED_STRING() antlr.TerminalNode

	// IsFileSystemContext differentiates from other interfaces.
	IsFileSystemContext()
}

type FileSystemContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFileSystemContext() *FileSystemContext {
	var p = new(FileSystemContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FstabParserRULE_fileSystem
	return p
}

func InitEmptyFileSystemContext(p *FileSystemContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FstabParserRULE_fileSystem
}

func (*FileSystemContext) IsFileSystemContext() {}

func NewFileSystemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FileSystemContext {
	var p = new(FileSystemContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FstabParserRULE_fileSystem

	return p
}

func (s *FileSystemContext) GetParser() antlr.Parser { return s.parser }

func (s *FileSystemContext) ADFS() antlr.TerminalNode {
	return s.GetToken(FstabParserADFS, 0)
}

func (s *FileSystemContext) AFFS() antlr.TerminalNode {
	return s.GetToken(FstabParserAFFS, 0)
}

func (s *FileSystemContext) BTRFS() antlr.TerminalNode {
	return s.GetToken(FstabParserBTRFS, 0)
}

func (s *FileSystemContext) EXFAT() antlr.TerminalNode {
	return s.GetToken(FstabParserEXFAT, 0)
}

func (s *FileSystemContext) STRING() antlr.TerminalNode {
	return s.GetToken(FstabParserSTRING, 0)
}

func (s *FileSystemContext) QUOTED_STRING() antlr.TerminalNode {
	return s.GetToken(FstabParserQUOTED_STRING, 0)
}

func (s *FileSystemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FileSystemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FileSystemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FstabListener); ok {
		listenerT.EnterFileSystem(s)
	}
}

func (s *FileSystemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FstabListener); ok {
		listenerT.ExitFileSystem(s)
	}
}

func (p *FstabParser) FileSystem() (localctx IFileSystemContext) {
	localctx = NewFileSystemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, FstabParserRULE_fileSystem)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(59)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1008) != 0) {
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

// IMountOptionsContext is an interface to support dynamic dispatch.
type IMountOptionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	QUOTED_STRING() antlr.TerminalNode
	STRING() antlr.TerminalNode

	// IsMountOptionsContext differentiates from other interfaces.
	IsMountOptionsContext()
}

type MountOptionsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMountOptionsContext() *MountOptionsContext {
	var p = new(MountOptionsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FstabParserRULE_mountOptions
	return p
}

func InitEmptyMountOptionsContext(p *MountOptionsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FstabParserRULE_mountOptions
}

func (*MountOptionsContext) IsMountOptionsContext() {}

func NewMountOptionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MountOptionsContext {
	var p = new(MountOptionsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FstabParserRULE_mountOptions

	return p
}

func (s *MountOptionsContext) GetParser() antlr.Parser { return s.parser }

func (s *MountOptionsContext) QUOTED_STRING() antlr.TerminalNode {
	return s.GetToken(FstabParserQUOTED_STRING, 0)
}

func (s *MountOptionsContext) STRING() antlr.TerminalNode {
	return s.GetToken(FstabParserSTRING, 0)
}

func (s *MountOptionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MountOptionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MountOptionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FstabListener); ok {
		listenerT.EnterMountOptions(s)
	}
}

func (s *MountOptionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FstabListener); ok {
		listenerT.ExitMountOptions(s)
	}
}

func (p *FstabParser) MountOptions() (localctx IMountOptionsContext) {
	localctx = NewMountOptionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, FstabParserRULE_mountOptions)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(61)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FstabParserSTRING || _la == FstabParserQUOTED_STRING) {
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

// IFreqContext is an interface to support dynamic dispatch.
type IFreqContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DIGITS() antlr.TerminalNode

	// IsFreqContext differentiates from other interfaces.
	IsFreqContext()
}

type FreqContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFreqContext() *FreqContext {
	var p = new(FreqContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FstabParserRULE_freq
	return p
}

func InitEmptyFreqContext(p *FreqContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FstabParserRULE_freq
}

func (*FreqContext) IsFreqContext() {}

func NewFreqContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FreqContext {
	var p = new(FreqContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FstabParserRULE_freq

	return p
}

func (s *FreqContext) GetParser() antlr.Parser { return s.parser }

func (s *FreqContext) DIGITS() antlr.TerminalNode {
	return s.GetToken(FstabParserDIGITS, 0)
}

func (s *FreqContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FreqContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FreqContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FstabListener); ok {
		listenerT.EnterFreq(s)
	}
}

func (s *FreqContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FstabListener); ok {
		listenerT.ExitFreq(s)
	}
}

func (p *FstabParser) Freq() (localctx IFreqContext) {
	localctx = NewFreqContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, FstabParserRULE_freq)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(63)
		p.Match(FstabParserDIGITS)
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

// IPassContext is an interface to support dynamic dispatch.
type IPassContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DIGITS() antlr.TerminalNode

	// IsPassContext differentiates from other interfaces.
	IsPassContext()
}

type PassContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPassContext() *PassContext {
	var p = new(PassContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FstabParserRULE_pass
	return p
}

func InitEmptyPassContext(p *PassContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FstabParserRULE_pass
}

func (*PassContext) IsPassContext() {}

func NewPassContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PassContext {
	var p = new(PassContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FstabParserRULE_pass

	return p
}

func (s *PassContext) GetParser() antlr.Parser { return s.parser }

func (s *PassContext) DIGITS() antlr.TerminalNode {
	return s.GetToken(FstabParserDIGITS, 0)
}

func (s *PassContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PassContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PassContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FstabListener); ok {
		listenerT.EnterPass(s)
	}
}

func (s *PassContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FstabListener); ok {
		listenerT.ExitPass(s)
	}
}

func (p *FstabParser) Pass() (localctx IPassContext) {
	localctx = NewPassContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, FstabParserRULE_pass)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(65)
		p.Match(FstabParserDIGITS)
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
