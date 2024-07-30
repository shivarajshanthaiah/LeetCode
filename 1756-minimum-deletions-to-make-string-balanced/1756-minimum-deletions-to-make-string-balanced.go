func minimumDeletions(s string) int {
	countB := 0
	deletions := 0

	for _, char := range s {
		if char == 'a' && countB > 0 {
			countB--
			deletions++
		} else if char == 'b' {
			countB++
		}
	}
	return deletions
}