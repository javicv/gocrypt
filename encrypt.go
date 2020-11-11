package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"

	"golang.org/x/crypto/pbkdf2"
)

func encryptString(text string, passwd string) (b64enc string, err error) {
	salt := make([]byte, _saltSize)
	_, err = rand.Read(salt)
	if err != nil {
		return
	}

	key := pbkdf2.Key([]byte(passwd), salt, _iterations, _keySize, sha256.New)

	ciphertextNonce, err := encrypt([]byte(text), key)
	if err != nil {
		return
	}

	ciphertextNonceSalt := make([]byte, 0)
	ciphertextNonceSalt = append(ciphertextNonceSalt, salt...)
	ciphertextNonceSalt = append(ciphertextNonceSalt, ciphertextNonce...)

	b64enc = base64.StdEncoding.EncodeToString(ciphertextNonceSalt)

	return
}

func encrypt(plaintext []byte, key []byte) (ciphertextNonce []byte, err error) {
	nonce := make([]byte, _nonceSize)
	_, err = rand.Read(nonce)
	if err != nil {
		return
	}

	// Create the cipher and block.
	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	cipher, err := cipher.NewGCM(block)
	if err != nil {
		return
	}

	ciphertext := cipher.Seal(nil, nonce, plaintext, nil)
	ciphertextNonce = make([]byte, 0)

	ciphertextNonce = append(ciphertextNonce, nonce...)
	ciphertextNonce = append(ciphertextNonce, ciphertext...)

	return
}
