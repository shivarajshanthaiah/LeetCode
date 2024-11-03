func makeFancyString(s string) string {
	ans := []byte{}
	for i := 0; i < len(s); i++ {
		if i < len(s)-2 && s[i] == s[i+1] && s[i+1] == s[i+2] {
			continue
		} else {
			ans = append(ans, s[i])
		}
	}
	return string(ans)
}