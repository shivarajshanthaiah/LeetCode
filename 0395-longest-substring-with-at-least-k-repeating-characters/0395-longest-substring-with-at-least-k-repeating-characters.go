func longestSubstring(s string, k int) int {
	freq := map[rune]int{}
	maxVal := 0

	for _, char := range s {
		freq[char]++
	}

	chars := []rune{}
	for key, val := range freq {
		if val < k {
			chars = append(chars, key)
		}
	}

	if len(chars) > 0 {
		for _, str := range strings.Split(s, string(chars[0])) {
			maxVal = max(longestSubstring(str, k), maxVal)
		}
		return maxVal
	}
	return len(s)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}