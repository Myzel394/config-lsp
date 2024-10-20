package roothandler

import (
	"fmt"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func fetchAddLanguageActions(uri protocol.DocumentUri) ([]protocol.CodeAction, error) {
	actions := make([]protocol.CodeAction, 0, len(AllSupportedLanguages))

	kind := protocol.CodeActionKindQuickFix
	isPreferred := true

	for _, language := range AllSupportedLanguages {
		actions = append(actions, protocol.CodeAction{
			Title:       fmt.Sprintf("Use %s for this file", language),
			Kind:        &kind,
			IsPreferred: &isPreferred,
			Edit: &protocol.WorkspaceEdit{
				Changes: map[protocol.DocumentUri][]protocol.TextEdit{
					uri: {
						{
							Range: protocol.Range{
								Start: protocol.Position{
									Line:      0,
									Character: 0,
								},
								End: protocol.Position{
									Line:      0,
									Character: 0,
								},
							},
							NewText: fmt.Sprintf("#?lsp.language=%s\n", language),
						},
					},
				},
			},
		})
	}

	return actions, nil
}
