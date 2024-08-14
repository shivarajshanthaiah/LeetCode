func maxSubarrayLength(nums []int, k int) int {
	i, j := 0, 0
	ans := 0
	freq := make(map[int]int)

	for i < len(nums) {
		freq[nums[i]]++
		for freq[nums[i]] > k {
			freq[nums[j]]--
			j++
		}
		if i-j+1 > ans {
			ans = i - j + 1
		}
        i++
	}
	return ans
}