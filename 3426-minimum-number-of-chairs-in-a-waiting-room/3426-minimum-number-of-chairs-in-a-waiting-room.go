func minimumChairs(s string) int {
	count, maxChairs := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'E' {
			count++
		} else {
			count--
		}
		if count > maxChairs {
			maxChairs = count
		}
	}
	return maxChairs
}