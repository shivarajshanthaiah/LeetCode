func detectCapitalUse(word string) bool {
	upper, lower := 0, 0
	for i := 0; i < len(word); i++ {
		if word[i] > 59 && word[i] < 91 {
			upper++
		} else {
			lower++
		}
	}

	if upper == len(word) || lower == len(word) {
		return true
	} else if word[0] < 91 && lower == len(word)-1 {
		return true
	}
	return false
}