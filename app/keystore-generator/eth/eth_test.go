// Copyright (c) 2025 The bel2 developers
package eth

import (
	"os"
	"testing"
)

func TestParseKeystore(t *testing.T) {
	// Create test keystore file
	testKeystore := `{
		"address": "3f5ce5fbfe3e9af3971dd833d26ba9b5c936f0be",
		"crypto": {
			"cipher": "aes-128-ctr",
			"ciphertext": "5318b4d5bcd28de64ee5559e671353e16f075ecae9f99c7a79a38af5f869aa46",
			"cipherparams": {
				"iv": "6087dab2f9fdbbfaddc31a909735c1e6"
			},
			"kdf": "scrypt",
			"kdfparams": {
				"dklen": 32,
				"n": 262144,
				"p": 1,
				"r": 8,
				"salt": "ae3cd4e7013836a3df6bd7241b12db061dbe2c6785853cce422d148a624ce0bd"
			},
			"mac": "517ead924a9d0dc3124507e3393d175ce3ff7c1e96529c6c555ce9e51205e9b2"
		},
		"id": "e13b209c-3b2f-4327-bab0-3bef2e51630d",
		"version": 3
	}`
	
	err := os.WriteFile("test_keystore.json", []byte(testKeystore), 0600)
	if err != nil {
		t.Fatalf("Failed to create test keystore file: %v", err)
	}
	defer os.Remove("test_keystore.json")

	// Test parsing
	expectedPrivateKey := "4c0883a69102937d6231471b5dbb6204fe5129617082792ae468d01a3f362318"
	got, err := ParseKeystore("test_keystore.json", "testpassword")
	if err != nil {
		t.Errorf("ParseKeystore() error = %v", err)
		return
	}
	if got != expectedPrivateKey {
		t.Errorf("ParseKeystore() = %v, want %v", got, expectedPrivateKey)
	}
}
