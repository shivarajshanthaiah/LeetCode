func commonChars(words []string) []string {
	charCount := make(map[rune]int)
	for _, char := range words[0] {
		charCount[char]++
	}

	for _, word := range words[1:] {
		wordCount := make(map[rune]int)
		for _, char := range word {
			wordCount[char]++
		}
		for char, count := range charCount {
			if wordCount[char] < count {
				charCount[char] = wordCount[char]
			}
		}
	}

	res := []string{}
	for char, count := range charCount {
		for i := 0; i < count; i++ {
			res = append(res, string(char))
		}
	}
	return res
}