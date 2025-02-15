package fields

import (
	docvalues "config-lsp/doc-values"
	"os"
	"path"
	"strings"
)

type PathnameValue struct {
	*docvalues.PathValue
}

// TODO: Check if this works
func (v PathnameValue) createSystemPath(value string) string {
	if strings.HasPrefix(value, "~/") {
		// Path of current user

		home, err := os.UserHomeDir()

		if err != nil {
			return value
		}

		relativePath := strings.TrimPrefix(value, "~/")
		return path.Join(home, relativePath)
	}

	if strings.HasPrefix(value, "~") {
		// Path of another user

		// TODO: Check how this is supposed to work.
		// Why would you want to get the home directory of another user?

		return value

		// nextSlash := strings.Index(value, "/")
		//
		// if nextSlash == -1 {
		// 	// Path missing
		// 	return value
		// }
		//
		// user := value[1:nextSlash]
		// relativePath := value[nextSlash+1:]
	}

	if strings.HasPrefix(value, "%(prefix)/") {
		// yeah hell no
	}

	return value
}
