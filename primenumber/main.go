/*
Prime Number Checker: Write a program that takes a number as input and checks whether it is a prime number or not.
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	n := 17
	if isPrime(n) {
		fmt.Printf("%d is a prime number", n)
	} else {
		fmt.Printf("%d is not a prime", n)
	}
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	} else if n == 2 {
		return true
	} else if n%2 == 0 {
		return false
	}

	limit := int(math.Sqrt(float64(n)))
	for i := 3; i < limit; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}
