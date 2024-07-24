func doesAliceWin(s string) bool {
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' {
			count++
		}
	}
	return count > 0
}