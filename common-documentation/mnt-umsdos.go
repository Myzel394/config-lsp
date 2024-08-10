package commondocumentation

import (
	docvalues "config-lsp/doc-values"
	"config-lsp/utils"
)

var UmsdosDocumentationAssignable = utils.FilterMap(MsdosDocumentationAssignable, func(key docvalues.EnumString, value docvalues.Value) bool {
	// `dotsOK` is explicitly not supported
	if key.InsertText == "dotsOK" {
		return false
	}

	return true
})

var UmsdosDocumentationEnums = MsdosDocumentationEnums
