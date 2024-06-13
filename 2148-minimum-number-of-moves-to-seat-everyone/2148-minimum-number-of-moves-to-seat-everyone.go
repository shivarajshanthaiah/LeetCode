func minMovesToSeat(seats []int, students []int) int {
	ans := 0
	sort.Ints(seats)
	sort.Ints(students)

	for i := 0; i < len(students); i++ {
		ans += abs(students[i] - seats[i])
	}
	return ans
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}