func rangeSum(nums []int, n int, left int, right int) int {
	sum := []int{}
	mod := 1_000_000_007
	for i := 0; i < len(nums); i++ {
		currSum := nums[i]
		sum = append(sum, currSum)
		for j := i + 1; j < len(nums); j++ {
			currSum += nums[j]
			sum = append(sum, currSum)
		}
	}
	sort.Ints(sum)
	ans := 0
	for i := left - 1; i < right; i++ {
		ans = (ans + sum[i]) % mod
	}
	return ans
}