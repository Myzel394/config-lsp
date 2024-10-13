package docvalues

import (
	"regexp"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

var color24BitPattern = regexp.MustCompile(`^#[0-9a-fA-F]{6}$`)
var color12BitPattern = regexp.MustCompile(`^#[0-9a-fA-F]{3}$`)

type InvalidColorError struct{}

func (e InvalidColorError) Error() string {
	return "Color is invalid. It must be in the form of: #RRGGBB"
}

type HexColorValue struct {
	Allow12Bit bool
}

func (v HexColorValue) GetTypeDescription() []string {
	return []string{"Color in HEX-Format (e.g. #RRGGBB)"}
}

func (v HexColorValue) DeprecatedCheckIsValid(value string) []*InvalidValue {
	if color24BitPattern.MatchString(value) || (v.Allow12Bit && color12BitPattern.MatchString(value)) {
		return nil
	}

	return []*InvalidValue{
		{
			Err:   InvalidColorError{},
			Start: 0,
			End:   uint32(len(value)),
		},
	}
}

var characters = []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F'}

func (v HexColorValue) DeprecatedFetchCompletions(line string, cursor uint32) []protocol.CompletionItem {
	completions := make([]protocol.CompletionItem, 0)

	if line == "" {
		kind := protocol.CompletionItemKindValue
		completions = append(completions, protocol.CompletionItem{
			Label: "#",
			Kind:  &kind,
		})
	} else if !color24BitPattern.MatchString(line) {
		for _, c := range characters {
			kind := protocol.CompletionItemKindValue
			completions = append(completions, protocol.CompletionItem{
				Label: string(c),
				Kind:  &kind,
			})
		}
	}

	return completions
}

func (v HexColorValue) DeprecatedFetchHoverInfo(line string, cursor uint32) []string {
	return []string{"Color in HEX-Format (e.g. #RRGGBB)"}
}
