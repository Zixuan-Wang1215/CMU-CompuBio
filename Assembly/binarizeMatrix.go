package main

// BinarizeMatrix takes a matrix of values and a threshold.
// It binarizes the matrix accordin g to the threshold.
// If entries across main diagonal are both above threshold, only retain the bigger one.
func BinarizeMatrix(mtx [][]float64, threshold float64) [][]int {

	int_mtx := Make2D_2[int]()

	for i := 0; i < len(mtx); i++ {
		for j := 0; j < len(mtx[0]); j++ {
			if mtx[i][j] >= threshold && mtx[j][i] >= threshold {
				max_val := max(mtx[i][j], mtx[j][i])

				if max_val >= threshold {
					mtx[i][j] = 1
				} else {
					mtx[i][j] = 0
				}
			} else {
				if mtx[i][j] >= threshold {
					mtx[i][j] = 1
				} else {
					mtx[i][j] = 0
				}
			}
		}
	}

	return mtx
}
