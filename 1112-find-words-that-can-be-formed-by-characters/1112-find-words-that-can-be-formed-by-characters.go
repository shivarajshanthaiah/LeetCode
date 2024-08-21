func countCharacters(words []string, chars string) int {
	freq := make(map[rune]int)

	for _, char := range chars {
		freq[char]++
	}

	totLen := 0

	for _, word := range words {

		dupFreq := make(map[rune]int)

		for key, val := range freq {
			dupFreq[key] = val
		}
		currLen := 0
		canForm := true
		for _, c := range word {
			if dupFreq[c] > 0 {
				dupFreq[c]--
				currLen++
			} else {
				canForm = false
				break
			}
		}
		if canForm {
			totLen += currLen
		}
	}
	return totLen
}