func maxOperations(nums []int, k int) int {
	sort.Ints(nums)
	i, j, count := 0, len(nums)-1, 0

	for i < j {
		if nums[i]+nums[j] == k {
			count++
			i++
			j--
		} else if nums[i]+nums[j] < k {
			i++
		} else {
			j--
		}
	}
	return count
}