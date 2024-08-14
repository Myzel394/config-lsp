// Documentation taken from https://github.com/pirate/wireguard-docs
package wireguard

import (
	docvalues "config-lsp/doc-values"
)

var HeaderField = docvalues.EnumValue{
	EnforceValues: true,
	Values: []docvalues.EnumString{
		docvalues.CreateEnumStringWithDoc(
			"Interface",
			"Defines the VPN settings for the local node.",
		),
		docvalues.CreateEnumStringWithDoc(
			"Peer",
			`Defines the VPN settings for a remote peer capable of routing traffic for one or more addresses (itself and/or other peers). Peers can be either a public bounce server that relays traffic to other peers, or a directly accessible client via LAN/internet that is not behind a NAT and only routes traffic for itself.

All clients must be defined as peers on the public bounce server. Simple clients that only route traffic for themselves, only need to define peers for the public relay, and any other nodes directly accessible. Nodes that are behind separate NATs should not be defined as peers outside of the public server config, as no direct route is available between separate NATs. Instead, nodes behind NATs should only define the public relay servers and other public clients as their peers, and should specify AllowedIPs = 192.0.2.1/24 on the public server that accept routes and bounce traffic for the VPN subnet to the remote NAT-ed peers.

In summary, all nodes must be defined on the main bounce server. On client servers, only peers that are directly accessible from a node should be defined as peers of that node, any peers that must be relayed by a bounce server should be left out and will be handled by the relay server's catchall route.`,
		),
	},
}

var minPortValue = 1
var maxPortValue = 65535

// https://www.rfc-editor.org/rfc/rfc791
var minMTUValue = 68
var maxMTUValue = 1500

