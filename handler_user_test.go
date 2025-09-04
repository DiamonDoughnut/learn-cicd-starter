package main

import (
	"strings"
	"testing"
)

func TestGenerateRandomSHA256Hash(t *testing.T) {
	hash1, err := generateRandomSHA256Hash()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	hash2, err := generateRandomSHA256Hash()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if hash1 == hash2 {
		t.Error("expected different hashes, got same hash")
	}

	if len(hash1) != 64 {
		t.Errorf("expected hash length 64, got %d", len(hash1))
	}

	// Check if hash contains only hex characters
	for _, char := range hash1 {
		if !strings.ContainsRune("0123456789abcdef", char) {
			t.Errorf("hash contains non-hex character: %c", char)
		}
	}
}