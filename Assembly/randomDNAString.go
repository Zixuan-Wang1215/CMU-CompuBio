package main

import (
	"math/rand" // this should be helpful!
)

// this should be helpful!

// this should be helpful!

// GenerateRandomGenome takes a parameter length and returns
// a random DNA string of this length where each nucleotide has equal probability.
func GenerateRandomGenome(length int) string {
	letters := []string{"A", "C", "T", "G"}
	genome := ""
	for length > 0 {
		genome += letters[rand.Intn(3)]
		length--
	}
	return genome
}
