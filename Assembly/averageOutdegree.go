package main

// AverageOutDegree takes the adjacency list of a directed network.
// It returns the average outdegree of each node in the network.
// It does not include nodes with no outgoing edges in the average.
func AverageOutDegree(adjList map[string][]string) float64 {
	count := 0

	for i := range adjList {
		count += len(adjList[i])
	}

	return float64(count) / float64(len(adjList))

}

// AverageOutDegreeAllNodes takes as input the adjacency list of a directed network.
// It returns the average outdegree of each node in the network, including nodes with outdegree zero.
func AverageOutDegreeAllNodes(adjList map[string][]string) float64 {
	return 0.0
}
