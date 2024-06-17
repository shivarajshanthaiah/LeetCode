func judgeSquareSum(c int) bool {
	sum := 0
	i := 0
    j := int(math.Sqrt(float64(c)))
	for i <= j {
		sum = i*i + j*j
		if sum == c {
			return true
		} else if sum > c {
			j--
		} else if sum < c {
			i++
		}
	}
	return false
}