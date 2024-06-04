func frequencySort(s string) string {
	freq := make(map[rune]int)

	for _, char := range s {
		freq[char]++
	}
	var unique []rune
	var frequency []int

	for key, val := range freq {
		unique = append(unique, key)
		frequency = append(frequency, val)
	}

	sort.Slice(unique, func(i, j int) bool {
		if freq[unique[i]] == freq[unique[j]] {
			return unique[i] < unique[j]
		}
		return freq[unique[i]] > freq[unique[j]]
	})

	var res []rune
	for _, char := range unique {
		for i := 0; i < freq[char]; i++ {
			res = append(res, char)
		}
	}
	return string(res)
}