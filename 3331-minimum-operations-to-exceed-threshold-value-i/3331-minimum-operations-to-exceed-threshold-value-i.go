func minOperations(nums []int, k int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] < k {
			count++
		}
	}
	return count
}