package analyzer

import (
	"config-lsp/common"
	"config-lsp/handlers/ssh_config/fields"
	"fmt"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

var tagOption = fields.CreateNormalizedName("Tag")

func analyzeTagOptions(
	ctx *analyzerContext,
) {
	// Check if the specified tags actually exist
	for _, options := range ctx.document.Indexes.AllOptionsPerName[tagOption] {
		for _, option := range options {
			tag, found := ctx.document.Indexes.Tags[option.OptionValue.Value.Value]

			if found && tag.Block.Start.Line > option.Start.Line {
				continue
			}

			ctx.diagnostics = append(ctx.diagnostics, protocol.Diagnostic{
				Range:    option.OptionValue.ToLSPRange(),
				Message:  fmt.Sprintf("Unknown tag: %s", option.OptionValue.Value.Value),
				Severity: &common.SeverityError,
			})
		}
	}
}
