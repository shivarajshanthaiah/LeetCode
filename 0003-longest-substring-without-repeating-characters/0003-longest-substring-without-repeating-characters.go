func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	for i := 0; i < len(s); i++ {
		unique := make(map[byte]bool)
		curr := 0
		for j := i; j < len(s); j++ {
			if unique[s[j]] {
				break
			}
			unique[s[j]] = true
			curr++
		}
		if curr > maxLen {
			maxLen = curr
		}
	}
	return maxLen
}