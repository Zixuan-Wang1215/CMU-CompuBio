package main

// LCSLength takes two strings as input. It returns the length of a longest common
// subsequence of the two strings.
func EditDistance(str1, str2 string) int {

	matrix := InitializeMatrix(Make2D_2[int](len(str1)+1, len(str2)+1)) //str1 is the col indicies, and str2 is the row indicies

	// the value of the cell, if they are not the same, then take the bigger of the two adjacent cells

	//i is the col
	//j is the row

	for i := 1; i <= len(str1); i++ {
		for j := 1; j <= len(str2); j++ {
			if str1[i-1] == str2[j-1] {
				//this checked of there is an alignment match
				matrix[i][j] = matrix[i-1][j-1]
			} else {
				//this is the second
				matrix[i][j] = min(matrix[i-1][j-1]+1, min(matrix[i][j-1]+1, matrix[i-1][j]+1))
			}
			//fmt.Println(matrix)
		}
	}

	return matrix[len(str1)][len(str2)]
}

// Input: indices i and j which you want to find --> input the sink
// Output: longest (int terms of weight) path

func InitializeMatrix(mtx [][]int) [][]int {
	for row := 1; row < len(mtx); row++ {
		mtx[row][0] = row
	}

	for col := 1; col < len(mtx[0]); col++ {
		mtx[0][col] = col
	}
	return mtx
}

func Make2D_2[T any](n, m int) [][]T {
	matrix := make([][]T, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]T, m)
	}

	return matrix
}
func EditDistanceMatrix(patterns []string) [][]int {

	matrix := Make2D_2[int](len(patterns), len(patterns)) //str1 is the col indicies, and str2 is the row indicies

	// the value of the cell, if they are not the same, then take the bigger of the two adjacent cells

	//i is the col
	//j is the row

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if j != i && matrix[i][j] == 0 {
				matrix[i][j] = EditDistance(patterns[i], patterns[j])
				matrix[j][i] = EditDistance(patterns[i], patterns[j])
			}
		}
	}

	return matrix
}
