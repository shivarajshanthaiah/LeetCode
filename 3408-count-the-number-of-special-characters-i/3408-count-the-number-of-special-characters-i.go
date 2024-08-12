func numberOfSpecialChars(word string) int {
	charMap := make(map[rune]int)
	count := 0

	for _, char := range word {
		charMap[char]++
	}

	countMap := make(map[rune]bool)
	for _, char := range word {
		if 'a' <= char && char <= 'z' {
			if _, ok := charMap[char-('a'-'A')]; ok {
				if !countMap[char] {
					count++
					countMap[char] = true
				}
			}
		}
	}
	return count
}