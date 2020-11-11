package main

import (
	"testing"
)

func TestEncryptDecryptString(t *testing.T) {
	expected := "Hello World"

	passwd := "testkey"
	enc, err := encryptString(expected, passwd)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	got, err := decryptString(enc, passwd)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if got != expected {
		t.Errorf("Expected: %v, got: %v", expected, got)
	}
}
