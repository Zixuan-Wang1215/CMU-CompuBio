package main

// LocalScoreTable takes two strings and alignment penalties. It returns a 2-D array
// holding dynamic programming scores for local alignment with these penalties.

func Make2D_2[T any](n, m int) [][]T {
	matrix := make([][]T, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]T, m)
	}

	return matrix
}

func LocalScoreTable(str1, str2 string, match, mismatch, gap float64) [][]float64 {

	matrix := Make2D_2[float64](len(str1)+1, len(str2)+1) //str1 is the col indicies, and str2 is the row indicies

	//fmt.Println("LCS Matrix", matrix)

	// for each cell in the matrix, check first if the two letters are equal to eachother, then add one to
	// the value of the cell, if they are not the same, then take the bigger of the two adjacent cells

	//i is the col
	//j is the row

	for i := 1; i < len(str1)+1; i++ {
		for j := 1; j < len(str2)+1; j++ {
			if str1[i-1] == str2[j-1] {
				//this checked of there is an alignment match
				matrix[i][j] = matrix[i-1][j-1] + match
				continue
			}
			//this is the second case which takes the largest of the two previous
			matrix[i][j] = max(matrix[i][j-1]-gap, matrix[i-1][j]-gap, matrix[i-1][j-1]-mismatch)

			if matrix[i][j] < 0 {
				matrix[i][j] = 0
			}

		}
		//fmt.Println(matrix[i])
	}

	return matrix
}
