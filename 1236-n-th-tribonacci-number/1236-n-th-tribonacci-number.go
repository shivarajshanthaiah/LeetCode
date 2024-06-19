func tribonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	t1, t2, t3 := 0, 1, 1

	for i := 3; i <= n; i++ {
		sum := t1 + t2 + t3
		t1, t2, t3 = t2, t3, sum
	}
	return t3
}