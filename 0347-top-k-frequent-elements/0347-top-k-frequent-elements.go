func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		freq[nums[i]]++
	}

	unique := make([]int, 0, len(freq))
	for num := range freq {
		unique = append(unique, num)
	}

	sort.Slice(unique, func(i, j int) bool {
		return freq[unique[i]] > freq[unique[j]]
	})

	return unique[:k]
}