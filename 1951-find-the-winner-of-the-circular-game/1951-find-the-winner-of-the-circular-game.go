func findTheWinner(n int, k int) int {
	if k == 1 {
		return n
	}

	queue := []int{}
	for i := 1; i <= n; i++ {
		queue = append(queue, i)
	}

	for len(queue) > 1 {
		for i := 1; i < k; i++ {
			front := queue[0]
			queue = queue[1:]
			queue = append(queue, front)
		}
		queue = queue[1:]
	}
	return queue[0]
}