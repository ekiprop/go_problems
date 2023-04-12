/*
Reverse a string: Write a function that takes a string as input and returns the string reversed.
*/

package main

import "fmt"

func main() {
	str := "hello world"
	reversed := reverseString(str)
	fmt.Println(reversed)

}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
