package fields

import docvalues "config-lsp/doc-values"

var MatchAllowedOptions = map[string]struct{}{
	"AcceptEnv":                       {},
	"AllowAgentForwarding":            {},
	"AllowGroups":                     {},
	"AllowStreamLocalForwarding":      {},
	"AllowTcpForwarding":              {},
	"AllowUsers":                      {},
	"AuthenticationMethods":           {},
	"AuthorizedKeysCommand":           {},
	"AuthorizedKeysCommandUser":       {},
	"AuthorizedKeysFile":              {},
	"AuthorizedPrincipalsCommand":     {},
	"AuthorizedPrincipalsCommandUser": {},
	"AuthorizedPrincipalsFile":        {},
	"Banner":                          {},
	"CASignatureAlgorithms":           {},
	"ChannelTimeout":                  {},
	"ChrootDirectory":                 {},
	"ClientAliveCountMax":             {},
	"ClientAliveInterval":             {},
	"DenyGroups":                      {},
	"DenyUsers":                       {},
	"DisableForwarding":               {},
	"ExposeAuthInfo":                  {},
	"ForceCommand":                    {},
	"GatewayPorts":                    {},
	"GSSAPIAuthentication":            {},
	"HostbasedAcceptedAlgorithms":     {},
	"HostbasedAuthentication":         {},
	"HostbasedUsesNameFromPacketOnly": {},
	"IgnoreRhosts":                    {},
	"Include":                         {},
	"IPQoS":                           {},
	"KbdInteractiveAuthentication":    {},
	"KerberosAuthentication":          {},
	"LogLevel":                        {},
	"MaxAuthTries":                    {},
	"MaxSessions":                     {},
	"PasswordAuthentication":          {},
	"PermitEmptyPasswords":            {},
	"PermitListen":                    {},
	"PermitOpen":                      {},
	"PermitRootLogin":                 {},
	"PermitTTY":                       {},
	"PermitTunnel":                    {},
	"PermitUserRC":                    {},
	"PubkeyAcceptedAlgorithms":        {},
	"PubkeyAuthentication":            {},
	"PubkeyAuthOptions":               {},
	"RekeyLimit":                      {},
	"RevokedKeys":                     {},
	"RDomain":                         {},
	"SetEnv":                          {},
	"StreamLocalBindMask":             {},
	"StreamLocalBindUnlink":           {},
	"TrustedUserCAKeys":               {},
	"UnusedConnectionTimeout":         {},
	"X11DisplayOffset":                {},
	"X11Forwarding":                   {},
	"X11UseLocalhos":                  {},
}

var MatchUserField = docvalues.UserValue("", false)
var MatchGroupField = docvalues.GroupValue("", false)
var MatchHostField = docvalues.DomainValue()
var MatchLocalAddressField = docvalues.StringValue{}
var MatchLocalPortField = docvalues.StringValue{}
var MatchRDomainField = docvalues.StringValue{}
var MatchAddressField = docvalues.IPAddressValue{
	AllowIPv4:  true,
	AllowIPv6:  true,
	AllowRange: true,
}
