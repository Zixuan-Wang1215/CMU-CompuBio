package main

// RichnessMap takes a map of frequency maps as input.  It returns a map
// whose values are the richness of each sample.
func RichnessMap(allMaps map[string](map[string]int)) map[string]int {
	outputMap := make(map[string]int)
	for name, val := range allMaps {
		outputMap[name] = Richness(val)
	}

	return outputMap
}
