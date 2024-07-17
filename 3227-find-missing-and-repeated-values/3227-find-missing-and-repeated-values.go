func findMissingAndRepeatedValues(grid [][]int) []int {
	nums := make(map[int]int)
	res := []int{0, 0}

	for _, row := range grid {
		for _, val := range row {
			nums[val]++
		}
	}

	n := len(grid)
	for i := 1; i <= n*n; i++ {
		if nums[i] == 2 {
			res[0] = i
		}
		if nums[i] == 0 {
			res[1] = i
		}
	}
	return res
}