package handlers

import "config-lsp/common"

// TODO: Cache options in a map like: EnumValues -> []Option
// for faster lookup

func AnalyzeValue() []common.ValueError {
	errors := make([]common.ValueError, 0)

	for optionName, line := range Parser.Lines {
		documentationOption := Options[optionName]

		err := documentationOption.Value.CheckIsValid(line.Value)

		if err != nil {
			errors = append(errors, common.ValueError{
				Line:  line.Position.Line,
				Start: len(optionName) + len(" "),
				End:   len(optionName) + len(" ") + len(line.Value),
				Error: err,
			})
		}
	}

	return errors
}

// func AnalyzeSSHConfigIssues() []common.ParserError {}
