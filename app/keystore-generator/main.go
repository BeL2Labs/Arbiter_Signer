// Copyright (c) 2025 The bel2 developers
package main

import (
	"bufio"
	"encoding/hex"
	"flag"
	"fmt"
	"os"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/BeL2Labs/Arbiter_Signer/app/keystore-generator/btc"
	"github.com/BeL2Labs/Arbiter_Signer/app/keystore-generator/eth"
)

func main() {
	chain := flag.String("c", "eth", "Chain type (eth or btc)")
	privateKey := flag.String("s", "", "Private key (64 hex characters)")
	password := flag.String("p", "", "Password for keystore")
	outputFile := flag.String("o", "", "Output filename (optional)")
	fileToParse := flag.String("f", "", "Keystore file to parse")
	flag.Parse()

	if *fileToParse != "" {
		privateKey, err := eth.ParseKeystore(*fileToParse, *password)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		fmt.Println("Private key:", privateKey)
	} else {
		switch *chain {
		case "eth":
			GenerateETHKeystore(*privateKey, *password, *outputFile)
		case "btc":
			GenerateBTCKeystore(*privateKey, *password, *outputFile)
		default:
			fmt.Println("Error: Invalid chain type")
			os.Exit(1)
		}
	}
}

func GenerateETHKeystore(privateKeyHex string, password string, outputFile string) {
	if privateKeyHex == "" {
		fmt.Println("Error: Private key is required")
		os.Exit(1)
	}

	if len(privateKeyHex) != 64 {
		fmt.Println("Error: Private key must be exactly 64 hex characters")
		os.Exit(1)
	}

	if password == "" {
		fmt.Print("Enter password: ")
		reader := bufio.NewReader(os.Stdin)
		inputPassword, _ := reader.ReadString('\n')
		password = inputPassword[:len(inputPassword)-1]
	}

	privateKeyBytes, err := hex.DecodeString(privateKeyHex)
	if err != nil {
		fmt.Println("Error: Invalid private key format - must be 64 hex characters")
		os.Exit(1)
	}

	privateKey, err := crypto.ToECDSA(privateKeyBytes)
	if err != nil {
		fmt.Println("Error: Failed to create private key:", err)
		os.Exit(1)
	}

	ks := keystore.NewKeyStore(".", keystore.StandardScryptN, keystore.StandardScryptP)

	account, err := ks.ImportECDSA(privateKey, password)
	if err != nil {
		fmt.Println("Error: Failed to import private key:", err)
		os.Exit(1)
	}

	if outputFile != "" {
		err := os.Rename(account.URL.Path, outputFile)
		if err != nil {
			fmt.Println("Error: Failed to rename keystore file:", err)
			os.Exit(1)
		}
		fmt.Println("Ethereum keystore created successfully at:", outputFile)
	} else {
		fmt.Println("Ethereum keystore created successfully at:", account.URL.Path)
	}
}

func GenerateBTCKeystore(privateKeyHex string, password string, outputFile string) {
	if privateKeyHex == "" {
		fmt.Println("Error: Private key is required")
		os.Exit(1)
	}

	if len(privateKeyHex) != 64 {
		fmt.Println("Error: Private key must be exactly 64 hex characters")
		os.Exit(1)
	}

	if password == "" {
		fmt.Print("Enter password: ")
		reader := bufio.NewReader(os.Stdin)
		inputPassword, _ := reader.ReadString('\n')
		password = inputPassword[:len(inputPassword)-1]
	}

	privateKeyBytes, err := hex.DecodeString(privateKeyHex)
	if err != nil {
		fmt.Println("Error: Invalid private key format - must be 64 hex characters")
		os.Exit(1)
	}

	privKey, _ := btcec.PrivKeyFromBytes(privateKeyBytes)
	wif, err := btcutil.NewWIF(privKey, &chaincfg.MainNetParams, true)
	if err != nil {
		fmt.Println("Error: Failed to create Bitcoin WIF:", err)
		os.Exit(1)
	}

	filename := fmt.Sprintf("btc_keystore_%s.txt", wif.String()[:8])
	if outputFile != "" {
		filename = outputFile
	}

	// Encrypt WIF with password
	encrypted, err := btc.Encrypt([]byte(wif.String()), password)
	if err != nil {
		fmt.Println("Error: Failed to encrypt keystore:", err)
		os.Exit(1)
	}

	if err := os.WriteFile(filename, encrypted, 0600); err != nil {
		fmt.Println("Error: Failed to save Bitcoin keystore:", err)
		os.Exit(1)
	}

	fmt.Println("Bitcoin keystore created successfully at:", filename)
}
