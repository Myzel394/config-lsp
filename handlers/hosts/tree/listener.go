package tree

import (
	"config-lsp/common"
	"config-lsp/handlers/hosts/parser"

	"github.com/antlr4-go/antlr/v4"
)

type hostsListenerContext struct {
	line uint32
}

type hostsParserListener struct {
	*parser.BaseHostsListener
	Parser       *HostsParser
	hostsContext hostsListenerContext
}

func (s *hostsParserListener) EnterComment(ctx *parser.CommentContext) {
	line := uint32(s.hostsContext.line)
	s.Parser.CommentLines[line] = struct{}{}
}

func (s *hostsParserListener) EnterEntry(ctx *parser.EntryContext) {
	location := characterRangeFromCtx(ctx.BaseParserRuleContext)
	location.ChangeBothLines(s.hostsContext.line)

	s.Parser.Tree.Entries[location.Start.Line] = &HostsEntry{
		Location: location,
	}
}

func (s *hostsParserListener) EnterIpAddress(ctx *parser.IpAddressContext) {
	location := characterRangeFromCtx(ctx.BaseParserRuleContext)
	location.ChangeBothLines(s.hostsContext.line)

	entry := s.Parser.Tree.Entries[location.Start.Line]

	entry.IPAddress = &HostsIPAddress{
		Location: location,
		Value:    ctx.GetText(),
	}
}

func (s *hostsParserListener) EnterHostname(ctx *parser.HostnameContext) {
	location := characterRangeFromCtx(ctx.BaseParserRuleContext)
	location.ChangeBothLines(s.hostsContext.line)

	entry := s.Parser.Tree.Entries[location.Start.Line]

	entry.Hostname = &HostsHostname{
		Location: location,
		Value:    ctx.GetText(),
	}

	s.Parser.Tree.Entries[location.Start.Line] = entry
}

func (s *hostsParserListener) EnterAliases(ctx *parser.AliasesContext) {
	location := characterRangeFromCtx(ctx.BaseParserRuleContext)
	location.ChangeBothLines(s.hostsContext.line)

	entry := s.Parser.Tree.Entries[location.Start.Line]

	aliases := make([]*HostsHostname, 0)

	entry.Aliases = aliases
}

func (s *hostsParserListener) EnterAlias(ctx *parser.AliasContext) {
	location := characterRangeFromCtx(ctx.BaseParserRuleContext)
	location.ChangeBothLines(s.hostsContext.line)

	entry := s.Parser.Tree.Entries[location.Start.Line]

	alias := HostsHostname{
		Location: location,
		Value:    ctx.GetText(),
	}

	entry.Aliases = append(entry.Aliases, &alias)
}

func createHostsFileListener(
	parser *HostsParser,
	line uint32,
) hostsParserListener {
	return hostsParserListener{
		Parser: parser,
		hostsContext: hostsListenerContext{
			line: line,
		},
	}
}

type errorListener struct {
	*antlr.DefaultErrorListener
	Errors       []common.LSPError
	hostsContext hostsListenerContext
}

func createErrorListener(
	line uint32,
) errorListener {
	return errorListener{
		Errors: make([]common.LSPError, 0),
		hostsContext: hostsListenerContext{
			line: line,
		},
	}
}

func (d *errorListener) SyntaxError(
	recognizer antlr.Recognizer,
	offendingSymbol interface{},
	_ int,
	character int,
	message string,
	error antlr.RecognitionException,
) {
	line := d.hostsContext.line
	d.Errors = append(d.Errors, common.LSPError{
		Range: common.CreateSingleCharRange(uint32(line), uint32(character)),
		Err: common.SyntaxError{
			Message: message,
		},
	})
}
