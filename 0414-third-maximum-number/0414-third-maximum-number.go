func thirdMax(nums []int) int {
	var num1, num2, num3 *int

	for _, num := range nums {
		if (num1 != nil && num == *num1) || (num2 != nil && num == *num2 || (num3 != nil && num == *num3)) {
			continue
		}

		if num1 == nil || num > *num1 {
			num3 = num2
			num2 = num1
			num1 = &num
		} else if num2 == nil || num > *num2 {
			num3 = num2
			num2 = &num
		} else if num3 == nil || num > *num3 {
			num3 = &num
		}
	}

	if num3 == nil {
		return *num1
	}
	return *num3
}