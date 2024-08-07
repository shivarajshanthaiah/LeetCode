func findPermutationDifference(s string, t string) int {
	ans := 0
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(t); j++ {
			if s[i] == t[j] {
				ans += diff(i, j)
			}
		}
	}
	return ans
}

func diff(a, b int) int {
	if a >= b {
		return a - b
	}
	return b - a
}