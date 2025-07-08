package docvalues

import (
	"config-lsp/common"
	"config-lsp/utils"
	"errors"
	"fmt"
	"strings"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

type KeyEnumAssignmentValue struct {
	Values    map[EnumString]DeprecatedValue
	Separator string
	// If true, the value is optional, so this won't return an error if the value is not present
	// This is useful for cases where they key can stand on its own without a value
	ValueIsOptional bool
}

func (v KeyEnumAssignmentValue) GetTypeDescription() []string {
	if len(v.Values) == 1 {
		firstKey := utils.KeysOfMap(v.Values)[0]
		valueDescription := v.Values[firstKey].GetTypeDescription()

		if len(valueDescription) == 1 {
			return []string{
				fmt.Sprintf("Key-DeprecatedValue pair in form of '<%s>%s<%s>'", firstKey.DescriptionText, v.Separator, valueDescription[0]),
			}
		}
	}

	var result []string
	for key, value := range v.Values {
		result = append(result, key.Documentation)
		result = append(result, value.GetTypeDescription()...)
	}

	return append([]string{
		"Key-DeprecatedValue pair in form of 'key%svalue'", v.Separator,
	}, result...)
}

func (v KeyEnumAssignmentValue) getValue(findKey string) (*DeprecatedValue, bool) {
	for key, value := range v.Values {
		if key.InsertText == findKey {
			switch value.(type) {
			case CustomValue:
				customValue := value.(CustomValue)
				context := KeyValueAssignmentContext{
					SelectedKey: findKey,
				}

				fetchedValue := customValue.FetchValue(context)

				return &fetchedValue, true
			default:
				return &value, true
			}
		}
	}

	return nil, false
}

func (v KeyEnumAssignmentValue) DeprecatedCheckIsValid(value string) []*InvalidValue {
	parts := strings.Split(value, v.Separator)

	if len(parts) == 0 || parts[0] == "" {
		// Nothing to check for
		return nil
	}

	if len(parts) < 2 {
		if v.ValueIsOptional {
			return nil
		}

		return []*InvalidValue{
			{
				Err:   KeyValueAssignmentError{},
				Start: 0,
				End:   uint32(len(parts[0]) + len(v.Separator)),
			},
		}
	}

	if len(parts) > 2 {
		return []*InvalidValue{
			{
				Err:   errors.New("Key-DeprecatedValue pair must be in the form of 'key<separator>value'"),
				Start: 0,
				End:   uint32(len(value)),
			},
		}
	}

	checkValue, found := v.getValue(parts[0])

	if !found {
		return []*InvalidValue{
			{
				Err: ValueNotInEnumError{
					AvailableValues: utils.Map(utils.KeysOfMap(v.Values), func(key EnumString) string { return key.InsertText }),
					ProvidedValue:   parts[0],
				},
				Start: 0,
				End:   uint32(len(parts[0])),
			},
		}
	}

	errors := (*checkValue).DeprecatedCheckIsValid(parts[1])

	if len(errors) > 0 {
		ShiftInvalidValues(uint32(len(parts[0])+len(v.Separator)), errors)
		return errors
	}

	return nil
}

func (v KeyEnumAssignmentValue) FetchEnumCompletions() []protocol.CompletionItem {
	completions := make([]protocol.CompletionItem, 0)

	for enumKey := range v.Values {
		textFormat := protocol.InsertTextFormatPlainText
		kind := protocol.CompletionItemKindField
		val := v.Values[enumKey]
		description := val.GetTypeDescription()
		insertText := enumKey.InsertText + v.Separator

		var documentation string

		if len(description) == 1 {
			documentation = fmt.Sprintf("%s%s<%s> \n\n%s", enumKey.InsertText, v.Separator, description[0], enumKey.Documentation)
		} else {
			documentation = fmt.Sprintf("%s%s<value> \n\n%s", enumKey.InsertText, v.Separator, enumKey.Documentation)
		}

		completions = append(completions, protocol.CompletionItem{
			Label:            enumKey.InsertText,
			InsertText:       &insertText,
			InsertTextFormat: &textFormat,
			Kind:             &kind,
			Documentation:    documentation,
		})
	}

	return completions
}

type selectedValue string

const (
	keySelected   selectedValue = "key"
	valueSelected selectedValue = "value"
)

func (v KeyEnumAssignmentValue) getValueAtCursor(line string, cursor uint32) (string, *selectedValue, uint32) {
	relativePosition, found := utils.FindPreviousCharacter(line, v.Separator, int(cursor))

	if found {
		// DeprecatedValue found
		selected := valueSelected
		return line[uint32(relativePosition+1):], &selected, cursor - uint32(relativePosition)
	}

	selected := keySelected

	// Key, let's check for the separator
	relativePosition, found = utils.FindNextCharacter(line, v.Separator, int(cursor))

	if found {
		return line[:uint32(relativePosition)], &selected, cursor
	}

	// No separator, so we can just return the whole line
	return line, &selected, cursor
}

func (v KeyEnumAssignmentValue) FetchCompletions(value string, cursor common.CursorPosition) []protocol.CompletionItem {
	// Possible scenarios:
	// <empty> -> Fetch enum completions
	// enu| -> Fetch enum completions
	// enu|=value -> Fetch enum completions
	// enum=| -> Fetch enum completions
	// enum=val|ue -> Fetch value
	// This boils down to checking for the separator
	if value == "" {
		return v.FetchEnumCompletions()
	}

	index := common.DeprecatedImprovedCursorToIndex(cursor, value, 0)

	foundPosition, found := utils.FindPreviousCharacter(
		value,
		v.Separator,
		int(index),
	)

	if found {
		selectedEnum := value[:uint32(foundPosition)]

		enumValue, found := v.getValue(selectedEnum)

		if !found {
			// The user typed in an enum that is not in the list
			return nil
		}

		start := uint32(foundPosition + len(v.Separator))
		remainingValue := value[start:]
		remainingCursor := cursor.ShiftHorizontal(-start)

		return (*enumValue).FetchCompletions(remainingValue, remainingCursor)
	} else {
		// No separator found, so we can just return the enum completions
		return v.FetchEnumCompletions()
	}
}

func (v KeyEnumAssignmentValue) DeprecatedFetchCompletions(line string, cursor uint32) []protocol.CompletionItem {
	return nil
}

func (v KeyEnumAssignmentValue) DeprecatedFetchHoverInfo(line string, cursor uint32) []string {
	if len(v.DeprecatedCheckIsValid(line)) != 0 {
		return []string{}
	}

	value, selected, cursor := v.getValueAtCursor(line, cursor)

	if selected == nil {
		return []string{}
	}

	if *selected == keySelected {
		// Search for enum documentation
		enums := utils.KeysOfMap(v.Values)
		key := value

		for _, enum := range enums {
			if enum.InsertText == value {
				return []string{
					fmt.Sprintf("## `%s%s<value>`", key, v.Separator),
					enum.Documentation,
				}
			}
		}
	} else if *selected == valueSelected {
		// Search for value documentation
		// - 1 to remove the separator
		key := strings.SplitN(line, v.Separator, 2)[0]
		checkValue, found := v.getValue(key)

		if !found {
			return []string{}
		}

		info := (*checkValue).DeprecatedFetchHoverInfo(value, cursor)

		return append(
			[]string{
				fmt.Sprintf("## `%s%s%s`", key, v.Separator, value),
			},
			info...,
		)
	}

	return []string{}
}
