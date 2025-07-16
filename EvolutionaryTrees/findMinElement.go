package main

// FindMinElement takes a distance matrix as input.
// It returns a pair (row, col, val) where (row, col) corresponds to the minimum
// value of the matrix, and val is the minimum value.
// NOTE: you should have that row < col.
func FindMinElement(mtx DistanceMatrix) (int, int, float64) {
	var element float64
	min_row := 0
	min_col := 0
	for row := 0; row < len(mtx); row++ {
		for col := row + 1; col < len(mtx[row]); col++ {
			if mtx[row][col] < element {
				element = min(element, mtx[row][col])
				min_row = row
				min_col = col
			}

		}
	}

	return min_row, min_col, element
}
