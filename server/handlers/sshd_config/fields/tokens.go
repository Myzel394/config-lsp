package fields

var AvailableTokens = map[string]string{
	"%%": "A literal ‘%’.",
	"%C": "Identifies the connection endpoints, containing four space-separated values: client address, client port number, server address, and server port number.",
	"%D": "The routing domain in which the incoming connection was received.",
	"%F": "The fingerprint of the CA key.",
	"%f": "The fingerprint of the key or certificate.",
	"%h": "The home directory of the user.",
	"%i": "The key ID in the certificate.",
	"%K": "The base64-encoded CA key.",
	"%k": "The base64-encoded key or certificate for authentication.",
	"%s": "The serial number of the certificate.",
	"%T": "The type of the CA key.",
	"%t": "The key or certificate type.",
	"%U": "The numeric user ID of the target user.",
	"%u": "The username.",
}

// A map of <option name> to <list of supported tokens>
// This is derived from the TOKENS section of ssh_config
var OptionsTokensMap = map[NormalizedOptionName][]string{
	"authorizedkeyscommand": {
		"%%",
		"%C",
		"%D",
		"%f",
		"%h",
		"%k",
		"%t",
		"%U",
		"%u",
	},
	"authorizedkeysfile": {
		"%%",
		"%h",
		"%U",
		"%u",
	},
	"authorizedprincipalscommand": {
		"%%",
		"%C",
		"%D",
		"%F",
		"%f",
		"%h",
		"%i",
		"%K",
		"%k",
		"%s",
		"%T",
		"%t",
		"%U",
		"%u",
	},
	"authorizedprincipalsfile": {
		"%%",
		"%h",
		"%U",
		"%u",
	},
	"chrootdirectory": {
		"%%",
		"%h",
		"%U",
		"%u",
	},
	"routingdomain": {
		"%D",
	},
}
