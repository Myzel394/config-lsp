package fields

import (
	commondocumentation "config-lsp/common-documentation"
	docvalues "config-lsp/doc-values"
)

var UserField = docvalues.UserValue("", false)

var PathField = docvalues.PathValue{
	RequiredType: docvalues.PathTypeFile,
}

var CommandField = docvalues.StringValue{}

var EmailField = docvalues.RegexValue{
	Regex: *commondocumentation.EmailRegex,
}
