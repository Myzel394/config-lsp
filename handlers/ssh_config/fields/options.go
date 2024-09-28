package fields

var AllowedDuplicateOptions = map[string]struct{}{
	"CertificateFile": {},
	"Match": {},
	"Host": {},
}

// A list of 
// <Option name> -> <List of fields that need to be present for the option>
var DependentFields = map[string][]string{
	"CanonicalDomains": {"CanonicalizeHostname"},
	"ControlPersist": {"ControlMaster"},
}

var HostDisallowedOptions = map[string]struct{}{
	"EnableSSHKeysign": {},
}

var GlobalDisallowedOptions = map[string]struct{}{}

