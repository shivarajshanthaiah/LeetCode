func uncommonFromSentences(s1 string, s2 string) []string {

	words1 := strings.Fields(s1)
	words2 := strings.Fields(s2)

	wordFreq := make(map[string]int)

	for _, word := range words1 {
		wordFreq[word]++
	}

	for _, word := range words2 {
		wordFreq[word]++
	}

	ans := []string{}

	for word, count := range wordFreq {
		if count == 1 {
			ans = append(ans, word)
		}
	}
	return ans
}