func longestPalindrome(s string) int {

	freq := make(map[rune]int)
	for _, char := range s {
		freq[char]++
	}

	foundOdd := false
	length := 0

	for _, char := range freq {
		if char%2 == 0 {
			length += char
		} else {
			length += char - 1
			foundOdd = true
		}
	}

	if foundOdd {
		length++
	}
	return length
}