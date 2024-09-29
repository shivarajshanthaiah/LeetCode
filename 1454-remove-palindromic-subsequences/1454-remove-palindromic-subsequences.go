func removePalindromeSub(s string) int {
	if len(s) == 0 {
		return 0
	}
	if isPal(s) {
		return 1
	}
	return 2
}

func isPal(s string) bool {
	i := 0
	j := len(s) - 1

	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}