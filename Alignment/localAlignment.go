package main

// LocalAlignment takes two strings, along with match, mismatch, and gap scores.
// It returns a maximum score local alignment of the strings corresponding to these penalties.
func LocalAlignment(str1, str2 string, match, mismatch, gap float64) (Alignment, int, int, int, int) {
	var alignment Alignment
	mtx := LocalScoreTable(str1, str2, match, mismatch, gap)

	// create the first indicies based on the length of the two strings
	max_val := 0.0
	row := 0
	col := 0

	for i := 1; i <= len(str1); i++ {
		for j := 1; j <= len(str2); j++ {
			if mtx[i][j] > max_val {
				row = i
				col = j
				max_val = mtx[i][j]
			}
		}
	}

	endStr1 := row
	endStr2 := col

	// loop until you hit the left or top edge
	for mtx[row][col] > 0 {
		// find the values up, left, and diagonal from the current index
		left := -1000.0
		if InMatrix(mtx, row, col-1) {
			left = mtx[row][col-1] - gap
		}

		up := -1000.0
		if InMatrix(mtx, row-1, col) {
			up = mtx[row-1][col] - gap
		}

		diag := -1000.0
		if InMatrix(mtx, row-1, col-1) {
			diag = mtx[row-1][col-1]
			if str1[row-1] == str2[col-1] {
				diag += match
			} else {
				diag -= mismatch
			}
		}

		if diag == mtx[row][col] {
			alignment[0] = string(str1[row-1]) + alignment[0]
			alignment[1] = string(str2[col-1]) + alignment[1]
			row--
			col--
		} else if up == mtx[row][col] {
			alignment[0] = string(str1[row-1]) + alignment[0]
			alignment[1] = "-" + alignment[1]
			row--
		} else if left == mtx[row][col] {
			alignment[0] = "-" + alignment[0]
			alignment[1] = string(str2[col-1]) + alignment[1]
			col--
		}

	}
	return alignment, row, endStr1, col, endStr2
}

func InMatrix(matrix [][]float64, row int, col int) bool {
	if row < 0 || col < 0 {
		return false
	}
	if row >= len(matrix) {
		return false
	}
	if col >= len(matrix[row]) {
		return false
	}
	return true
}
