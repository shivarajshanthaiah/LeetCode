func longestCommonPrefix(arr1 []int, arr2 []int) int {
	prefix := make(map[int]int)
	maxLen := 0

	for _, x := range arr1 {
		for x > 0 {
			prefix[x] = 1
			x = x / 10
		}
	}

	for _, y := range arr2 {
		for y > 0 {
			if prefix[y] == 1 && y > maxLen {
				maxLen = y
			}
			y = y / 10
		}
	}

	if maxLen > 0 {
		maxLen = len(strconv.Itoa(maxLen))
	}
	return maxLen
}