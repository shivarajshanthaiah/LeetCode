func minimumAverage(nums []int) float64 {
	sort.Ints(nums)
	res := math.MaxFloat64
	for i := 0; i < len(nums)/2; i++ {
		res = math.Min(res, float64(nums[i]+nums[len(nums)-i-1]))
	}
	return res / 2
}