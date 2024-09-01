package openssh

import (
	docvalues "config-lsp/doc-values"
	"strings"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

type Option struct {
	Documentation string
	Value         docvalues.Value
}

func GetDocumentation(o *Option) protocol.MarkupContent {
	typeDescription := strings.Join(o.Value.GetTypeDescription(), "\n")

	return protocol.MarkupContent{
		Kind:  protocol.MarkupKindPlainText,
		Value: "### Type\n" + typeDescription + "\n\n---\n\n### Documentation\n" + o.Documentation,
	}
}

func NewOption(documentation string, value docvalues.Value) Option {
	return Option{documentation, value}
}
