func maxArea(height []int) int {
	left, right := 0, len(height)-1
	res := 0

	for left < right {
		smallest := 0
		if height[right] < height[left] {
			smallest = height[right]
			right--
		} else {
			smallest = height[left]
			left++
		}
		curr := (right - left + 1) * smallest
		if curr > res {
			res = curr
		}
	}
	return res
}