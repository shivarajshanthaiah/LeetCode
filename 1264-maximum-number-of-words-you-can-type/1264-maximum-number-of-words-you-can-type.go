func canBeTypedWords(text string, brokenLetters string) int {
	words := []string{}
	idx := 0
	for i := 0; i < len(text); i++ {
		if text[i] == ' ' {
			if idx != i {
				words = append(words, text[idx:i])
			}
			idx = i + 1
		}
	}

	if idx < len(text) {
		words = append(words, text[idx:])
	}

	count := 0
	for _, word := range words {
		isBroken := false
		for _, ch := range word {
			if strings.ContainsRune(brokenLetters, ch) {
				isBroken = true
				break
			}
		}
		if isBroken != true {
			count++
		}
	}
	return count
}