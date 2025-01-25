// Copyright (c) 2025 The bel2 developers
package eth

import (
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func ParseKeystore(filePath string, password string) (string, error) {
	// Read keystore file
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read keystore file: %v", err)
	}

	// Decrypt keystore
	key, err := keystore.DecryptKey(data, password)
	if err != nil {
		return "", fmt.Errorf("failed to decrypt keystore: %v", err)
	}

	// Get private key
	privateKeyBytes := crypto.FromECDSA(key.PrivateKey)
	return common.Bytes2Hex(privateKeyBytes), nil
}
