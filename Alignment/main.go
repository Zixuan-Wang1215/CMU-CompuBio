package main

import (
	"fmt"
)

func main() {
	fmt.Println("Sequence alignment!")

	//Debug()

	ShortSARSDemo()
}

func Debug() {
	str1 := "GACT"
	str2 := "ATG"

	lcs := LongestCommonSubsequence(str1, str2)
	fmt.Println(lcs)
}

func ShortSARSDemo() {
	sars := ReadFASTAFile("Data/Coronaviruses/SARS-CoV_genome.fasta")
	sars2 := ReadFASTAFile("Data/Coronaviruses/SARS-CoV-2_genome.fasta")

	fmt.Println("LCS length between genomes:", LCSLength(sars, sars2))

	fmt.Println("Edit distance between genomes:", EditDistance(sars, sars2))

	fmt.Println("LCS between genomes:", LongestCommonSubsequence(sars, sars2))
}
