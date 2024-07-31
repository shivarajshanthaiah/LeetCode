func pivotInteger(n int) int {
	sum1, sum2 := 0, 0
	i, j := 1, n
	for i != j {
		if sum1 < sum2 {
			sum1 += i
			i++
		} else {
			sum2 += j
			j--
		}
	}
	if sum1 == sum2 {
		return i
	}
	return -1
}