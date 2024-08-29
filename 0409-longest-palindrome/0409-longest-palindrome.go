// func longestPalindrome(s string) int {

// 	freq := make(map[rune]int)
// 	for _, char := range s {
// 		freq[char]++
// 	}

// 	foundOdd := false
// 	length := 0

// 	for _, char := range freq {
// 		if char%2 == 0 {
// 			length += char
// 		} else {
// 			length += char - 1
// 			foundOdd = true
// 		}
// 	}

// 	if foundOdd {
// 		length++
// 	}
// 	return length
// }

func longestPalindrome(s string) int {
	freq := make(map[rune]bool)
	res := 0

	for _, c := range s {
		if freq[c] {
			delete(freq, c)
			res += 2
		} else {
			freq[c] = true
		}
	}
	if len(freq) > 0 {
		return res + 1
	}
	return res
}