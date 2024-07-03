func repeatedNTimes(nums []int) int {
	n := len(nums) / 2
	freq := make(map[int]int)
	for _, val := range nums {
		freq[val]++
	}

	for key, val := range freq {
		if val == n {
			return key
		}
	}
	return -1
}