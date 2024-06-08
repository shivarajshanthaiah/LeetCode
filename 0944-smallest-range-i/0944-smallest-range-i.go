func smallestRangeI(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	min := nums[0]
	max := nums[len(nums)-1]
	newMin := min + k
	newMax := max - k
	if newMin > newMax {
		return 0
	}
	return newMax - newMin
}