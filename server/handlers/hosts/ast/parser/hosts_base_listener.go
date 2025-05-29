// Code generated from Hosts.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Hosts

import "github.com/antlr4-go/antlr/v4"

// BaseHostsListener is a complete listener for a parse tree produced by HostsParser.
type BaseHostsListener struct{}

var _ HostsListener = &BaseHostsListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseHostsListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseHostsListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseHostsListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseHostsListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterLineStatement is called when production lineStatement is entered.
func (s *BaseHostsListener) EnterLineStatement(ctx *LineStatementContext) {}

// ExitLineStatement is called when production lineStatement is exited.
func (s *BaseHostsListener) ExitLineStatement(ctx *LineStatementContext) {}

// EnterEntry is called when production entry is entered.
func (s *BaseHostsListener) EnterEntry(ctx *EntryContext) {}

// ExitEntry is called when production entry is exited.
func (s *BaseHostsListener) ExitEntry(ctx *EntryContext) {}

// EnterAliases is called when production aliases is entered.
func (s *BaseHostsListener) EnterAliases(ctx *AliasesContext) {}

// ExitAliases is called when production aliases is exited.
func (s *BaseHostsListener) ExitAliases(ctx *AliasesContext) {}

// EnterAlias is called when production alias is entered.
func (s *BaseHostsListener) EnterAlias(ctx *AliasContext) {}

// ExitAlias is called when production alias is exited.
func (s *BaseHostsListener) ExitAlias(ctx *AliasContext) {}

// EnterHostname is called when production hostname is entered.
func (s *BaseHostsListener) EnterHostname(ctx *HostnameContext) {}

// ExitHostname is called when production hostname is exited.
func (s *BaseHostsListener) ExitHostname(ctx *HostnameContext) {}

// EnterDomain is called when production domain is entered.
func (s *BaseHostsListener) EnterDomain(ctx *DomainContext) {}

// ExitDomain is called when production domain is exited.
func (s *BaseHostsListener) ExitDomain(ctx *DomainContext) {}

// EnterIpAddress is called when production ipAddress is entered.
func (s *BaseHostsListener) EnterIpAddress(ctx *IpAddressContext) {}

// ExitIpAddress is called when production ipAddress is exited.
func (s *BaseHostsListener) ExitIpAddress(ctx *IpAddressContext) {}

// EnterIpv4Address is called when production ipv4Address is entered.
func (s *BaseHostsListener) EnterIpv4Address(ctx *Ipv4AddressContext) {}

// ExitIpv4Address is called when production ipv4Address is exited.
func (s *BaseHostsListener) ExitIpv4Address(ctx *Ipv4AddressContext) {}

// EnterIpv6Address is called when production ipv6Address is entered.
func (s *BaseHostsListener) EnterIpv6Address(ctx *Ipv6AddressContext) {}

// ExitIpv6Address is called when production ipv6Address is exited.
func (s *BaseHostsListener) ExitIpv6Address(ctx *Ipv6AddressContext) {}

// EnterIpRange is called when production ipRange is entered.
func (s *BaseHostsListener) EnterIpRange(ctx *IpRangeContext) {}

// ExitIpRange is called when production ipRange is exited.
func (s *BaseHostsListener) ExitIpRange(ctx *IpRangeContext) {}

// EnterIpRangeBits is called when production ipRangeBits is entered.
func (s *BaseHostsListener) EnterIpRangeBits(ctx *IpRangeBitsContext) {}

// ExitIpRangeBits is called when production ipRangeBits is exited.
func (s *BaseHostsListener) ExitIpRangeBits(ctx *IpRangeBitsContext) {}

// EnterIpPort is called when production ipPort is entered.
func (s *BaseHostsListener) EnterIpPort(ctx *IpPortContext) {}

// ExitIpPort is called when production ipPort is exited.
func (s *BaseHostsListener) ExitIpPort(ctx *IpPortContext) {}

// EnterComment is called when production comment is entered.
func (s *BaseHostsListener) EnterComment(ctx *CommentContext) {}

// ExitComment is called when production comment is exited.
func (s *BaseHostsListener) ExitComment(ctx *CommentContext) {}

// EnterLeadingComment is called when production leadingComment is entered.
func (s *BaseHostsListener) EnterLeadingComment(ctx *LeadingCommentContext) {}

// ExitLeadingComment is called when production leadingComment is exited.
func (s *BaseHostsListener) ExitLeadingComment(ctx *LeadingCommentContext) {}
