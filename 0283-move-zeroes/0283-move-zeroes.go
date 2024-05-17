func moveZeroes(nums []int) {
	foundZero := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[foundZero], nums[i] = nums[i], nums[foundZero]
			foundZero++
		}
	}
}