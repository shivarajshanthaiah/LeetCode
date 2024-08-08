func sortString(s string) string {
	freq := make(map[rune]int)
	for _, v := range s {
		freq[v]++
	}

	count := len(s)
	res := []rune{}
	for count != 0 {
		for i := 'a'; i <= 'z'; i++ {
			if freq[i] != 0 {
				freq[i]--
				count--
				res = append(res, i)
			}
		}
		for j := 'z'; j >= 'a'; j-- {
			if freq[j] != 0 {
				freq[j]--
				count--
				res = append(res, j)
			}
		}
	}
	return string(res)
}