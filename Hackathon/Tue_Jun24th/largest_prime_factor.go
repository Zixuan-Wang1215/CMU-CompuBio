package main

import (
	"fmt"
)

// Please do not remove package declarations because these are used by the autograder. If you need additional packages, then you may declare them above.

// Insert your LargestPrimeFactor() function here, along with any subroutines that you need.
func main() {
	fmt.Println(LargestPrimeFactor(2))
}

// LargestPrimeFactor: takes in an int, and returns the greatest prime factor of n as an int
func LargestPrimeFactor(n int) int {
	// counting down, so that it returns the largest number first
	for i := n; i > 1; i-- {
		// if n is divisible by i and is prime return that number
		if isPrime(i) && n%i == 0 {
			return i
		}
	}
	return 0
}

// isPrime: takes in an int, and returns a bool value of whether n is prime or not
func isPrime(n int) bool {
	if n == 2 {
		return true
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
