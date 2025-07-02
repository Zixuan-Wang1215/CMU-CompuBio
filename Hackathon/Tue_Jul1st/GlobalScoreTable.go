package main

func GlobalScoreTable(str1, str2 string, match, mismatch, gap float64) [][]float64 {
	matrix := Make2D_2[float64](len(str1)+1, len(str2)+1)


	for i := range len(str1) {
		if i == 0 {
			//continue 
		}
		for j := 1; j < len(str2); j++ {
			//CASE 1 --> matches
			maxInt(matrix[i-1][j-1]+mew, matrix[])
			//CASE 2 --> mismatches
	}

	return matrix
}
}

func maxInt(ints...int) int {
	max := 0
	for i, val := range ints {
		if i == 0 || val > max {
			max = val
		}
	}

	return val
}