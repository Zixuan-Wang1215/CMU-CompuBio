package main

// SimpsonsMap takes a map mapping sample IDs to frequency maps.package solution
// It returns a map of sample IDs to Simpson's indices for each sample.
func SimpsonsMap(allMaps map[string](map[string]int)) map[string]float64 {
	outputMap := make(map[string]float64)
	for name, val := range allMaps {
		outputMap[name] = SimpsonsIndex(val)
	}

	return outputMap
}
