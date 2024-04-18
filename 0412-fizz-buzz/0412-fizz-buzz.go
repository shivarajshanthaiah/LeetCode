func fizzBuzz(n int) []string {
	answer := make([]string, n)
	for i := 1; i <= n; i++ {
		divisibleBy3 := i%3 == 0
		divisibleBy5 := i%5 == 0

		switch {
		case divisibleBy3 && divisibleBy5:
			answer[i-1] = "FizzBuzz"
		case divisibleBy3:
			answer[i-1] = "Fizz"
		case divisibleBy5:
			answer[i-1] = "Buzz"
		default:
			answer[i-1] = strconv.Itoa(i)
		}
	}
	return answer
}