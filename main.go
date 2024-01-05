package main

import (
	"bytes"
	"fmt"
	"github.com/EthanHosier/go-file-encrypt-decrypt/filecrypt"
	"golang.org/x/term"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(0)
	}
	function := os.Args[1]

	switch function {
	case "help":
		printHelp()
	case "encrypt":
		encryptHandle()
	case "decrypt":
		decryptHandle()
	default:
		fmt.Println("Run 'encrypt' to encrypt a file or 'decrypt' to decrypt a file.")
		os.Exit(1)
	}
}

func printHelp() {
	fmt.Printf("file encryption\nSimple file encryptor for your day to day use.\n\n Usage:\n\n go run . encrypt <file> <password>\n\tEncrypts a file given a password\n go run . decrypt <file> <password>\tTries to decrypt a file using a password\n help\tDisplays help text \n\n")
}

func encryptHandle() {
	if len(os.Args) < 3 {
		fmt.Println("Missing a path to the file. For more info run go . help")
		os.Exit(0)
	}

	file := os.Args[2]
	if !validateFile(file) {
		panic("File not found")
	}
	password := getPassword()
	fmt.Println("Encrypting file...")
	filecrypt.Encrypt(file, password)
	fmt.Println("File sucessfully encrypted")
}

func decryptHandle() {
	if len(os.Args) < 3 {
		fmt.Println("Missing a path to the file. For more info run go . help")
		os.Exit(0)
	}
	file := os.Args[2]
	if !validateFile(file) {
		panic("File not found")
	}

	fmt.Println("Enter password:")
	password, _ := term.ReadPassword(0)
	fmt.Println("Decrypting file...")
	filecrypt.Decrypt(file, password)
	fmt.Print("File sucessfully decrypted")
}

func getPassword() []byte {
	fmt.Println("Enter a password")
	password, _ := term.ReadPassword(0)
	fmt.Println("Confirm password")
	password2, _ := term.ReadPassword(0)
	if !validatePassword(password, password2) {
		fmt.Println("Passwords do not match")
		return getPassword()
	}
	return password
}

func validatePassword(password1 []byte, password2 []byte) bool {
	if bytes.Equal(password1, password2) {
		return true
	}
	return false
}

func validateFile(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}
	return true
}
