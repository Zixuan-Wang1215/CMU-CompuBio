package main

// Min takes an arbitrary number of integers and returns their minimum.
func Min(nums ...int) int {
	if len(nums) == 0 {
		panic("Error: no inputs given to Max.")
	}
	m := nums[0]
	// nums gets converted to an array
	for _, val := range nums {
		if val < m {
			// update m
			m = val
		}
	}
	return m
}

// MaxInts
// Input: an arbitrary number of integers
// Output: their maximum
func MaxInts(nums ...int) int {
	//nums is a slice
	if len(nums) == 0 {
		panic("No integers given.")
	}
	m := 0

	for i, val := range nums {
		if i == 0 || val > m {
			// update m if we are at the first integer
			// OR if current element is bigger than known max
			m = val
		}
	}

	return m
}
