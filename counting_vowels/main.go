/*
Counting vowels in a string:
Write a program that takes a string as input
and counts the number of vowels in the string.
*/

package main

import "fmt"

func main() {
	var input string
	fmt.Print("Enter a string: ")
	fmt.Scanln(&input)

	count := 0
	for _, char := range input {
		switch char {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			count++
		}
	}

	fmt.Printf("The string \"%s\" contains %d vowels.\n", input, count)
}
