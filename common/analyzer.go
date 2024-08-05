package common

import (
	docvalues "config-lsp/doc-values"
)

func AnalyzeValues(
	parser SimpleConfigParser,
	availableOptions map[string]Option,
) []docvalues.ValueError {
	errors := make([]docvalues.ValueError, 0)

	for optionName, line := range parser.Lines {
		documentationOption := availableOptions[optionName]

		err := documentationOption.Value.CheckIsValid(line.Value)

		if err != nil {
			errors = append(errors, docvalues.ValueError{
				Line:   line.Position.Line,
				Option: optionName,
				Value:  line.Value,
			})
		}
	}

	return errors
}
