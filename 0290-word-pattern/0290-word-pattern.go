func wordPattern(pattern string, s string) bool {
	pat := make(map[string]string)
	fPat := make(map[string]string)

	str := strings.Split(s, " ")
	if len(pattern) != len(str) {
		return false
	}

	for key, val := range pattern {
		if pat[str[key]] == "" && fPat[string(val)] == "" {
			pat[str[key]] = string(val)
			fPat[string(val)] = str[key]
		} else if pat[str[key]] != string(val) {
			return false
		}
	}
	return true
}