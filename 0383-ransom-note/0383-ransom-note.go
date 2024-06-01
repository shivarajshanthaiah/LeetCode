func canConstruct(ransomNote string, magazine string) bool {

	freq := make(map[rune]int)
	for _, char := range magazine {
		freq[char]++
	}

	for _, char := range ransomNote {
		if freq[char] == 0 {
			return false
		}
		freq[char]--
	}
	return true
}