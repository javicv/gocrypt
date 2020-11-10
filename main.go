package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Fprintf(os.Stderr, `To encrypt/decrypt values
use: %s [-e|-d] PASSWORD VALUE
	-e: Encrypt value
	-d: Decrypt value`, os.Args[0])
		os.Exit(1)
	}
	if os.Args[1] != "-e" && os.Args[1] != "-d" {
		fmt.Fprintf(os.Stderr, "Wrong Method %s. Only -e or -d are allowed.\n", os.Args[1])
	}
	method := os.Args[1]
	password := os.Args[2]
	value := []byte(os.Args[3])
	hasher := sha256.New()
	hasher.Write([]byte(password))
	key := hasher.Sum(nil)

	fmt.Fprintf(os.Stderr, "METHOD: %s\nPASSWORD: %s\nVALUE: %s\n", method, password, value)

	var result string
	if method == "-e" {
		ciph, err := encrypt(value, key)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err.Error())
			os.Exit(2)
		}
		result = base64.StdEncoding.EncodeToString(ciph)
	} else {
		ciph := make([]byte, base64.StdEncoding.EncodedLen(len(value)))
		n, err := base64.StdEncoding.Decode(ciph, value)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err.Error())
			os.Exit(3)
		}
		ciph, err = decrypt(ciph[:n], key)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err.Error())
			os.Exit(4)
		}
		result = string(ciph)
	}
	fmt.Printf("Result: %s\n", result)
}

func encrypt(plaintext []byte, key []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, plaintext, nil), nil
}

func decrypt(ciphertext []byte, key []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("new cipher")
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		fmt.Println("new gcm")
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	return gcm.Open(nil, nonce, ciphertext, nil)
}
