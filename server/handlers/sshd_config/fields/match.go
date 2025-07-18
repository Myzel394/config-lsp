package fields

import (
	docvalues "config-lsp/doc-values"
	matchparser "config-lsp/handlers/sshd_config/match-parser"
)

var MatchAllowedOptions = map[NormalizedOptionName]struct{}{
	"acceptenv":                       {},
	"allowagentforwarding":            {},
	"allowgroups":                     {},
	"allowstreamlocalforwarding":      {},
	"allowtcpforwarding":              {},
	"allowusers":                      {},
	"authenticationmethods":           {},
	"authorizedkeyscommand":           {},
	"authorizedkeyscommanduser":       {},
	"authorizedkeysfile":              {},
	"authorizedprincipalscommand":     {},
	"authorizedprincipalscommanduser": {},
	"authorizedprincipalsfile":        {},
	"banner":                          {},
	"casignaturealgorithms":           {},
	"channeltimeout":                  {},
	"chrootdirectory":                 {},
	"clientalivecountmax":             {},
	"clientaliveinterval":             {},
	"denygroups":                      {},
	"denyusers":                       {},
	"disableforwarding":               {},
	"exposeauthinfo":                  {},
	"forcecommand":                    {},
	"gatewayports":                    {},
	"gssapiauthentication":            {},
	"hostbasedacceptedalgorithms":     {},
	"hostbasedauthentication":         {},
	"hostbasedusesnamefrompacketonly": {},
	"ignorerhosts":                    {},
	"include":                         {},
	"ipqos":                           {},
	"kbdinteractiveauthentication":    {},
	"kerberosauthentication":          {},
	"loglevel":                        {},
	"maxauthtries":                    {},
	"maxsessions":                     {},
	"passwordauthentication":          {},
	"permitemptypasswords":            {},
	"permitlisten":                    {},
	"permitopen":                      {},
	"permitrootlogin":                 {},
	"permittty":                       {},
	"permittunnel":                    {},
	"permituserrc":                    {},
	"pubkeyacceptedalgorithms":        {},
	"pubkeyauthentication":            {},
	"pubkeyauthoptions":               {},
	"rekeylimit":                      {},
	"revokedkeys":                     {},
	"rdomain":                         {},
	"setenv":                          {},
	"streamlocalbindmask":             {},
	"streamlocalbindunlink":           {},
	"trustedusercakeys":               {},
	"unusedconnectiontimeout":         {},
	"x11displayoffset":                {},
	"x11forwarding":                   {},
	"x11uselocalhos":                  {},
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
	AllowRange: docvalues.AllowedStatusRequired,
}

var MatchValueFieldMap = map[matchparser.MatchCriteriaType]docvalues.DeprecatedValue{
	matchparser.MatchCriteriaTypeUser:         MatchUserField,
	matchparser.MatchCriteriaTypeGroup:        MatchGroupField,
	matchparser.MatchCriteriaTypeHost:         MatchHostField,
	matchparser.MatchCriteriaTypeLocalAddress: MatchLocalAddressField,
	matchparser.MatchCriteriaTypeLocalPort:    MatchLocalPortField,
	matchparser.MatchCriteriaTypeRDomain:      MatchRDomainField,
	matchparser.MatchCriteriaTypeAddress:      MatchAddressField,
}
