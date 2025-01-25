// Copyright (c) 2025 The bel2 developers
package btc

import (
	"os"
	"testing"
)

func TestParseKeystore(t *testing.T) {
	// Create test keystore file
	testWIF := "cV1F7Nfk8ZNZ1PjYz3z7J7q1z7J7q1z7J7q1z7J7q1z7J7q1z7J7q1z7J7q"
	err := os.WriteFile("test_keystore.txt", []byte(testWIF), 0600)
	if err != nil {
		t.Fatalf("Failed to create test keystore file: %v", err)
	}
	defer os.Remove("test_keystore.txt")

	// Test parsing
	got, err := ParseKeystore("test_keystore.txt", "password")
	if err != nil {
		t.Errorf("ParseKeystore() error = %v", err)
		return
	}
	if got != testWIF {
		t.Errorf("ParseKeystore() = %v, want %v", got, testWIF)
	}
}
