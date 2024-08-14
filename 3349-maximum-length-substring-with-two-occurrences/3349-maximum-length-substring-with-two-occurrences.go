func maximumLengthSubstring(s string) int {
	i, j := 0, 0
	ans := 0
	freq := make(map[byte]int)

	for i < (len(s)) {
		freq[s[i]]++
		for freq[s[i]] > 2 {
			freq[s[j]]--
			j++
		}
		ans = max(i-j+1, ans)
        i++
	}
	return ans
}