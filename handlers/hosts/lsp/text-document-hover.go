package lsp

import (
	"config-lsp/common"
	"config-lsp/handlers/hosts"
	"config-lsp/handlers/hosts/ast"
	"config-lsp/handlers/hosts/fields"
	"config-lsp/handlers/hosts/handlers"
	"strings"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentHover(
	context *glsp.Context,
	params *protocol.HoverParams,
) (*protocol.Hover, error) {
	document := hosts.DocumentParserMap[params.TextDocument.URI]

	line := params.Position.Line
	index := common.LSPCharacterAsIndexPosition(params.Position.Character)

	if _, found := document.Parser.CommentLines[line]; found {
		// Comment
		return nil, nil
	}

	rawEntry, found := document.Parser.Tree.Entries.Get(line)

	if !found {
		// Empty line
		return nil, nil
	}

	entry := rawEntry.(*ast.HostsEntry)
	target := handlers.GetHoverTargetInEntry(index, *entry)

	var hostname *ast.HostsHostname

	switch *target {
	case handlers.HoverTargetIPAddress:
		line := entry.IPAddress.Value.String()
		relativeCursor := uint32(entry.IPAddress.Location.Start.GetRelativeIndexPosition(index))
		hover := fields.IPAddressField.DeprecatedFetchHoverInfo(line, relativeCursor)

		return &protocol.Hover{
			Contents: hover,
		}, nil
	case handlers.HoverTargetHostname:
		hostname = entry.Hostname
	case handlers.HoverTargetAlias:
		for _, alias := range entry.Aliases {
			if alias.Location.ContainsPosition(index) {
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
			[]string{
				"",
			}...,
		)
		contents = append(
			contents,
			handlers.GetHoverInfoForHostname(index, *document, *hostname)...,
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
