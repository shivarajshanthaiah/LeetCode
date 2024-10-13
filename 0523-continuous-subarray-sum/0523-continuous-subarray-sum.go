func checkSubarraySum(nums []int, k int) bool {
	if len(nums) < 2 {
		return false
	}

	Map := make(map[int]int)
	Map[0] = -1
	sum := 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if k != 0 {
			sum %= k
		}
		if prev, found := Map[sum]; found {
			if i-prev > 1 {
				return true
			}
		} else {
			Map[sum] = i
		}
	}
	return false
}