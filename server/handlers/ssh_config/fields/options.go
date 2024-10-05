package fields

var AllowedDuplicateOptions = map[NormalizedOptionName]struct{}{
	"certificatefile": {},
	"match":           {},
	"host":            {},
}

// A list of
// <Option name> -> <List of fields that need to be present for the option>
var DependentFields = map[NormalizedOptionName][]NormalizedOptionName{
	"canonicaldomains": {"canonicalizehostname"},
	"controlpersist":   {"controlmaster"},
}

var HostDisallowedOptions = map[NormalizedOptionName]struct{}{
	"enablesshkeysign": {},
}

var GlobalDisallowedOptions = map[NormalizedOptionName]struct{}{}
