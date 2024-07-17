func findMissingAndRepeatedValues(grid [][]int) []int {
	nums := make(map[int]int)

	for _, row := range grid {
		for _, val := range row {
			nums[val]++
		}
	}

	n := len(grid) * len(grid)
	repeat, missing := 0, 0
	for i := 1; i <= n; i++ {
		if nums[i] == 2 {
			repeat = i
		}
		if nums[i] == 0 {
			missing = i
		}
	}
	return []int{repeat, missing}
}