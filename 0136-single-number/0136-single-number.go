func singleNumber(nums []int) int {
	c := 0
	for _, n := range nums {
		c ^= n
	}
	return c
}