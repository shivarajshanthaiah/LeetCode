func distributeCandies(candyType []int) int {
	n := len(candyType) / 2

    sort.Ints(candyType)
	variety := 1
	for i := 1; i < len(candyType); i++ {
		if candyType[i] != candyType[i-1] {
			variety++
		}
	}
	if variety > n {
		return n
	}
	return variety
}