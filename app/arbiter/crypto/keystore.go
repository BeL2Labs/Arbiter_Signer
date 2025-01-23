// Copyright (c) 2025 The bel2 developers

package crypto

import (
	"fmt"
	"os/exec"
)

func GetKeyFromKeystore(encryptedFile string, password string) (string, error) {
	cmd := exec.Command("openssl", "enc", "-d", "-aes-256-cbc", "-in", encryptedFile, "-pass", fmt.Sprintf("pass:%s", password))
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(output), nil
}
