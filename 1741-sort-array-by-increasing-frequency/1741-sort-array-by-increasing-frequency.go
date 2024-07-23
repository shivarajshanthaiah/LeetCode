func frequencySort(nums []int) []int {
	freq := map[int]int{}
	for i := 0; i < len(nums); i++ {
		freq[nums[i]]++
	}

	sort.Slice(nums, func(i, j int) bool {
		if freq[nums[i]] == freq[nums[j]] {
			return nums[i] > nums[j]
		}
		return freq[nums[i]] < freq[nums[j]]
	})
	return nums
}