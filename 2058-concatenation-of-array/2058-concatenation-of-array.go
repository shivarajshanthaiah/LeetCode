func getConcatenation(nums []int) []int {
	// return append(nums, nums...)
	for _, num := range nums {
		nums = append(nums, num)
	}
	return nums
}