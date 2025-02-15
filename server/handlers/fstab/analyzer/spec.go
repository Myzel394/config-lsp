package analyzer

import (
	"config-lsp/common"
	"config-lsp/handlers/fstab/ast"
	"regexp"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

var volatileBlockFields = regexp.MustCompile(`^/dev/(sd|nvme|mmcblk|sr|vd|loop|cdrom)[a-zA-Z0-9]*$`)

func analyzeSpecField(
	ctx *analyzerContext,
	field *ast.FstabField,
) {
	if field == nil {
		return
	}

	if field.Value.Value == "" {
		return
	}

	if !volatileBlockFields.MatchString(field.Value.Value) {
		return
	}

	codeDescription := protocol.CodeDescription{
		HRef: protocol.URI("https://wiki.archlinux.org/title/Persistent_block_device_naming"),
	}
	ctx.diagnostics = append(ctx.diagnostics, protocol.Diagnostic{
		Range:           field.ToLSPRange(),
		Message:         "Kernel name descriptors for block devices are not persistent and can change each boot, they should not be used in configuration files. Prefer device UUIDs or LABELs instead.",
		CodeDescription: &codeDescription,
		Severity:        &common.SeverityWarning,
	})
}
