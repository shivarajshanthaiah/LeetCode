func smallerNumbersThanCurrent(nums []int) []int {
	res := []int{}

	for i := 0; i < len(nums); i++ {
		count := 0
		for j := 0; j < len(nums); j++ {
			if j != i && nums[j] < nums[i] {
				count++
			}
		}
		res = append(res, count)
	}
	return res
}