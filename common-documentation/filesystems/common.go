package commondocumentation

import docvalues "config-lsp/doc-values"

type AssignableOption struct {
	Documentation string
	Handler       func(context docvalues.KeyValueAssignmentContext) docvalues.Value
}
