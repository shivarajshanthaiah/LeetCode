func checkRecord(s string) bool {
	ab := 0
	late := 0

	for i := 0; i < len(s); i++ {
		if s[i] == 'A' {
			ab++
			if ab > 1 {
				return false
			}
		}
		if s[i] == 'L' {
			late++
			if late > 2 {
				return false
			}
		} else {
			late = 0
		}
	}

	return true
}