func minAddToMakeValid(s string) int {
	count := 0
	neg := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			count++
		} else if s[i] == ')' && count > 0 {
			count--
		} else if s[i] == ')' {
			neg++
		}
	}
	return abs(count) + abs(neg)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}