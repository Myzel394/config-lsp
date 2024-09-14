package docvalues

import (
	"config-lsp/utils"
	"fmt"
	"strings"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

type KeyEnumAssignmentValue struct {
	Values          map[EnumString]Value
	Separator       string
	ValueIsOptional bool
}

func (v KeyEnumAssignmentValue) GetTypeDescription() []string {
	if len(v.Values) == 1 {
		firstKey := utils.KeysOfMap(v.Values)[0]
		valueDescription := v.Values[firstKey].GetTypeDescription()

		if len(valueDescription) == 1 {
			return []string{
				fmt.Sprintf("Key-Value pair in form of '<%s>%s<%s>'", firstKey.DescriptionText, v.Separator, valueDescription[0]),
			}
		}
	}

	var result []string
	for key, value := range v.Values {
		result = append(result, key.Documentation)
		result = append(result, value.GetTypeDescription()...)
	}

	return append([]string{
		"Key-Value pair in form of 'key%svalue'", v.Separator,
	}, result...)
}

func (v KeyEnumAssignmentValue) getValue(findKey string) (*Value, bool) {
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

func (v KeyEnumAssignmentValue) CheckIsValid(value string) []*InvalidValue {
	parts := strings.Split(value, v.Separator)

	if len(parts) == 0 || parts[0] == "" {
		// Nothing to check for
		return nil
	}

	if len(parts) != 2 {
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

	errors := (*checkValue).CheckIsValid(parts[1])

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
		// Value found
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

func (v KeyEnumAssignmentValue) FetchCompletions(line string, cursor uint32) []protocol.CompletionItem {
	if cursor == 0 {
		return v.FetchEnumCompletions()
	}

	relativePosition, found := utils.FindPreviousCharacter(
		line,
		v.Separator,
		int(cursor),
	)

	if found {
		selectedKey := line[:uint32(relativePosition)]
		line = line[uint32(relativePosition+len(v.Separator)):]
		cursor -= uint32(relativePosition)

		keyValue, found := v.getValue(selectedKey)

		if !found {
			// Hmm... weird
			return v.FetchEnumCompletions()
		}

		return (*keyValue).FetchCompletions(line, cursor)
	} else {
		return v.FetchEnumCompletions()
	}
}

func (v KeyEnumAssignmentValue) FetchHoverInfo(line string, cursor uint32) []string {
	if len(v.CheckIsValid(line)) != 0 {
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

		info := (*checkValue).FetchHoverInfo(value, cursor)

		return append(
			[]string{
				fmt.Sprintf("## `%s%s%s`", key, v.Separator, value),
			},
			info...,
		)
	}

	return []string{}
}
