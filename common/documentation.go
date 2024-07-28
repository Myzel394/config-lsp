package common

import (
	"fmt"
	"strings"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

type Value interface {
	getTypeDescription() []string
}

type EnumValue struct {
	Values []string
}
func (v EnumValue) getTypeDescription() []string {
	lines := make([]string, len(v.Values) + 1)
	lines[0] = "Enum of:"

	for index, value := range v.Values {
		lines[index + 1] += "\t* " + value
	}

	return lines
}

type PositiveNumberValue struct {}
func (v PositiveNumberValue) getTypeDescription() []string {
	return []string{ "Positive number" }
}

type ArrayValue struct {
	SubValue Value
	Separator string
	AllowDuplicates bool
}
func (v ArrayValue) getTypeDescription() []string {
	subValue := v.SubValue.(Value)

	return append(
		[]string{ "An Array separated by " + v.Separator + " of:" },
		subValue.getTypeDescription()...
	)
}

type OrValue struct {
	Values []Value
}
func (v OrValue) getTypeDescription() []string {
	lines := make([]string, 0)

	for _, subValueRaw := range v.Values {
		subValue := subValueRaw.(Value)
		subLines := subValue.getTypeDescription()

		for index, line := range subLines {
			if strings.HasPrefix(line, "\t*") {
				subLines[index] = "\t" + line
			} else {
				subLines[index] = "\t* " + line
			}
		}

		lines = append(lines, subLines...)
	}

	return append(
		[]string{ "One of:" },
		lines...
	)
}

type StringValue struct {}
func (v StringValue) getTypeDescription() []string {
	return []string{ "String" }
}

type CustomValue struct {
	FetchValue func() Value
}
func (v CustomValue) getTypeDescription() []string {
	return []string{ "Custom" }
}

type Prefix struct {
	Prefix string
	Meaning string
}
type PrefixWithMeaningValue struct {
	Prefixes []Prefix
	SubValue Value
}
func (v PrefixWithMeaningValue) getTypeDescription() []string {
	subDescription := v.SubValue.getTypeDescription()

	prefixDescription := Map(v.Prefixes, func(prefix Prefix) string {
		return fmt.Sprintf("_%s_ -> %s", prefix.Prefix, prefix.Meaning)
	})

	return append(subDescription,
		append(
			[]string{ "The following prefixes are allowed:" },
			prefixDescription...,
		)...,
	)
}


type Option struct {
	Documentation string
	Value Value
}

func GetDocumentation(o *Option) protocol.MarkupContent {
	typeDescription := strings.Join(o.Value.getTypeDescription(), "\n")

	return protocol.MarkupContent{
		Kind: protocol.MarkupKindPlainText,
		Value: "### Type\n" + typeDescription + "\n\n---\n\n### Documentation\n" + o.Documentation,
	}
}

func NewOption(documentation string, value Value) Option {
	return Option{documentation, value}
}

