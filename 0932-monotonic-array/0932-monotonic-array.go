func isMonotonic(nums []int) bool {
	if len(nums) == 1 {
		return true
	}

	inc := true
	dec := true

	for i := 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			inc = false
			break
		}
	}

	for i := 1; i < len(nums); i++ {
		if nums[i-1] < nums[i] {
			dec = false
			break
		}
	}
	return inc || dec
}