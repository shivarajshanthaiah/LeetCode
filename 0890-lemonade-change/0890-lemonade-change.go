func lemonadeChange(bills []int) bool {
	if bills[0] != 5 {
		return false
	}
	collection5 := 0
	collection10 := 0

	for i := 0; i < len(bills); i++ {
		if bills[i] == 5 {
			collection5++
		}
		if bills[i] == 10 {
			if collection5 < 1 {
				return false
			} else {
				collection10++
				collection5--
			}
		}

		if bills[i] == 20 {
			if collection5 > 0 && collection10 > 0 {
				collection5--
				collection10--
			} else if collection5 > 2 {
				collection5 -= 3
			} else {
				return false
			}
		}
	}
	return true
}