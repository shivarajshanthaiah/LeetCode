func uniqueOccurrences(arr []int) bool {
	freq := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		freq[arr[i]]++
	}

	exists := make(map[int]bool)
	for _, val := range freq {
		if exists[val] {
			return false
		}
		exists[val] = true
	}
	return true
}