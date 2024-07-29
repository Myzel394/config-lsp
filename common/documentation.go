package common

import (
	"fmt"
	"strconv"
	"strings"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

type Value interface {
	GetTypeDescription() []string
	CheckIsValid(value string) error
}

type EnumValue struct {
	Values []string
	// If `true`, the value MUST be one of the values in the Values array
	// Otherwise an error is shown
	// If `false`, the value is just a hint
	EnforceValues bool
}

func (v EnumValue) GetTypeDescription() []string {
	if len(v.Values) == 1 {
		return []string{fmt.Sprintf("'%s'", v.Values[0])}
	}

	lines := make([]string, len(v.Values)+1)
	lines[0] = "Enum of:"

	for index, value := range v.Values {
		lines[index+1] += "\t* " + value
	}

	return lines
}
func (v EnumValue) CheckIsValid(value string) error {
	if !v.EnforceValues {
		return nil
	}

	for _, validValue := range v.Values {
		if validValue == value {
			return nil
		}

	}

	return ValueNotInEnumError{
		ProvidedValue:   value,
		AvailableValues: v.Values,
	}
}

type PositiveNumberValue struct{}

func (v PositiveNumberValue) GetTypeDescription() []string {
	return []string{"Positive number"}
}
func (v PositiveNumberValue) CheckIsValid(value string) error {
	number, err := strconv.Atoi(value)

	if err != nil {
		return NotANumberError{}
	}

	if number < 0 {
		return NumberIsNotPositiveError{}
	}

	return nil
}

var SimpleDuplicatesExtractor = func(value string) string {
	return value
}

var ExtractKeyDuplicatesExtractor = func(separator string) func(string) string {
	return func(value string) string {
		splitted := strings.Split(value, separator)

		if len(splitted) == 0 {
			return ""
		}

		return splitted[0]
	}
}

var DuplicatesAllowedExtractor func(string) string = nil

type ArrayValue struct {
	SubValue        Value
	Separator       string
	// If this function is nil, no duplicate check is done
	// (value) => Extracted value
	// This is used to extract the value from the user input,
	// because you may want to preprocess the value before checking for duplicates
	DuplicatesExtractor *(func(string) string)
}

func (v ArrayValue) GetTypeDescription() []string {
	subValue := v.SubValue.(Value)

	return append(
		[]string{fmt.Sprintf("An Array separated by '%s' of:", v.Separator)},
		subValue.GetTypeDescription()...,
	)
}
func (v ArrayValue) CheckIsValid(value string) error {
	values := strings.Split(value, v.Separator)

	if v.DuplicatesExtractor != nil {
		valuesOccurrences := SliceToMap(
			Map(values, *v.DuplicatesExtractor),
			0,
		)

		// Only continue if there are actually duplicate values
		if len(values) != len(valuesOccurrences) {
			for _, duplicateRawValue := range values {
				duplicateValue := (*v.DuplicatesExtractor)(duplicateRawValue)
				valuesOccurrences[duplicateValue]++
			}

			duplicateValues := FilterMapWhere(valuesOccurrences, func(_ string, value int) bool {
				return value > 1
			})

			return ArrayContainsDuplicatesError{
				Duplicates: KeysOfMap(duplicateValues),
			}
		}
	}


	for _, subValue := range values {
		err := v.SubValue.CheckIsValid(subValue)

		if err != nil {
			return err
		}
	}

	return nil
}

type OrValue struct {
	Values []Value
}

func (v OrValue) GetTypeDescription() []string {
	lines := make([]string, 0)

	for _, subValueRaw := range v.Values {
		subValue := subValueRaw.(Value)
		subLines := subValue.GetTypeDescription()

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
		[]string{"One of:"},
		lines...,
	)
}
func (v OrValue) CheckIsValid(value string) error {
	var firstError error = nil

	for _, subValue := range v.Values {
		err := subValue.CheckIsValid(value)

		if err == nil {
			return nil
		} else if firstError == nil {
			firstError = err
		}
	}

	return firstError
}

type StringValue struct{}

func (v StringValue) GetTypeDescription() []string {
	return []string{"String"}
}

func (v StringValue) CheckIsValid(value string) error {
	if value == "" {
		return EmptyStringError{}
	}

	return nil
}

type CustomValue struct {
	FetchValue func() Value
}

func (v CustomValue) GetTypeDescription() []string {
	return []string{"Custom"}
}

func (v CustomValue) CheckIsValid(value string) error {
	return v.FetchValue().CheckIsValid(value)
}

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

	prefixDescription := Map(v.Prefixes, func(prefix Prefix) string {
		return fmt.Sprintf("_%s_ -> %s", prefix.Prefix, prefix.Meaning)
	})

	return append(subDescription,
		append(
			[]string{"The following prefixes are allowed:"},
			prefixDescription...,
		)...,
	)
}

func (v PrefixWithMeaningValue) CheckIsValid(value string) error {
	return v.SubValue.CheckIsValid(value)
}

type PathType uint8

const (
	PathTypeExistenceOptional PathType = 0
	PathTypeFile              PathType = 1
	PathTypeDirectory         PathType = 2
)

type PathValue struct {
	RequiredType PathType
}

func (v PathValue) GetTypeDescription() []string {
	hints := make([]string, 0)

	switch v.RequiredType {
	case PathTypeExistenceOptional:
		hints = append(hints, "Optional")
		break
	case PathTypeFile:
		hints = append(hints, "File")
	case PathTypeDirectory:
		hints = append(hints, "Directory")
	}

	return []string{strings.Join(hints, ", ")}
}

func (v PathValue) CheckIsValid(value string) error {
	if !DoesPathExist(value) {
		return PathDoesNotExistError{}
	}

	isValid := false

	if (v.RequiredType & PathTypeFile) == PathTypeFile {
		isValid = isValid && IsPathFile(value)
	}

	if (v.RequiredType & PathTypeDirectory) == PathTypeDirectory {
		isValid = isValid && IsPathDirectory(value)
	}

	if isValid {
		return nil
	}

	return PathInvalidError{}
}

type KeyValueAssignmentValue struct {
	Key Value
	Value Value
	Separator string
}

func (v KeyValueAssignmentValue) GetTypeDescription() []string {
	return []string{
		fmt.Sprintf("Key-Value pair in form of 'key%svalue'", v.Separator),
		fmt.Sprintf("#### Key\n%s", strings.Join(v.Key.GetTypeDescription(), "\n")),
		fmt.Sprintf("#### Value:\n%s", strings.Join(v.Value.GetTypeDescription(), "\n")),
	}
}
func (v KeyValueAssignmentValue) CheckIsValid(value string) error {
	parts := strings.Split(value, v.Separator)

	if len(parts) == 1 && parts[0] == "" {
		return nil
	}

	if len(parts) != 2 {
		return KeyValueAssignmentError{}
	}

	err := v.Key.CheckIsValid(parts[0])

	if err != nil {
		return err
	}

	err = v.Value.CheckIsValid(parts[1])

	if err != nil {
		return err
	}

	return nil
}


type Option struct {
	Documentation string
	Value         Value
}

func GetDocumentation(o *Option) protocol.MarkupContent {
	typeDescription := strings.Join(o.Value.GetTypeDescription(), "\n")

	return protocol.MarkupContent{
		Kind:  protocol.MarkupKindPlainText,
		Value: "### Type\n" + typeDescription + "\n\n---\n\n### Documentation\n" + o.Documentation,
	}
}

func NewOption(documentation string, value Value) Option {
	return Option{documentation, value}
}
