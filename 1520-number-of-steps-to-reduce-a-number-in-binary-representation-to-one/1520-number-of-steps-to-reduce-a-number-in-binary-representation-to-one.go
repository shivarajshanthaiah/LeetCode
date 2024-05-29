func numSteps(s string) int {

	steps := 0
	carry := 0

	for i := len(s) - 1; i > 0; i-- {
		bit := int(s[i]) - '0'

		if bit+carry == 1 {
			carry = 1
			steps += 2
		} else {
			steps += 1
		}
	}

	if carry == 1 {
		steps += 1
	}
	return steps
}