var interfaceOptions map[docvalues.EnumString]docvalues.Value = map[docvalues.EnumString]docvalues.Value{
	docvalues.CreateEnumStringWithDoc(
		"Address",
		`Defines what address range the local node should route traffic for. Depending on whether the node is a simple client joining the VPN subnet, or a bounce server that's relaying traffic between multiple clients, this can be set to a single IP of the node itself (specified with CIDR notation), e.g. 192.0.2.3/32), or a range of IPv4/IPv6 subnets that the node can route traffic for.

## Examples
Node is a client that only routes traffic for itself

	Address = 192.0.2.3/32

Node is a public bounce server that can relay traffic to other peers
When the node is acting as the public bounce server, it should set this to be the entire subnet that it can route traffic, not just a single IP for itself.

	Address = 192.0.2.1/24

You can also specify multiple subnets or IPv6 subnets like so:

	Address = 192.0.2.1/24,2001:DB8::/64
`,
	): docvalues.IPAddressValue{
		AllowIPv4: true,
		AllowIPv6: true,
	},
	docvalues.CreateEnumStringWithDoc(
		"ListenPort",
		`When the node is acting as a public bounce server, it should hardcode a port to listen for incoming VPN connections from the public internet. Clients not acting as relays should not set this value.

## Examples
Using default WireGuard port

	ListenPort = 51820

Using custom WireGuard port

	ListenPort = 7000
`): docvalues.NumberValue{
		Min: &minPortValue,
		Max: &maxPortValue,
	},
	docvalues.CreateEnumStringWithDoc(
		"PrivateKey",
		`This is the private key for the local node, never shared with other servers. All nodes must have a private key set, regardless of whether they are public bounce servers relaying traffic, or simple clients joining the VPN.

This key can be generated with [wg genkey > example.key]
`,
	): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"DNS",
		`The DNS server(s) to announce to VPN clients via DHCP, most clients will use this server for DNS requests over the VPN, but clients can also override this value locally on their nodes

The value can be left unconfigured to use the system's default DNS servers

## Examples
A single DNS server can be provided

	DNS = 9.9.9.9

or multiple DNS servers can be provided

	DNS = 9.9.9.9,1.1.1.1,8.8.8.8
`,
	): docvalues.ArrayValue{
		Separator:           ",",
		DuplicatesExtractor: &docvalues.SimpleDuplicatesExtractor,
		SubValue: docvalues.IPAddressValue{
			AllowIPv4:  true,
			AllowIPv6:  true,
			AllowRange: false,
		},
	},
	docvalues.CreateEnumStringWithDoc(
		"Table",
		`Optionally defines which routing table to use for the WireGuard routes, not necessary to configure for most setups.

There are two special values: ‘off’ disables the creation of routes altogether, and ‘auto’ (the default) adds routes to the default table and enables special handling of default routes.

https://git.zx2c4.com/WireGuard/about/src/tools/man/wg-quick.8

## Examples

	Table = 1234
	`): docvalues.OrValue{
		Values: []docvalues.Value{
			docvalues.EnumValue{
				EnforceValues: false,
				Values: []docvalues.EnumString{
					docvalues.CreateEnumStringWithDoc(
						"off",
						"Disable the creation of routes altogether",
					),
					docvalues.CreateEnumStringWithDoc(
						"auto",
						"Adds routes to the default table and enables special handling of default routes",
					),
				},
			},
			docvalues.StringValue{},
		},
	},
	docvalues.CreateEnumStringWithDoc(
		"MTU",
		`Optionally defines the maximum transmission unit (MTU, aka packet/frame size) to use when connecting to the peer, not necessary to configure for most setups.

The MTU is automatically determined from the endpoint addresses or the system default route, which is usually a sane choice.

https://git.zx2c4.com/WireGuard/about/src/tools/man/wg-quick.8

## Examples

	MTU = 1500
	`): docvalues.NumberValue{
		Min: &minMTUValue,
		Max: &maxMTUValue,
	},
	docvalues.CreateEnumStringWithDoc(
		"PreUp",
		`Optionally run a command before the interface is brought up. This option can be specified multiple times, with commands executed in the order they appear in the file.

## Examples

Add an IP route 

	PreUp = ip rule add ipproto tcp dport 22 table 1234
	`): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"PostUp",
		`Optionally run a command after the interface is brought up. This option can appear multiple times, as with PreUp

## Examples
Read in a config value from a file or some command's output

    PostUp = wg set %i private-key /etc/wireguard/wg0.key <(some command here)

Log a line to a file

    PostUp = echo "$(date +%s) WireGuard Started" >> /var/log/wireguard.log

Hit a webhook on another server

    PostUp = curl https://events.example.dev/wireguard/started/?key=abcdefg

Add a route to the system routing table

    PostUp = ip rule add ipproto tcp dport 22 table 1234

Add an iptables rule to enable packet forwarding on the WireGuard interface

    PostUp   = iptables -A FORWARD -i %i -j ACCEPT; iptables -A FORWARD -o %i -j ACCEPT; iptables -t nat -A POSTROUTING -o eth0 -j MASQUERADE

Force WireGuard to re-resolve IP address for peer domain

    PostUp = resolvectl domain %i "~."; resolvectl dns %i 192.0.2.1; resolvectl dnssec %i yes
	`): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"PreDown",
		`Optionally run a command before the interface is brought down. This option can appear multiple times, as with PreUp

## Examples
Log a line to a file
    
	PostDown = echo "$(date +%s) WireGuard Going Down" >> /var/log/wireguard.log

Hit a webhook on another server

    PostDown = curl https://events.example.dev/wireguard/stopping/?key=abcdefg
	`): docvalues.StringValue{},
	docvalues.CreateEnumStringWithDoc(
		"PostDown",
		`Optionally run a command after the interface is brought down. This option can appear multiple times, as with PreUp

## Examples

Log a line to a file

    PostDown = echo "$(date +%s) WireGuard Stopped" >> /var/log/wireguard.log

Hit a webhook on another server

    PostDown = curl https://events.example.dev/wireguard/stopped/?key=abcdefg

Remove the iptables rule that forwards packets on the WireGuard interface

    PostDown = iptables -D FORWARD -i %i -j ACCEPT; iptables -D FORWARD -o %i -j ACCEPT; iptables -t nat -D POSTROUTING -o eth0 -j MASQUERADE
	`): docvalues.StringValue{},
}

var interfaceAllowedDuplicateFields = map[string]struct{}{
	"PreUp":    {},
	"PostUp":   {},
	"PreDown":  {},
	"PostDown": {},
}
