func countKDifference(nums []int, k int) int {
	pairs := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if abs(nums[i]-nums[j]) == k {
				pairs++
			}
		}
	}
	return pairs
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}