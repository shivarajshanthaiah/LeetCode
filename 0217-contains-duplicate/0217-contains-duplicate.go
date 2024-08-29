func containsDuplicate(nums []int) bool {
	dup := make(map[int]int)
	for _, v := range nums {
		dup[v]++
	}

	for _, v := range dup {
		if v > 1 {
			return true
		}
	}
	return false
}