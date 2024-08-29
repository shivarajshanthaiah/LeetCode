func containsDuplicate(nums []int) bool {
	dup := make(map[int]int, len(nums))
	for _, v := range nums {
		dup[v]++
	}

	for _, v := range dup {
		if v > 1 {
			return true
            break
		}
	}
	return false
}