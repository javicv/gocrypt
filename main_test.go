package main

import (
	"crypto/sha256"
	"testing"
)

func TestEncryptDecrypt(t *testing.T) {
	expected := "Hello World"

	passwd := "testkey"
	hasher := sha256.New()
	hasher.Write([]byte(passwd))
	key := hasher.Sum(nil)
	enc, err := encrypt([]byte(expected), key)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	dec, err := decrypt(enc, key)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	got := string(dec)
	if got != expected {
		t.Errorf("Expected: %v, got: %v", expected, got)
	}
}
