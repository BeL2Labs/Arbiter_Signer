package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/ethereum/go-ethereum/accounts/keystore"
)

// ReadBTCKeystore reads a BTC keystore file
func ReadBTCKeystore(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// BTC keystore is binary format, return raw data
	return data, nil
}

// ReadETHKeystore reads an ETH keystore file
func ReadETHKeystore(path string) (*keystore.Key, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// Parse ETH keystore JSON
	var key keystore.Key
	if err := json.Unmarshal(data, &key); err != nil {
		return nil, err
	}

	return &key, nil
}

// ReadKeystore automatically detects and reads either BTC or ETH keystore
func ReadKeystore(path string) (interface{}, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// Try to parse as ETH keystore first
	var ethKey keystore.Key
	if err := json.Unmarshal(data, &ethKey); err == nil {
		return &ethKey, nil
	}

	// If not ETH keystore, treat as BTC keystore
	return data, nil
}

// GetKeyType determines if a keystore is BTC or ETH format
func GetKeyType(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	// Check if data is valid JSON
	var temp interface{}
	if json.Unmarshal(data, &temp) == nil {
		return "eth", nil
	}

	// Check if data is valid hex (BTC WIF)
	if _, err := hex.DecodeString(string(data)); err == nil {
		return "btc", nil
	}

	return "", errors.New("unknown keystore format")
}

// GetEthKeyFromKeystore reads an ETH keystore file and returns the private key as hex string
func GetEthKeyFromKeystore(path, password string) (string, error) {
	// Read ETH keystore file
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	// Decrypt ETH keystore
	privateKey, err := keystore.DecryptKey(data, password)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(privateKey.PrivateKey.D.Bytes()), nil
}

// GetBtcKeyFromKeystore reads a BTC keystore file and returns the private key as hex string
func GetBtcKeyFromKeystore(path, password string) (string, error) {
	data, err := ReadBTCKeystore(path)
	if err != nil {
		return "", err
	}
	// Decrypt BTC keystore using password
	decryptedData, err := decryptBTCKeystore(data, password)
	if err != nil {
		return "", err
	}
	hexPrivateKey, err := convertWIFToHex(string(decryptedData))
	if err != nil {
		return "", fmt.Errorf("failed to convert WIF to hex: %v", err)
	}
	return hexPrivateKey, nil
}

func convertWIFToHex(wif string) (string, error) {
	decodedWIF, err := btcutil.DecodeWIF(wif)
	if err != nil {
		return "", err
	}

	privateKeyBytes := decodedWIF.PrivKey.Serialize()
	hexPrivateKey := hex.EncodeToString(privateKeyBytes)
	return hexPrivateKey, nil
}

// decryptBTCKeystore decrypts BTC keystore data using password
func decryptBTCKeystore(data []byte, password string) ([]byte, error) {
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
