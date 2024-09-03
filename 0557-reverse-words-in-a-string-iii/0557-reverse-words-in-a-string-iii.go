func reverseWords(s string) string {
	s2 := []string{}
	idx := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			if idx != i {
				s2 = append(s2, s[idx:i])
			}
			idx = i + 1

		}
	}

	if idx < len(s) {
		s2 = append(s2, s[idx:])
	}

	for i, word := range s2 {
		reversed := reverseStr(word)
		s2[i] = reversed
	}

	res := ""
	for i := 0; i < len(s2); i++ {
		if i != 0 {
			res += " "
		}
		res += s2[i]
	}

	return res

}

func reverseStr(s string) string {
	runes := []rune(s)
	left, right := 0, len(runes)-1
	for left < right {
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}
	return string(runes)
}