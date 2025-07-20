package analyzer

import (
	"config-lsp/common"
	"config-lsp/handlers/bitcoin_conf/fields"
	"config-lsp/utils"
	"fmt"
	"strings"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func analyzeSectionsNamesAreValid(ctx *analyzerContext) {
	for _, section := range ctx.document.Config.Sections {
		if section.Header == nil {
			// Root section
			continue
		}

		if _, found := fields.AvailableSections[section.Header.Name]; !found {
			availableSections := strings.Join(
				utils.MapMapToSlice(fields.AvailableSections, func(name string, _ string) string {
					return name
				}),
				", ",
			)

			severity := common.SeverityError
			ctx.diagnostics = append(ctx.diagnostics, protocol.Diagnostic{
				Message:  fmt.Sprintf("Unknown section, it must be one of: %s", availableSections),
				Range:    section.Header.ToLSPRange(),
				Severity: &severity,
			})
		}
	}
}
