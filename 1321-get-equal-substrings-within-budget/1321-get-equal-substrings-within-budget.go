func equalSubstring(s string, t string, maxCost int) int {
	start := 0
	curCost := 0
	maxLen := 0

	for end := 0; end < len(s); end++ {
		cost := int(s[end]) - int(t[end])
		if cost < 0 {
			cost = -cost
		}
		curCost += cost

		for curCost > maxCost {
			cost = int(s[start]) - int(t[start])
			if cost < 0 {
				cost = -cost
			}
			curCost -= cost
			start++
		}

		if end-start+1 > maxLen {
			maxLen = end - start + 1
		}
	}
	return maxLen

}