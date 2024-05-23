func addStrings(num1 string, num2 string) string {

	num1 = reverseString(num1)
	num2 = reverseString(num2)

	sum := ""
	carry := 0

	maxLen := len(num1)
	if len(num2) > maxLen {
		maxLen = len(num2)
	}

	for i := 0; i < maxLen; i++ {
		digit1 := 0
		digit2 := 0

		if i < len(num1) {
			digit1 = int(num1[i] - '0')
		}
		if i < len(num2) {
			digit2 = int(num2[i] - '0')
		}

		digitSum := digit1 + digit2 + carry
		carry = digitSum / 10

		sum = fmt.Sprintf("%d%s", digitSum%10, sum)
	}
	if carry > 0 {
		sum = fmt.Sprintf("%d%s", carry, sum)

	}
	return sum

}

func reverseString(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}