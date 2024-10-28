func longestSquareStreak(nums []int) int {
	sort.Ints(nums)

	values := make(map[int]bool)
	for _, num := range nums {
		values[num] = true
	}

	maxStreak := 0
	for _, num := range nums {
		streak := 0
		curr := num

		for values[curr] {
			streak++
			curr = curr * curr
		}
		if streak > maxStreak {
			maxStreak = streak
		}
	}
	if maxStreak > 1 {
		return maxStreak
	}
	return -1
}