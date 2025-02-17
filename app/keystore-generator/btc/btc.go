// Copyright (c) 2025 The bel2 developers
package btc

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/btcsuite/btcd/btcutil"
)

func convertWIFToHex(wif string) (string, error) {
	decodedWIF, err := btcutil.DecodeWIF(wif)
	if err != nil {
		return "", err
	}

	privateKeyBytes := decodedWIF.PrivKey.Serialize()
	hexPrivateKey := hex.EncodeToString(privateKeyBytes)
	return hexPrivateKey, nil
}

func ParseKeystore(filePath string, password string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read keystore file: %v", err)
	}

	// Decrypt the data
	decrypted, err := decrypt(data, password)
	if err != nil {
		return "", fmt.Errorf("failed to decrypt keystore: %v", err)
	}
	hexPrivateKey, err := convertWIFToHex(string(decrypted))
	if err != nil {
		return "", fmt.Errorf("failed to convert WIF to hex: %v", err)
	}
	return hexPrivateKey, nil
}

func Encrypt(data []byte, password string) ([]byte, error) {
	key := sha256.Sum256([]byte(password))
	block, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext, nil
}

func decrypt(data []byte, password string) ([]byte, error) {
	key := sha256.Sum256([]byte(password))
	block, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(data) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}

	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
