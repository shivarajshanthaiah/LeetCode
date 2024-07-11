func maxVowels(s string, k int) int {
	str := []rune(s)
	curVow:= 0

	for i := 0; i < k; i++ {
		if isVowel(str[i]) {
			curVow++
		}
	}

	maxVow := curVow
	for i := k; i < len(str); i++ {
		if isVowel(str[i]) {
			curVow++
		}
		if isVowel(str[i-k]) {
			curVow--
		}
		if curVow > maxVow {
			maxVow = curVow
		}
	}
	return maxVow
}

func isVowel(c rune) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}