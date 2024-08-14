package wireguard

import (
	docvalues "config-lsp/doc-values"
	"config-lsp/utils"
	"regexp"
	"strings"
)

var linePattern = regexp.MustCompile(`^\s*(?P<key>.+?)\s*(?P<separator>=)\s*(?P<value>\S.*?)?\s*(?:(?:;|#).*)?\s*$`)

type wireguardPropertyKey struct {
	Location characterLocation
	Name     string
}

type wireguardPropertyValue struct {
	Location characterLocation
	Value    string
}

type wireguardPropertySeparator struct {
	Location characterLocation
}

type wireguardProperty struct {
	Key       wireguardPropertyKey
	Separator *wireguardPropertySeparator
	Value     *wireguardPropertyValue
}

func (p wireguardProperty) String() string {
	if p.Value == nil {
		return p.Key.Name
	}

	return p.Key.Name + "=" + p.Value.Value
}

func createWireguardProperty(line string) (*wireguardProperty, error) {
	if !strings.Contains(line, "=") {
		indexes := utils.GetTrimIndex(line)

		if indexes == nil {
			// weird, should not happen
			return nil, &docvalues.MalformedLineError{}
		}

		return &wireguardProperty{
			Key: wireguardPropertyKey{
				Location: characterLocation{
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
	key := wireguardPropertyKey{
		Location: characterLocation{
			Start: keyStart,
			End:   keyEnd,
		},
		Name: line[keyStart:keyEnd],
	}

	separatorStart := uint32(indexes[4])
	separatorEnd := uint32(indexes[5])
	separator := wireguardPropertySeparator{
		Location: characterLocation{
			Start: separatorStart,
			End:   separatorEnd,
		},
	}

	var value *wireguardPropertyValue

	if indexes[6] != -1 && indexes[7] != -1 {
		// value exists
		valueStart := uint32(indexes[6])
		valueEnd := uint32(indexes[7])

		value = &wireguardPropertyValue{
			Location: characterLocation{
				Start: valueStart,
				End:   valueEnd,
			},
			Value: line[valueStart:valueEnd],
		}
	}

	return &wireguardProperty{
		Key:       key,
		Separator: &separator,
		Value:     value,
	}, nil
}

// [<line number>]: <property>
type wireguardProperties map[uint32]wireguardProperty

func (p *wireguardProperties) AddLine(lineNumber uint32, line string) error {
	property, err := createWireguardProperty(line)

	if err != nil {
		return err
	}

	(*p)[lineNumber] = *property

	return nil
}
