func maxDistance(arrays [][]int) int {
	if len(arrays) == 0 {
		return 0
	}
	ans := 0
	minVal := arrays[0][0]
	maxVal := arrays[0][len(arrays[0])-1]

	for i := 1; i < len(arrays); i++ {
		n := len(arrays[i])

		curDiff := max(maxVal-arrays[i][0], arrays[i][n-1]-minVal)

		if curDiff > ans {
			ans = curDiff
		}
		minVal = min(minVal, arrays[i][0])
		maxVal = max(maxVal, arrays[i][n-1])

	}
	return ans
}