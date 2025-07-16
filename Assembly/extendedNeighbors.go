package main

import (
	"slices"
)

// GetExtendedNeighbors takes in a pattern (read), the overlap graph and maxK.
// It returns the extendedNeighbors list. For each neighbor *n* in this list,
// distance between n and pattern is between 2 to maxK.
func GetExtendedNeighbors(pattern string, adjList map[string][]string, maxK int) []string {
	neighbors := make([]string, 0)
	depth := 0

	parentNodes := make([]string, 1)
	parentNodes[0] = pattern

	for depth < maxK {
		childNodes := findNeighbors(adjList, parentNodes)

		if slices.Contains(childNodes, pattern) {
			return neighbors
		}

		parentNodes = childNodes
		depth++
		for _, node := range childNodes {
			if !slices.Contains(neighbors, node) && depth > 1 {
				neighbors = append(neighbors, node)
			}
		}
	}

	return neighbors
}

func findNeighbors(adjList map[string][]string, parentNodes []string) []string {
	p := make([]string, 0)
	for _, node := range parentNodes {
		p = append(p, adjList[node]...)
	}

	return p
}
