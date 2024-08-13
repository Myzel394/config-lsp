package docvalues

import (
	"config-lsp/utils"
	"fmt"
	"strings"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

type Prefix struct {
	Prefix  string
	Meaning string
}
type PrefixWithMeaningValue struct {
	Prefixes []Prefix
	SubValue Value
}

func (v PrefixWithMeaningValue) GetTypeDescription() []string {
	subDescription := v.SubValue.GetTypeDescription()

	prefixDescription := utils.Map(v.Prefixes, func(prefix Prefix) string {
		return fmt.Sprintf("_%s_ -> %s", prefix.Prefix, prefix.Meaning)
	})

	return append(subDescription,
		append(
			[]string{"The following prefixes are allowed:"},
			prefixDescription...,
		)...,
	)
}

func (v PrefixWithMeaningValue) CheckIsValid(value string) []*InvalidValue {
	for _, prefix := range v.Prefixes {
		if strings.HasPrefix(value, prefix.Prefix) {
			return v.SubValue.CheckIsValid(value[len(prefix.Prefix):])
		}
	}

	return v.SubValue.CheckIsValid(value)
}

func (v PrefixWithMeaningValue) FetchCompletions(line string, cursor uint32) []protocol.CompletionItem {
	textFormat := protocol.InsertTextFormatPlainText
	kind := protocol.CompletionItemKindText

	prefixCompletions := utils.Map(v.Prefixes, func(prefix Prefix) protocol.CompletionItem {
		return protocol.CompletionItem{
			Label:            prefix.Prefix,
			Detail:           &prefix.Meaning,
			InsertTextFormat: &textFormat,
			Kind:             &kind,
		}
	})

	return append(prefixCompletions, v.SubValue.FetchCompletions(line, cursor)...)
}

func (v PrefixWithMeaningValue) FetchHoverInfo(line string, cursor uint32) []string {
	for _, prefix := range v.Prefixes {
		if strings.HasPrefix(line, prefix.Prefix) {
			return append([]string{
				fmt.Sprintf("Prefix: _%s_ -> %s", prefix.Prefix, prefix.Meaning),
			},
				v.SubValue.FetchHoverInfo(line[1:], cursor)...,
			)
		}
	}

	return v.SubValue.FetchHoverInfo(line[1:], cursor)
}
