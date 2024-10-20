func findMaxLength(nums []int) int {
	maxLen := 0
	sum := 0
	mp := make(map[int]int)
	for idx, val := range nums {
		if val == 0 {
			sum--
		} else {
			sum++
		}

		if sum == 0 {
			maxLen = idx + 1
		} else if _, ok := mp[sum]; ok {
			maxLen = max(maxLen, idx-mp[sum])
		} else {
			mp[sum] = idx
		}
	}
	return maxLen
}