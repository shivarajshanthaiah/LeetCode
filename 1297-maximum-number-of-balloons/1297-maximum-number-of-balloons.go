func maxNumberOfBalloons(text string) int {
	char := make(map[byte]int, len(text))
	for i := 0; i < len(text); i++ {
		char[text[i]]++
	}

	res := 0
	word := "balloon"
	for {
		for i := 0; i < len(word); i++ {
			char[word[i]]--
			if char[word[i]] < 0 {
				return res
			}
		}
		res++
	}
}