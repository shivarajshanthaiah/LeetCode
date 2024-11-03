func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i:]+s[:i] == goal {
			return true
		}
	}
	return false
}