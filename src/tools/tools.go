package tools

// Max returns the maximum integer from a slice of integers.
// If the slice is empty, then a panic occurs.
func Max(nums ...int) int {
	if len(nums) == 0 {
		panic("max: cannot find maximum of an empty slice")
	}

	maximum := nums[0]
	for _, num := range nums[1:] {
		if num > maximum {
			maximum = num
		}
	}
	return maximum
}

