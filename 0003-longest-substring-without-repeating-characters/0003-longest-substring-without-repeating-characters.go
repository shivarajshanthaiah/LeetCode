func lengthOfLongestSubstring(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		unique := make(map[byte]bool)
		curLen := 0
		for j := i; j < len(s); j++ {
			if unique[s[j]] {
				break
			}
			unique[s[j]] = true
			curLen++
		}
		if curLen > res {
			res = curLen
		}
	}
	return res
}