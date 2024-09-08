package fields

import (
	commondocumentation "config-lsp/common-documentation"
	docvalues "config-lsp/doc-values"
)

var UserField = docvalues.DocumentationValue{
	Documentation: "A user on the host machine. The user must have a valid entry in the passwd(5) database file.",
	Value:         docvalues.UserValue("", false),
}

var UserDeclaration = "`user`"

var PathField = docvalues.DocumentationValue{
	Documentation: "Append messages to file, specified by its absolute pathname",
	Value: docvalues.PathValue{
		RequiredType: docvalues.PathTypeFile,
	},
}

var PathDeclaration = "`/path/to/file`"

var CommandField = docvalues.DocumentationValue{
	Documentation: "Pipe the message to command on its standard input. The command is run under the privileges of the daemon's unprivileged account.",
	Value:         docvalues.StringValue{},
}

var CommandDeclaration = "`|command`"

var EmailField = docvalues.DocumentationValue{
	Documentation: "An email address in RFC 5322 format. If an address extension is appended to the user-part, it is first compared for an exact match.  It is then stripped so that an address such as user+ext@example.com will only use the part that precedes ‘+’ as a key.",
	Value: docvalues.RegexValue{
		Regex: *commondocumentation.EmailRegex,
	},
}

var EmailDeclaration = "`user-part@domain-part`"

var IncludeField = docvalues.DocumentationValue{
	Documentation: "Include any definitions in file as alias entries. The format of the file is identical to this one.",
	Value: docvalues.PathValue{
		RequiredType: docvalues.PathTypeFile,
	},
}

var IncludeDeclaration = "`include:/path/to/file`"

var ErrorMessageField = docvalues.DocumentationValue{
	Documentation: "A status code and message to return. The code must be 3 digits, starting 4XX (TempFail) or 5XX (PermFail). The message must be present and can be freely chosen.",
	Value:         docvalues.StringValue{},
}

var ErrorDeclaration = "`error:code message`"
