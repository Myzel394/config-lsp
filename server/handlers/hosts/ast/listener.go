package ast

import (
	"config-lsp/common"
	docvalues "config-lsp/doc-values"
	parser2 "config-lsp/handlers/hosts/ast/parser"
	"errors"
	"net"
	"regexp"

	"github.com/antlr4-go/antlr/v4"
)

type hostsListenerContext struct {
	line uint32
}

type hostsParserListener struct {
	*parser2.BaseHostsListener
	Parser       *HostsParser
	Errors       []common.LSPError
	hostsContext hostsListenerContext
}

func (s *hostsParserListener) EnterComment(ctx *parser2.CommentContext) {
	line := s.hostsContext.line
	s.Parser.CommentLines[line] = struct{}{}
}

func (s *hostsParserListener) EnterEntry(ctx *parser2.EntryContext) {
	location := common.CharacterRangeFromCtx(ctx.BaseParserRuleContext)
	location.ChangeBothLines(s.hostsContext.line)

	s.Parser.Tree.Entries.Put(location.Start.Line, &HostsEntry{
		Location: location,
	})
}

var containsPortPattern = regexp.MustCompile(`:[0-9]+$`)

func (s *hostsParserListener) EnterIpAddress(ctx *parser2.IpAddressContext) {
	location := common.CharacterRangeFromCtx(ctx.BaseParserRuleContext)
	location.ChangeBothLines(s.hostsContext.line)

	ip := net.ParseIP(ctx.GetText())

	if ip == nil {
		if containsPortPattern.MatchString(ctx.GetText()) {
			s.Errors = append(s.Errors, common.LSPError{
				Range: location,
				Err:   errors.New("Port numbers are not allowed in IP addresses"),
			})
		} else {
			s.Errors = append(s.Errors, common.LSPError{
				Range: location,
				Err:   docvalues.InvalidIPAddress{},
			})
		}

		return
	}

	ipAddr, err := net.ResolveIPAddr("ip", ip.String())

	if err != nil {
		s.Errors = append(s.Errors, common.LSPError{
			Range: location,
			Err:   docvalues.InvalidIPAddress{},
		})
	}

	rawEntry, _ := s.Parser.Tree.Entries.Get(location.Start.Line)
	entry := rawEntry.(*HostsEntry)

	entry.IPAddress = &HostsIPAddress{
		Location: location,
		Value:    *ipAddr,
	}
}

func (s *hostsParserListener) EnterHostname(ctx *parser2.HostnameContext) {
	location := common.CharacterRangeFromCtx(ctx.BaseParserRuleContext)
	location.ChangeBothLines(s.hostsContext.line)

	rawEntry, _ := s.Parser.Tree.Entries.Get(location.Start.Line)
	entry := rawEntry.(*HostsEntry)

	entry.Hostname = &HostsHostname{
		Location: location,
		Value:    ctx.GetText(),
	}
}

func (s *hostsParserListener) EnterAliases(ctx *parser2.AliasesContext) {
	location := common.CharacterRangeFromCtx(ctx.BaseParserRuleContext)
	location.ChangeBothLines(s.hostsContext.line)

	rawEntry, _ := s.Parser.Tree.Entries.Get(location.Start.Line)
	entry := rawEntry.(*HostsEntry)

	aliases := make([]*HostsHostname, 0)

	entry.Aliases = aliases
}

func (s *hostsParserListener) EnterAlias(ctx *parser2.AliasContext) {
	location := common.CharacterRangeFromCtx(ctx.BaseParserRuleContext)
	location.ChangeBothLines(s.hostsContext.line)

	rawEntry, _ := s.Parser.Tree.Entries.Get(location.Start.Line)
	entry := rawEntry.(*HostsEntry)

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
		Errors: make([]common.LSPError, 0),
	}
}

type errorListener struct {
	*antlr.DefaultErrorListener
	Errors       []common.LSPError
	hostsContext hostsListenerContext
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
