package docvalues

import (
	"config-lsp/common"
	"config-lsp/utils"
	"errors"
	"fmt"
	"strings"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

type Prefix struct {
	Prefix  string
	Meaning string
}

/*
Create a new PrefixValue that allows the given prefixes.
@example:
```

	PrefixValue{
		Prefixes: []Prefix{
			{Prefix: "debug_", Meaning: "Enable debug mode"},
			{Prefix: "test_", Meaning: "Enable test mode"},
		},
		SubValue: StringValue{},
	}

```
*/
type PrefixValue struct {
	Prefixes []Prefix
	// A prefix is required
	Required bool
	SubValue DeprecatedValue
}

func (v PrefixValue) GetTypeDescription() []string {
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

func (v PrefixValue) DeprecatedCheckIsValid(value string) []*InvalidValue {
	for _, prefix := range v.Prefixes {
		if strings.HasPrefix(value, prefix.Prefix) {
			return v.SubValue.DeprecatedCheckIsValid(value[len(prefix.Prefix):])
		}
	}

	if v.Required {
		return []*InvalidValue{{
			Err:   errors.New("A prefix is required"),
			Start: 0,
			End:   uint32(len(value)),
		}}
	}

	return v.SubValue.DeprecatedCheckIsValid(value)
}

func (v PrefixValue) FetchCompletions(value string, cursor common.CursorPosition) []protocol.CompletionItem {
	textFormat := protocol.InsertTextFormatPlainText
	kind := protocol.CompletionItemKindText

	// Check if the line starts with a prefix
	startsWithPrefix := false
	for _, prefix := range v.Prefixes {
		if strings.HasPrefix(value, prefix.Prefix) {
			startsWithPrefix = true
			break
		}
	}

	var prefixCompletions []protocol.CompletionItem
	if !startsWithPrefix {
		prefixCompletions = utils.Map(v.Prefixes, func(prefix Prefix) protocol.CompletionItem {
			return protocol.CompletionItem{
				Label:            prefix.Prefix,
				Detail:           &prefix.Meaning,
				InsertTextFormat: &textFormat,
				Kind:             &kind,
			}
		})
	}

	return append(prefixCompletions, v.SubValue.FetchCompletions(value, cursor)...)
}

func (v PrefixValue) DeprecatedFetchHoverInfo(line string, cursor uint32) []string {
	for _, prefix := range v.Prefixes {
		if strings.HasPrefix(line, prefix.Prefix) {
			return append([]string{
				fmt.Sprintf("Prefix: _%s_ -> %s", prefix.Prefix, prefix.Meaning),
			},
				v.SubValue.DeprecatedFetchHoverInfo(line[1:], cursor)...,
			)
		}
	}

	return v.SubValue.DeprecatedFetchHoverInfo(line[1:], cursor)
}
