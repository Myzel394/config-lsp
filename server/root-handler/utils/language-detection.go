package utils

import (
	"config-lsp/common"
	"config-lsp/root-handler/shared"
	"config-lsp/utils"
	"fmt"
	"path"
	"regexp"
	"strings"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

type UnsupportedLanguageError struct {
	SuggestedLanguage string
}

func (e UnsupportedLanguageError) Error() string {
	return fmt.Sprintf("Language '%s' is not supported. Choose one of: %s", e.SuggestedLanguage, strings.Join(shared.AllSupportedLanguages, ", "))
}

type LanguageUndetectableError struct{}

func (e LanguageUndetectableError) Error() string {
	return "Please add: '#?lsp.language=<language>' to the top of the file. config-lsp was unable to detect the appropriate language for this file."
}

var valueToLanguageMap = map[string]shared.SupportedLanguage{
	"sshd":        shared.LanguageSSHDConfig,
	"sshd_config": shared.LanguageSSHDConfig,
	"sshdconfig":  shared.LanguageSSHDConfig,

	"ssh":        shared.LanguageSSHConfig,
	"ssh_config": shared.LanguageSSHConfig,
	"sshconfig":  shared.LanguageSSHConfig,

	".ssh/config":   shared.LanguageSSHConfig,
	"~/.ssh/config": shared.LanguageSSHConfig,

	"fstab":     shared.LanguageFstab,
	"etc/fstab": shared.LanguageFstab,

	"wireguard":         shared.LanguageWireguard,
	"wg":                shared.LanguageWireguard,
	"languagewireguard": shared.LanguageWireguard,
	"host":              shared.LanguageHosts,
	"hosts":             shared.LanguageHosts,
	"etc/hosts":         shared.LanguageHosts,

	"aliases":     shared.LanguageAliases,
	"mailaliases": shared.LanguageAliases,
	"etc/aliases": shared.LanguageAliases,

	"bitcoin_conf":  shared.LanguageBitcoinConf,
	"bitcoinconf":   shared.LanguageBitcoinConf,
	"bitcoin":       shared.LanguageBitcoinConf,
	"bitcoin.conf":  shared.LanguageBitcoinConf,
	"btc":           shared.LanguageBitcoinConf,
	"btcd":          shared.LanguageBitcoinConf,
	"bitcoind":      shared.LanguageBitcoinConf,
	"bitcoind.conf": shared.LanguageBitcoinConf,
	"bitcoind_conf": shared.LanguageBitcoinConf,
	"bitcoindconf":  shared.LanguageBitcoinConf,
}

var filenameToLanguageMap = map[string]shared.SupportedLanguage{
	"sshd_config": shared.LanguageSSHDConfig,
	"sshdconfig":  shared.LanguageSSHDConfig,
	"sshd":        shared.LanguageSSHDConfig,
	"sshd_conf":   shared.LanguageSSHDConfig,
	"sshdconf":    shared.LanguageSSHDConfig,

	"ssh_config": shared.LanguageSSHConfig,
	"sshconfig":  shared.LanguageSSHConfig,
	"ssh":        shared.LanguageSSHConfig,
	"ssh_conf":   shared.LanguageSSHConfig,
	"sshconf":    shared.LanguageSSHConfig,

	"fstab": shared.LanguageFstab,

	"hosts": shared.LanguageHosts,

	"aliases":     shared.LanguageAliases,
	"mailaliases": shared.LanguageAliases,

	"bitcoin_conf":  shared.LanguageBitcoinConf,
	"bitcoin.conf":  shared.LanguageBitcoinConf,
	"bitcoind.conf": shared.LanguageBitcoinConf,
	"bitcoin":       shared.LanguageBitcoinConf,
	"bitcoind":      shared.LanguageBitcoinConf,
	"btc":           shared.LanguageBitcoinConf,
	"btc.conf":      shared.LanguageBitcoinConf,
	"btcd":          shared.LanguageBitcoinConf,
	"btcd.conf":     shared.LanguageBitcoinConf,
}

var typeOverwriteRegex = regexp.MustCompile(`#\s*\?\s*lsp\.language\s*=\s*(\w+)\s*`)
var wireguardPattern = regexp.MustCompile(`wg(\d+)?(\.conf)?$`)

func DetectLanguage(
	content string,
	advertisedLanguage string,
	uri protocol.DocumentUri,
) (shared.SupportedLanguage, error) {
	if match := typeOverwriteRegex.FindStringSubmatchIndex(content); match != nil {
		language := content[match[2]:match[3]]
		suggestedLanguage := strings.ToLower(language)

		foundLanguage, ok := valueToLanguageMap[suggestedLanguage]

		contentUntilMatch := content[:match[0]]

		if ok {
			return foundLanguage, nil
		}

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
		return shared.LanguageSSHDConfig, nil

	case "file:///etc/fstab":
		return shared.LanguageFstab, nil

	// Darwin
	case "file:///private/etc/hosts":
		fallthrough
	case "file:///etc/hosts":
		return shared.LanguageHosts, nil

	// Darwin
	case "file:///private/etc/aliases":
		fallthrough
	case "file:///etc/aliases":
		return shared.LanguageAliases, nil
	}

	filename := path.Base(string(uri))

	if language, found := filenameToLanguageMap[filename]; found {
		return language, nil
	}

	if strings.HasPrefix(uri, "file:///etc/wireguard/") || wireguardPattern.MatchString(uri) {
		return shared.LanguageWireguard, nil
	}

	if strings.HasSuffix(uri, ".ssh/config") {
		return shared.LanguageSSHConfig, nil
	}

	return "", LanguageUndetectableError{}
}
