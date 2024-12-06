func maxCount(banned []int, n int, maxSum int) int {
	ban := make(map[int]bool)

	for i := 0; i < len(banned); i++ {
		ban[banned[i]] = true
	}

	count := 0
	sum := 0
	for i := 1; i <= n; i++ {
		if !ban[i] && i <= n {
			sum += i
			if sum <= maxSum {
				count++
			}
		}
	}
	return count
}