package analyzer

import (
	"config-lsp/common"
	"config-lsp/handlers/sshd_config/fields"
	"config-lsp/utils"
	"fmt"
	"strings"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func analyzeTokens(
	ctx *analyzerContext,
) {
	for _, option := range ctx.document.Config.GetAllOptions() {
		if option.Key == nil || option.OptionValue == nil {
			continue
		}

		key := option.Key.Key
		text := option.OptionValue.Value.Value
		var tokens []string

		if foundTokens, found := fields.OptionsTokensMap[key]; found {
			tokens = foundTokens
		} else {
			tokens = []string{}
		}

		disallowedTokens := utils.Without(utils.KeysOfMap(fields.AvailableTokens), tokens)

		for _, token := range disallowedTokens {
			if strings.Contains(text, token) {
				optionName := string(key)

				if formatted, found := fields.FieldsNameFormattedMap[key]; found {
					optionName = formatted
				}

				ctx.diagnostics = append(ctx.diagnostics, protocol.Diagnostic{
					Range:    option.OptionValue.ToLSPRange(),
					Message:  fmt.Sprintf("Token '%s' is not allowed for option '%s'", token, optionName),
					Severity: &common.SeverityError,
				})
			}
		}
	}
}
