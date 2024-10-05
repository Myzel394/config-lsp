package parser

import (
	docvalues "config-lsp/doc-values"
	"config-lsp/utils"
	"regexp"
	"strings"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

var linePattern = regexp.MustCompile(`^\s*(?P<key>.+?)\s*(?P<separator>=)\s*(?P<value>\S.*?)?\s*(?:(?:;|#).*)?\s*$`)

type WireguardPropertyKey struct {
	Location CharacterLocation
	Name     string
}

type WireguardPropertyValue struct {
	Location CharacterLocation
	Value    string
}

type WireguardPropertySeparator struct {
	Location CharacterLocation
}

type WireguardProperty struct {
	Key       WireguardPropertyKey
	Separator *WireguardPropertySeparator
	Value     *WireguardPropertyValue
}

func (p WireguardProperty) String() string {
	if p.Value == nil {
		return p.Key.Name
	}

	return p.Key.Name + "=" + p.Value.Value
}

func (p WireguardProperty) GetLineRange(line uint32) protocol.Range {
	return protocol.Range{
		Start: protocol.Position{
			Line:      line,
			Character: p.Key.Location.Start,
		},
		End: protocol.Position{
			Line:      line,
			Character: p.Key.Location.End,
		},
	}
}

func (p WireguardProperty) GetInsertRange(line uint32) protocol.Range {
	var insertPosition uint32 = p.Separator.Location.End
	var length uint32 = 0

	if p.Value != nil {
		insertPosition = p.Value.Location.Start - 1
		// Length of the value; +1 because of the starting space
		length = (p.Value.Location.End - p.Value.Location.Start) + 1
	}

	return protocol.Range{
		Start: protocol.Position{
			Line:      line,
			Character: insertPosition,
		},
		End: protocol.Position{
			Line:      line,
			Character: insertPosition + length,
		},
	}
}

// WireguardProperties [<line number>]: <property>
type WireguardProperties map[uint32]WireguardProperty

func (p *WireguardProperties) AddLine(lineNumber uint32, line string) error {
	property, err := CreateWireguardProperty(line)

	if err != nil {
		return err
	}

	(*p)[lineNumber] = *property

	return nil
}

func CreateWireguardProperty(line string) (*WireguardProperty, error) {
	if !strings.Contains(line, "=") {
		indexes := utils.GetTrimIndex(line)

		if indexes == nil {
			// weird, should not happen
			return nil, &docvalues.MalformedLineError{}
		}

		return &WireguardProperty{
			Key: WireguardPropertyKey{
				Name: line[indexes[0]:indexes[1]],
				Location: CharacterLocation{
					Start: uint32(indexes[0]),
					End:   uint32(indexes[1]),
				},
			},
		}, nil
	}

	indexes := linePattern.FindStringSubmatchIndex(line)

	if indexes == nil || len(indexes) == 0 {
		return nil, &docvalues.MalformedLineError{}
	}

	keyStart := uint32(indexes[2])
	keyEnd := uint32(indexes[3])
	key := WireguardPropertyKey{
		Location: CharacterLocation{
			Start: keyStart,
			End:   keyEnd,
		},
		Name: line[keyStart:keyEnd],
	}

	separatorStart := uint32(indexes[4])
	separatorEnd := uint32(indexes[5])
	separator := WireguardPropertySeparator{
		Location: CharacterLocation{
			Start: separatorStart,
			End:   separatorEnd,
		},
	}

	var value *WireguardPropertyValue

	if indexes[6] != -1 && indexes[7] != -1 {
		// value exists
		valueStart := uint32(indexes[6])
		valueEnd := uint32(indexes[7])

		value = &WireguardPropertyValue{
			Location: CharacterLocation{
				Start: valueStart,
				End:   valueEnd,
			},
			Value: line[valueStart:valueEnd],
		}
	}

	return &WireguardProperty{
		Key:       key,
		Separator: &separator,
		Value:     value,
	}, nil
}
