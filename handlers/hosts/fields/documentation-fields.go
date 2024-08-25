package fields

import docvalues "config-lsp/doc-values"

var IPAddressField = docvalues.IPAddressValue{
	AllowIPv4: true,
	AllowIPv6: true,
}

var HostnameField = docvalues.DomainValue()
