package shared

type SupportedLanguage string

const (
	LanguageSSHConfig   SupportedLanguage = "ssh_config"
	LanguageSSHDConfig  SupportedLanguage = "sshd_config"
	LanguageFstab       SupportedLanguage = "fstab"
	LanguageWireguard   SupportedLanguage = "wireguard"
	LanguageHosts       SupportedLanguage = "hosts"
	LanguageAliases     SupportedLanguage = "aliases"
	LanguageBitcoinConf SupportedLanguage = "bitcoin_conf"
)

var AllSupportedLanguages = []string{
	string(LanguageSSHConfig),
	string(LanguageSSHDConfig),
	string(LanguageFstab),
	string(LanguageWireguard),
	string(LanguageHosts),
	string(LanguageAliases),
	string(LanguageBitcoinConf),
}
