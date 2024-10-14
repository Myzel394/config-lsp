package analyzer

import (
	"config-lsp/common"
	sshconfig "config-lsp/handlers/ssh_config"
	"config-lsp/handlers/ssh_config/ast"
	"config-lsp/handlers/ssh_config/indexes"
	"config-lsp/utils"
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

var whitespacePattern = regexp.MustCompile(`\S+`)
var environmtalVariablePattern = regexp.MustCompile(`\${.+?}`)
var availableTokens = []string{
	"%%",
	"%C",
	"%d",
	"%f",
	"%H",
	"%h",
	"%l",
	"%i",
	"%j",
	"%K",
	"%k",
	"%L",
	"%l",
	"%n",
	"%p",
	"%r",
	"%T",
	"%t",
	"%u",
}

func analyzeIncludeValues(
	ctx *analyzerContext,
) {
	for _, include := range ctx.document.Indexes.Includes {
		for _, value := range include.Values {
			if isImpossibleToVerify(value.Value) {
				continue
			}

			validPaths, err := createIncludePaths(value.Value)

			if err != nil {
				ctx.diagnostics = append(ctx.diagnostics, protocol.Diagnostic{
					Range:    value.LocationRange.ToLSPRange(),
					Message:  err.Error(),
					Severity: &common.SeverityError,
				})
			} else {
				value.Paths = validPaths
			}
		}
	}
}

// We can't evaluate environmental variables or tokens as we don't know the actual
// values
func isImpossibleToVerify(
	path string,
) bool {
	if environmtalVariablePattern.MatchString(path) {
		return true
	}

	for _, token := range availableTokens {
		if strings.Contains(path, token) {
			return true
		}
	}

	return false
}

func createIncludePaths(
	suggestedPath string,
) ([]indexes.ValidPath, error) {
	var absolutePath string

	if path.IsAbs(suggestedPath) {
		absolutePath = suggestedPath
	} else if strings.HasPrefix(suggestedPath, "~") {
		homeFolder, err := os.UserHomeDir()

		if err != nil {
			return nil, errors.New(fmt.Sprintf("Could not find home folder (error: %s)", err))
		}

		absolutePath = path.Join(homeFolder, suggestedPath[1:])
	} else {
		homeFolder, err := os.UserHomeDir()

		if err != nil {
			return nil, errors.New(fmt.Sprintf("Could not find home folder (error: %s)", err))
		}

		absolutePath = path.Join(homeFolder, ".ssh", suggestedPath)
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
) (*sshconfig.SSHDocument, error) {
	if d, ok := sshconfig.DocumentParserMap[filePath]; ok {
		return d, nil
	}

	c := ast.NewSSHConfig()

	content, err := os.ReadFile(filePath)

	if err != nil {
		return nil, err
	}

	parseErrors := c.Parse(string(content))

	if len(parseErrors) > 0 {
		return nil, errors.New(fmt.Sprintf("Errors in %s", filePath))
	}

	d := &sshconfig.SSHDocument{
		Config: c,
	}

	errs := Analyze(d)

	if len(errs) > 0 {
		return nil, errors.New(fmt.Sprintf("Errors in %s", filePath))
	}

	sshconfig.DocumentParserMap[filePath] = d

	return d, nil
}
