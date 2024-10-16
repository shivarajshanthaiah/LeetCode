func secondHighest(s string) int {
	nums := make(map[int]bool)

	for i := 0; i < len(s); i++ {
		if unicode.IsDigit(rune(s[i])) {
			num := int(s[i] - '0')
			nums[num] = true
		}
	}

	Max := math.MinInt
	SecMax := math.MinInt

	for key, _ := range nums {
		if key > Max {
			Max, SecMax = key, Max
		} else if key > SecMax && key < Max {
			SecMax = key
		}
	}
	if SecMax == math.MinInt {
		return -1
	}
	return SecMax
}