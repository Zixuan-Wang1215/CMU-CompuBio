package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	AnalyzeYear("2019")

}

// input: string representing a year or other identifier
// output: nothing, but we will run a metagenomics analysis and write the results of this analysis to files
func AnalyzeYear(year string) {
	//step 1: lets read in a single file
	filename := fmt.Sprintf("Data/%s_Samples/Fall_Allegheny_1.txt", year)
	filename2 := fmt.Sprintf("Data/%s_Samples/Spring_Monongahela_1.txt", year)

	freqMap := ReadFrequencyMapFromFile(filename)
	freqMap2 := ReadFrequencyMapFromFile(filename2)

	fmt.Println("File Contains: ", len(freqMap), "total keys. ")
	fmt.Println("File Conatins: ", len(freqMap2), "total keys. ")
	fmt.Println("Bray Curtis distance: ", BrayCurtisDistance(freqMap, freqMap2))

	// We want all of the samples
	// step 2: read all samples from a dir

	path := fmt.Sprintf("Data/%s_Samples", year)
	allMaps := ReadSamplesFromDirectory(path)

	fmt.Println("We have ", len(allMaps), "total samples. ")

	// step 3: process richness and eveness of all samples
	richness := RichnessMap(allMaps)
	richnessFile := fmt.Sprintf("Matrices/RichnessMatrix_%s.csv", year)
	WriteRichnessMapToFile(richness, richnessFile)

	simpson := SimpsonsMap(allMaps)
	simpsonFile := fmt.Sprintf("Matrices/SimpsonMatrix_%s.csv", year)
	WriteSimpsonsMapToFile(simpson, simpsonFile)

	distanceMetric := "Bray-Curtis"
	sampleNames, mtxBC := BetaDiversityMatrix(allMaps, distanceMetric)
	brayCurtisFile := fmt.Sprintf("Matrices/BrayCurtisBetaDiversityMatrix__%s.csv", year)
	WriteBetaDiversityMatrixToFile(mtxBC, sampleNames, brayCurtisFile)

	distanceMetric2 := "Jaccard"
	sampleNames2, mtxBC2 := BetaDiversityMatrix(allMaps, distanceMetric2)
	jaccardFile := fmt.Sprintf("Matrices/JaccardBetaDiversityMatrix__%s.csv", year)
	WriteBetaDiversityMatrixToFile(mtxBC2, sampleNames2, jaccardFile)
}
