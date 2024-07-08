type Queue struct {
	items []int
}

func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

func (q *Queue) Dequeue() int {
	deq := q.items[0]
	q.items = q.items[1:]
	return deq
}

func findTheWinner(n int, k int) int {
	q := &Queue{[]int{}}
	for i := 1; i <= n; i++ {
		q.Enqueue(i)
	}

	for len(q.items) > 1 {
		for j := 0; j < k-1; j++ {
			q.Enqueue(q.Dequeue())
		}
		q.Dequeue()
	}
	return q.items[0]
}