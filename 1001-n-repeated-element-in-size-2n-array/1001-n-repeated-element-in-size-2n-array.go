func repeatedNTimes(nums []int) int {
	n := len(nums) / 2
	freq := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		freq[nums[i]]++
	}

	for key, val := range freq {
		if val == n {
			return key
		}
	}
    return n
}