package docvalues

import (
	"config-lsp/utils"
	"fmt"
	"strings"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

type KeyValueAssignmentError struct{}

func (e KeyValueAssignmentError) Error() string {
	return "This is not valid key-value assignment"
}

type KeyValueAssignmentValue struct {
	Key             Value
	Value           Value
	ValueIsOptional bool
	Separator       string
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

	if len(parts) == 0 || parts[0] == "" {
		// Nothing to check for
		return nil
	}

	err := v.Key.CheckIsValid(parts[0])

	if err != nil {
		return err
	}

	if len(parts) != 2 {
		if v.ValueIsOptional {
			return nil
		}

		return KeyValueAssignmentError{}
	}

	err = v.Value.CheckIsValid(parts[1])

	if err != nil {
		return err
	}

	return nil
}

func (v KeyValueAssignmentValue) FetchCompletions(line string, cursor uint32) []protocol.CompletionItem {
	if cursor == 0 {
		return v.Key.FetchCompletions(line, cursor)
	}

	relativePosition, found := utils.FindPreviousCharacter(line, v.Separator, int(cursor-1))

	if found {
		line = line[uint32(relativePosition+len(v.Separator)):]
		cursor -= uint32(relativePosition)

		return v.Value.FetchCompletions(line, cursor)
	} else {
		return v.Key.FetchCompletions(line, cursor)
	}
}
