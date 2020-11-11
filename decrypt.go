package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"

	"golang.org/x/crypto/pbkdf2"
)

func decryptString(b64CiphertextNonceSalt string, passwd string) (text string, err error) {
	ciphertextNonceSalt, err := base64.StdEncoding.DecodeString(b64CiphertextNonceSalt)
	if err != nil {
		return
	}

	salt := ciphertextNonceSalt[:_saltSize]
	ciphertextNonce := ciphertextNonceSalt[_saltSize:]

	key := pbkdf2.Key([]byte(passwd), salt, _iterations, _keySize, sha256.New)

	plaintext, err := decrypt(ciphertextNonce, key)
	if err != nil {
		return
	}

	text = string(plaintext)

	return
}

func decrypt(ciphertextNonce []byte, key []byte) (plaintext []byte, err error) {
	// Create slices pointing to the ciphertext and nonce.
	nonce := ciphertextNonce[:_nonceSize]
	ciphertext := ciphertextNonce[_nonceSize:]

	// Create the cipher and block.
	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	cipher, err := cipher.NewGCM(block)
	if err != nil {
		return
	}

	// Decrypt and return result.
	plaintext, err = cipher.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return
	}

	return
}
