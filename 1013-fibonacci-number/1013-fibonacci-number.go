func fib(n int) int {
	if n < 0 {
		return 0
	}
	if n == 0 || n == 1 {
		return n
	}

	fib1 := fib(n - 1)
	fib2 := fib(n - 2)

	fibunacci := fib1 + fib2
	return fibunacci
}