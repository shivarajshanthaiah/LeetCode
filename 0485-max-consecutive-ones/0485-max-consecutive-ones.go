func findMaxConsecutiveOnes(nums []int) int {
	maxWin := 0
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			count++
		} else if nums[i] == 0 {
			count = 0
		}

		if count > maxWin {
			maxWin = count
		}
	}
	return maxWin
}