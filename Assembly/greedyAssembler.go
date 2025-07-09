package main

/**
GreedyAssembler(reads)
reads2 ß copy of reads
k ß length of each string in reads2
genome ß reads2[0]
reads2 ß remove reads2[0]
while some read in reads2 has suffix = genome[0, k-1]
genome ß read[0] + genome
reads2 ß remove read
while some read in reads2 has prefix = genome[n–k+1, n]
// n is length of genome
genome ß genome + read[length(read) - 1]
reads2 ß remove read
return genome
**/

func GreedyAssembler(reads []string) string {
	reads2 := reads
	genome := reads2[0]
	len_genome := len(genome)
	reads2 = reads[1:] //?? why does it remove the first str?
	for len(reads2) > 0 {
		for i := 0; i < len(reads2); i++ {
			if HasPrefix(reads2[i], genome, len_genome, len(reads2[i])) { //has suffix is not the same thing as equal to
				genome += string(reads2[i][len(reads2[i])-1])
				len_genome++
				reads2 = RemoveRead(reads2, reads2[i])

			}
		}

		for i := 0; i < len(reads2); i++ {
			if HasSuffix(reads2[i], genome, len(reads2[i])) {

				genome = string(reads2[i][0]) + genome
				len_genome++
				reads2 = RemoveRead(reads2, reads2[i])

			}
		}
	}

	return genome

}

func RemoveRead(reads []string, read string) []string {
	for i, r := range reads {
		if r == read {
			reads = append(reads[:i], reads[i+1:]...)
			break
		}
	}
	return reads
}

// I love this
func HasSuffix(read string, genome string, len int) bool {
	if read[1:] == genome[0:(len-1)] {
		return true
	}
	return false
}

func HasPrefix(read string, genome string, len_genome int, len_read int) bool {
	if read[:len_read-1] == genome[len_genome-len_read+1:] {
		return true
	}

	return false
}
