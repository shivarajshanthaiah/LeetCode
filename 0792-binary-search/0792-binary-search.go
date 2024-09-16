func search(nums []int, target int) int {

	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}

	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}