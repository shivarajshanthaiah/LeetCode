func nthUglyNumber(n int) int {
	if n <= 0 {
		return 0
	}
	uglyNum := make([]int, n)
	uglyNum[0] = 1
	i2, i3, i5 := 0, 0, 0
	next2, next3, next5 := 2, 3, 5

	for i := 1; i < n; i++ {
		nextUgly := min(next2, next3, next5)
		uglyNum[i] = nextUgly

		if nextUgly == next2 {
			i2++
			next2 = uglyNum[i2] * 2
		}

		if nextUgly == next3 {
			i3++
			next3 = uglyNum[i3] * 3
		}
		if nextUgly == next5 {
			i5++
			next5 = uglyNum[i5] * 5
		}
	}
	return uglyNum[n-1]
}

func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}