package docvalues

import (
	"config-lsp/utils"
	"strconv"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

type NotAPowerOfTwoError struct{}

func (e NotAPowerOfTwoError) Error() string {
	return "This must be a power of 2 (e.g. 32, 64, 128, 265, 512, 1024,, 1024, ...)"
}

type PowerOfTwoValue struct{}

func (v PowerOfTwoValue) GetTypeDescription() []string {
	return []string{"A power of 2"}
}

func isPowerOfTwo(number int) bool {
	count := 0

	for number > 0 {
		count += number & 1
		number >>= 1

		if count > 1 {
			return false
		}
	}

	return true
}

func (v PowerOfTwoValue) CheckIsValid(value string) []*InvalidValue {
	number, err := strconv.Atoi(value)

	if err != nil {
		return []*InvalidValue{{
			Err:   NotANumberError{},
			Start: 0,
			End:   uint32(len(value)),
		},
		}
	}

	if number <= 0 || !isPowerOfTwo(number) {
		return []*InvalidValue{{
			Err:   NotAPowerOfTwoError{},
			Start: 0,
			End:   uint32(len(value)),
		},
		}
	}

	return nil
}

var powers = []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192, 16384, 32768, 65536}

func (v PowerOfTwoValue) FetchCompletions(line string, cursor uint32) []protocol.CompletionItem {
	textFormat := protocol.InsertTextFormatPlainText
	kind := protocol.CompletionItemKindValue

	return utils.Map(
		powers,
		func(power int) protocol.CompletionItem {
			return protocol.CompletionItem{
				Label:            strconv.Itoa(power),
				InsertTextFormat: &textFormat,
				Kind:             &kind,
			}
		},
	)
}
