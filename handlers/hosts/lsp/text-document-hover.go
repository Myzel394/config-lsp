package lsp

import (
	"config-lsp/handlers/hosts/fields"
	"config-lsp/handlers/hosts/handlers"
	"config-lsp/handlers/hosts/handlers/analyzer"
	"strings"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentHover(
	context *glsp.Context,
	params *protocol.HoverParams,
) (*protocol.Hover, error) {
	parser := documentParserMap[params.TextDocument.URI]

	line := params.Position.Line
	character := params.Position.Character

	if _, found := parser.CommentLines[line]; found {
		// Comment
		return nil, nil
	}

	entry, found := parser.Tree.Entries[line]

	if !found {
		// Empty line
		return nil, nil
	}

	target := handlers.GetHoverTargetInEntry(*entry, character)

	var hostname *analyzer.HostsHostname

	switch *target {
	case handlers.HoverTargetIPAddress:
		relativeCursor := character - entry.IPAddress.Location.Start.Character
		hover := fields.IPAddressField.FetchHoverInfo(entry.IPAddress.Value.String(), relativeCursor)

		return &protocol.Hover{
			Contents: hover,
		}, nil
	case handlers.HoverTargetHostname:
		hostname = entry.Hostname
	case handlers.HoverTargetAlias:
		for _, alias := range entry.Aliases {
			if alias.Location.Start.Character <= character && character <= alias.Location.End.Character {
				hostname = alias
				break
			}
		}
	}

	if hostname != nil {
		contents := []string{
			"## Hostname",
		}
		contents = append(
			contents,
			fields.HostnameField.GetTypeDescription()...,
		)
		contents = append(
			contents,
			[]string{
				"",
			}...,
		)
		contents = append(
			contents,
			fields.HostnameField.Documentation,
		)
		contents = append(
			contents,
			handlers.GetHoverInfoForHostname(*parser, *hostname, character)...,
		)

		return &protocol.Hover{
			Contents: &protocol.MarkupContent{
				Kind:  protocol.MarkupKindMarkdown,
				Value: strings.Join(contents, "\n"),
			},
		}, nil
	}

	return nil, nil
}
