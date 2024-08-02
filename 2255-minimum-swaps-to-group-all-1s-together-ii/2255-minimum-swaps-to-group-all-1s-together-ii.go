func minSwaps(nums []int) int {
	window := 0
	for _, num := range nums {
		if num == 1 {
			window++
		}
	}
	swaps := 0
	for i := len(nums) - window; i < len(nums); i++ {
		if nums[i] == 0 {
			swaps++
		}
	}

	minSwaps := swaps
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			swaps++
		}
		last := i - window
		if last < 0 {
			last += len(nums)
		}
		if nums[last] == 0 {
			swaps--
		}
		if swaps < minSwaps {
			minSwaps = swaps
		}
	}
	return minSwaps
}