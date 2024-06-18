func countCompleteDayPairs(hours []int) int64 {
	seen := make(map[int]int)
	var count int64

	for _, h := range hours {
		r := h % 24
		comp := (24 - r) % 24

		if compCount := seen[comp]; compCount > 0 {
			count += int64(compCount)
		}
		seen[r]++
	}
	return count
}