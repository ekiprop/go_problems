package main

import "fmt"

func isPalindrome(x int) bool {
	// Special case: Negative numbers are not palindromes
	if x < 0 {
		return false
	}

	// Variables for reversing the digits
	reversed := 0
	original := x

	// Reverse the digits of the number
	for x > 0 {
		digit := x % 10
		reversed = reversed*10 + digit
		x = x / 10
	}

	// Check if the reversed number is equal to the original number
	return original == reversed
}

func main() {
	// Example usage
	number := 1213
	result := isPalindrome(number)

	if result {
		fmt.Printf("%d is a palindrome.\n", number)
	} else {
		fmt.Printf("%d is not a palindrome.\n", number)
	}
}
