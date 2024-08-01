func countSeniors(details []string) int {
	res := 0
	for _, d := range details {
		if d[11] > '6' || d[11] == '6' && d[12] > '0' {
			res++
		}
	}
	return res
}