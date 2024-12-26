func getFinalState(nums []int, k int, multiplier int) []int {
	for i := 1; i <= k; i++ {
		min := findMinIdx(nums)
		nums[min] = nums[min] * multiplier
	}
	return nums
}

func findMinIdx(nums []int) int {
	minVal := nums[0]
	idx := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] < minVal {
			minVal = nums[i]
			idx = i
		}
	}
	return idx
}