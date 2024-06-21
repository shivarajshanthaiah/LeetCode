func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	n := len(customers)
	if n == 0 || n != len(grumpy) || minutes > n {
		return 0
	}

	satisfied := 0
	for i := 0; i < n; i++ {
		if grumpy[i] == 0 {
			satisfied += customers[i]
		}
	}

	addSatisfied := 0
	for i := 0; i < minutes; i++ {
		if grumpy[i] == 1 {
			addSatisfied += customers[i]
		}
	}
	maxSatisfied := addSatisfied
	for i := minutes; i < n; i++ {
		if grumpy[i] == 1 {
			addSatisfied += customers[i]
		}
		if grumpy[i-minutes] == 1 {
			addSatisfied -= customers[i-minutes]
		}
		maxSatisfied = max(maxSatisfied, addSatisfied)
	}
	return satisfied + maxSatisfied
}