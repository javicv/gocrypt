package main

import (
	"fmt"
	"os"
)

const _nonceSize int = 12
const _keySize int = 16
const _saltSize int = 16
const _iterations int = 32767

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
	value := os.Args[3]

	var result string
	var err error
	if method == "-e" {
		result, err = encryptString(value, password)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err.Error())
			os.Exit(2)
		}
	} else {
		result, err = decryptString(value, password)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err.Error())
			os.Exit(4)
		}
	}
	fmt.Printf("%s: %s\n", value, result)
}
