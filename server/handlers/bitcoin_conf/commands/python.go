package commands

import "os/exec"

func isPythonAvailable() bool {
	_, err := exec.LookPath("python3")

	if err == nil {
		return true
	}

	_, err = exec.LookPath("python")

	// Check if python is python3
	if err != nil {
		return false
	}

	cmd := exec.Command("python", "--version")
	output, err := cmd.Output()
	if err == nil && string(output) == "Python 3" {
		return true
	}

	return false
}
