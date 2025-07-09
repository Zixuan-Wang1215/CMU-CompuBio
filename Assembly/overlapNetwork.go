package main

// MakeOverlapNetwork() takes a slice of reads with match, mismatch, gap and a threshold.
// It returns adjacency lists of reads; edges are only included
// in the overlap graph is the alignment score of the two reads is at least the threshold.
func MakeOverlapNetwork(reads []string, match, mismatch, gap, threshold float64) map[string][]string {
	//Initialize an adjacency list to represent the overlap graph.
	adjacencyList := make(map[string][]string)

	scoringMatrix := OverlapScoringMatrix(reads, match, mismatch, gap)

	binarized_mtx := BinarizeMatrix(scoringMatrix, threshold)

	for i := 0; i < len(binarized_mtx); i++ {
		for j := 0; j < len(binarized_mtx[0]); j++ {
			if binarized_mtx[i][j] == 1 {
				adjacencyList[reads[i]] = append(adjacencyList[reads[i]], reads[j])
			}

		}
	}

	return adjacencyList
}
