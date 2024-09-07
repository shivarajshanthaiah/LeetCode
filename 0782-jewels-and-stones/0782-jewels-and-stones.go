func numJewelsInStones(jewels string, stones string) int {
	count := make(map[rune]rune)
	for _, j := range jewels {
		count[j] = j
	}
	res := 0
	for _, s := range stones {
		if s == count[s] {
			res++
		}
	}
	return res
}