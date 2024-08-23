func sumBase(n int, k int) int {
	if n == 0 {
		return 0
	}
	sum := 0
	for n > 0 {
		sum += n % k
		n /= k
	}
	return sum
}