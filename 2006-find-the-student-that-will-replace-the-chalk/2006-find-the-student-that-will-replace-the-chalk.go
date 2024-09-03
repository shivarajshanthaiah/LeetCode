func chalkReplacer(chalk []int, k int) int {
	totChalk := 0
	for _, c := range chalk {
		totChalk += c
	}

	k %= totChalk
	for i, c := range chalk {
		if k < c {
			return i
		}
		k -= c
	}
	return -1
}