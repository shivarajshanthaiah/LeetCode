func minimumOperations(nums []int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i]%3 != 0 {
			count++
		}
	}
	return count
}