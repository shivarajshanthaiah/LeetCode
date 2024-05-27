func specialArray(nums []int) int {
	sort.Ints(nums)

	for i := 0; i <= len(nums); i++ {
		count := 0

		for _, num := range nums {
			if num >= i {
				count++
			}
		}
		if count == i {
			return i
		}
	}
	return -1
}