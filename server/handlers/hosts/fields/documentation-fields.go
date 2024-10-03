package fields

import docvalues "config-lsp/doc-values"

var IPAddressField = docvalues.IPAddressValue{
	AllowIPv4: true,
	AllowIPv6: true,
}

var HostnameField = docvalues.DocumentationValue{
	Documentation: `Host names may contain only alphanumeric characters, minus signs ("-"), and periods ("."). They must begin with an alphabetic character and end with an alphanumeric character.  Optional aliases provide for name changes, alternate spellings, shorter hostnames, or generic hostnames (for example, localhost).`,
	Value:         docvalues.StringValue{},
}
