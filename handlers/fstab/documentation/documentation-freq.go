package fstabdocumentation

import docvalues "config-lsp/doc-values"

var minValue = 0
var maxValue = 9

var FreqField = docvalues.NumberValue{
	Min: &minValue,
	Max: &maxValue,
}
