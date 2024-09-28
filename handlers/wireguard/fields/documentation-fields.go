package fields

import (
	docvalues "config-lsp/doc-values"
)

// Documentation taken from https://github.com/pirate/wireguard-docs
var HeaderInterfaceEnum = docvalues.CreateEnumStringWithDoc(
	"[Interface]",
	"Defines the VPN settings for the local node.",
)
var HeaderPeerEnum = docvalues.CreateEnumStringWithDoc(
	"[Peer]",
	`Defines the VPN settings for a remote peer capable of routing traffic for one or more addresses (itself and/or other peers). Peers can be either a public bounce server that relays traffic to other peers, or a directly accessible client via LAN/internet that is not behind a NAT and only routes traffic for itself.

All clients must be defined as peers on the public bounce server. Simple clients that only route traffic for themselves, only need to define peers for the public relay, and any other nodes directly accessible. Nodes that are behind separate NATs should not be defined as peers outside of the public server config, as no direct route is available between separate NATs. Instead, nodes behind NATs should only define the public relay servers and other public clients as their peers, and should specify AllowedIPs = 192.0.2.1/24 on the public server that accept routes and bounce traffic for the VPN subnet to the remote NAT-ed peers.

In summary, all nodes must be defined on the main bounce server. On client servers, only peers that are directly accessible from a node should be defined as peers of that node, any peers that must be relayed by a bounce server should be left out and will be handled by the relay server's catchall route.`,
)

var minPortValue = 1
var maxPortValue = 65535

// https://www.rfc-editor.org/rfc/rfc791
var minMTUValue = 68
var maxMTUValue = 1500

