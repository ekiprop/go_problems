/*
Finding the nth Fibonacci number: Write a function that takes
an integer n and returns the nth Fibonacci number.
*/

package main

import "fmt"

func main() {
	n := 4
	fib := fibonacci(n)
	fmt.Printf("The %dth Fibonacci number is: %d\n", n, fib)
}

func fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}
