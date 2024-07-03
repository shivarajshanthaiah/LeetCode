func minDifference(nums []int) int {
	n := len(nums)
	a, b := 0, 0

	if n < 5 {
		return 0
	}

	ans := math.MaxInt64
	sort.Ints(nums)
	a = nums[1]
	b = nums[n-3]
	ans = min(ans, b-a)

	a = nums[2]
	b = nums[n-2]
	ans = min(ans, b-a)

	a = nums[3]
	b = nums[n-1]
	ans = min(ans, b-a)

	a = nums[0]
	b = nums[n-4]
	ans = min(ans, b-a)

	return ans
}