package diagnostics

import (
	"config-lsp/common"
	"fmt"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func GenerateUnknownOption(
	diagnosticRange protocol.Range,
	optionName string,
) protocol.Diagnostic {
	return protocol.Diagnostic{
		Range:    diagnosticRange,
		Message:  fmt.Sprintf("Unknown option: %s", optionName),
		Severity: &common.SeverityError,
	}
}