var InterfaceOptions = map[string]docvalues.DocumentationValue{
	"Address": {
		Documentation: `Defines what address range the local node should route traffic for. Depending on whether the node is a simple client joining the VPN subnet, or a bounce server that's relaying traffic between multiple clients, this can be set to a single IP of the node itself (specified with CIDR notation), e.g. 192.0.2.3/32), or a range of IPv4/IPv6 subnets that the node can route traffic for.

## Examples
Node is a client that only routes traffic for itself

	Address = 192.0.2.3/32

Node is a public bounce server that can relay traffic to other peers
When the node is acting as the public bounce server, it should set this to be the entire subnet that it can route traffic, not just a single IP for itself.

	Address = 192.0.2.1/24

You can also specify multiple subnets or IPv6 subnets like so:

	Address = 192.0.2.1/24,2001:DB8::/64
`,
		Value: docvalues.IPAddressValue{
			AllowIPv4:  true,
			AllowIPv6:  true,
			AllowRange: true,
		},
	},
	"ListenPort": {
		Documentation: `When the node is acting as a public bounce server, it should hardcode a port to listen for incoming VPN connections from the public internet. Clients not acting as relays should not set this value. If not specified, chosen randomly.

## Examples
Using default WireGuard port

	ListenPort = 51820

Using custom WireGuard port

	ListenPort = 7000
`,
		Value: docvalues.NumberValue{
			Min: &minPortValue,
			Max: &maxPortValue,
		},
	},
	"PrivateKey": {
		Documentation: `This is the private key for the local node, never shared with other servers. All nodes must have a private key set, regardless of whether they are public bounce servers relaying traffic, or simple clients joining the VPN.

This key can be generated with [wg genkey > example.key]
`,
		Value: docvalues.StringValue{},
	},
	"DNS": {
		Documentation: `The DNS server(s) to announce to VPN clients via DHCP, most clients will use this server for DNS requests over the VPN, but clients can also override this value locally on their nodes

The value can be left unconfigured to use the system's default DNS servers

## Examples
A single DNS server can be provided

	DNS = 9.9.9.9

or multiple DNS servers can be provided

	DNS = 9.9.9.9,1.1.1.1,8.8.8.8
`,
		Value: docvalues.ArrayValue{
			Separator:           ",",
			DuplicatesExtractor: &docvalues.SimpleDuplicatesExtractor,
			SubValue: docvalues.IPAddressValue{
				AllowIPv4:  true,
				AllowIPv6:  true,
				AllowRange: false,
			},
		},
	},
	"Table": {
		Documentation: `Optionally defines which routing table to use for the WireGuard routes, not necessary to configure for most setups.

There are two special values: ‘off’ disables the creation of routes altogether, and ‘auto’ (the default) adds routes to the default table and enables special handling of default routes.

https://git.zx2c4.com/WireGuard/about/src/tools/man/wg-quick.8

## Examples

	Table = 1234
	`,
		Value: docvalues.OrValue{
			Values: []docvalues.DeprecatedValue{
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
	},
	"MTU": {
		Documentation: `Optionally defines the maximum transmission unit (MTU, aka packet/frame size) to use when connecting to the peer, not necessary to configure for most setups.

The MTU is automatically determined from the endpoint addresses or the system default route, which is usually a sane choice.

https://git.zx2c4.com/WireGuard/about/src/tools/man/wg-quick.8

## Examples

	MTU = 1500
	`, Value: docvalues.NumberValue{
			Min: &minMTUValue,
			Max: &maxMTUValue,
		},
	},
	"PreUp": {
		Documentation: `Optionally run a command before the interface is brought up. This option can be specified multiple times, with commands executed in the order they appear in the file.

## Examples

Add an IP route 

	PreUp = ip rule add ipproto tcp dport 22 table 1234
	`, Value: docvalues.StringValue{},
	},
	"PostUp": {
		Documentation: `Optionally run a command after the interface is brought up. This option can appear multiple times, as with PreUp

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
	`,
		Value: docvalues.StringValue{},
	},
	"PreDown": {
		Documentation: `Optionally run a command before the interface is brought down. This option can appear multiple times, as with PreUp

## Examples
Log a line to a file
    
	PostDown = echo "$(date +%s) WireGuard Going Down" >> /var/log/wireguard.log

Hit a webhook on another server

    PostDown = curl https://events.example.dev/wireguard/stopping/?key=abcdefg
	`,
		Value: docvalues.StringValue{},
	},
	"PostDown": {
		Documentation: `Optionally run a command after the interface is brought down. This option can appear multiple times, as with PreUp

## Examples

Log a line to a file

    PostDown = echo "$(date +%s) WireGuard Stopped" >> /var/log/wireguard.log

Hit a webhook on another server

    PostDown = curl https://events.example.dev/wireguard/stopped/?key=abcdefg

Remove the iptables rule that forwards packets on the WireGuard interface

    PostDown = iptables -D FORWARD -i %i -j ACCEPT; iptables -D FORWARD -o %i -j ACCEPT; iptables -t nat -D POSTROUTING -o eth0 -j MASQUERADE
	`,
		Value: docvalues.StringValue{},
	},
	"FwMark": {
		Documentation: "a 32-bit fwmark for outgoing packets. If set to 0 or \"off\", this option is disabled. May be specified in hexadecimal by prepending \"0x\". Optional",
		Value:         docvalues.StringValue{},
	},
}

var InterfaceAllowedDuplicateFields = map[string]struct{}{
	"PreUp":    {},
	"PostUp":   {},
	"PreDown":  {},
	"PostDown": {},
}

var PeerOptions = map[string]docvalues.DocumentationValue{
	"Endpoint": {
		Documentation: `Defines the publicly accessible address for a remote peer. This should be left out for peers behind a NAT or peers that don't have a stable publicly accessible IP:PORT pair. Typically, this only needs to be defined on the main bounce server, but it can also be defined on other public nodes with stable IPs like public-server2 in the example config below.

## Examples
Endpoint is an IP address

	[Endpoint = 123.124.125.126:51820] (IPv6 is also supported)

Endpoint is a hostname/FQDN

	 Endpoint = public-server1.example-vpn.tld:51820
	 `,
		Value: docvalues.StringValue{},
	},
	"AllowedIPs": {
		Documentation: `This defines the IP ranges for which a peer will route traffic. On simple clients, this is usually a single address (the VPN address of the simple client itself). For bounce servers this will be a range of the IPs or subnets that the relay server is capable of routing traffic for. Multiple IPs and subnets may be specified using comma-separated IPv4 or IPv6 CIDR notation (from a single /32 or /128 address, all the way up to 0.0.0.0/0 and ::/0 to indicate a default route to send all internet and VPN traffic through that peer). This option may be specified multiple times.

When deciding how to route a packet, the system chooses the most specific route first, and falls back to broader routes. So for a packet destined to 192.0.2.3, the system would first look for a peer advertising 192.0.2.3/32 specifically, and would fall back to a peer advertising 192.0.2.1/24 or a larger range like 0.0.0.0/0 as a last resort.

## Examples

Peer is a simple client that only accepts traffic to/from itself

    AllowedIPs = 192.0.2.3/32

Peer is a relay server that can bounce VPN traffic to all other peers

    AllowedIPs = 192.0.2.1/24

Peer is a relay server that bounces all internet & VPN traffic (like a proxy), including IPv6

    AllowedIPs = 0.0.0.0/0,::/0

Peer is a relay server that routes to itself and only one other peer

    AllowedIPs = 192.0.2.3/32,192.0.2.4/32

Peer is a relay server that routes to itself and all nodes on its local LAN

    AllowedIPs = 192.0.2.3/32,192.168.1.1/24
`,
		Value: docvalues.ArrayValue{
			Separator:           ",",
			DuplicatesExtractor: &docvalues.SimpleDuplicatesExtractor,
			SubValue: docvalues.IPAddressValue{
				AllowIPv4:  true,
				AllowIPv6:  true,
				AllowRange: true,
			},
		},
	},
	"PublicKey": {
		Documentation: `This is the public key for the remote node, shareable with all peers. All nodes must have a public key set, regardless of whether they are public bounce servers relaying traffic, or simple clients joining the VPN.

This key can be generated with wg pubkey < example.key > example.key.pub. (see above for how to generate the private key example.key)

## Examples

	PublicKey = somePublicKeyAbcdAbcdAbcdAbcd=
	`,
		Value: docvalues.StringValue{},
	},
	"PersistentKeepalive": {
		Documentation: `If the connection is going from a NAT-ed peer to a public peer, the node behind the NAT must regularly send an outgoing ping in order to keep the bidirectional connection alive in the NAT router's connection table.

## Examples

Local public node to remote public node

    This value should be left undefined as persistent pings are not needed.

Local public node to remote NAT-ed node

    This value should be left undefined as it's the client's responsibility to keep the connection alive because the server cannot reopen a dead connection to the client if it times out.

Oocal NAT-ed node to remote public node

    PersistentKeepalive = 25 this will send a ping to every 25 seconds keeping the connection open in the local NAT router's connection table.
`,
		Value: docvalues.PositiveNumberValue(),
	},
	"PresharedKey": {
		Documentation: "Optionally defines a pre-shared key for the peer, used to authenticate the connection. This is not necessary, but strongly recommended for security.",
		Value:         docvalues.StringValue{},
	},
}

var PeerAllowedDuplicateFields = map[string]struct{}{
	"AllowedIPs": {},
}

var OptionsHeaderMap = map[string](map[string]docvalues.DocumentationValue){
	"Interface": InterfaceOptions,
	"Peer":      PeerOptions,
}
