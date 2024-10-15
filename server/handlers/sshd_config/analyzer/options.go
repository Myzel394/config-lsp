package analyzer

import (
	"config-lsp/common"
	docvalues "config-lsp/doc-values"
	"config-lsp/handlers/sshd_config/ast"
	"config-lsp/handlers/sshd_config/fields"
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
			checkOption(ctx, entry.(*ast.SSHDOption), false)
		case *ast.SSHDMatchBlock:
			matchBlock := entry.(*ast.SSHDMatchBlock)
			checkMatchBlock(ctx, matchBlock)
		}
	}
}

func checkOption(
	ctx *analyzerContext,
	option *ast.SSHDOption,
	isInMatchBlock bool,
) {
	if option.Key == nil {
		return
	}

	checkIsUsingDoubleQuotes(ctx, option.Key.Value, option.Key.LocationRange)
	checkQuotesAreClosed(ctx, option.Key.Value, option.Key.LocationRange)

	key := fields.CreateNormalizedName(option.Key.Key)
	docOption, found := fields.Options[key]

	if !found {
		ctx.diagnostics = append(ctx.diagnostics, protocol.Diagnostic{
			Range:    option.Key.ToLSPRange(),
			Message:  fmt.Sprintf("Unknown option: %s", option.Key.Key),
			Severity: &common.SeverityError,
		})

		return
	}

	if _, found := fields.MatchAllowedOptions[key]; !found && isInMatchBlock {
		ctx.diagnostics = append(ctx.diagnostics, protocol.Diagnostic{
			Range:    option.Key.ToLSPRange(),
			Message:  fmt.Sprintf("Option '%s' is not allowed inside Match blocks", option.Key.Key),
			Severity: &common.SeverityError,
		})
	}

	if option.OptionValue != nil {
		checkIsUsingDoubleQuotes(ctx, option.OptionValue.Value, option.OptionValue.LocationRange)
		checkQuotesAreClosed(ctx, option.OptionValue.Value, option.OptionValue.LocationRange)

		invalidValues := docOption.DeprecatedCheckIsValid(option.OptionValue.Value.Value)

		for _, invalidValue := range invalidValues {
			err := docvalues.LSPErrorFromInvalidValue(option.Start.Line, *invalidValue)
			err.ShiftCharacter(option.OptionValue.Start.Character)

			ctx.diagnostics = append(ctx.diagnostics, protocol.Diagnostic{
				Range:    err.Range.ToLSPRange(),
				Message:  err.Err.Error(),
				Severity: &common.SeverityError,
			})
		}
	}

	if option.Separator == nil || option.Separator.Value.Value == "" {
		ctx.diagnostics = append(ctx.diagnostics, protocol.Diagnostic{
			Range:    option.Key.LocationRange.ToLSPRange(),
			Message:  fmt.Sprintf("There should be a separator between an option and its value"),
			Severity: &common.SeverityError,
		})
	} else {
		checkIsUsingDoubleQuotes(ctx, option.Separator.Value, option.Separator.LocationRange)
		checkQuotesAreClosed(ctx, option.Separator.Value, option.Separator.LocationRange)
	}
}

func checkMatchBlock(
	ctx *analyzerContext,
	matchBlock *ast.SSHDMatchBlock,
) {
	it := matchBlock.Options.Iterator()

	for it.Next() {
		option := it.Value().(*ast.SSHDOption)

		checkOption(ctx, option, true)
	}
}
