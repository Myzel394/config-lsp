package analyzer

import (
	"config-lsp/common"
	"config-lsp/handlers/ssh_config/ast"
	"config-lsp/handlers/ssh_config/fields"
	"config-lsp/utils"
	"fmt"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func analyzeStructureIsValid(
	ctx *analyzerContext,
) {
	it := ctx.document.Config.Options.Iterator()

	for it.Next() {
		entry := it.Value().(ast.SSHEntry)

		switch entry.(type) {
		case *ast.SSHOption:
			checkOption(ctx, entry.(*ast.SSHOption), nil)
		case *ast.SSHMatchBlock:
			matchBlock := entry.(*ast.SSHMatchBlock)
			checkBlock(ctx, matchBlock)
		case *ast.SSHHostBlock:
			hostBlock := entry.(*ast.SSHHostBlock)
			checkBlock(ctx, hostBlock)
		}

	}
}

func checkOption(
	ctx *analyzerContext,
	option *ast.SSHOption,
	block ast.SSHBlock,
) {
	checkIsUsingDoubleQuotes(ctx, option.Key.Value, option.Key.LocationRange)
	checkQuotesAreClosed(ctx, option.Key.Value, option.Key.LocationRange)

	_, found := fields.Options[option.Key.Key]

	if !found {
		// Diagnostics will be handled by `values.go`
		return
	}

	// Check for values that are not allowed in Host blocks
	if block != nil && block.GetBlockType() == ast.SSHBlockTypeHost {
		if utils.KeyExists(fields.HostDisallowedOptions, option.Key.Key) {
			ctx.diagnostics = append(ctx.diagnostics, protocol.Diagnostic{
				Range:    option.Key.LocationRange.ToLSPRange(),
				Message:  fmt.Sprintf("Option '%s' is not allowed in Host blocks", option.Key.Key),
				Severity: &common.SeverityError,
			})
		}
	}

	if option.OptionValue != nil {
		checkIsUsingDoubleQuotes(ctx, option.OptionValue.Value, option.OptionValue.LocationRange)
		checkQuotesAreClosed(ctx, option.OptionValue.Value, option.OptionValue.LocationRange)
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

func checkBlock(
	ctx *analyzerContext,
	block ast.SSHBlock,
) {
	checkOption(ctx, block.GetEntryOption(), block)

	it := block.GetOptions().Iterator()
	for it.Next() {
		option := it.Value().(*ast.SSHOption)

		checkOption(ctx, option, block)
	}
}
