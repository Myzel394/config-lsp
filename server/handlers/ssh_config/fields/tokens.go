package fields

import "config-lsp/utils"

var AvailableTokens = map[string]string{
	"%%": "A literal ‘%’.",
	"%C": "Hash of %l%h%p%r%j.",
	"%d": "Local user's home directory.",
	"%f": "The fingerprint of the server's host key.",
	"%H": "The known_hosts hostname or address that is being searched for.",
	"%h": "The remote hostname.",
	"%I": "A string describing the reason for a KnownHostsCommand execution: either ADDRESS when looking up a host by address (only when CheckHostIP is enabled), HOSTNAME when searching by hostname, or ORDER when preparing the host key algorithm preference list to use for the destination host.",
	"%i": "The local user ID.",
	"%j": "The contents of the ProxyJump option, or the empty string if this option is unset.",
	"%K": "The base64 encoded host key.",
	"%k": "The host key alias if specified, otherwise the original remote hostname given on the command line.",
	"%L": "The local hostname.",
	"%l": "The local hostname, including the domain name.",
	"%n": "The original remote hostname, as given on the command line.",
	"%p": "The remote port.",
	"%r": "The remote username.",
	"%T": "The local tun(4) or tap(4) network interface assigned if tunnel forwarding was requested, or \"NONE\" otherwise.",
	"%t": "The type of the server host key, e.g. ssh-ed25519.",
	"%u": "The local username.",
}

// A map of <option name> to <list of supported tokens>
// This is derived from the TOKENS section of ssh_config
var OptionsTokensMap = map[NormalizedOptionName][]string{
	"certificatefile":    firstTokens,
	"controlpath":        firstTokens,
	"identityagent":      firstTokens,
	"identityfile":       firstTokens,
	"include":            firstTokens,
	"localforward":       firstTokens,
	"match":              firstTokens,
	"exec":               firstTokens,
	"remotecommand":      firstTokens,
	"remoteforward":      firstTokens,
	"revokedhostkeys":    firstTokens,
	"userknownhostsfile": firstTokens,

	"knownhostscommand": append(firstTokens, []string{
		"%f", "%H", "%I", "%K", "%t",
	}...),

	"hostname": {
		"%%",
		"%h",
	},

	"localcommand": utils.KeysOfMap(AvailableTokens),

	"proxycommand": {
		"%%", "%h", "%n", "%p", "%r",
	},
}

var firstTokens = []string{
	"%%",
	"%C",
	"%d",
	"%h",
	"%i",
	"%j",
	"%k",
	"%L",
	"%l",
	"%n",
	"%p",
	"%r",
	"%u",
}
