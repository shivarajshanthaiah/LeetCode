func maxDistance(position []int, m int) int {
	sort.Ints(position)
	left, right := 1, position[len(position)-1]-position[0]
	ans := 0
	for left <= right {
		mid := (left + right) / 2
		lastPos, balls := position[0], 1
		for i := 1; i < len(position); i++ {
			if position[i]-lastPos >= mid {
				lastPos = position[i]
				balls++
			}
		}
		if balls >= m {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return ans
}