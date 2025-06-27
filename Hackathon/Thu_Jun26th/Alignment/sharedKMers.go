package main

// Please do not remove package declarations because these are used by the autograder. If you need additional packages, then you may declare them above.

// Insert your CountSharedKmers() function here, along with any subroutines that you need.
func CountSharedKmers(str1, str2 string, k int) int {
	p := 0

	map1 := MakeMap(str1, k)
	map2 := MakeMap(str2, k)

	for i := range map1 {
		p += Min2(map1[i], map2[i])
	}

	return p
}

// makeMap takes in a string and an int, and makes a frequency map of the k-m
func MakeMap(input string, k int) map[string]int {
	table := make(map[string]int)

	for i := 0; i <= len(input)-k+1; i++ {
		table[input[i:i+k]]++
	}

	return table
}

func Min2(x, y int) int {
	if x < y {
		return x
	}

	return y
}
