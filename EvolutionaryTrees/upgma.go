package main

/*
func UPGMA (D, speciesNames) {
	numLeaves <- CountRows(D)
	//count the number of rows (create the cluster array based on this)
	t <- InitializeTree(speciesNames)
	clusters <- InitializeClusters(t)
	for every integer p from numLeaves to 2*numLeaves-2 (n-1 internal nodes theorem)
		row, col, value <- FindMinElt(D)
		t.nodes[p].age <- val/2
		t.nodes[p].child1 <- clusters[row]
		t.nodes[p].child2 <- clusters[col]
		clusterSize1 <- CountLeaves(t.nodes[p].child1) // reference to some row --> counts the number of leaves beneath --> use Recursion!
		clusterSize2 <- CountLeaves(t.nodes[p].child2)

		D <- AddRowCol(D, clusterSize1, clusterSize2, row, col)
		D <- DelRowCol(D, row, col)
		clusters <- append(clusters, t[k])
		clusters <- DelCluste
s(clusters, row, col)
		return t}
*/
// UPGMA takes a distance matrix and a collection of species names as input.
// It returns a Tree (an array of nodes) resulting from applying
// UPGMA to this dataset.
func UPGMA(mtx DistanceMatrix, speciesNames []string) Tree {
	var t Tree

	numLeaves := CountRows(mtx)
	t = InitializeTree(speciesNames)
	clusters := InitializeClusters(t) // make a cluster array of len()

	//
	for p := numLeaves; p < 2*numLeaves-2; p++ {
		row, col, val := FindMinElt(D)
		t.nodes[p].age = val / 2
		t.nodes[p].child1 = clusters[row]
		t.nodes[p].child2 = clusters[col]
	}
	clusterSize1 := CountLeaves(t.nodes[p].child)
	clusterSize2 := CountLeaves(t.nodes[p].child2)

	D := AddRowCol(D, clusterSize1, clusterSize2, row, col)
	D = DelRowCol(D, row, col)

	clusters = append(clusters, t[k])
	clusters = DelClusters(clusters, row, col)

	return t
}
