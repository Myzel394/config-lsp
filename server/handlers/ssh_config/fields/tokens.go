package fields

var AvailableTokens = []string{
	"%%",
	"%C",
	"%d",
	"%f",
	"%H",
	"%h",
	"%l",
	"%i",
	"%j",
	"%K",
	"%k",
	"%L",
	"%l",
	"%n",
	"%p",
	"%r",
	"%T",
	"%t",
	"%u",
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

	"localcommand": AvailableTokens,

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
