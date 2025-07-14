package docvalues

import (
	"config-lsp/common"
	"config-lsp/utils"
	"errors"
	"fmt"
	"strings"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

type Suffix struct {
	Suffix  string
	Meaning string
}

/*
Create a new SuffixValue that allows the given prefixes.
@example:
```

	SuffixValue{
		Suffixes: []Suffix{
			{Prefix: "debug_", Meaning: "Enable debug mode"},
			{Prefix: "test_", Meaning: "Enable test mode"},
		},
		SubValue: StringValue{},
	}

```
*/
type SuffixValue struct {
	Suffixes []Suffix
	// A suffix is required
	Required bool
	SubValue DeprecatedValue
}

func (v SuffixValue) GetTypeDescription() []string {
	subDescription := v.SubValue.GetTypeDescription()

	suffixDescription := utils.Map(v.Suffixes, func(suffix Suffix) string {
		return fmt.Sprintf("_%s_ -> %s", suffix.Suffix, suffix.Meaning)
	})

	return append(subDescription,
		append(
			[]string{"The following suffixes are allowed:"},
			suffixDescription...,
		)...,
	)
}

func (v SuffixValue) DeprecatedCheckIsValid(value string) []*InvalidValue {
	for _, suffix := range v.Suffixes {
		if strings.HasSuffix(value, suffix.Suffix) {
			return v.SubValue.DeprecatedCheckIsValid(value[:len(value)-len(suffix.Suffix)])
		}
	}

	if v.Required {
		return []*InvalidValue{{
			Err:   errors.New("A suffix is required"),
			Start: 0,
			End:   uint32(len(value)),
		}}
	}

	return v.SubValue.DeprecatedCheckIsValid(value)
}

func (v SuffixValue) FetchCompletions(value string, cursor common.CursorPosition) []protocol.CompletionItem {
	textFormat := protocol.InsertTextFormatPlainText
	kind := protocol.CompletionItemKindText

	suffixCompletions := utils.Map(v.Suffixes, func(suffix Suffix) protocol.CompletionItem {
		return protocol.CompletionItem{
			Label:            suffix.Suffix,
			Detail:           &suffix.Meaning,
			InsertTextFormat: &textFormat,
			Kind:             &kind,
		}
	})

	return append(suffixCompletions, v.SubValue.FetchCompletions(value, cursor)...)
}

func (v SuffixValue) DeprecatedFetchHoverInfo(line string, cursor uint32) []string {
	for _, suffix := range v.Suffixes {
		if strings.HasSuffix(line, suffix.Suffix) {
			return append([]string{
				fmt.Sprintf("Suffix: _%s_ -> %s", suffix.Suffix, suffix.Meaning),
			},
				v.SubValue.DeprecatedFetchHoverInfo(line[:len(line)-len(suffix.Suffix)], cursor)...,
			)
		}
	}

	return v.SubValue.DeprecatedFetchHoverInfo(line, cursor)
}
