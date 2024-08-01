func countSeniors(details []string) int {
	res := 0
	for _, v := range details {
		l := len(v)
		ageStr := string(v[l-4]) + string(v[l-3])
		age, _ := strconv.Atoi(ageStr)
		if age > 60 {
			res++
		}
	}
	return res
}