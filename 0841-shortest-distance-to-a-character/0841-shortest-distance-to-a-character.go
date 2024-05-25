func shortestToChar(s string, c byte) []int {
	n := len(s)
	result := make([]int, n)
	position := -n

	for i := 0; i < n; i++ {
		if s[i] == c {
			position = i
		}
		result[i] = i - position
	}

	for i := position - 1; i >= 0; i-- {
		if s[i] == c {
			position = i
		}
		if position-i < result[i] {
			result[i] = position - i
		}
	}
	return result
}