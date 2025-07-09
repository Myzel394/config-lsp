package docvalues

import (
	"config-lsp/common"
	"config-lsp/utils"
	"fmt"
	"strings"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

type KeyValueAssignmentError struct{}

func (e KeyValueAssignmentError) Error() string {
	return "This is not valid key-value assignment"
}

type KeyValueAssignmentContext struct {
	SelectedKey string
}

func (KeyValueAssignmentContext) GetIsContext() bool {
	return true
}

type KeyValueAssignmentValue struct {
	Key DeprecatedValue
	// If this is a `CustomValue`, it will receive a `KeyValueAssignmentContext`
	Value           DeprecatedValue
	ValueIsOptional bool
	Separator       string
}

func (v KeyValueAssignmentValue) GetTypeDescription() []string {
	keyDescription := v.Key.GetTypeDescription()
	valueDescription := v.Value.GetTypeDescription()

	if len(keyDescription) == 1 && len(valueDescription) == 1 {
		return []string{
			fmt.Sprintf("Key-DeprecatedValue pair in form of '<%s>%s<%s>'", keyDescription[0], v.Separator, valueDescription[0]),
		}
	} else {
		return []string{
			fmt.Sprintf("Key-DeprecatedValue pair in form of 'key%svalue'", v.Separator),
			fmt.Sprintf("#### Key\n%s", strings.Join(v.Key.GetTypeDescription(), "\n")),
			fmt.Sprintf("#### DeprecatedValue:\n%s", strings.Join(v.Value.GetTypeDescription(), "\n")),
		}
	}
}

func (v KeyValueAssignmentValue) getValue(selectedKey string) DeprecatedValue {
	switch v.Value.(type) {
	case CustomValue:
		{
			customValue := v.Value.(CustomValue)
			context := KeyValueAssignmentContext{
				SelectedKey: selectedKey,
			}

			return customValue.FetchValue(context)
		}
	default:
		{
			return v.Value
		}
	}
}

func (v KeyValueAssignmentValue) DeprecatedCheckIsValid(value string) []*InvalidValue {
	parts := strings.Split(value, v.Separator)

	if len(parts) == 0 || parts[0] == "" {
		// Nothing to check for
		return nil
	}

	err := v.Key.DeprecatedCheckIsValid(parts[0])

	if err != nil {
		return err
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

	errors := v.getValue(parts[0]).DeprecatedCheckIsValid(parts[1])

	if len(errors) > 0 {
		ShiftInvalidValues(uint32(len(parts[0])+len(v.Separator)), errors)
		return errors
	}

	return nil
}

func (v KeyValueAssignmentValue) FetchCompletions(value string, cursor common.CursorPosition) []protocol.CompletionItem {
	if value == "" {
		return v.Key.FetchCompletions(value, 0)
	}

	index := common.DeprecatedImprovedCursorToIndex(
		cursor,
		value,
		0,
	)

	foundPosition, found := utils.FindPreviousCharacter(
		value,
		v.Separator,
		int(index),
	)

	if found {
		selectedKey := value[:uint32(foundPosition)]

		start := uint32(foundPosition + len(v.Separator))
		remainingValue := value[start:]
		remainingCursor := cursor.ShiftHorizontal(-start)

		return v.getValue(selectedKey).FetchCompletions(remainingValue, remainingCursor)
	} else {
		return v.Key.FetchCompletions(value, cursor)
	}
}

func (v KeyValueAssignmentValue) getValueAtCursor(line string, cursor uint32) (string, *selectedValue, uint32) {
	relativePosition, found := utils.FindPreviousCharacter(line, v.Separator, int(cursor))

	if found {
		// DeprecatedValue found
		selected := valueSelected
		return line[:uint32(relativePosition)], &selected, cursor - uint32(relativePosition)
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

func (v KeyValueAssignmentValue) DeprecatedFetchHoverInfo(line string, cursor uint32) []string {
	if len(v.DeprecatedCheckIsValid(line)) != 0 {
		return []string{}
	}

	value, selected, cursor := v.getValueAtCursor(line, cursor)

	if selected == nil {
		return []string{}
	}

	if *selected == keySelected {
		// Get key documentation
		return v.Key.DeprecatedFetchHoverInfo(value, cursor)
	} else if *selected == valueSelected {
		// Get for value documentation
		key := strings.SplitN(line, v.Separator, 2)[0]

		return v.getValue(key).DeprecatedFetchHoverInfo(value, cursor)
	}

	return []string{}
}
