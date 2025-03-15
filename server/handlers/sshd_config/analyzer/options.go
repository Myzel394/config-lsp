package analyzer

import (
	"config-lsp/common"
	docvalues "config-lsp/doc-values"
	"config-lsp/handlers/sshd_config/ast"
	"config-lsp/handlers/sshd_config/diagnostics"
	"config-lsp/handlers/sshd_config/fields"
	"config-lsp/utils"
	"fmt"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func analyzeStructureIsValid(
	ctx *analyzerContext,
) {
	it := ctx.document.Config.Options.Iterator()

	for it.Next() {
		entry := it.Value().(ast.SSHDEntry)

		switch entry.(type) {
		case *ast.SSHDOption:
			checkOption(ctx, entry.(*ast.SSHDOption), nil)
		case *ast.SSHDMatchBlock:
			matchBlock := entry.(*ast.SSHDMatchBlock)
			checkMatchBlock(ctx, matchBlock)
		}
	}
}

func checkOption(
	ctx *analyzerContext,
	option *ast.SSHDOption,
	matchBlock *ast.SSHDMatchBlock,
) {
	if option.Key == nil {
		return
	}

	///// General checks
	checkIsUsingDoubleQuotes(ctx, option.Key.Value, option.Key.LocationRange)
	checkQuotesAreClosed(ctx, option.Key.Value, option.Key.LocationRange)

	if option.Separator == nil || option.Separator.Value.Value == "" {
		ctx.diagnostics = append(ctx.diagnostics, protocol.Diagnostic{
			Range:    option.Key.LocationRange.ToLSPRange(),
			Message:  "There should be a separator between an option and its value",
			Severity: &common.SeverityError,
		})
	} else {
		checkIsUsingDoubleQuotes(ctx, option.Separator.Value, option.Separator.LocationRange)
		checkQuotesAreClosed(ctx, option.Separator.Value, option.Separator.LocationRange)
	}

	///// Check if the key is valid
	docOption, optionFound := fields.Options[option.Key.Key]

	if !optionFound {
		ctx.diagnostics = append(ctx.diagnostics, diagnostics.GenerateUnknownOption(
			option.Key.ToLSPRange(),
			option.Key.Value.Value,
		))
		ctx.document.Indexes.UnknownOptions[option.Start.Line] = ast.SSHDOptionInfo{
			Option:     option,
			MatchBlock: matchBlock,
		}

		// Since we don't know the option, we can't verify the value
		return
	} else {
		// Check for values that are not allowed in Match blocks
		if matchBlock != nil && !utils.KeyExists(fields.MatchAllowedOptions, option.Key.Key) {
			ctx.diagnostics = append(ctx.diagnostics, protocol.Diagnostic{
				Range:    option.Key.ToLSPRange(),
				Message:  fmt.Sprintf("Option '%s' is not allowed in Match blocks", option.Key.Key),
				Severity: &common.SeverityError,
			})
		}
	}

	///// Check if the value is valid
	if option.OptionValue != nil {
		checkIsUsingDoubleQuotes(ctx, option.OptionValue.Value, option.OptionValue.LocationRange)
		checkQuotesAreClosed(ctx, option.OptionValue.Value, option.OptionValue.LocationRange)

		invalidValues := docOption.DeprecatedCheckIsValid(option.OptionValue.Value.Value)

		for _, invalidValue := range invalidValues {
			err := docvalues.LSPErrorFromInvalidValue(option.Start.Line, *invalidValue).ShiftCharacter(option.OptionValue.Start.Character)

			ctx.diagnostics = append(ctx.diagnostics, protocol.Diagnostic{
				Range:    err.Range.ToLSPRange(),
				Message:  err.Err.Error(),
				Severity: &common.SeverityError,
			})
		}
	}

}

func checkMatchBlock(
	ctx *analyzerContext,
	matchBlock *ast.SSHDMatchBlock,
) {
	it := matchBlock.Options.Iterator()

	for it.Next() {
		option := it.Value().(*ast.SSHDOption)

		checkOption(ctx, option, matchBlock)
	}
}
