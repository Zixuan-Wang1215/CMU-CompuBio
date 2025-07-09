package main

// Please do not remove package declarations because these are used by the autograder. If you need additional packages, then you may declare them above.

// Insert your BinarizeMatrix() function here, along with any subroutines that you need.
func BinarizeMatrix(mtx [][]float64, threshold float64) [][]int {
	matrix := Make2D_2[int](len(mtx), len(mtx[0]))

	for i := 0; i < len(mtx); i++ {
		for j := 0; j < len(mtx[0]); j++ {
			if mtx[i][j] >= threshold {

				matrix[i][j] = 1
			} else {
				matrix[i][j] = 0
			}

		}

	}
	for i := 0; i < len(mtx); i++ {
		for j := 0; j < len(mtx[0]); j++ {
			if matrix[i][j] == 1 {
				if (mtx[i][j] >= threshold && (mtx[i][j] > mtx[j][i])) || (mtx[i][j] > threshold && mtx[i][j] == mtx[j][i] && i < j) {
					/**
					if i == 2 && j == 0 {

						fmt.Println(mtx[i][j] >= threshold && mtx[i][j] > mtx[j][i])
						fmt.Println(mtx[i][j], mtx[j][i])
						fmt.Println(mtx[i][j] > threshold && mtx[i][j] == mtx[j][i] && i < j)

					}**/
					matrix[i][j] = 1
				} else {
					matrix[i][j] = 0
				}
			}
		}
	}
	return matrix
}
