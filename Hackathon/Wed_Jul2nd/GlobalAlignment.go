package main

import (
	"fmt"
)

type Alignment [2]string

func main() {
	fmt.Println("Global Alignment")
}

func InitializeMatrix(mtx [][]float64, gap float64) [][]float64 {
	for row := 1; row < len(mtx); row++ {
		mtx[row][0] = float64(row) * gap
	}

	for col := 1; col < len(mtx[0]); col++ {
		mtx[0][col] = float64(col) * gap
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

func GlobalScoreTable(str1, str2 string, match, mismatch, gap float64) [][]float64 {

	matrix := InitializeMatrix(Make2D_2[float64](len(str1)+1, len(str2)+1), gap*-1) //str1 is the col indicies, and str2 is the row indicies

	// for each cell in the matrix, check first if the two letters are equal to eachother, then add one to
	// the value of the cell, if they are not the same, then take the bigger of the two adjacent cells

	//i is the col
	//j is the row
	for i := 1; i < len(str1)+1; i++ {
		for j := 1; j < len(str2)+1; j++ {
			if str1[i-1] == str2[j-1] {
				//this checked of there is an alignment match
				matrix[i][j] = matrix[i-1][j-1] + match
			}
			//this is the second case which takes the largest of the two previous
			matrix[i][j] = max(matrix[i][j-1]-gap, matrix[i-1][j]-gap, matrix[i-1][j-1]-mismatch)

		}
		//fmt.Println(matrix[i])
	}

	return matrix
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

func GlobalAlignment(str1, str2 string, match, mismatch, gap float64) Alignment {
	var alignment Alignment
	mtx := GlobalScoreTable(str1, str2, match, mismatch, gap)

	// create the first indicies based on the length of the two strings
	row := len(str1)
	col := len(str2)

	// loop until you hit the left or top edge
	for row != 0 || col != 0 {
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

		// if there's a match use it! Otherwise go up or left if they are the same
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
	return alignment
}
