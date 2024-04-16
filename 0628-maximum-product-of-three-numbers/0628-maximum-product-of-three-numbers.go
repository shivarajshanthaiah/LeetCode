func maximumProduct(nums []int) int {
	max1, max2, max3 := math.MinInt64, math.MinInt64, math.MinInt64
	min1, min2 := math.MaxInt64, math.MaxInt64

	for _, num := range nums {
		if num > max1 {
			max3, max2, max1 = max2, max1, num
		} else if num > max2 {
			max3, max2 = max2, num
		} else if num > max3 {
			max3 = num
		}

		if num < min1 {
			min2, min1 = min1, num
		} else if num < min2 {
			min2 = num
		}
	}
	return max(max1*max2*max3, max1*min1*min2)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}