func numIdenticalPairs(nums []int) int {
	pairs := 0
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if nums[i] == nums[j] && i < j {
				pairs++
			}
		}
	}
	return pairs
}