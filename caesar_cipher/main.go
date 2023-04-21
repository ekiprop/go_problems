package main

import (
	"fmt"
)

func caesarCipher(text string, shift int, action string) string {
	result := ""
	for _, char := range text {
		if char >= 'a' && char <= 'z' {
			if action == "encrypt" {
				result += string((int(char-'a')+shift)%26 + 'a')
			} else if action == "decrypt" {
				result += string((int(char-'a')-shift+26)%26 + 'a')
			}
		} else if char >= 'A' && char <= 'Z' {
			if action == "encrypt" {
				result += string((int(char-'A')+shift)%26 + 'A')
			} else if action == "decrypt" {
				result += string((int(char-'A')-shift+26)%26 + 'A')
			}
		} else {
			result += string(char)
		}
	}
	return result
}

func main() {
	var message string
	fmt.Print("Enter the message to be encrypted/decrypted: ")
	fmt.Scanln(&message)

	var key int
	fmt.Print("Enter the key (an integer between 1 and 25): ")
	fmt.Scanln(&key)

	var action string
	fmt.Print("Enter 'encrypt' or 'decrypt': ")
	fmt.Scanln(&action)

	if action == "encrypt" {
		encryptedMessage := caesarCipher(message, key, action)
		fmt.Printf("The encrypted message is: %s\n", encryptedMessage)
	} else if action == "decrypt" {
		decryptedMessage := caesarCipher(message, key, action)
		fmt.Printf("The decrypted message is: %s\n", decryptedMessage)
	} else {
		fmt.Println("Invalid action. Please enter 'encrypt' or 'decrypt'.")
	}
}
