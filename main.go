package main

import (
	"flag"
	"fmt"
	"os"
)

const _nonceSize int = 12
const _keySize int = 16
const _saltSize int = 16
const _iterations int = 32767

var method string
var password string
var value string

func init() {
	const (
		methodDefault   = "encrypt"
		methodUsage     = "Method to use encrypt/decrypt"
		passwordDefault = ""
		passwordUsage   = "Password to encrypt/decrypt the value"
		valueDefault    = ""
		valueUsage      = "Value to encrypt/decrypt"
	)
	flag.StringVar(&method, "m", methodDefault, methodUsage)
	flag.StringVar(&method, "method", methodDefault, methodUsage)
	flag.StringVar(&password, "p", passwordDefault, passwordUsage)
	flag.StringVar(&password, "password", passwordDefault, passwordUsage)
	flag.StringVar(&value, "v", valueDefault, valueUsage)
	flag.StringVar(&value, "value", valueDefault, valueUsage)
}

func main() {
	flag.Parse()

	if method != "encrypt" && method != "decrypt" {
		fmt.Fprintf(os.Stderr, "Wrong Method '%s'. Only 'encrypt' or 'decrypt' are allowed.\n", method)
		os.Exit(1)
	}
	if len(password) == 0 {
		fmt.Fprintf(os.Stderr, "You must provide the encryption password.\n")
		flag.PrintDefaults()
		os.Exit(2)
	}
	if len(value) == 0 {
		fmt.Fprintf(os.Stderr, "One value must be provided to be encrypted/decrypted.\n")
		flag.PrintDefaults()
		os.Exit(3)
	}

	var result string
	var err error
	switch method {
	case "encrypt":
		result, err = encryptString(value, password)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err.Error())
			os.Exit(4)
		}
		break
	case "decrypt":
		result, err = decryptString(value, password)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err.Error())
			os.Exit(5)
		}
		break
	}
	fmt.Printf("%s: %s\n", value, result)
}
