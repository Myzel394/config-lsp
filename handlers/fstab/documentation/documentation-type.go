package fstabdocumentation

import (
	docvalues "config-lsp/doc-values"
	"config-lsp/utils"
)

var FileSystemTypeField = docvalues.ArrayValue{
	Separator:           ",",
	DuplicatesExtractor: &docvalues.SimpleDuplicatesExtractor,
	SubValue: docvalues.EnumValue{
		EnforceValues: false,
		Values: utils.Map(
			utils.KeysOfMap(MountOptionsMapField),
			func(key string) docvalues.EnumString {
				return docvalues.CreateEnumString(key)
			},
		),
	},
}
