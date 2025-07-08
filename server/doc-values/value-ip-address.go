package docvalues

import (
	"config-lsp/common"
	"config-lsp/utils"
	"fmt"
	net "net/netip"
	"strconv"
	"strings"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

var NonRoutableNetworks = []net.Prefix{
	net.MustParsePrefix("240.0.0.0/4"),
	net.MustParsePrefix("2001:db8::/32"),
}

type InvalidIPAddress struct{}

func (e InvalidIPAddress) Error() string {
	return "This is not a valid IP Address"
}

type IP4AddressNotAllowedError struct{}

func (e IP4AddressNotAllowedError) Error() string {
	return "IPv4 Addresses are not allowed"
}

type IP6AddressNotAllowedError struct{}

func (e IP6AddressNotAllowedError) Error() string {
	return "IPv6 Addresses are not allowed"
}

type IpRangeNotAllowedError struct{}

func (e IpRangeNotAllowedError) Error() string {
	return "IP Ranges are not allowed; Use a single IP Address instead"
}

type IPAddressNotAllowedError struct{}

func (e IPAddressNotAllowedError) Error() string {
	return "This IP Address is not allowed"
}

type IPAddressValue struct {
	AllowIPv4     bool
	AllowIPv6     bool
	AllowRange    bool
	AllowedIPs    *[]net.Prefix
	DisallowedIPs *[]net.Prefix
}

func (v IPAddressValue) GetTypeDescription() []string {
	if v.AllowedIPs != nil && len(*v.AllowedIPs) != 0 {
		return append(
			[]string{"One of the following IP Addresses (in range):"},
			utils.Map(*v.AllowedIPs, func(ip net.Prefix) string {
				return ip.String()
			})...,
		)
	}

	return []string{"An IP Address"}
}

func (v IPAddressValue) DeprecatedCheckIsValid(value string) []*InvalidValue {
	var ip net.Prefix

	if v.AllowRange {
		rawIP, err := net.ParsePrefix(value)

		if err != nil {
			return []*InvalidValue{
				{
					Err:   InvalidIPAddress{},
					Start: 0,
					End:   uint32(len(value)),
				},
			}
		}

		ip = rawIP
	} else {
		rawIP, err := net.ParseAddr(value)

		if err != nil {
			return []*InvalidValue{{
				Err:   InvalidIPAddress{},
				Start: 0,
				End:   uint32(len(value)),
			},
			}
		}

		ip = net.PrefixFrom(rawIP, 32)
	}

	if !ip.IsValid() {
		return []*InvalidValue{{
			Err:   InvalidIPAddress{},
			Start: 0,
			End:   uint32(len(value)),
		},
		}
	}

	if v.AllowedIPs != nil {
		for _, allowedIP := range *v.AllowedIPs {
			if allowedIP.Contains(ip.Addr()) {
				return nil
			}
		}

		return []*InvalidValue{{
			Err:   IPAddressNotAllowedError{},
			Start: 0,
			End:   uint32(len(value)),
		},
		}
	}

	if v.DisallowedIPs != nil {
		for _, disallowedIP := range *v.DisallowedIPs {
			if disallowedIP.Contains(ip.Addr()) {
				return []*InvalidValue{{
					Err:   IPAddressNotAllowedError{},
					Start: 0,
					End:   uint32(len(value)),
				},
				}
			}
		}
	}

	if v.AllowIPv4 && ip.Addr().Is4() {
		return nil
	}

	if v.AllowIPv6 && ip.Addr().Is6() {
		return nil
	}

	return []*InvalidValue{{
		Err:   InvalidIPAddress{},
		Start: 0,
		End:   uint32(len(value)),
	},
	}
}

func (v IPAddressValue) DeprecatedFetchCompletions(line string, cursor uint32) []protocol.CompletionItem {
	if v.AllowedIPs != nil && len(*v.AllowedIPs) != 0 {
		kind := protocol.CompletionItemKindValue

		return utils.Map(*v.AllowedIPs, func(ip net.Prefix) protocol.CompletionItem {
			return protocol.CompletionItem{
				Label: ip.Addr().String(),
				Kind:  &kind,
			}
		})
	}

	if v.AllowRange {
		slashIndex := strings.LastIndex(line, "/")

		if slashIndex > -1 && cursor >= uint32(slashIndex) {
			completions := make([]protocol.CompletionItem, 33)

			for i := 0; i < len(completions); i++ {
				kind := protocol.CompletionItemKindValue
				completions[i] = protocol.CompletionItem{
					Label: strconv.Itoa(i),
					Kind:  &kind,
				}
			}

			return completions
		}
	}

	return []protocol.CompletionItem{}
}

func (v IPAddressValue) FetchCompletions(value string, cursor common.CursorPosition) []protocol.CompletionItem {
	return v.DeprecatedFetchCompletions(
		value,
		common.DeprecatedImprovedCursorToIndex(
			cursor,
			value,
			0,
		),
	)
}

func (v IPAddressValue) DeprecatedFetchHoverInfo(line string, cursor uint32) []string {
	if v.AllowRange {
		ip, err := net.ParsePrefix(line)

		if err != nil {
			return []string{}
		}

		if ip.Addr().IsPrivate() {
			return []string{
				fmt.Sprintf("%s (Private IP Address)", ip.String()),
			}
		} else if ip.Addr().IsLoopback() {
			return []string{
				fmt.Sprintf("%s (Loopback IP Address)", ip.String()),
			}
		} else {
			return []string{
				fmt.Sprintf("%s (Public IP Address)", ip.String()),
			}
		}
	} else {
		ip, err := net.ParseAddr(line)

		if err != nil {
			return []string{}
		}

		if ip.IsPrivate() {
			return []string{
				fmt.Sprintf("%s (Private IP Address)", ip.String()),
			}
		} else if ip.IsLoopback() {
			return []string{
				fmt.Sprintf("%s (Loopback IP Address)", ip.String()),
			}
		} else {
			return []string{
				fmt.Sprintf("%s (Public IP Address)", ip.String()),
			}
		}
	}
}
