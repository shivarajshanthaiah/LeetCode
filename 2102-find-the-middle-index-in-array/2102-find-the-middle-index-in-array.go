func findMiddleIndex(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	leftSum := 0
	for key, val := range nums {
		if sum-leftSum-val == leftSum {
			return key
		} else {
			leftSum += val
		}
	}
	return -1
}