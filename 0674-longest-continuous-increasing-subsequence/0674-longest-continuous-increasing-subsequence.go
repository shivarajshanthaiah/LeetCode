func findLengthOfLCIS(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	longest := 1
	curr := 1
	for i := 1; i < len(nums); i++ {
		if nums[i-1] < nums[i] {
			curr++
		} else {
			curr = 1
		}
		if curr > longest {
			longest = curr
		}
	}
	return longest
}