package analyzer

import (
	"config-lsp/common"
	"config-lsp/utils"
	"context"
	"fmt"
	"net/netip"
	"strings"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func analyzeDNSPropertyContainsFallback(
	ctx *analyzerContext,
) {
	sections, found := ctx.document.Indexes.SectionsByName["Interface"]

	if !found {
		return
	}

	interfaceSection := sections[0]

	_, property := interfaceSection.FindFirstPropertyByName("DNS")

	if property == nil {
		return
	}

	dnsAmount := len(strings.Split(property.Value.Value, ","))

	if dnsAmount == 1 {
		ctx.diagnostics = append(ctx.diagnostics, protocol.Diagnostic{
			Message:  "There is only one DNS server specified. It is recommended to set up fallback DNS servers",
			Severity: &common.SeverityHint,
			Range:    property.Value.ToLSPRange(),
		})
	}
}

func analyzeKeepAlivePropertyIsSet(
	ctx *analyzerContext,
) {
	for _, section := range ctx.document.Indexes.SectionsByName["Peer"] {
		// If an endpoint is set, then we should only check for the keepalive property
		_, endpoint := section.FindFirstPropertyByName("Endpoint")
		_, persistentKeepAlive := section.FindFirstPropertyByName("PersistentKeepalive")

		if endpoint != nil && persistentKeepAlive == nil {
			ctx.diagnostics = append(ctx.diagnostics, protocol.Diagnostic{
				Message:  "PersistentKeepalive is not set. It is recommended to set this property, as it helps to maintain the connection when users are behind NAT",
				Severity: &common.SeverityHint,
				Range:    section.Header.ToLSPRange(),
			})
		}
	}
}

func analyzeSymmetricPropertiesSet(
	ctx *analyzerContext,
) {
	for section, info := range ctx.document.Indexes.AsymmetricRules {
		if info.PreMissing {
			properties := section.FindPropertiesByName("PreUp")

			// TODO: Fix later
			if len(properties) != 0 {
				for _, property := range properties {
					ctx.diagnostics = append(ctx.diagnostics, protocol.Diagnostic{
						Message:  "PreUp is set, but PreDown is not. It is recommended to set both properties symmetrically",
						Range:    property.Key.ToLSPRange(),
						Severity: &common.SeverityHint,
					})
				}
			}
		}

		if info.PostMissing {
			properties := section.FindPropertiesByName("PostUp")

			// TODO: Fix later
			if len(properties) == 0 {
				for _, property := range properties {
					ctx.diagnostics = append(ctx.diagnostics, protocol.Diagnostic{
						Message:  "PostUp is set, but PostDown is not. It is recommended to set both properties symmetrically",
						Range:    property.Key.ToLSPRange(),
						Severity: &common.SeverityHint,
					})
				}
			}
		}
	}
}

type key uint8

const (
	lineKey key = iota
)

// Strategy
// Simply compare the host bits of the IP addresses.
// Use a binary tree to store the host bits.
func analyzeDuplicateAllowedIPs(
	ctx *analyzerContext,
) {
	ipHostSet := utils.CreateIPv4HostSet()

	for _, section := range ctx.document.Indexes.SectionsByName["Peer"] {
		_, property := section.FindFirstPropertyByName("AllowedIPs")

		if property == nil {
			continue
		}

		ipAddress, err := netip.ParsePrefix(property.Value.Value)

		if err != nil {
			// This should not happen...
			continue
		}

		if ipContext, _ := ipHostSet.ContainsIP(ipAddress); ipContext != nil {
			ctxx := *ipContext
			definedLineRaw := ctxx.Value(lineKey)

			definedLine := definedLineRaw.(uint32)

			ctx.diagnostics = append(ctx.diagnostics, protocol.Diagnostic{
				Message:  fmt.Sprintf("This IP range is already covered on line %d", definedLine+1),
				Severity: &common.SeverityError,
				Range:    property.Value.ToLSPRange(),
			})
		} else {
			ipContext := context.WithValue(
				context.Background(),
				lineKey,
				property.Start.Line,
			)

			ipHostSet.AddIP(
				ipAddress,
				ipContext,
			)
		}
	}
}
