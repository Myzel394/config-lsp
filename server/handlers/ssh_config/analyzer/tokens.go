package analyzer

import (
	"config-lsp/common"
	"config-lsp/handlers/ssh_config/fields"
	"config-lsp/utils"
	"fmt"
	"strings"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func analyzeTokens(
	ctx *analyzerContext,
) {
	for _, info := range ctx.document.Config.GetAllOptions() {
		if info.Option.Key == nil || info.Option.OptionValue == nil {
			continue
		}

		key := info.Option.Key.Key
		text := info.Option.OptionValue.Value.Value
		var tokens []string

		if foundTokens, found := fields.OptionsTokensMap[key]; found {
			tokens = foundTokens
		} else {
			tokens = []string{}
		}

		disallowedTokens := utils.Without(fields.AvailableTokens, tokens)

		for _, token := range disallowedTokens {
			if strings.Contains(text, token) {
				optionName := string(key)

				if formatted, found := fields.FieldsNameFormattedMap[key]; found {
					optionName = formatted
				}

				ctx.diagnostics = append(ctx.diagnostics, protocol.Diagnostic{
					Range:    info.Option.OptionValue.ToLSPRange(),
					Message:  fmt.Sprintf("Token '%s' is not allowed for option '%s'", token, optionName),
					Severity: &common.SeverityError,
				})
			}
		}
	}
}
