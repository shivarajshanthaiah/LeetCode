func findTheDifference(s string, t string) byte {
	sCount := make(map[rune]int)
	tCount := make(map[rune]int)

	for _, char := range s {
		sCount[char]++
	}

	for _, char := range t {
		tCount[char]++
	}

	for char, count := range tCount {
		if count > sCount[char] {
			return byte(char)
		}
	}
	return 0
}