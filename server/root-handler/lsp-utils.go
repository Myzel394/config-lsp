package roothandler

import (
	"config-lsp/common"
	"config-lsp/utils"
	"fmt"
	"regexp"
	"strings"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

type SupportedLanguage string

const (
	LanguageSSHConfig  SupportedLanguage = "ssh_config"
	LanguageSSHDConfig SupportedLanguage = "sshd_config"
	LanguageFstab      SupportedLanguage = "fstab"
	LanguageWireguard  SupportedLanguage = "languagewireguard"
	LanguageHosts      SupportedLanguage = "hosts"
	LanguageAliases    SupportedLanguage = "aliases"
)

var AllSupportedLanguages = []string{
	string(LanguageSSHConfig),
	string(LanguageSSHDConfig),
	string(LanguageFstab),
	string(LanguageWireguard),
	string(LanguageHosts),
	string(LanguageAliases),
}

type FatalFileNotReadableError struct {
	FileURI protocol.DocumentUri
	Err     error
}

func (e FatalFileNotReadableError) Error() string {
	return fmt.Sprintf("Fatal error! config-lsp was unable to read the file (%s); error: %s", e.FileURI, e.Err.Error())
}

type UnsupportedLanguageError struct {
	SuggestedLanguage string
}

func (e UnsupportedLanguageError) Error() string {
	return fmt.Sprintf("Language '%s' is not supported. Choose one of: %s", e.SuggestedLanguage, strings.Join(AllSupportedLanguages, ", "))
}

type LanguageUndetectableError struct{}

func (e LanguageUndetectableError) Error() string {
	return "Please add: '#?lsp.language=<language>' to the top of the file. config-lsp was unable to detect the appropriate language for this file."
}

var valueToLanguageMap = map[string]SupportedLanguage{
	"sshd_config": LanguageSSHDConfig,
	"sshdconfig":  LanguageSSHDConfig,

	"ssh_config": LanguageSSHConfig,
	"sshconfig":  LanguageSSHConfig,

	".ssh/config":   LanguageSSHConfig,
	"~/.ssh/config": LanguageSSHConfig,

	"fstab":     LanguageFstab,
	"etc/fstab": LanguageFstab,

	"wireguard":         LanguageWireguard,
	"wg":                LanguageWireguard,
	"languagewireguard": LanguageWireguard,
	"host":              LanguageHosts,
	"hosts":             LanguageHosts,
	"etc/hosts":         LanguageHosts,

	"aliases":     LanguageAliases,
	"mailaliases": LanguageAliases,
	"etc/aliases": LanguageAliases,
}

var typeOverwriteRegex = regexp.MustCompile(`#\?\s*lsp\.language\s*=\s*(\w+)\s*`)
var wireguardPattern = regexp.MustCompile(`/wg\d+\.conf$`)

var undetectableError = common.ParseError{
	Line: 0,
	Err:  LanguageUndetectableError{},
}

func DetectLanguage(
	content string,
	advertisedLanguage string,
	uri protocol.DocumentUri,
) (SupportedLanguage, error) {
	if match := typeOverwriteRegex.FindStringSubmatch(content); match != nil {
		suggestedLanguage := strings.ToLower(match[1])

		foundLanguage, ok := valueToLanguageMap[suggestedLanguage]

		if ok {
			return foundLanguage, nil
		}

		matchIndex := strings.Index(content, match[0])
		contentUntilMatch := content[:matchIndex]

		return "", common.ParseError{
			Line: uint32(utils.CountCharacterOccurrences(contentUntilMatch, '\n')),
			Err: UnsupportedLanguageError{
				SuggestedLanguage: suggestedLanguage,
			},
		}
	}

	if language, ok := valueToLanguageMap[advertisedLanguage]; ok {
		return language, nil
	}

	switch uri {
	case "file:///etc/ssh/sshd_config":
		fallthrough
	case "file:///etc/ssh/ssh_config":
		return LanguageSSHDConfig, nil

	case "file:///etc/fstab":
		return LanguageFstab, nil

	// Darwin
	case "file:///private/etc/hosts":
		fallthrough
	case "file:///etc/hosts":
		return LanguageHosts, nil

	// Darwin
	case "file:///private/etc/aliases":
		fallthrough
	case "file:///etc/aliases":
		return LanguageAliases, nil
	}

	if strings.HasPrefix(uri, "file:///etc/wireguard/") || wireguardPattern.MatchString(uri) {
		return LanguageWireguard, nil
	}

	if strings.HasSuffix(uri, ".ssh/config") {
		return LanguageSSHConfig, nil
	}

	return "", undetectableError
}
