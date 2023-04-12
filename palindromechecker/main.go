/*
Palindrome Checker: Write a program that
takes a string as input and checks
whether it is a palindrome or not.
*/

package main

import "fmt"

func main() {
	str := "racecar"
	if isPalindrome(str) {
		fmt.Printf("%s is a palindrome", str)
	} else {
		fmt.Printf("%s is not a palindrome", str)
	}
}

func isPalindrome(s string) bool {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}
	return true
}
