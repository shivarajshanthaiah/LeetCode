func captureForts(forts []int) int {
	prev := -1
	max := 0

	for i := 0; i < len(forts); i++ {
		if forts[i] != 0 {
			if prev != -1 && forts[prev] != forts[i] {
				max = int(math.Max(float64(max), float64(i-prev-1)))
			}
			prev = i
		}
	}
	return max
}