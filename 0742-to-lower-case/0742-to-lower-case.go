func toLowerCase(s string) string {
	newStr := []rune(s)

	newStr = convert(newStr)
	return string(newStr)
}

func convert(s []rune) []rune {
	for i := 0; i < len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			s[i] += 'a' - 'A'
		}
	}
	return s
}