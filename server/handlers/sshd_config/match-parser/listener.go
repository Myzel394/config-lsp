package matchparser

import (
	"config-lsp/common"
	commonparser "config-lsp/common/parser"
	"config-lsp/handlers/sshd_config/match-parser/parser"
	"config-lsp/utils"
	"errors"
	"fmt"
	"strings"
)

func createMatchListenerContext(
	line uint32,
	startCharacter uint32,
) *matchListenerContext {
	return &matchListenerContext{
		currentEntry:   nil,
		line:           line,
		startCharacter: startCharacter,
	}
}

type matchListenerContext struct {
	currentEntry   *MatchEntry
	line           uint32
	startCharacter uint32
}

func createListener(
	match *Match,
	context *matchListenerContext,
) matchParserListener {
	return matchParserListener{
		match:        match,
		Errors:       make([]common.LSPError, 0),
		matchContext: context,
	}
}

type matchParserListener struct {
	*parser.BaseMatchListener
	match        *Match
	Errors       []common.LSPError
	matchContext *matchListenerContext
}

func (s *matchParserListener) EnterMatchEntry(ctx *parser.MatchEntryContext) {
	location := common.CharacterRangeFromCtx(ctx.BaseParserRuleContext).ShiftHorizontal(s.matchContext.startCharacter).ChangeBothLines(s.matchContext.line)

	entry := &MatchEntry{
		LocationRange: location,
		Value:         commonparser.ParseRawString(ctx.GetText(), commonparser.FullFeatures),
	}

	s.match.Entries = append(s.match.Entries, entry)
	s.matchContext.currentEntry = entry
}

func (s *matchParserListener) ExitMatchEntry(ctx *parser.MatchEntryContext) {
	s.matchContext.currentEntry = nil
}

var availableCriteria = map[string]MatchCriteriaType{
	string(MatchCriteriaTypeUser):         MatchCriteriaTypeUser,
	string(MatchCriteriaTypeGroup):        MatchCriteriaTypeGroup,
	string(MatchCriteriaTypeHost):         MatchCriteriaTypeHost,
	string(MatchCriteriaTypeLocalAddress): MatchCriteriaTypeLocalAddress,
	string(MatchCriteriaTypeLocalPort):    MatchCriteriaTypeLocalPort,
	string(MatchCriteriaTypeRDomain):      MatchCriteriaTypeRDomain,
	string(MatchCriteriaTypeAddress):      MatchCriteriaTypeAddress,
}

func (s *matchParserListener) EnterCriteria(ctx *parser.CriteriaContext) {
	location := common.CharacterRangeFromCtx(ctx.BaseParserRuleContext).ShiftHorizontal(s.matchContext.startCharacter).ChangeBothLines(s.matchContext.line)

	value := commonparser.ParseRawString(ctx.GetText(), commonparser.FullFeatures)

	criteria, found := availableCriteria[value.Value]

	if !found {
		s.Errors = append(s.Errors, common.LSPError{
			Range: location,
			Err:   errors.New(fmt.Sprintf("Unknown criteria: %s; It must be one of: %s", ctx.GetText(), strings.Join(utils.KeysOfMap(availableCriteria), ", "))),
		})
		return
	}

	s.matchContext.currentEntry.Criteria = MatchCriteria{
		LocationRange: location,
		Type:          criteria,
		Value:         value,
	}
}

func (s *matchParserListener) EnterSeparator(ctx *parser.SeparatorContext) {
	location := common.CharacterRangeFromCtx(ctx.BaseParserRuleContext).ShiftHorizontal(s.matchContext.startCharacter).ChangeBothLines(s.matchContext.line)

	s.matchContext.currentEntry.Separator = &MatchSeparator{
		LocationRange: location,
	}
}

func (s *matchParserListener) EnterValues(ctx *parser.ValuesContext) {
	location := common.CharacterRangeFromCtx(ctx.BaseParserRuleContext).ShiftHorizontal(s.matchContext.startCharacter).ChangeBothLines(s.matchContext.line)

	s.matchContext.currentEntry.Values = &MatchValues{
		LocationRange: location,
		Values:        make([]*MatchValue, 0),
	}
}

func (s *matchParserListener) EnterValue(ctx *parser.ValueContext) {
	location := common.CharacterRangeFromCtx(ctx.BaseParserRuleContext).ShiftHorizontal(s.matchContext.startCharacter).ChangeBothLines(s.matchContext.line)

	value := &MatchValue{
		LocationRange: location,
		Value:         commonparser.ParseRawString(ctx.GetText(), commonparser.FullFeatures),
	}

	s.matchContext.currentEntry.Values.Values = append(s.matchContext.currentEntry.Values.Values, value)
}
