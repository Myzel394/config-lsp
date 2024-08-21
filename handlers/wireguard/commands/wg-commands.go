package wgcommands

import (
	"os/exec"
	"regexp"
	"strings"
)

var whitespacePattern = regexp.MustCompile(`[\s\n]+`)

func AreWireguardToolsAvailable() bool {
	_, err := exec.LookPath("wg")

	return err == nil
}

func CreateNewPrivateKey() (string, error) {
	cmd := exec.Command("wg", "genkey")

	bytes, err := cmd.Output()

	if err != nil {
		return "", err
	}

	return string(whitespacePattern.ReplaceAll(bytes, []byte(""))), nil
}

func CreatePublicKey(privateKey string) (string, error) {
	cmd := exec.Command("wg", "pubkey")
	cmd.Stdin = strings.NewReader(privateKey)

	bytes, err := cmd.Output()

	if err != nil {
		return "", err
	}

	return string(whitespacePattern.ReplaceAll(bytes, []byte(""))), nil
}

func CreatePresharedKey() (string, error) {
	cmd := exec.Command("wg", "genpsk")

	bytes, err := cmd.Output()

	if err != nil {
		return "", err
	}

	return string(whitespacePattern.ReplaceAll(bytes, []byte(""))), nil
}
