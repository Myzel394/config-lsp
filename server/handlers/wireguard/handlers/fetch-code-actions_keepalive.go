package handlers

import (
	"config-lsp/handlers/wireguard"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func GetKeepaliveCodeActions(
	d *wireguard.WGDocument,
	params *protocol.CodeActionParams,
) []protocol.CodeAction {
	line := params.Range.Start.Line

	for _, section := range d.Indexes.SectionsByName["Peer"] {
		if section.Start.Line >= line && line <= section.End.Line {
			_, endpoint := section.FindFirstPropertyByName("Endpoint")
			_, persistentKeepAlive := section.FindFirstPropertyByName("PersistentKeepalive")

			if endpoint != nil && persistentKeepAlive == nil {
				var insertionLine uint32
				lastProperty := section.GetLastProperty()

				if lastProperty == nil {
					insertionLine = section.End.Line
				} else {
					insertionLine = lastProperty.End.Line + 1
				}

				return []protocol.CodeAction{
					{
						Title: "Add PersistentKeepalive",
						Edit: &protocol.WorkspaceEdit{
							Changes: map[protocol.DocumentUri][]protocol.TextEdit{
								params.TextDocument.URI: {
									{
										Range: protocol.Range{
											Start: protocol.Position{
												Line:      insertionLine,
												Character: 0,
											},
											End: protocol.Position{
												Line:      insertionLine,
												Character: 0,
											},
										},
										NewText: "PersistentKeepalive = 25\n",
									},
								},
							},
						},
					},
				}
			}
		}
	}

	return nil
}
