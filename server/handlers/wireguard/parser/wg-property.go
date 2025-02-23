package parser

import (
	docvalues "config-lsp/doc-values"
	"config-lsp/handlers/wireguard/ast"
	"config-lsp/utils"
	"regexp"
	"strings"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

var linePattern = regexp.MustCompile(`^\s*(?P<key>.+?)\s*(?P<separator>=)\s*(?P<value>\S.*?)?\s*(?:(?:;|#).*)?\s*$`)

// WireguardProperties [<line number>]: <property>
type WireguardProperties map[uint32]ast.WGProperty

func (p *WireguardProperties) AddLine(lineNumber uint32, line string) error {
	property, err := CreateWireguardProperty(line)

	if err != nil {
		return err
	}

	(*p)[lineNumber] = *property

	return nil
}

func CreateWireguardProperty(line string) (*ast.WGProperty, error) {
	if !strings.Contains(line, "=") {
		indexes := utils.GetTrimIndex(line)

		if indexes == nil {
			// weird, should not happen
			return nil, &docvalues.MalformedLineError{}
		}

		return &ast.WGProperty{
			Key: ast.WGPropertyKey{
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
	key := ast.WGPropertyKey{
		Location: CharacterLocation{
			Start: keyStart,
			End:   keyEnd,
		},
		Name: line[keyStart:keyEnd],
	}

	separatorStart := uint32(indexes[4])
	separatorEnd := uint32(indexes[5])
	separator := ast.WGPropertySeparator{
		Location: CharacterLocation{
			Start: separatorStart,
			End:   separatorEnd,
		},
	}

	var value *ast.WGPropertyValue

	if indexes[6] != -1 && indexes[7] != -1 {
		// value exists
		valueStart := uint32(indexes[6])
		valueEnd := uint32(indexes[7])

		value = &ast.WGPropertyValue{
			Location: CharacterLocation{
				Start: valueStart,
				End:   valueEnd,
			},
			Value: line[valueStart:valueEnd],
		}
	}

	return &ast.WGProperty{
		Key:       key,
		Separator: &separator,
		Value:     value,
	}, nil
}
