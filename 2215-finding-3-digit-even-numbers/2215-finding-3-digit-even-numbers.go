func findEvenNumbers(digits []int) []int {
	digCount := make(map[int]int)
	for _, digit := range digits {
		digCount[digit]++
	}

	ans := []int{}
	for i := 100; i < 999; i += 2 {
		if canform(i, digCount) {
			ans = append(ans, i)
		}
	}
	return ans
}

func canform(num int, digCount map[int]int) bool {
	tempCount := make(map[int]int)

	for num > 0 {
		digit := num % 10
		tempCount[digit]++
		num /= 10
	}

	for digit, _ := range tempCount {
		if tempCount[digit] > digCount[digit] {
			return false
		}
	}
	return true
}