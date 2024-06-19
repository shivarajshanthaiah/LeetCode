func findPeakElement(nums []int) int {
	for i := 1; i < len(nums)-1; i++ {
		if nums[i-1] < nums[i] && nums[i] > nums[i+1] {
			return i
		}
	}
	var maxIn int
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] {
			maxIn = i + 1
		}
	}
	return maxIn
}