func longestOnes(nums []int, k int) int {
	maxOnes, left, zeros := 0, 0, 0

	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			zeros++
		}
		for zeros > k {
			if nums[left] == 0 {
				zeros--
			}
			left++
		}
		maxOnes = max(right-left+1, maxOnes)
	}
	return maxOnes
}