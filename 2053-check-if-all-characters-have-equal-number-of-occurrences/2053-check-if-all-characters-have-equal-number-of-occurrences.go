func areOccurrencesEqual(s string) bool {
	freq := make(map[rune]int)
	for _, c := range s {
		freq[c]++
	}

	freqVal := 0
	for _, v := range freq {
		freqVal = v
		break
	}

	for _, v := range freq {
		if v != freqVal {
			return false
		}
	}
	return true
}