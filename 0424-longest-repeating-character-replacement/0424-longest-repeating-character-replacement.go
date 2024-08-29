func characterReplacement(s string, k int) int {
	freq := make([]int, 26)
	count := 0
	res := 0
	for i, j := 0, 0; i < len(s); i++ {
		freq[s[i]-'A']++
		count++

		for count-maxFreq(freq) > k {
			freq[s[j]-'A']--
			count--
			j++
		}
		res = max(res, count)
	}
	return res
}

func maxFreq(freq []int) int {
	ans := 0
	for _, v := range freq {
		ans = max(ans, v)
	}
	return ans
}