package analyzer

import (
	"config-lsp/common"
	sshdconfig "config-lsp/handlers/sshd_config"
	"config-lsp/handlers/sshd_config/ast"
	"config-lsp/handlers/sshd_config/indexes"
	"config-lsp/utils"
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"
)

var whitespacePattern = regexp.MustCompile(`\S+`)

func analyzeIncludeValues(
	d *sshdconfig.SSHDDocument,
) []common.LSPError {
	errs := make([]common.LSPError, 0)

	for _, include := range d.Indexes.Includes {
		for _, value := range include.Values {
			validPaths, err := createIncludePaths(value.Value)

			if err != nil {
				errs = append(errs, common.LSPError{
					Range: value.LocationRange,
					Err:   err,
				})
			} else {
				value.Paths = validPaths
			}
		}
	}

	return errs
}

func createIncludePaths(
	suggestedPath string,
) ([]indexes.ValidPath, error) {
	var absolutePath string

	if path.IsAbs(suggestedPath) {
		absolutePath = suggestedPath
	} else {
		absolutePath = path.Join("/etc", "ssh", suggestedPath)
	}

	files, err := filepath.Glob(absolutePath)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Could not find file %s (error: %s)", absolutePath, err))
	}

	if len(files) == 0 {
		return nil, errors.New(fmt.Sprintf("Could not find file %s", absolutePath))
	}

	return utils.Map(
		files,
		func(file string) indexes.ValidPath {
			return indexes.ValidPath(file)
		},
	), nil
}

func parseFile(
	filePath string,
) (*sshdconfig.SSHDDocument, error) {
	if d, ok := sshdconfig.DocumentParserMap[filePath]; ok {
		return d, nil
	}

	c := ast.NewSSHDConfig()

	content, err := os.ReadFile(filePath)

	if err != nil {
		return nil, err
	}

	parseErrors := c.Parse(string(content))

	if len(parseErrors) > 0 {
		return nil, errors.New(fmt.Sprintf("Errors in %s", filePath))
	}

	d := &sshdconfig.SSHDDocument{
		Config: c,
	}

	errs := Analyze(d)

	if len(errs) > 0 {
		return nil, errors.New(fmt.Sprintf("Errors in %s", filePath))
	}

	sshdconfig.DocumentParserMap[filePath] = d

	return d, nil
}
