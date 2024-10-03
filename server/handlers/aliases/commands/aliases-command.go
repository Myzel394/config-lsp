package commands

import (
	"os/exec"
	"strings"
)

func IsNewAliasesCommandAvailable() bool {
	_, err := exec.LookPath("newaliases")

	return err == nil
}

func IsSendmailCommandAvailable() bool {
	_, err := exec.LookPath("sendmail")

	return err == nil
}

func IsPostfixCommandAvailable() bool {
	_, err := exec.LookPath("postfix")

	return err == nil
}

func CanSendTestMails() bool {
	return IsSendmailCommandAvailable() && IsPostfixCommandAvailable()
}

func UpdateAliasesDatabase() error {
	cmd := exec.Command("newaliases")

	err := cmd.Run()

	return err
}

func SendTestMail(address string, content string) error {
	cmd := exec.Command("sendmail", address)
	cmd.Stdin = strings.NewReader(content)

	err := cmd.Run()

	return err
}
