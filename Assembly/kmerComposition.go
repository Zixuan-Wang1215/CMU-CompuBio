package main

// KmerComposition returns the k-mer composition (all k-mer substrings) of a given genome.
func KmerComposition(genome string, k int) []string {

	checkViability(k, genome)

	kmers := make([]string, 0)
	for i := 0; i < len(genome)-k+1; i++ {
		kmers = append(kmers, genome[i:i+k])
	}

	return kmers
}

func checkViability(k int, genome string) {
	if k <= 0 {
		panic("k must be a positive integer.")
	}

	if k >= len(genome) {
		panic("k cannot be longer than or equal to the lenth of genome.")
	}
}
