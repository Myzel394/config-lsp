// Code generated from Hosts.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Hosts

import "github.com/antlr4-go/antlr/v4"

// HostsListener is a complete listener for a parse tree produced by HostsParser.
type HostsListener interface {
	antlr.ParseTreeListener

	// EnterLineStatement is called when entering the lineStatement production.
	EnterLineStatement(c *LineStatementContext)

	// EnterEntry is called when entering the entry production.
	EnterEntry(c *EntryContext)

	// EnterAliases is called when entering the aliases production.
	EnterAliases(c *AliasesContext)

	// EnterAlias is called when entering the alias production.
	EnterAlias(c *AliasContext)

	// EnterHostname is called when entering the hostname production.
	EnterHostname(c *HostnameContext)

	// EnterDomain is called when entering the domain production.
	EnterDomain(c *DomainContext)

	// EnterIpAddress is called when entering the ipAddress production.
	EnterIpAddress(c *IpAddressContext)

	// EnterIpv4Address is called when entering the ipv4Address production.
	EnterIpv4Address(c *Ipv4AddressContext)

	// EnterSingleIPv4Address is called when entering the singleIPv4Address production.
	EnterSingleIPv4Address(c *SingleIPv4AddressContext)

	// EnterIpv6Address is called when entering the ipv6Address production.
	EnterIpv6Address(c *Ipv6AddressContext)

	// EnterSingleIPv6Address is called when entering the singleIPv6Address production.
	EnterSingleIPv6Address(c *SingleIPv6AddressContext)

	// EnterIpv4Digit is called when entering the ipv4Digit production.
	EnterIpv4Digit(c *Ipv4DigitContext)

	// EnterIpv6Octet is called when entering the ipv6Octet production.
	EnterIpv6Octet(c *Ipv6OctetContext)

	// EnterIpRange is called when entering the ipRange production.
	EnterIpRange(c *IpRangeContext)

	// EnterIpRangeBits is called when entering the ipRangeBits production.
	EnterIpRangeBits(c *IpRangeBitsContext)

	// EnterComment is called when entering the comment production.
	EnterComment(c *CommentContext)

	// EnterLeadingComment is called when entering the leadingComment production.
	EnterLeadingComment(c *LeadingCommentContext)

	// ExitLineStatement is called when exiting the lineStatement production.
	ExitLineStatement(c *LineStatementContext)

	// ExitEntry is called when exiting the entry production.
	ExitEntry(c *EntryContext)

	// ExitAliases is called when exiting the aliases production.
	ExitAliases(c *AliasesContext)

	// ExitAlias is called when exiting the alias production.
	ExitAlias(c *AliasContext)

	// ExitHostname is called when exiting the hostname production.
	ExitHostname(c *HostnameContext)

	// ExitDomain is called when exiting the domain production.
	ExitDomain(c *DomainContext)

	// ExitIpAddress is called when exiting the ipAddress production.
	ExitIpAddress(c *IpAddressContext)

	// ExitIpv4Address is called when exiting the ipv4Address production.
	ExitIpv4Address(c *Ipv4AddressContext)

	// ExitSingleIPv4Address is called when exiting the singleIPv4Address production.
	ExitSingleIPv4Address(c *SingleIPv4AddressContext)

	// ExitIpv6Address is called when exiting the ipv6Address production.
	ExitIpv6Address(c *Ipv6AddressContext)

	// ExitSingleIPv6Address is called when exiting the singleIPv6Address production.
	ExitSingleIPv6Address(c *SingleIPv6AddressContext)

	// ExitIpv4Digit is called when exiting the ipv4Digit production.
	ExitIpv4Digit(c *Ipv4DigitContext)

	// ExitIpv6Octet is called when exiting the ipv6Octet production.
	ExitIpv6Octet(c *Ipv6OctetContext)

	// ExitIpRange is called when exiting the ipRange production.
	ExitIpRange(c *IpRangeContext)

	// ExitIpRangeBits is called when exiting the ipRangeBits production.
	ExitIpRangeBits(c *IpRangeBitsContext)

	// ExitComment is called when exiting the comment production.
	ExitComment(c *CommentContext)

	// ExitLeadingComment is called when exiting the leadingComment production.
	ExitLeadingComment(c *LeadingCommentContext)
}
