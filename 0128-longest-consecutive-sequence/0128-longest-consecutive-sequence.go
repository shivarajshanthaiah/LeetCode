func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	numMap := make(map[int]bool)
	for _, num := range nums {
		numMap[num] = true
	}

	longestStreak := 0
	for num := range numMap {
		if !numMap[num-1] {
			currentNum := num
			currentStreak := 1

			for numMap[currentNum+1] {
				currentNum++
				currentStreak++
			}
			if currentStreak > longestStreak {
				longestStreak = currentStreak
			}
		}
	}
	return longestStreak
}