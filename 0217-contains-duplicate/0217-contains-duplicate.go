func containsDuplicate(nums []int) bool {
	dup := make(map[int]int, len(nums))
	for _, v := range nums {
		dup[v]++
		if dup[v] > 1 {
			return true
		}
	}
	return false
}