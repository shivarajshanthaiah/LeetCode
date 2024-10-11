func findLHS(nums []int) int {
	maxLen := 0
	count := make(map[int]int)

	for _, num := range nums {
		count[num]++
	}

	for key, val := range count {
		curLen := 0
		if count[key+1] > 0 {
			curLen = val + count[key+1]
		}

		if curLen > maxLen {
			maxLen = curLen
		}

	}
	return maxLen
}