package wireguard

import (
	"os/exec"
	"regexp"
	"strings"
)

var whitespacePattern = regexp.MustCompile(`[\s\n]+`)

func areWireguardToolsAvailable() bool {
	_, err := exec.LookPath("wg")

	return err == nil
}

func createNewPrivateKey() (string, error) {
	cmd := exec.Command("wg", "genkey")

	bytes, err := cmd.Output()

	if err != nil {
		return "", err
	}

	return string(whitespacePattern.ReplaceAll(bytes, []byte(""))), nil
}

func createPublicKey(privateKey string) (string, error) {
	cmd := exec.Command("wg", "pubkey")
	cmd.Stdin = strings.NewReader(privateKey)

	bytes, err := cmd.Output()

	if err != nil {
		return "", err
	}

	return string(whitespacePattern.ReplaceAll(bytes, []byte(""))), nil
}

func createPresharedKey() (string, error) {
	cmd := exec.Command("wg", "genpsk")

	bytes, err := cmd.Output()

	if err != nil {
		return "", err
	}

	return string(whitespacePattern.ReplaceAll(bytes, []byte(""))), nil
}
