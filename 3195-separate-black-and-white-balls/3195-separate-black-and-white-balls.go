func minimumSteps(s string) int64 {
	one, ans := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			one++
		} else {
			ans += one
		}
	}
	return int64(ans)
}