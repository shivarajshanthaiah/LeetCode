func pivotIndex(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	leftSum := 0
	for i := range nums {
		if sum-leftSum-nums[i] == leftSum {
			return i
		} else {
			leftSum += nums[i]
		}
	}
	return -1
}