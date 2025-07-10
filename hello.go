package cmucompubio
package main

import (
	"fmt"
	"math"
)

func main() {
	IsPrime(4)
}

func IsPrime(p int) bool {
	sqrtP := int(math.Sqrt(float64(p)))
	if p == 1 {
		return false
	}
	fmt.Println("Square root:", sqrtP)
	for i := 2; i <= sqrtP; i++ {
		fmt.Println("Checking:", i)
		if p%i == 0 {
			return false
		}
	}
	return true
}

