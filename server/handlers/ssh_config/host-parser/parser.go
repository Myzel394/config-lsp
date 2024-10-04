package hostparser

import (
	"config-lsp/common"
	commonparser "config-lsp/common/parser"
	"regexp"
)

func NewHost() *Host {
	match := new(Host)
	match.Clear()

	return match
}

func (h *Host) Clear() {
	h.Hosts = make([]*HostValue, 0)
}

var textPattern = regexp.MustCompile(`\S+`)

func (h *Host) Parse(
	input string,
	line uint32,
	startCharacter uint32,
) []common.LSPError {
	hostsIndexes := textPattern.FindAllStringIndex(input, -1)

	for _, hostIndex := range hostsIndexes {
		startIndex := hostIndex[0]
		endIndex := hostIndex[1]

		rawHost := input[startIndex:endIndex]

		value := commonparser.ParseRawString(rawHost, commonparser.FullFeatures)
		host := HostValue{
			LocationRange: common.LocationRange{
				Start: common.Location{
					Line:      line,
					Character: startCharacter + uint32(startIndex),
				},
				End: common.Location{
					Line:      line,
					Character: startCharacter + uint32(endIndex),
				},
			},
			Value: value,
		}

		h.Hosts = append(h.Hosts, &host)
	}

	return nil
}
