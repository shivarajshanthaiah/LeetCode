func minIncrementForUnique(nums []int) int {
	sort.Ints(nums)
	count := 0
	temp := 0
	for _, num := range nums {
		if temp < num {
			temp = num
		}
		count += temp - num
		temp++
	}
	return count
}