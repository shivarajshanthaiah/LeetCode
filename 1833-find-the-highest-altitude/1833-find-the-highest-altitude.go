func largestAltitude(gain []int) int {
	maxAlt, curAlt := 0, 0

	for _, value := range gain {
		curAlt += value

		if curAlt > maxAlt {
			maxAlt = curAlt
		}
	}
	return maxAlt
}