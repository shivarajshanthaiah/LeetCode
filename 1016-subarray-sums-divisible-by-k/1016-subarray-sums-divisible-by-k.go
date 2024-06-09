func subarraysDivByK(nums []int, k int) int {
	ans := 0
	sum := 0
	count := make(map[int]int)
	count[0] = 1
	for i := 0; i < len(nums); i++ {
		sum = (((sum + nums[i]) % k) + k) % k
		ans += count[sum]
		count[sum]++
	}
	return ans
}