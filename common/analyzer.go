package common

func AnalyzeValues(
	parser SimpleConfigParser,
	availableOptions map[string]Option,
) []ValueError {
	errors := make([]ValueError, 0)

	for optionName, line := range parser.Lines {
		documentationOption := availableOptions[optionName]

		err := documentationOption.Value.CheckIsValid(line.Value)

		if err != nil {
			errors = append(errors, ValueError{
				Line:   line.Position.Line,
				Option: optionName,
				Value: line.Value,
				DocError: err,
			})
		}
	}

	return errors
}

