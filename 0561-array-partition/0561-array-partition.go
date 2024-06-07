func arrayPairSum(nums []int) int {
	res := 0
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if i == 0 || i%2 == 0 {
			res += nums[i]
		}
	}
	return res
}