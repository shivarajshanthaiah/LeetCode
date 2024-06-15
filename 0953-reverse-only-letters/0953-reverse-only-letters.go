func reverseOnlyLetters(s string) string {
	temp := []rune(s)
	i, j := 0, len(s)-1

	isLetter := func(c rune) bool {
		return ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z')
	}

	for i < j {
		if !isLetter(temp[i]) {
			i++
			continue
		}
		if !isLetter(temp[j]) {
			j--
			continue
		}

		temp[i], temp[j] = temp[j], temp[i]
		i++
		j--
	}
	return string(temp)
}